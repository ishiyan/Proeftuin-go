package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"

	_ "net/http/pprof"

	"github.com/pkg/errors"
	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"

	"time"

	// "cifar_cnn/cifar"
	"./cifar"

	"gopkg.in/cheggaaa/pb.v1"
)

var (
	epochs  = flag.Int("epochs", 1, "Number of epochs to train for")
	dataset = flag.String("dataset", "train", "Which dataset to train on? Valid options are \"train\" or \"test\"")
	// loc        = flag.String("path", "/pfs/data",
	dtype      = flag.String("dtype", "float64", "Which dtype to use")
	batchsize  = flag.Int("batchsize", 100, "Batch size")
	cpuprofile = flag.String("cpuprofile", "", "CPU profiling")
)

const loc = "./data/"

// const loc = "/pfs/data/train" // Path for Pachyderm data mapping inside container

var dt tensor.Dtype

func parseDtype() {
	switch *dtype {
	case "float64":
		dt = tensor.Float64
	case "float32":
		dt = tensor.Float32
	default:
		log.Fatalf("Unknown dtype: %v", *dtype)
	}
}

type sli struct {
	start, end int
}

func (s sli) Start() int { return s.start }
func (s sli) End() int   { return s.end }
func (s sli) Step() int  { return 1 }

type convnet struct {
	g                  *gorgonia.ExprGraph
	w0, w1, w2, w3, w4 *gorgonia.Node // weights. the number at the back indicates which layer it's used for
	d0, d1, d2, d3     float64        // dropout probabilities

	out     *gorgonia.Node
	predVal gorgonia.Value
}

func newConvNet(g *gorgonia.ExprGraph) *convnet {
	w0 := gorgonia.NewTensor(g, dt, 4, gorgonia.WithShape(32, 3, 5, 5), gorgonia.WithName("w0"), gorgonia.WithInit(gorgonia.GlorotN(1.0)))
	w1 := gorgonia.NewTensor(g, dt, 4, gorgonia.WithShape(64, 32, 5, 5), gorgonia.WithName("w1"), gorgonia.WithInit(gorgonia.GlorotN(1.0)))
	w2 := gorgonia.NewTensor(g, dt, 4, gorgonia.WithShape(128, 64, 5, 5), gorgonia.WithName("w2"), gorgonia.WithInit(gorgonia.GlorotN(1.0)))
	w3 := gorgonia.NewMatrix(g, dt, gorgonia.WithShape(512, 256), gorgonia.WithName("w3"), gorgonia.WithInit(gorgonia.GlorotN(1.0)))
	w4 := gorgonia.NewMatrix(g, dt, gorgonia.WithShape(256, 10), gorgonia.WithName("w4"), gorgonia.WithInit(gorgonia.GlorotN(1.0)))
	return &convnet{
		g:  g,
		w0: w0,
		w1: w1,
		w2: w2,
		w3: w3,
		w4: w4,

		d0: 0.2,
		d1: 0.2,
		d2: 0.2,
		d3: 0.55,
	}
}

func (m *convnet) learnables() gorgonia.Nodes {
	return gorgonia.Nodes{m.w0, m.w1, m.w2, m.w3, m.w4}
}

