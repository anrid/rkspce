package domain

// Product is a product offered at the local farmer’s market
// every week.
type Product struct {
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// ToItem creates a new basket item from this product.
func (p Product) ToItem() Item {
	return Item{Code: p.Code, Price: p.Price}
}
