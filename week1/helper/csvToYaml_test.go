package helper

import (
	"reflect"
	"testing"
)

func TestConvertCsvToYaml(t *testing.T) {
	type args struct {
		records [][]string
	}
	tests := []struct {
		name         string
		args         args
		wantYamlText []*City
	}{
		{
			name: "Test nil records string",
			args: args{records: [][]string{
				{"Thành phố Hà Nội", "01", "Quận Ba Đình", "001", "Phường Kim Mã", "00028", "Phường,"},
				{"Thành phố Hà Nội", "01", "Quận Ba Đình", "001", "Phường Giảng Võ", "00031", "Phường,"},
			}},
			wantYamlText: []*City{
				{
					ID: "01",
					Name: "Thành phố Hà Nội",
					Districts: []*District{
						{
							ID: "001",
							Name: "Quận Ba Đình",
							Wards: []*Ward{
								{
									ID: "00028",
									Name: "Phường Kim Mã",
								},
								{
									ID: "00031",
									Name: "Phường Giảng Võ",
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotYamlText := ConvertCsvToYaml(tt.args.records); !reflect.DeepEqual(gotYamlText, tt.wantYamlText) {
				t.Errorf("ConvertCsvToYaml() = %v, want %v", gotYamlText, tt.wantYamlText)
			}
		})
	}
}
