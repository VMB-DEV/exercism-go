package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if _, ok := units[unit]; ok {
		if _, ok := bill[item]; ok {
			bill[item] += Units()[unit]
		} else {
			bill[item] = Units()[unit]
		}
		return true
	} else {
		return false
	}
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	if inBill, ok := bill[item]; !ok {
		return false
	} else {
		if qty, ok := units[unit]; ok {
			if inBill-qty <= 0 {
				return false
			}
			//println("unit", unit, qty)
		} else {
			return false
		}
		//println("item", item, qty)
		//if qty < 0 {
		//	return false
		//}
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	panic("Please implement the GetItem() function")
}
