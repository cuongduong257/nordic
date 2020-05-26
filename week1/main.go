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
	ID    string  `yaml:"id,omitempty"`
	Name  string  `yaml:"name"`
	Wards []*Ward `yaml:"wards"`
}

// City : Define City
type City struct {
	ID        string      `yaml:"id"`
	Name      string      `yaml:"name"`
	Districts []*District `yaml:"districts"`
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

	data := []*City{}
	cityMap := make(map[interface{}]*City)
	districtMap := make(map[interface{}]*District)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		check(err)

		if record[0] == "Tỉnh Thành Phố" {
			continue
		}

		city, ok := cityMap[record[1]]
		if !ok {
			city = &City{
				ID:   record[1],
				Name: record[0],
			}
			cityMap[record[1]] = city
			data = append(data, city)
		}

		district, ok := districtMap[record[3]]
		if !ok {
			district = &District{
				ID:   record[3],
				Name: record[2],
			}
			districtMap[record[3]] = district
			city.Districts = append(city.Districts, district)
		}

		ward := &Ward{
			ID:   record[5],
			Name: record[4],
		}
		district.Wards = append(district.Wards, ward)
	}

	marshal, error := yaml.Marshal(data)
	check(error)
	w.WriteString(string(marshal))

	w.Flush()
	fmt.Printf("Done...\n")
}
