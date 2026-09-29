package model

type Wine struct {
	ProductDetails
	Year string
	Kind string
}

func (w Wine) CalculatePrice() int64 {
	return w.Price + w.Price*23/100
}
