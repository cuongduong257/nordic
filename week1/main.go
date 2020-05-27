package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"nordic/week1/helper"
	"os"

	"gopkg.in/yaml.v2"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Start...")
	input, err := os.Open("data.csv")
	check(err)
	defer input.Close()

	output, err := os.Create("output.yml")
	check(err)
	defer output.Close()

	r := csv.NewReader(input)
	w := bufio.NewWriter(output)

	records, err := r.ReadAll()
	check(err)

	yamlText := helper.ConvertCsvToYaml(records)
	marshal, err := yaml.Marshal(yamlText)
	check(err)

	w.WriteString(string(marshal))
	w.Flush()
	fmt.Printf("Done...\n")
}
