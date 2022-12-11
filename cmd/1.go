package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/lukeab/adventofcode-2022/pkg/config"
)

func main() {

	// input file 1
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	//TODO: make this possibly env/arg driven from config with module constructor style.
	cfg.Inputfile = "inputs/1_test"
	f, err := os.Open(cfg.Inputfile)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
