package translocation

func leftTranslocation(num int, bit int) int {
	return num << bit
}

func rightTranslocation(num int, bit int) int {
	return num >> bit
}
