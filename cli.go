package main

import (
	"encoding/json"
	"fmt"
	"io"
)

func execute_list(stdout io.Writer, stderr io.Writer, p inventoryProvider, pretty bool) int {
	groups := p.list()
	return output(stdout, stderr, groups, pretty)
}

func execute_host(stdout io.Writer, stderr io.Writer, p inventoryProvider, pretty bool, hostName string) int {
	groups := p.host(hostName)
	return output(stdout, stderr, groups, pretty)
}

func output(stdout io.Writer, stderr io.Writer, content interface{}, pretty bool) int {
	var err error
	var b []byte

	if pretty {
		b, err = json.MarshalIndent(content, "", "  ")	
	} else {
		b, err = json.Marshal(content)		
	}

	if err != nil {
		fmt.Fprintf(stderr, "Error encoding JSON: %s\n", err)
		return 1
	}

	_, err = stdout.Write(b)
	if err != nil {
		fmt.Fprintf(stderr, "Error writing JSON: %s\n", err)
		return 1
	}

	return 0
}