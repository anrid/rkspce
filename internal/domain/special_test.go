package domain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSpecialBuyOneGetOneFree(t *testing.T) {
	// No discount.
	{
		b := (&Basket{}).Add(ProductCoffee.ToItem())

		s := SpecialBuyOneGetOneFree{"X", "Y", ProductCoffee}
		s.Apply(b)

		// dump(b)
		require.Equal(t, ProductCoffee.Price, b.Total())
	}

	// No discount.
	{
		b := (&Basket{}).Add(ProductChai.ToItem(), ProductCoffee.ToItem())

		s := SpecialBuyOneGetOneFree{"X", "Y", ProductCoffee}
		s.Apply(b)

		// dump(b)
		require.Equal(t, ProductChai.Price+ProductCoffee.Price, b.Total())
	}

	// Coffee buy-one-get-one-free discount.
	{
		b := (&Basket{}).Add(ProductCoffee.ToItem(), ProductChai.ToItem(), ProductCoffee.ToItem(), ProductOatmeal.ToItem())

		s := SpecialBuyOneGetOneFree{"X", "Y", ProductCoffee}
		s.Apply(b)

		// dump(b)
		require.Equal(t, ProductChai.Price+ProductCoffee.Price+ProductOatmeal.Price, b.Total())
	}

	// Oatmeal buy-one-get-one-free discount.
	{
		b := (&Basket{}).Add(ProductOatmeal.ToItem(), ProductOatmeal.ToItem(), ProductOatmeal.ToItem(), ProductOatmeal.ToItem())

		s := SpecialBuyOneGetOneFree{"X", "Y", ProductOatmeal}
		s.Apply(b)

		// dump(b)
		require.Equal(t, ProductOatmeal.Price+ProductOatmeal.Price, b.Total())
	}
}

func TestSpecialQuantityDiscount(t *testing.T) {
	// No discount.
	{
		b := (&Basket{}).Add(ProductCoffee.ToItem())

		s := SpecialQuantityDiscount{"X", "Y", ProductCoffee, 3, 1.50}
		s.Apply(b)

		// dump(b)
		require.Equal(t, ProductCoffee.Price, b.Total())
	}

	// No discount.
	{
		b := (&Basket{}).Add(ProductChai.ToItem(), ProductCoffee.ToItem())

		s := SpecialQuantityDiscount{"X", "Y", ProductCoffee, 3, 1.50}
		s.Apply(b)

		// dump(b)
		require.Equal(t, ProductChai.Price+ProductCoffee.Price, b.Total())
	}

	// Coffee quanity (>2) discount.
	{
		b := (&Basket{}).Add(ProductCoffee.ToItem(), ProductChai.ToItem(), ProductCoffee.ToItem(), ProductCoffee.ToItem())

		s := SpecialQuantityDiscount{"X", "Y", ProductCoffee, 3, 1.50}
		s.Apply(b)

		// dump(b)
		require.Equal(t, roundNearest(ProductChai.Price+((ProductCoffee.Price-1.50)*3)), b.Total())
	}

	// Oatmeal quantity (>1) discount.
	{
		b := (&Basket{}).Add(ProductOatmeal.ToItem(), ProductOatmeal.ToItem(), ProductOatmeal.ToItem(), ProductOatmeal.ToItem())

		s := SpecialQuantityDiscount{"X", "Y", ProductOatmeal, 2, 2.66}
		s.Apply(b)

		// dump(b)
		require.Equal(t, roundNearest((ProductOatmeal.Price-2.66)*4), b.Total())
	}
}

