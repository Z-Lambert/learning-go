package iteration

func Repeat(character string, count int) string {
	accumulator := ""
	for i := 0; i < count; i++ {
		accumulator += character
	}
	return accumulator
}
