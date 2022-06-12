package matematika

func ganjilGenap(number int) string {
	if number%2 == 0 {
		return "genap"
	} else {
		return "ganjil"
	}
}

func CekGanjilGenap(number ...int) string {
	result := ""
	for index, val := range number {
		if index != 0 {
			result += ", "
		}
		result += ganjilGenap(val)
	}

	return result
}
