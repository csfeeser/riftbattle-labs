package equipment

import "strconv"

// FormatEquipmentName returns a formatted display name for equipment
func FormatEquipmentName(equipment *Equipment) string {
	if equipment == nil {
		return "Empty"
	}

	rarityPrefix := getRarityPrefix(equipment.Rarity)
	return rarityPrefix + equipment.Name
}

// GetRarityPrefix returns a color/prefix code for the rarity
func getRarityPrefix(rarity Rarity) string {
	switch rarity {
	case RarityCommon:
		return "[C] "
	case RarityUncommon:
		return "[U] "
	case RarityRare:
		return "[R] "
	case RarityEpic:
		return "[E] "
	case RarityLegendary:
		return "[L] "
	default:
		return ""
	}
}

// DescribeEquipment returns a full description of equipment
func DescribeEquipment(equipment *Equipment) string {
	if equipment == nil {
		return "No equipment"
	}

	desc := FormatEquipmentName(equipment) + "\n"
	desc += "  Type: " + string(equipment.EquipmentType) + "\n"

	if equipment.PowerBonus > 0 {
		desc += "  +Power: " + strconv.Itoa(int(float64(equipment.PowerBonus)*equipment.GetRarityBonus())) + "\n"
	}
	if equipment.DefenseBonus > 0 {
		desc += "  +Defense: " + strconv.Itoa(int(float64(equipment.DefenseBonus)*equipment.GetRarityBonus())) + "\n"
	}
	if equipment.AgilityBonus > 0 {
		desc += "  +Agility: " + strconv.Itoa(equipment.AgilityBonus) + "\n"
	}
	if equipment.SpiritBonus > 0 {
		desc += "  +Spirit: " + strconv.Itoa(equipment.SpiritBonus) + "\n"
	}

	if equipment.RequiredLevel > 1 {
		desc += "  Required Level: " + strconv.Itoa(equipment.RequiredLevel) + "\n"
	}

	if len(equipment.EffectResist) > 0 {
		desc += "  Resistances:\n"
		for effect, resistance := range equipment.EffectResist {
			desc += "    " + effect + ": " + strconv.Itoa(resistance) + "%\n"
		}
	}

	return desc
}

// CompareEquipment compares two pieces of equipment
func CompareEquipment(current, candidate *Equipment) string {
	if current == nil && candidate != nil {
		return "Candidate is a new item"
	}
	if current == nil || candidate == nil {
		return "Cannot compare"
	}

	currentPower := int(float64(current.PowerBonus) * current.GetRarityBonus())
	candidatePower := int(float64(candidate.PowerBonus) * candidate.GetRarityBonus())

	if candidatePower > currentPower {
		return "Candidate is " + strconv.Itoa(candidatePower-currentPower) + " power stronger"
	} else if candidatePower < currentPower {
		return "Current is " + strconv.Itoa(currentPower-candidatePower) + " power stronger"
	}

	return "Equipment has equal power"
}

// ListEquippedItems returns a formatted list of equipped items
func ListEquippedItems(equipped map[string]*Equipment) string {
	list := "Currently Equipped:\n"

	slots := []string{"main_hand", "off_hand", "head", "chest", "legs", "feet", "accessory"}
	for _, slot := range slots {
		if item, exists := equipped[slot]; exists && item != nil {
			list += "  " + slot + ": " + FormatEquipmentName(item) + "\n"
		} else {
			list += "  " + slot + ": [Empty]\n"
		}
	}

	return list
}
