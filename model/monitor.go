package model

type Monitor struct {
	ProductDetails
	Size      string
	Resolutin string
}

func (m Monitor) CalculatePrice() int64 {
	return m.Price + m.Price*30/100
}
