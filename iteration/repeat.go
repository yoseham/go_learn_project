package iteration

func Repeat(character string, repeatnum int) string{
	var repeated string
	for i := 0; i<repeatnum; i++{
		repeated += character
	}
	return repeated
}
