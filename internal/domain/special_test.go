package domain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSpecialBuyOneGetOneFree(t *testing.T) {
	// No discount.
	{
		b := NewBasket().Add(ProductCoffee.ToItem())

		s := SpecialBuyOneGetOneFree{"X", "Y", ProductCoffee}
		s.Apply(b)

		require.Equal(t, ProductCoffee.Price, b.Total())
	}

	// No discount.
	{
		b := NewBasket().Add(ProductChai.ToItem(), ProductCoffee.ToItem())

		s := SpecialBuyOneGetOneFree{"X", "Y", ProductCoffee}
		s.Apply(b)

		require.Equal(t, ProductChai.Price+ProductCoffee.Price, b.Total())
	}

	// Coffee buy-one-get-one-free discount.
	{
		b := NewBasket().Add(ProductCoffee.ToItem(), ProductChai.ToItem(), ProductCoffee.ToItem(), ProductOatmeal.ToItem())

		s := SpecialBuyOneGetOneFree{"X", "Y", ProductCoffee}
		s.Apply(b)
		s.Apply(b) // Double apply here on purpose!

		require.Equal(t, ProductChai.Price+ProductCoffee.Price+ProductOatmeal.Price, b.Total())
	}

	// Oatmeal buy-one-get-one-free discount.
	{
		b := NewBasket().Add(ProductOatmeal.ToItem(), ProductOatmeal.ToItem(), ProductOatmeal.ToItem(), ProductOatmeal.ToItem())

		s := SpecialBuyOneGetOneFree{"X", "Y", ProductOatmeal}
		s.Apply(b)

		require.Equal(t, ProductOatmeal.Price+ProductOatmeal.Price, b.Total())
	}
}

func TestSpecialQuantityDiscount(t *testing.T) {
	// No discount.
	{
		b := NewBasket().Add(ProductCoffee.ToItem())

		s := SpecialQuantityDiscount{"X", "Y", ProductCoffee, 3, 1.50}
		s.Apply(b)

		require.Equal(t, ProductCoffee.Price, b.Total())
	}

	// No discount.
	{
		b := NewBasket().Add(ProductChai.ToItem(), ProductCoffee.ToItem())

		s := SpecialQuantityDiscount{"X", "Y", ProductCoffee, 3, 1.50}
		s.Apply(b)

		require.Equal(t, ProductChai.Price+ProductCoffee.Price, b.Total())
	}

	// Coffee quanity (>2) discount.
	{
		b := NewBasket().Add(ProductCoffee.ToItem(), ProductChai.ToItem(), ProductCoffee.ToItem(), ProductCoffee.ToItem())

		s := SpecialQuantityDiscount{"X", "Y", ProductCoffee, 3, 1.50}
		s.Apply(b)
		s.Apply(b) // Double apply here on purpose!

		require.Equal(t, roundNearest(ProductChai.Price+((ProductCoffee.Price-1.50)*3)), b.Total())
	}

	// Oatmeal quantity (>1) discount.
	{
		b := NewBasket().Add(ProductOatmeal.ToItem(), ProductOatmeal.ToItem(), ProductOatmeal.ToItem(), ProductOatmeal.ToItem())

		s := SpecialQuantityDiscount{"X", "Y", ProductOatmeal, 2, 2.66}
		s.Apply(b)

		require.Equal(t, roundNearest((ProductOatmeal.Price-2.66)*4), b.Total())
	}
}

func TestSpecialBuyOneGetOtherDiscounted(t *testing.T) {
	// No discount.
	{
		b := NewBasket().Add(ProductChai.ToItem(), ProductCoffee.ToItem())

		s := SpecialBuyOneGetOtherDiscounted{"X", "Y", ProductChai, ProductMilk, 1, 1.0}
		s.Apply(b)

		require.Equal(t, ProductChai.Price+ProductCoffee.Price, b.Total())
	}

	// Buy chai get milk free discount.
	{
		b := NewBasket().Add(ProductChai.ToItem(), ProductMilk.ToItem(), ProductMilk.ToItem())

		s := SpecialBuyOneGetOtherDiscounted{"X", "Y", ProductChai, ProductMilk, 1, 1.0}
		s.Apply(b)

		require.Equal(t, roundNearest(ProductChai.Price+ProductMilk.Price), b.Total())
	}

	// Buy apples get oatmeal free discount.
	{
		b := NewBasket().Add(ProductApples.ToItem(), ProductOatmeal.ToItem(), ProductOatmeal.ToItem(), ProductChai.ToItem())

		s := SpecialBuyOneGetOtherDiscounted{"X", "Y", ProductApples, ProductOatmeal, 2, 1.0}
		s.Apply(b)
		s.Apply(b) // Double apply here on purpose!

		require.Equal(t, roundNearest(ProductApples.Price+ProductChai.Price), b.Total())
	}

	// Buy oatmeal get apples (unlimited) at a 50% discount.
	{
		b := NewBasket().Add(ProductApples.ToItem(), ProductOatmeal.ToItem(), ProductApples.ToItem(), ProductOatmeal.ToItem())

		s := SpecialBuyOneGetOtherDiscounted{"X", "Y", ProductOatmeal, ProductApples, 0, 0.5}
		s.Apply(b)

		require.Equal(t, roundNearest((ProductApples.Price*0.5*2)+(ProductOatmeal.Price*2)), b.Total())
	}

	// Buy oatmeal get apples (limit 1) at a 35% discount.
	{
		b := NewBasket().Add(ProductApples.ToItem(), ProductOatmeal.ToItem(), ProductApples.ToItem(), ProductOatmeal.ToItem())

		s := SpecialBuyOneGetOtherDiscounted{"X", "Y", ProductOatmeal, ProductApples, 1, 0.35}
		s.Apply(b)

		require.Equal(t, roundNearest(ProductApples.Price+(ProductApples.Price*0.65)+(ProductOatmeal.Price*2)), b.Total())
	}
}

