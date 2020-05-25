package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// Ward : Define ward
type Ward struct {
	ID   string
	Name string
}

// District : Define Districts
type District struct {
	ID    string
	Name  string
	Wards []Ward
}

// City : Define City
type City struct {
	ID        string
	Name      string
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

		city := &City{
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
		if citys.record[1] == record[0] {
			fmt.Printf('ddd')
		}
		citys = citys.append(city)

		
		fmt.Print(city)
		s := fmt.Sprintf("- id: %s\n  name: %s\n  districts:\n", record[1], record[0])

		_, wErr := w.WriteString(s)
		check(wErr)
	}

	w.Flush()
	fmt.Printf("Write done...")
}
