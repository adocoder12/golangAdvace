package model

type Shirt struct {
	ProductDetails
	Size  string
	Color string
}

func (s Shirt) CalculatePrice() int64 {
	return s.Price - s.Price*20/100
}
