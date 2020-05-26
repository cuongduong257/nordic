package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

// Ward : Define ward
type Ward struct {
	ID   string `yaml:"id,omitempty"`
	Name string `yaml:"name"`
}

// District : Define Districts
type District struct {
	ID    string `yaml:"id,omitempty"`
	Name  string `yaml:"name"`
	Wards []Ward
}

// City : Define City
type City struct {
	ID        string `yaml:"id,omitempty"`
	Name      string `yaml:"name"`
	Districts []District
}

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
	citys := []City{}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		check(err)

		city := City{
			ID:   record[1],
			Name: record[0],
			Districts: []District{
				{
					ID:   record[3],
					Name: record[2],
					Wards: []Ward{
						{
							ID:   record[5],
							Name: record[4],
						},
					},
				},
			},
		}
		citys = append(citys, city)
	}

	marshal, error := yaml.Marshal(citys)
	check(error)
	w.WriteString(string(marshal))

	w.Flush()
	fmt.Printf("Write done...")
}
