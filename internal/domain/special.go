package domain

// Special is a special offer.
type Special interface {
	Code() string
	Description() string
	Apply(*Basket)
}

// FromSpecial creates a new basket item from a special.
func FromSpecial(s Special, price float64) Item {
	return Item{Code: s.Code(), Price: price, IsSpecial: true}
}

// SpecialBuyOneGetOneFree makes every other instance of
// a given product free.
type SpecialBuyOneGetOneFree struct {
	code        string
	description string
	p           Product
}

// Code ...
func (s SpecialBuyOneGetOneFree) Code() string {
	return s.code
}

// Description ...
func (s SpecialBuyOneGetOneFree) Description() string {
	return s.description
}

// Apply this special to the given basket.
func (s SpecialBuyOneGetOneFree) Apply(b *Basket) {
	nb := &Basket{}
	var nextOneFree bool

	for _, i := range b.Items {
		nb.Add(i)
		if i.Code == s.p.Code {
			// Discount every other item of the same product code.
			if nextOneFree {
				// Make the current item free!
				nb.Add(FromSpecial(s, -s.p.Price))
				nextOneFree = false
			} else {
				nextOneFree = true
			}
		}
	}
	b.Items = nb.Items
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

// Code ...
func (s SpecialQuantityDiscount) Code() string {
	return s.code
}

// Description ...
func (s SpecialQuantityDiscount) Description() string {
	return s.description
}

// Apply this special to the given basket.
func (s SpecialQuantityDiscount) Apply(b *Basket) {
	// Skip this special if we cannot find enough
	// of the product in our basket.
	if b.Quantity(s.p.Code) < s.minimumQuantity {
		return
	}

	nb := &Basket{}
	for _, i := range b.Items {
		nb.Add(i)
		if i.Code == s.p.Code {
			// Add a discount to the current item.
			nb.Add(FromSpecial(s, -s.discount))
		}
	}
	b.Items = nb.Items
}

// SpecialBuyOneGetOtherDiscounted will make some (other) product
// free if a certain product is in the basket.
type SpecialBuyOneGetOtherDiscounted struct {
	code               string
	description        string
	p                  Product
	other              Product
	limit              int
	discountPercentage float64
}

// Code ...
func (s SpecialBuyOneGetOtherDiscounted) Code() string {
	return s.code
}

// Description ...
func (s SpecialBuyOneGetOtherDiscounted) Description() string {
	return s.description
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

	nb := &Basket{}
	for _, i := range b.Items {
		nb.Add(i)
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

			nb.Add(FromSpecial(s, -discountedPrice))
		}
	}
	b.Items = nb.Items
}
