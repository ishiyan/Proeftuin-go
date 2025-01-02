package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"euronext/euronext"
)

const configFileName = "eodarc2014.json"

type config struct {
	Inputs        string `json:"inputs"`
	Downloads     string `json:"downloads"`
	XmlInstrumnts string `json:"xmlInstruments"`
}

type instrument struct {
	Mnemonic string `json:"mnemonic"`
	Mep      string `json:"mep"`
	Mic      string `json:"mic"`
	Isin     string `json:"isin"`
	Type     string `json:"type"`
}

type download struct {
	FileName string
	Content  []byte
}

func main() {
	t := time.Now().Format("2006-01-02 15-04-05")
	fmt.Println("=======================================")
	fmt.Println(t)
	fmt.Println("=======================================")

	cfg, err := readConfig(configFileName)
	if err != nil {
		panic(fmt.Sprintf("cannot read configuration file %s: %s", configFileName, err))
	}

	err = euronext.EnsureDirectoryExists(cfg.Downloads)
	if err != nil {
		panic(fmt.Sprintf("cannot create directory %s: %s", cfg.Downloads, err))
	}

	fmt.Println("xml file: " + cfg.XmlInstrumnts)
	instruments, err := readInstruments(cfg.XmlInstrumnts)
	if err != nil {
		panic(fmt.Sprintf("cannot read instruments: %s", err))
	}

	l := len(instruments)
	for i, ins := range instruments {
		err := ins.archive(cfg, i, l)
		if err != nil {
			fmt.Printf("\n%s\n", err)
		}
	}

	fmt.Println("\nfinished " + time.Now().Format("2006-01-02 15-04-05"))
}

func readConfig(fileName string) (*config, error) {
	var conf config

	f, err := os.Open(fileName)
	if err != nil {
		return &conf, fmt.Errorf("cannot open '%s' file: %w", fileName, err)
	}
	defer f.Close()

	decoder := json.NewDecoder(f)

	err = decoder.Decode(&conf)
	if err != nil {
		return &conf, fmt.Errorf("cannot decode '%s' file: %w", fileName, err)
	}

	if !strings.HasSuffix(conf.Inputs, "/") {
		conf.Inputs += "/"
	}

	if !strings.HasSuffix(conf.Downloads, "/") {
		conf.Downloads += "/"
	}

	return &conf, nil
}

func readInstruments(fileName string) ([]instrument, error) {
	instruments := []instrument{}
	instrs, err := euronext.ReadXmlInstrumentsFile(fileName)
	if err != nil {
		return instruments, fmt.Errorf("cannot read instruments xml file '%s': %w", fileName, err)
	}

	//euronext.WriteJsonInstrumentsFile("output.json", instruments)
	//euronext.WriteXmlInstrumentsFile("output.xml", instruments)
	fmt.Println(len(instrs.Instrument), "instruments read from", fileName)
	for _, inst := range instrs.Instrument {
		ins := instrument{
			Mnemonic: strings.ToLower(inst.Symbol),
			Mep:      strings.ToLower(inst.Mep),
			Mic:      strings.ToLower(inst.Mic),
			Isin:     strings.ToLower(inst.Isin),
			Type:     strings.ToLower(inst.Type),
		}
		instruments = append(instruments, ins)
	}

	return instruments, nil
}

func (s *instrument) safeMnemonic() string {
	mnemonic := s.Mnemonic
	if mnemonic == "prn" || mnemonic == "com" || mnemonic == "lpt" || mnemonic == "aux" {
		mnemonic += "_"
	}

	return mnemonic
}

func (s *instrument) fileFolder() string {
	return fmt.Sprintf("%s/%s/%s/endofday/", s.Mic, s.Type, s.safeMnemonic())
}

func (s *instrument) fileOutputPrefix() string {
	return fmt.Sprintf("%s_%s_%s_", s.Mic, s.Mnemonic, s.Isin)
}

func (s *instrument) fileInputPrefix() string {
	return fmt.Sprintf("%s_%s_%s_", strings.ToUpper(s.Mic), strings.ToUpper(s.Mnemonic), strings.ToUpper(s.Isin))
}

func (s *instrument) archive(cfg *config, el, elen int) error {
	insFolder := s.fileFolder()
	insOutputPrefix := s.fileOutputPrefix()
	insInputPrefix := s.fileInputPrefix()
	downloads := make([]download, 0)
	log := fmt.Sprintf("(%d of %d) %s to %s ... ", el+1, elen, insInputPrefix, insFolder)

	startDate := time.Date(2013, time.December, 28, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC)
	for date := startDate; date.Before(endDate); date = date.AddDate(0, 0, 1) {
		sd := date.Format("20060102")
		file := fmt.Sprintf("%s%s/%s%s_eoh.js", cfg.Inputs, sd, insInputPrefix, sd)
		if _, err := os.Stat(file); err == nil {
			bs, err := os.ReadFile(file)
			if err != nil {
				return fmt.Errorf("%s cannot read file '%s': %w", log, file, err)
			}
			sd = date.Format("2006-01-02")
			file = fmt.Sprintf("%s%s_unadjusted.json", insOutputPrefix, sd)
			downloads = append(downloads, download{FileName: file, Content: bs})
		}
		sd = date.Format("20060102")
		file = fmt.Sprintf("%s%s/%s%s_eoh.js.adjusted", cfg.Inputs, sd, insInputPrefix, sd)
		if _, err := os.Stat(file); err == nil {
			bs, err := os.ReadFile(file)
			if err != nil {
				return fmt.Errorf("%s cannot read file '%s': %w", log, file, err)
			}
			sd = date.Format("2006-01-02")
			file = fmt.Sprintf("%s%s_adjusted.json", insOutputPrefix, sd)
			downloads = append(downloads, download{FileName: file, Content: bs})
		}
	}

	if len(downloads) == 0 {
		fmt.Println(log + "no files found")
		return nil
	}

	file := cfg.Downloads + insFolder
	err := euronext.EnsureDirectoryExists(file)
	if err != nil {
		return fmt.Errorf("%s cannot create instrument download directory '%s': %w", log, file, err)
	}

	file = fmt.Sprintf("%s%s%s2014.zip", cfg.Downloads, insFolder, insOutputPrefix)
	z, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("%s cannot create zip file '%s': %w", log, file, err)
	}
	defer z.Close()

	w := zip.NewWriter(z)
	defer w.Close()

	for _, d := range downloads {
		f, err := w.Create(d.FileName)
		if err != nil {
			return fmt.Errorf("%s cannot create zip entry '%s': %w", log, d.FileName, err)
		}

		_, err = f.Write(d.Content)
		if err != nil {
			return fmt.Errorf("%s cannot write zip entry '%s': %w", log, d.FileName, err)
		}
	}

	fmt.Println(log + " " + fmt.Sprint(len(downloads)) + " archived")
	return nil
}
