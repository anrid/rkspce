package domain

// Basket holds one or more products.
type Basket struct {
	Items []Item `json:"items"`
}

// Add adds an item to the basket.
func (b *Basket) Add(i ...Item) *Basket {
	b.Items = append(b.Items, i...)
	return b
}

// Total gets the total price of the basket
func (b *Basket) Total() (total float64) {
	for _, i := range b.Items {
		total += i.Price
	}
	total = roundNearest(total)
	return
}

// Quantity returns the quantity of a product in
// the basket.
func (b *Basket) Quantity(code string) (quantity int) {
	for _, i := range b.Items {
		if code == i.Code {
			quantity++
		}
	}
	return
}

// Item is an item in our basket.
type Item struct {
	Code      string  `json:"code"`
	Price     float64 `json:"price"`
	IsSpecial bool    `json:"is_special"`
}
