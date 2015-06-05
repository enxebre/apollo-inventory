package main

import (
	"flag"
	"os"
	"fmt"
)

type inventoryProvider interface {
    list() map[string]interface{}
	host(string) map[string]string
}

func main() {

	host := flag.String("host", "", "Inventory host mode")
	list := flag.Bool("list", true, "Inventory list mode")
	pretty := flag.Bool("pretty", false, "Inventory output format")
	path := flag.String("path", os.Getenv("TF_VAR_state_path"), "Terraform state file")
	inpr := flag.String("provider", "terraform", "Inventory host mode")
	flag.Parse()

	var p terraformProvider

	if *inpr == "terraform" {
		p = terraformProvider{ Path: *path}	

		if *path == "" {
			fmt.Printf("Usage: %s [options] --path\n", os.Args[0])
			os.Exit(1)
		}
	}

	if *host != "" {
		os.Exit(execute_host(os.Stdout, os.Stderr, p, *pretty, *host))
	} else if *list  {
		os.Exit(execute_list(os.Stdout, os.Stderr, p, *pretty))
	}
}
