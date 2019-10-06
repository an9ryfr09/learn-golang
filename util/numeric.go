package an9ryfr09

type Mydata struct {
	Num     int
	Numbers [6]int
}

func (data *Mydata) Add(n int) int {
	return data.Num + n
}

func (data *Mydata) Sum() int {
	n := 0
	for i := 0; i < len(data.Numbers); i++ {
		n += data.Numbers[i]
	}
	return n
}
