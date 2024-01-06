package iteration

func Iterate(character string, times int) string {
	var repeated string

	if times == 0 {
		times = 5
	}

	for i := 0; i < times; i++ {
		repeated += character
	}

	return repeated
}
