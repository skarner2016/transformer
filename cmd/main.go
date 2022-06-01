package main

import (
	"transformer/cmd/kernel"
	"os"
)

func main()  {
	if err := kernel.MainCmd.Execute(); err != nil {
		os.Exit(1)
	}
}