func (m *convnet) savemodel() (err error) {
	learnables := m.learnables()
	var f io.WriteCloser
	if f, err = os.OpenFile("model.bin", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644); err != nil {
		return
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	for _, l := range learnables {
		t := l.Value().(*tensor.Dense).Data() // []float32
		if err = enc.Encode(t); err != nil {
			return
		}
	}

	return nil
}

// This function is particularly verbose for educational reasons. In reality, you'd wrap up the layers within a layer struct type and perform per-layer activations
func (m *convnet) fwd(x *gorgonia.Node) (err error) {
	var c0, c1, c2, fc *gorgonia.Node
	var a0, a1, a2, a3 *gorgonia.Node
	var p0, p1, p2 *gorgonia.Node
	var l0, l1, l2, l3 *gorgonia.Node

	// LAYER 0
	// here we convolve with stride = (1, 1) and padding = (1, 1),
	// which is your bog standard convolution for convnet
	if c0, err = gorgonia.Conv2d(x, m.w0, tensor.Shape{5, 5}, []int{1, 1}, []int{1, 1}, []int{1, 1}); err != nil {
		return errors.Wrap(err, "Layer 0 Convolution failed")
	}
	if a0, err = gorgonia.Rectify(c0); err != nil {
		return errors.Wrap(err, "Layer 0 activation failed")
	}
	if p0, err = gorgonia.MaxPool2D(a0, tensor.Shape{2, 2}, []int{0, 0}, []int{2, 2}); err != nil {
		return errors.Wrap(err, "Layer 0 Maxpooling failed")
	}
	log.Printf("p0 shape %v", p0.Shape())
	if l0, err = gorgonia.Dropout(p0, m.d0); err != nil {
		return errors.Wrap(err, "Unable to apply a dropout")
	}

	// Layer 1
	if c1, err = gorgonia.Conv2d(l0, m.w1, tensor.Shape{5, 5}, []int{1, 1}, []int{1, 1}, []int{1, 1}); err != nil {
		return errors.Wrap(err, "Layer 1 Convolution failed")
	}
	if a1, err = gorgonia.Rectify(c1); err != nil {
		return errors.Wrap(err, "Layer 1 activation failed")
	}
	if p1, err = gorgonia.MaxPool2D(a1, tensor.Shape{2, 2}, []int{0, 0}, []int{2, 2}); err != nil {
		return errors.Wrap(err, "Layer 1 Maxpooling failed")
	}
	if l1, err = gorgonia.Dropout(p1, m.d1); err != nil {
		return errors.Wrap(err, "Unable to apply a dropout to layer 1")
	}

	// Layer 2
	if c2, err = gorgonia.Conv2d(l1, m.w2, tensor.Shape{5, 5}, []int{1, 1}, []int{1, 1}, []int{1, 1}); err != nil {
		return errors.Wrap(err, "Layer 2 Convolution failed")
	}
	if a2, err = gorgonia.Rectify(c2); err != nil {
		return errors.Wrap(err, "Layer 2 activation failed")
	}
	if p2, err = gorgonia.MaxPool2D(a2, tensor.Shape{2, 2}, []int{0, 0}, []int{2, 2}); err != nil {
		return errors.Wrap(err, "Layer 2 Maxpooling failed")
	}
	log.Printf("p2 shape %v", p2.Shape())

	var r2 *gorgonia.Node
	b, c, h, w := p2.Shape()[0], p2.Shape()[1], p2.Shape()[2], p2.Shape()[3]
	if r2, err = gorgonia.Reshape(p2, tensor.Shape{b, c * h * w}); err != nil {
		return errors.Wrap(err, "Unable to reshape layer 2")
	}
	log.Printf("r2 shape %v", r2.Shape())
	if l2, err = gorgonia.Dropout(r2, m.d2); err != nil {
		return errors.Wrap(err, "Unable to apply a dropout on layer 2")
	}

	// ioutil.WriteFile("tmp.dot", []byte(m.g.ToDot()), 0644)

	// Layer 3
	log.Printf("l2 shape %v", l2.Shape())
	log.Printf("w3 shape %v", m.w3.Shape())
	if fc, err = gorgonia.Mul(l2, m.w3); err != nil {
		return errors.Wrapf(err, "Unable to multiply l2 and w3")
	}
	if a3, err = gorgonia.Rectify(fc); err != nil {
		return errors.Wrapf(err, "Unable to activate fc")
	}
	if l3, err = gorgonia.Dropout(a3, m.d3); err != nil {
		return errors.Wrapf(err, "Unable to apply a dropout on layer 3")
	}

	// output decode
	var out *gorgonia.Node
	if out, err = gorgonia.Mul(l3, m.w4); err != nil {
		return errors.Wrapf(err, "Unable to multiply l3 and w4")
	}
	m.out, err = gorgonia.SoftMax(out)
	gorgonia.Read(m.out, &m.predVal)
	return
}

func main() {
	flag.Parse()
	parseDtype()
	rand.Seed(1337)

	// intercept Ctrl+C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	doneChan := make(chan bool, 1)

	var inputs, targets tensor.Tensor
	var err error

	// go func() {
	// 	log.Println(http.ListenAndServe("localhost:6060", nil))
	// }()

	trainOn := *dataset
	if inputs, targets, err = cifar.Load(trainOn, loc); err != nil {
		log.Fatal(err)
	}

	numExamples := inputs.Shape()[0]
	bs := *batchsize

	if err := inputs.Reshape(numExamples, 3, 32, 32); err != nil {
		log.Fatal(err)
	}
	g := gorgonia.NewGraph()
	x := gorgonia.NewTensor(g, dt, 4, gorgonia.WithShape(bs, 3, 32, 32), gorgonia.WithName("x"))
	y := gorgonia.NewMatrix(g, dt, gorgonia.WithShape(bs, 10), gorgonia.WithName("y"))
	m := newConvNet(g)
	if err = m.fwd(x); err != nil {
		log.Fatalf("%+v", err)
	}
	losses := gorgonia.Must(gorgonia.HadamardProd(gorgonia.Must(gorgonia.Log(m.out)), y))
	cost := gorgonia.Must(gorgonia.Sum(losses))
	cost = gorgonia.Must(gorgonia.Neg(cost))

	// we wanna track costs
	var costVal gorgonia.Value
	gorgonia.Read(cost, &costVal)

	var accuracyGuess float64

	if _, err = gorgonia.Grad(cost, m.learnables()...); err != nil {
		log.Fatal(err)
	}

	// debug
	// ioutil.WriteFile("fullGraph.dot", []byte(g.ToDot()), 0644)
	// log.Printf("%v", prog)
	// logger := log.New(os.Stderr, "", 0)
	// vm := gorgonia.NewTapeMachine(g, gorgonia.BindDualValues(m.learnables()...), gorgonia.WithLogger(logger), gorgonia.WithWatchlist())

	prog, locMap, _ := gorgonia.Compile(g)
	log.Printf("%v", prog)

	vm := gorgonia.NewTapeMachine(g, gorgonia.WithPrecompiled(prog, locMap), gorgonia.BindDualValues(m.learnables()...))
	solver := gorgonia.NewRMSPropSolver(gorgonia.WithBatchSize(float64(bs)))
	defer vm.Close()
	// pprof
	// handlePprof(sigChan, doneChan)

	var profiling bool
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		profiling = true
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	go cleanup(sigChan, doneChan, profiling)

	batches := numExamples / bs
	log.Printf("Batches %d", batches)
	bar := pb.New(batches)
	bar.SetRefreshRate(time.Second)
	bar.SetMaxWidth(80)

	for i := 0; i < *epochs; i++ {
		bar.Prefix(fmt.Sprintf("Epoch %d", i))
		bar.Set(0)
		bar.Start()

		accuracyGuess = 0

		for b := 0; b < batches; b++ {
			start := b * bs
			end := start + bs
			if start >= numExamples {
				break
			}
			if end > numExamples {
				end = numExamples
			}

			var xVal, yVal tensor.Tensor
			if xVal, err = inputs.Slice(sli{start, end}); err != nil {
				log.Fatal("Unable to slice x")
			}

			if yVal, err = targets.Slice(sli{start, end}); err != nil {
				log.Fatal("Unable to slice y")
			}
			if err = xVal.(*tensor.Dense).Reshape(bs, 3, 32, 32); err != nil {
				log.Fatalf("Unable to reshape %v", err)
			}

			gorgonia.Let(x, xVal)
			gorgonia.Let(y, yVal)
			if err = vm.RunAll(); err != nil {
				log.Fatalf("Failed at epoch  %d: %v", i, err)
			}

			// temp for debug
			arrayOutput2 := m.predVal.Data().([]float64)
			yOutput2 := tensor.New(tensor.WithShape(bs, 10), tensor.WithBacking(arrayOutput2))

			for j := 0; j < yVal.Shape()[0]; j++ {

				// get label
				yRowT, _ := yVal.Slice(sli{j, j + 1})
				yRow := yRowT.Data().([]float64)
				var rowLabel int
				var yRowHigh float64

				for k := 0; k < 10; k++ {
					if k == 0 {
						rowLabel = 0
						yRowHigh = yRow[k]
					} else if yRow[k] > yRowHigh {
						rowLabel = k
						yRowHigh = yRow[k]
					}
				}

				// get prediction
				predRowT, _ := yOutput2.Slice(sli{j, j + 1})
				predRow := predRowT.Data().([]float64)
				var rowGuess int
				var predRowHigh float64

				// guess result
				for k := 0; k < 10; k++ {
					if k == 0 {
						rowGuess = 0
						predRowHigh = predRow[k]
					} else if predRow[k] > predRowHigh {
						rowGuess = k
						predRowHigh = predRow[k]
					}
				}

				if rowLabel == rowGuess {
					accuracyGuess += 1.0 / float64(numExamples)
				}
			}

			// end temp

			solver.Step(gorgonia.NodesToValueGrads(m.learnables()))
			vm.Reset()
			bar.Increment()
		}
		log.Printf("Epoch %d | cost %v | acc %v", i, costVal, accuracyGuess)

	}

	// import test data and run more loops
	if inputs, targets, err = cifar.Load("test", loc); err != nil {
		log.Fatal(err)
	}

	batches = inputs.Shape()[0] / bs
	bar = pb.New(batches)
	bar.SetRefreshRate(time.Second)
	bar.SetMaxWidth(80)

	var testActual, testPred []int

	for i := 0; i < 1; i++ {
		bar.Prefix(fmt.Sprintf("Epoch Test"))
		bar.Set(0)
		bar.Start()
		for b := 0; b < batches; b++ {
			start := b * bs
			end := start + bs
			if start >= numExamples {
				break
			}
			if end > numExamples {
				end = numExamples
			}

			var xVal, yVal tensor.Tensor
			if xVal, err = inputs.Slice(sli{start, end}); err != nil {
				log.Fatal("Unable to slice x")
			}

			if yVal, err = targets.Slice(sli{start, end}); err != nil {
				log.Fatal("Unable to slice y")
			}
			if err = xVal.(*tensor.Dense).Reshape(bs, 3, 32, 32); err != nil {
				log.Fatalf("Unable to reshape %v", err)
			}

			gorgonia.Let(x, xVal)
			gorgonia.Let(y, yVal)
			if err = vm.RunAll(); err != nil {
				log.Fatalf("Failed at epoch test: %v", err)
			}

			arrayOutput := m.predVal.Data().([]float64)
			yOutput := tensor.New(tensor.WithShape(bs, 10), tensor.WithBacking(arrayOutput))

			for j := 0; j < yVal.Shape()[0]; j++ {

				// get label
				yRowT, _ := yVal.Slice(sli{j, j + 1})
				yRow := yRowT.Data().([]float64)
				var rowLabel int
				var yRowHigh float64

				for k := 0; k < 10; k++ {
					if k == 0 {
						rowLabel = 0
						yRowHigh = yRow[k]
					} else if yRow[k] > yRowHigh {
						rowLabel = k
						yRowHigh = yRow[k]
					}
				}

				// get prediction
				predRowT, _ := yOutput.Slice(sli{j, j + 1})
				predRow := predRowT.Data().([]float64)
				var rowGuess int
				var predRowHigh float64

				// guess result
				for k := 0; k < 10; k++ {
					if k == 0 {
						rowGuess = 0
						predRowHigh = predRow[k]
					} else if predRow[k] > predRowHigh {
						rowGuess = k
						predRowHigh = predRow[k]
					}
				}

				testActual = append(testActual, rowLabel)
				testPred = append(testPred, rowGuess)
			}

			vm.Reset()
			bar.Increment()
		}
		if err = m.savemodel(); err != nil {
			return
		}
		log.Printf("Epoch Test | cost %v", costVal)

		printIntSlice("testActual.txt", testActual)
		printIntSlice("testPred.txt", testPred)

	}
}

func cleanup(sigChan chan os.Signal, doneChan chan bool, profiling bool) {
	select {
	case <-sigChan:
		log.Println("EMERGENCY EXIT!")
		if profiling {
			log.Println("Stop profiling")
			pprof.StopCPUProfile()
		}
		os.Exit(1)

	case <-doneChan:
		return
	}
}

func handlePprof(sigChan chan os.Signal, doneChan chan bool) {
	var profiling bool
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		profiling = true
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	go cleanup(sigChan, doneChan, profiling)
}

func printIntSlice(filePath string, values []int) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, value := range values {
		fmt.Fprintln(f, value)
	}
	return nil
}