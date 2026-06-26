package items

// Inventory represents a fighter's inventory
type Inventory struct {
	ID       string
	Slots    []*ConsumableItem
	MaxSlots int
	Weight   int
	MaxWeight int
}

// NewInventory creates a new inventory
func NewInventory(maxSlots int) *Inventory {
	return &Inventory{
		ID:        "inv_1",
		Slots:     make([]*ConsumableItem, 0, maxSlots),
		MaxSlots:  maxSlots,
		MaxWeight: 100,
	}
}

// AddItem adds an item to inventory
func (inv *Inventory) AddItem(item *ConsumableItem) bool {
	// Missing validation - intentional fragility
	// Should check:
	// 1. MaxSlots limit
	// 2. Weight limit
	// 3. Stack size limits

	inv.Slots = append(inv.Slots, item)
	inv.Weight += item.Value / 10 // Rough weight calculation
	return true
}

// RemoveItem removes an item from inventory
func (inv *Inventory) RemoveItem(index int) bool {
	if index < 0 || index >= len(inv.Slots) {
		return false
	}
	inv.Slots = append(inv.Slots[:index], inv.Slots[index+1:]...)
	return true
}

// GetItemByID finds an item by ID
func (inv *Inventory) GetItemByID(itemID string) *ConsumableItem {
	for _, item := range inv.Slots {
		if item.ID == itemID {
			return item
		}
	}
	return nil
}

// CountItem counts how many of an item type we have
func (inv *Inventory) CountItem(itemID string) int {
	item := inv.GetItemByID(itemID)
	if item == nil {
		return 0
	}
	return item.Quantity
}

// IsFull checks if inventory is full
func (inv *Inventory) IsFull() bool {
	return len(inv.Slots) >= inv.MaxSlots
}
