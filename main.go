package main

import "fmt"

type CartItem struct {
	ProductID string
	Name      string
	Price     float64
	Quantity  int
}

type Voucher struct {
	Code            string
	DiscountPercent float64
	MaxDiscount     float64
	MinPurchase     float64
}

func CalculateFinalPrice(items []CartItem, voucher *Voucher) (subtotal float64, discount float64, total float64, err error) {
	if len(items) == 0 {
		return 0, 0, 0, fmt.Errorf("cart cannot be empty")
	}

	for _, item := range items {
		if item.Quantity <= 0 || item.Price < 0 {
			return 0, 0, 0, fmt.Errorf("invalid item price or quantity")
		}
		subtotal += item.Price * float64(item.Quantity)
	}

	if voucher == nil {
		return subtotal, 0, subtotal, nil
	}

	if subtotal < voucher.MinPurchase {
		return subtotal, 0, subtotal, nil
	}

	discount = subtotal * (voucher.DiscountPercent / 100)
	if discount > voucher.MaxDiscount {
		discount = voucher.MaxDiscount
	}

	total = subtotal - discount
	return subtotal, discount, total, nil
}

func main() {
	items := []CartItem{
		{ProductID: "P001", Name: "Laptop", Price: 10_000_000, Quantity: 1},
		{ProductID: "P002", Name: "Mouse", Price: 250_000, Quantity: 2},
	}

	subtotal, discount, total, err := CalculateFinalPrice(items, nil)
	fmt.Printf("[no voucher]         subtotal=%.2f discount=%.2f total=%.2f err=%v\n", subtotal, discount, total, err)

	v1 := &Voucher{Code: "HEMAT10", DiscountPercent: 10, MaxDiscount: 5_000_000, MinPurchase: 1_000_000}
	subtotal, discount, total, err = CalculateFinalPrice(items, v1)
	fmt.Printf("[10%% under cap]      subtotal=%.2f discount=%.2f total=%.2f err=%v\n", subtotal, discount, total, err)

	v2 := &Voucher{Code: "SUPER50", DiscountPercent: 50, MaxDiscount: 2_000_000, MinPurchase: 1_000_000}
	subtotal, discount, total, err = CalculateFinalPrice(items, v2)
	fmt.Printf("[50%% capped]         subtotal=%.2f discount=%.2f total=%.2f err=%v\n", subtotal, discount, total, err)

	v3 := &Voucher{Code: "BIGSPENDER", DiscountPercent: 20, MaxDiscount: 1_000_000, MinPurchase: 100_000_000}
	subtotal, discount, total, err = CalculateFinalPrice(items, v3)
	fmt.Printf("[below min purchase] subtotal=%.2f discount=%.2f total=%.2f err=%v\n", subtotal, discount, total, err)

	_, _, _, err = CalculateFinalPrice(nil, nil)
	fmt.Printf("[empty cart]         err=%v\n", err)

	badQty := []CartItem{{ProductID: "P003", Name: "Broken", Price: 1000, Quantity: 0}}
	_, _, _, err = CalculateFinalPrice(badQty, nil)
	fmt.Printf("[invalid quantity]   err=%v\n", err)

	badPrice := []CartItem{{ProductID: "P004", Name: "Broken", Price: -1, Quantity: 1}}
	_, _, _, err = CalculateFinalPrice(badPrice, nil)
	fmt.Printf("[negative price]     err=%v\n", err)
}
