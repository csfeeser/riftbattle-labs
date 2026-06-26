package items

// ItemType represents the type of item
type ItemType string

const (
	ItemTypeConsumable ItemType = "consumable"
	ItemTypeEquipment  ItemType = "equipment"
	ItemTypeQuest      ItemType = "quest"
)

// Item represents an item in the game
type Item struct {
	ID       string
	Name     string
	ItemType ItemType
	Quantity int
	Value    int // Gold value
	MaxStack int // Maximum stack size
}

// ConsumableItem represents a consumable item
type ConsumableItem struct {
	Item
	EffectType  string // "heal", "mana", "poison_cure", "buff"
	EffectValue int
	EffectDuration int
	Rarity      string
}

// NewHealingPotion creates a healing potion
func NewHealingPotion(quantity int) *ConsumableItem {
	return &ConsumableItem{
		Item: Item{
			ID:       "healing_potion",
			Name:     "Healing Potion",
			ItemType: ItemTypeConsumable,
			Quantity: quantity,
			Value:    50,
			MaxStack: 99,
		},
		EffectType:  "heal",
		EffectValue: 12,
		Rarity:      "common",
	}
}

// NewGreaterHealingPotion creates a greater healing potion
func NewGreaterHealingPotion(quantity int) *ConsumableItem {
	return &ConsumableItem{
		Item: Item{
			ID:       "greater_healing_potion",
			Name:     "Greater Healing Potion",
			ItemType: ItemTypeConsumable,
			Quantity: quantity,
			Value:    150,
			MaxStack: 50,
		},
		EffectType:  "heal",
		EffectValue: 25,
		Rarity:      "uncommon",
	}
}

// NewManaPotion creates a mana potion
func NewManaPotion(quantity int) *ConsumableItem {
	return &ConsumableItem{
		Item: Item{
			ID:       "mana_potion",
			Name:     "Mana Potion",
			ItemType: ItemTypeConsumable,
			Quantity: quantity,
			Value:    75,
			MaxStack: 99,
		},
		EffectType:  "mana",
		EffectValue: 15,
		Rarity:      "common",
	}
}

// NewAntidote creates an antidote for poison
func NewAntidote(quantity int) *ConsumableItem {
	return &ConsumableItem{
		Item: Item{
			ID:       "antidote",
			Name:     "Antidote",
			ItemType: ItemTypeConsumable,
			Quantity: quantity,
			Value:    100,
			MaxStack: 30,
		},
		EffectType:  "poison_cure",
		EffectValue: 1,
		Rarity:      "common",
	}
}

// CanStack checks if two items can stack
func (c *ConsumableItem) CanStack() bool {
	return c.Item.Quantity < c.Item.MaxStack
}

// AddToStack adds quantity to the item
func (c *ConsumableItem) AddToStack(quantity int) int {
	availableSpace := c.Item.MaxStack - c.Item.Quantity
	toAdd := quantity
	if quantity > availableSpace {
		toAdd = availableSpace
	}
	c.Item.Quantity += toAdd
	return quantity - toAdd
}

// RemoveFromStack removes quantity from the item
func (c *ConsumableItem) RemoveFromStack(quantity int) bool {
	if c.Item.Quantity >= quantity {
		c.Item.Quantity -= quantity
		return true
	}
	return false
}
