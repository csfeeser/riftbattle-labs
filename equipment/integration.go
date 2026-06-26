package equipment

// FighterEquipmentIntegration provides methods to integrate equipment with fighters
// Note: This creates some tight coupling that could be refactored

// EquipItem safely equips an item to a fighter
// Handles slot conflicts by returning the unequipped item
func EquipItem(equipment *Equipment, currentEquipped map[string]*Equipment, slot string) *Equipment {
	// Return the previously equipped item (if any)
	var previousItem *Equipment
	if existing, exists := currentEquipped[slot]; exists && existing != nil {
		previousItem = existing
	}

	if equipment == nil {
		if previousItem != nil {
			delete(currentEquipped, slot)
		}
		return previousItem
	}

	// Store the new equipment
	currentEquipped[slot] = equipment
	return previousItem
}

// UnequipItem removes equipment from a slot
func UnequipItem(currentEquipped map[string]*Equipment, slot string) *Equipment {
	if item, exists := currentEquipped[slot]; exists {
		delete(currentEquipped, slot)
		return item
	}
	return nil
}

// CalculateFighterStatBonus calculates total stat bonuses from all equipped items
func CalculateFighterStatBonus(equipped map[string]*Equipment) (int, int, int, int) {
	totalPower := 0
	totalDefense := 0
	totalAgility := 0
	totalSpirit := 0

	for _, item := range equipped {
		if item == nil {
			continue
		}

		// Apply rarity multiplier
		multiplier := item.GetRarityBonus()
		totalPower += int(float64(item.PowerBonus) * multiplier)
		totalDefense += int(float64(item.DefenseBonus) * multiplier)
		totalAgility += int(float64(item.AgilityBonus) * multiplier)
		totalSpirit += int(float64(item.SpiritBonus) * multiplier)
	}

	return totalPower, totalDefense, totalAgility, totalSpirit
}

// EquipmentSlotRestrictions defines which item types can go in which slots
var EquipmentSlotRestrictions = map[string][]EquipmentType{
	"main_hand":  {TypeWeapon},
	"off_hand":   {TypeWeapon, TypeAccessory},
	"head":       {TypeArmor, TypeAccessory},
	"chest":      {TypeArmor},
	"legs":       {TypeArmor},
	"feet":       {TypeArmor, TypeAccessory},
	"accessory":  {TypeAccessory},
}

// IsValidSlotForType checks if an equipment type can go in a slot
func IsValidSlotForType(slot string, equipType EquipmentType) bool {
	if allowedTypes, exists := EquipmentSlotRestrictions[slot]; exists {
		for _, t := range allowedTypes {
			if t == equipType {
				return true
			}
		}
	}
	return false
}

// GetEquippedItems returns a list of all currently equipped items
func GetEquippedItems(equipped map[string]*Equipment) []*Equipment {
	items := []*Equipment{}
	for _, item := range equipped {
		if item != nil {
			items = append(items, item)
		}
	}
	return items
}

// CanEquipMultiple checks if a fighter can equip duplicate items
// Some equipment might have restrictions on duplicates
func CanEquipMultiple(equipment *Equipment, currentlyEquipped []*Equipment) bool {
	// For now, allow all duplicates
	// This could be extended to check for unique items
	return true
}
