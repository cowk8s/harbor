package main

import (
	"context"
	"flag"
)

func main() {

	configPath := flag.String("c", "", "Specify the yaml config file path")
	flag.Parse()

	// Missing config file
	if configPath == nil {
		flag.Usage()
		panic("no config file is specified")
	}

	// Load configurations

	// Append node ID
	vCtx := context.WithValue(context.Background(), )
	// Create the root context
	ctx, cancel := context.WithCancel(vCtx)
	defer cancel()

	
}