package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/grumpypixel/msfs2020-simconnect-go/filepacker"
)

func main() {
	var infile, outfile, template, packageName, functionName string

	flag.StringVar(&infile, "in", "", "Name of the input file")
	flag.StringVar(&outfile, "out", "", "Name of the output file")
	flag.StringVar(&template, "template", "", "Filename of the template")
	flag.StringVar(&packageName, "package", "main", "Name of the package")
	flag.StringVar(&functionName, "function", "GetData", "Name of the getter function")
	flag.Parse()

	if len(infile) == 0 {
		panic("No input file specified")
	}
	if len(outfile) == 0 {
		panic("No input file specified")
	}
	if len(template) == 0 {
		panic("No template file specified")
	}

	fmt.Println("Input file:", infile)
	fmt.Println("Ouput file:", outfile)
	fmt.Println("Template:", template)
	fmt.Println("Package:", packageName)
	fmt.Println("Function:", functionName)
	now := time.Now()
	timestamp := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d.%d",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second(), now.Nanosecond())
	fmt.Println("Timestamp:", timestamp)
	filepacker.Pack(infile, outfile, template, timestamp, packageName, functionName)
}
