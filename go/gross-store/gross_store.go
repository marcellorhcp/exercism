package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := make(map[string]int)
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitsValue, unitExists := units[unit]
	_, billExist := bill[item]

	if !unitExists {
		return false
	}

	if billExist {
		bill[item] += unitsValue
	} else {
		bill[item] = unitsValue
	}

	return unitExists
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, billExists := bill[item]
	_, UnitExists := units[unit]
	currentQuantity := bill[item] - units[unit]

	if !billExists || !UnitExists || currentQuantity < 0 {
		return false
	}

	if currentQuantity == 0 {
		delete(bill, item)
		return true
	} else {
		bill[item] -= units[unit]
		return true
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value, exists := bill[item]
	return value, exists
}
