package conversion

import "strconv"

func StringsToFloat(str []string) ([]float64, error) {
	values := []float64{}
	for _, stringVal := range str {
		floatVal, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {
			return nil, err
		}
		values = append(values, floatVal)
	}
	return values, nil
}
