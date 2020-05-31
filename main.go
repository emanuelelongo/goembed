package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: goembed [options] file|directory\n\n")
		fmt.Println("Embed static resources in a Go program")
		fmt.Println("For available options type: goembed -help")
		fmt.Println("Example:\n\tgoembed path/to/images > myResources.go")
		os.Exit(1)
	}
	args := os.Args[1:]

	var opt options
	flag.StringVar(&opt.name, "name", "", "The variable name to assign the resource (file mode only)")
	flag.StringVar(&opt.pkg, "package", "main", "The package to output in headers")
	flag.StringVar(&opt.prefix, "prefix", "resource", "A prefix for variables (ignored if name is set)")
	flag.BoolVar(&opt.noheader, "noheader", false, "Don't output headers (useful to append to existing file")

	flag.CommandLine.Parse(args)
	embed(flag.Args()[0], opt)
}
