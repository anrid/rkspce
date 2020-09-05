package domain

// SpecialBuyOneGetOneFree makes every other instance of
// a given product free.
type SpecialBuyOneGetOneFree struct {
	code        string
	description string
	p           Product
}

// Apply this special to the given basket.
func (s SpecialBuyOneGetOneFree) Apply(b *Basket) {
	var nextOneFree bool

	for _, i := range b.Items {
		if i.Code == s.p.Code {
			// Discount every other item of the same product code.
			if nextOneFree {
				i.AddDiscount(s.code, -s.p.Price)
				nextOneFree = false
			} else {
				nextOneFree = true
			}
		}
	}
}

// SpecialQuantityDiscount applies a discount to a product
// once a minimum quantity of that product is in the basket.
type SpecialQuantityDiscount struct {
	code            string
	description     string
	p               Product
	minimumQuantity int
	discount        float64
}

// Apply this special to the given basket.
func (s SpecialQuantityDiscount) Apply(b *Basket) {
	// Skip this special if we cannot find enough
	// of the product in our basket.
	if b.Quantity(s.p.Code) < s.minimumQuantity {
		return
	}

	for _, i := range b.Items {
		if i.Code == s.p.Code {
			// Add a discount to the current item.
			i.AddDiscount(s.code, -s.discount)
		}
	}
}

// SpecialBuyOneGetOtherDiscounted will discount some other
// product if a certain product is in the basket.
type SpecialBuyOneGetOtherDiscounted struct {
	code               string
	description        string
	p                  Product
	other              Product
	limit              int
	discountPercentage float64
}

// Apply this special to the given basket.
func (s SpecialBuyOneGetOtherDiscounted) Apply(b *Basket) {
	// Skip this special if we cannot find enough of
	// the product in our basket.
	if b.Quantity(s.p.Code) == 0 {
		return
	}

	isLimited := s.limit > 0
	limit := s.limit

	for _, i := range b.Items {
		if i.Code == s.other.Code {
			// Check if we need to limit the number of discounts given.
			if isLimited {
				if limit == 0 {
					continue
				}
				limit--
			}

			// Discount the current item.
			discountedPrice := roundNearest(s.other.Price * s.discountPercentage)

			i.AddDiscount(s.code, -discountedPrice)
		}
	}
}
