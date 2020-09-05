package domain

import (
	"fmt"
	"strings"

	"github.com/rs/xid"
)

// Basket holds one or more products.
type Basket struct {
	ID    string  `json:"id"`
	Items []*Item `json:"items"`
}

// NewBasket creates an new basket with a
// unique id.
func NewBasket() *Basket {
	return &Basket{
		ID: xid.New().String(),
	}
}

// Add adds an item to the basket.
func (b *Basket) Add(i ...*Item) *Basket {
	b.Items = append(b.Items, i...)
	return b
}

// Total gets the total price of the basket
func (b *Basket) Total() (total float64) {
	for _, i := range b.Items {
		total += i.Price
		for _, d := range i.Discounts {
			total += d.Price
		}
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

// Dump returns the basket as a pretty-printed string
// (same format as in the challenge readme).
func (b *Basket) Dump() string {
	row := func(code, special string, price float64) string {
		priceStr := fmt.Sprintf("%.02f", price)
		return fmt.Sprintf("%-12s%-12s%11s", code, special, priceStr)
	}

	// Add headers.
	rows := []string{
		"Item                          Price",
		"----                          -----",
	}

	// Add items.
	for _, i := range b.Items {
		rows = append(rows, row(i.Code, "", i.Price))
		// Add discounts.
		for _, d := range i.Discounts {
			rows = append(rows, row("", d.Code, d.Price))
		}
	}

	// Add footer.
	rows = append(rows, "-----------------------------------")
	rows = append(rows, row("", "", b.Total()))

	return strings.Join(rows, "\n") + "\n"
}

// Item is an item in our basket.
type Item struct {
	Code      string     `json:"code"`
	Price     float64    `json:"price"`
	Discounts []Discount `json:"discounts"`
}

// Discount lowers the price of an item in a basket.
type Discount struct {
	Code  string  `json:"code"`
	Price float64 `json:"price"`
}

// AddDiscount adds a discount on this item.
func (i *Item) AddDiscount(code string, price float64) {
	for _, d := range i.Discounts {
		if d.Code == code {
			return
		}
	}
	i.Discounts = append(i.Discounts, Discount{code, price})
}