func TestCurrentFarmersMarketSpecials(t *testing.T) {
	// Look up product just to get 100% code coverage.
	{
		p, err := GetProductByCode("OM1")
		require.Equal(t, "OM1", p.Code)

		_, err = GetProductByCode("WHATEVAH")
		require.Error(t, err)
		require.Equal(t, ErrNotFound, err)
	}

	CH1 := func() *Item { return ProductChai.ToItem() }
	AP1 := func() *Item { return ProductApples.ToItem() }
	CF1 := func() *Item { return ProductCoffee.ToItem() }
	MK1 := func() *Item { return ProductMilk.ToItem() }
	OM1 := func() *Item { return ProductOatmeal.ToItem() }

	// Basket: CH1, AP1
	// Total price expected: $9.11
	{
		b := NewBasket().Add(CH1(), AP1())

		ApplySpecials(b)

		require.Equal(t, 9.11, b.Total())

		require.Equal(t, `
Item                          Price
----                          -----
CH1                            3.11
AP1                            6.00
-----------------------------------
                               9.11
`[1:], b.Dump())
	}

	// Basket: CH1, AP1, CF1, MK1
	// Total price expected: $20.34
	{
		b := NewBasket().Add(CH1(), AP1(), CF1(), MK1())

		ApplySpecials(b)

		require.Equal(t, 20.34, b.Total())
	}

	// Basket: MK1, AP1
	// Total price expected: $10.75
	{
		b := NewBasket().Add(MK1(), AP1())

		ApplySpecials(b)

		require.Equal(t, 10.75, b.Total())
	}

	// Basket: CF1, CF1
	// Total price expected: $11.23
	{
		b := NewBasket().Add(CF1(), CF1())

		ApplySpecials(b)

		require.Equal(t, 11.23, b.Total())
	}

	// Basket: AP1, AP1, CH1, AP1
	// Total price expected: $16.61
	{
		b := NewBasket().Add(AP1(), AP1(), CH1(), AP1())

		ApplySpecials(b)

		require.Equal(t, 16.61, b.Total())
	}

	// Basket: CH1, AP1, AP1, AP1, MK1
	// Total price expected: $16.61
	{
		b := NewBasket().Add(CH1(), AP1(), AP1(), AP1(), MK1())

		ApplySpecials(b)

		require.Equal(t, 16.61, b.Total())
		require.Equal(t, `
Item                          Price
----                          -----
CH1                            3.11
AP1                            6.00
            APPL              -1.50
AP1                            6.00
            APPL              -1.50
AP1                            6.00
            APPL              -1.50
MK1                            4.75
            CHMK              -4.75
-----------------------------------
                              16.61
`[1:], b.Dump())
	}

	// Basket: OM1, AP1 -> OM1, AP1, AP1 -> OM1, AP1, AP1, AP1
	// Total price expected: $6.69 -> $12.69 -> $14.19
	{
		b := NewBasket().Add(OM1(), AP1())

		ApplySpecials(b)

		require.Equal(t, 6.69, b.Total())

		// Add one more pack of apples.
		b.Add(AP1())

		ApplySpecials(b)

		require.Equal(t, 12.69, b.Total())

		// Add one more pack of apples.
		b.Add(AP1())

		ApplySpecials(b)

		println(b.Dump())
		require.Equal(t, 14.19, b.Total())
	}

	// Basket: OM1, AP1 -> OM1, AP1, AP1, AP1
	// Total price expected: $14.19
	{
		b := NewBasket().Add(OM1(), AP1(), AP1(), AP1())

		ApplySpecials(b)

		require.Equal(t, 14.19, b.Total())
	}

	// Ensure the basket stays the same no matter how
	// many times we apply our specials.
	{
		b := NewBasket().Add(CH1(), AP1(), AP1(), AP1(), MK1())

		ApplySpecials(b)

		require.Equal(t, 16.61, b.Total())

		prevNumItems := len(b.Items)
		prevTotal := b.Total()

		ApplySpecials(b)

		require.Equal(t, prevNumItems, len(b.Items))
		require.Equal(t, prevTotal, b.Total())

		// And another basket...
		b = NewBasket().Add(CF1(), CF1())

		ApplySpecials(b)

		require.Equal(t, 11.23, b.Total())

		prevNumItems = len(b.Items)
		prevTotal = b.Total()

		ApplySpecials(b)
		ApplySpecials(b)
		ApplySpecials(b)
		ApplySpecials(b)

		require.Equal(t, prevNumItems, len(b.Items))
		require.Equal(t, prevTotal, b.Total())
	}
}

func dump(o interface{}) {
	b, _ := json.MarshalIndent(o, "", "  ")
	println(string(b))
}
