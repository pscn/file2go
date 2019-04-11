package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/pscn/file2go/encode"
	"github.com/pscn/file2go/template"
)

var (
	prefix = flag.String("prefix", "", "the prefix to use for the methods")
	pkg    = flag.String("pkg", "",
		"package name to use, defaults to the base directory of the file")
	verbose = flag.Bool("verbose", false, "be more verbose")
)

// Usage shows a short usage summary on stderr
func Usage(msg string) {
	if msg != "" {
		fmt.Fprintf(os.Stderr, "%s\n", msg)
	}
	fmt.Fprintf(os.Stderr, "Usage of file2go:\n")
	fmt.Fprintf(os.Stderr, "\tfile2go -prefix A filename\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

// FlagUsage shows a short usage summary on stderr
func FlagUsage() {
	Usage("")
}

func main() {
	log.SetFlags(0) // log.LstdFlags | log.Lshortfile)
	log.SetPrefix("file2go: ")
	flag.Usage = FlagUsage
	flag.Parse()
	if len(flag.Args()) != 1 {
		fmt.Println(len(flag.Args()))
		Usage("please specify a filename")
		os.Exit(2)
	}
	if *prefix == "" {
		Usage("please specify an prefix")
		os.Exit(2)
	}
	filename := flag.Args()[0]
	extension := path.Ext(filename)
	name := path.Base(filename[0 : len(filename)-len(extension)])
	directory := path.Dir(filename)
	if *pkg == "" {
		*pkg = path.Base(directory)
	}
	target := path.Join(path.Dir(filename), name+".go")
	if *verbose {
		log.Printf("converting: %s => %s\n", filename, target)
	}

	encoded, err := encode.File(filename, "encoded by file2go")
	if err != nil {
		log.Fatalf("failed to encode file: filename=%s, error=%s", filename, err)
	}
	tmpl, err := template.Generate(filename, *pkg, *prefix, &encoded)
	if err != nil {
		log.Fatalf("failed to generate code: %s\n", err)
	}
	if _, err := os.Stat(target); err == nil { // file exists
		data, err := ioutil.ReadFile(target)
		if err != nil {
			log.Fatalf("failed to read old file: %s; error=%s", target, err)
		}
		if string(data) == string(tmpl) {
			if *verbose {
				log.Printf("already up to date: %s", target)
			}
			os.Exit(0)
		}
	}
	file, err := os.Create(target)
	if err != nil {
		log.Fatalf("failed to open file: %s\n", err)
	}
	defer file.Close()
	_, err = file.Write(tmpl)
	if err != nil {
		log.Fatalf("failed to write code to file: %s; error=%s", target, err)
	}

}
