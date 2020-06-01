package helper

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

// ConvertCsvToYaml : conver csv data to yaml text
func ConvertCsvToYaml(records [][]string) (yamlText []City) {
	data := []*City{}
	cityMap := make(map[interface{}]*City)
	districtMap := make(map[interface{}]*District)

	for i, record := range records {
		// Though to second record

		if i == 0 {
			continue
		}
		// Though error record
		if len(record) < 6 {
			continue
		}
		cityName := record[0]
		cityCode := record[1]
		districtName := record[2]
		districtCode := record[3]
		wardName := record[4]
		wardCode := record[5]

		city, ok := cityMap[cityCode]
		if !ok {
			city = &City{
				ID:   string(cityCode),
				Name: string(cityName),
			}
			cityMap[cityCode] = city
			data = append(data, city)
		}

		district, ok := districtMap[districtCode]
		if !ok {
			district = &District{
				ID:   string(districtCode),
				Name: string(districtName),
			}
			districtMap[districtCode] = district
			city.Districts = append(city.Districts, district)
		}

		ward := &Ward{
			ID:   string(wardCode),
			Name: string(wardName),
		}
		district.Wards = append(district.Wards, ward)
	}

	cities := []City{}
	for _, c := range data {
		cities = append(cities, *c)
	}
	return cities
}