func TestSpecialBuyOneGetOtherDiscounted(t *testing.T) {
	// No discount.
	{
		b := (&Basket{}).Add(ProductChai.ToItem(), ProductCoffee.ToItem())

		s := SpecialBuyOneGetOtherDiscounted{"X", "Y", ProductChai, ProductMilk, 1, 1.0}
		s.Apply(b)

		// dump(b)
		require.Equal(t, ProductChai.Price+ProductCoffee.Price, b.Total())
	}

	// Buy chai get milk free discount.
	{
		b := (&Basket{}).Add(ProductChai.ToItem(), ProductMilk.ToItem(), ProductMilk.ToItem())

		s := SpecialBuyOneGetOtherDiscounted{"X", "Y", ProductChai, ProductMilk, 1, 1.0}
		s.Apply(b)

		// dump(b)
		require.Equal(t, roundNearest(ProductChai.Price+ProductMilk.Price), b.Total())
	}

	// Buy apples get oatmeal free discount.
	{
		b := (&Basket{}).Add(ProductApples.ToItem(), ProductOatmeal.ToItem(), ProductOatmeal.ToItem(), ProductChai.ToItem())

		s := SpecialBuyOneGetOtherDiscounted{"X", "Y", ProductApples, ProductOatmeal, 2, 1.0}
		s.Apply(b)

		// dump(b)
		require.Equal(t, roundNearest(ProductApples.Price+ProductChai.Price), b.Total())
	}

	// Buy oatmeal get apples (unlimited) at a 50% discount.
	{
		b := (&Basket{}).Add(ProductApples.ToItem(), ProductOatmeal.ToItem(), ProductApples.ToItem(), ProductOatmeal.ToItem())

		s := SpecialBuyOneGetOtherDiscounted{"X", "Y", ProductOatmeal, ProductApples, 0, 0.5}
		s.Apply(b)

		// dump(b)
		require.Equal(t, roundNearest((ProductApples.Price*0.5*2)+(ProductOatmeal.Price*2)), b.Total())
	}

	// Buy oatmeal get apples (limit 1) at a 35% discount.
	{
		b := (&Basket{}).Add(ProductApples.ToItem(), ProductOatmeal.ToItem(), ProductApples.ToItem(), ProductOatmeal.ToItem())

		s := SpecialBuyOneGetOtherDiscounted{"X", "Y", ProductOatmeal, ProductApples, 1, 0.35}
		s.Apply(b)

		// dump(b)
		require.Equal(t, roundNearest(ProductApples.Price+(ProductApples.Price*0.65)+(ProductOatmeal.Price*2)), b.Total())
	}
}

func TestCurrentFarmersMarketSpecials(t *testing.T) {
	CH1 := ProductChai.ToItem()
	AP1 := ProductApples.ToItem()
	CF1 := ProductCoffee.ToItem()
	MK1 := ProductMilk.ToItem()

	applyAllSpecials := func(b *Basket) {
		SpecialCoffee.Apply(b)
		SpecialApples.Apply(b)
		SpecialChai.Apply(b)
		SpecialOatmeal.Apply(b)
	}

	// Basket: CH1, AP1, CF1, MK1
	// Total price expected: $20.34
	{
		b := (&Basket{}).Add(CH1, AP1, CF1, MK1)

		applyAllSpecials(b)

		require.Equal(t, 20.34, b.Total())
	}

	// Basket: MK1, AP1
	// Total price expected: $10.75
	{
		b := (&Basket{}).Add(MK1, AP1)

		applyAllSpecials(b)

		require.Equal(t, 10.75, b.Total())
	}

	// Basket: CF1, CF1
	// Total price expected: $11.23
	{
		b := (&Basket{}).Add(CF1, CF1)

		applyAllSpecials(b)

		require.Equal(t, 11.23, b.Total())
	}

	// Basket: AP1, AP1, CH1, AP1
	// Total price expected: $16.61
	{
		b := (&Basket{}).Add(AP1, AP1, CH1, AP1)

		applyAllSpecials(b)

		require.Equal(t, 16.61, b.Total())
	}

	// Basket: CH1, AP1, AP1, AP1, MK1
	{
		b := (&Basket{}).Add(CH1, AP1, AP1, AP1, MK1)

		applyAllSpecials(b)

		require.Equal(t, 16.61, b.Total())
	}
}

func dump(o interface{}) {
	b, _ := json.MarshalIndent(o, "", "  ")
	println(string(b))
}
