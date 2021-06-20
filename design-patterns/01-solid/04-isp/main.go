package main

type document struct {
}

type machine interface {
	print(d document)
	fax(d document)
	scan(d document)
}

// ok if you need a multifunction device
type multiFunctionPrinter struct {
	// ...
}

func (m multiFunctionPrinter) print(d document) {

}

func (m multiFunctionPrinter) fax(d document) {

}

func (m multiFunctionPrinter) scan(d document) {

}

type oldFashionedPrinter struct {
	// ...
}

func (o oldFashionedPrinter) print(d document) {
	// ok
}

func (o oldFashionedPrinter) fax(d document) {
	panic("operation not supported")
}

// Deprecated: ...
func (o oldFashionedPrinter) scan(d document) {
	panic("operation not supported")
}

// better approach: split into several interfaces
type printer interface {
	print(d document)
}

type scanner interface {
	scan(d document)
}

// printer only
type myPrinter struct {
	// ...
}

func (m myPrinter) print(d document) {
	// ...
}

// combine interfaces
type photocopier struct{}

func (p photocopier) scan(d document) {
	// ...
}

func (p photocopier) print(d document) {
	// ...
}

type multiFunctionDevice interface {
	printer
	scanner
}

// interface combination + decorator
type multiFunctionMachine struct {
	printer printer
	scanner scanner
}

func (m multiFunctionMachine) print(d document) {
	m.printer.print(d)
}

func (m multiFunctionMachine) scan(d document) {
	m.scanner.scan(d)
}

func main() {
}
