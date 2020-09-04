package domain

import "math"

// All our products.
var (
	ProductChai    = Product{"CH1", "Chai", 3.11}
	ProductApples  = Product{"AP1", "Apples", 6.00}
	ProductCoffee  = Product{"CF1", "Coffee", 11.23}
	ProductMilk    = Product{"MK1", "Milk", 4.75}
	ProductOatmeal = Product{"OM1", "Oatmeal", 3.69}
)

// All our specials.
var (
	SpecialCoffee = SpecialBuyOneGetOneFree{
		"BOGO",
		"Buy-One-Get-One-Free Special on Coffee. (Unlimited)",
		ProductCoffee,
	}
	SpecialApples = SpecialQuantityDiscount{
		"APPL",
		"If you buy 3 or more bags of Apples, the price drops to $4.50.",
		ProductApples,
		3,
		1.50,
	}
	SpecialChai = SpecialBuyOneGetOtherDiscounted{
		"CHMK",
		"Purchase a box of Chai and get milk free. (Limit 1)",
		ProductChai,
		ProductMilk,
		1,
		1.0,
	}
	SpecialOatmeal = SpecialBuyOneGetOtherDiscounted{
		"APOM",
		"Purchase a bag of Oatmeal and get 50% off a bag of Apples",
		ProductOatmeal,
		ProductApples,
		0,
		0.50,
	}
)

// Round float64 to nearest.
func roundNearest(v float64) float64 {
	return math.Round(v*100) / 100
}
