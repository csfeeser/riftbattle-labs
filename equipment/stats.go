package equipment

// EquipmentSet represents a complete set of equipped items
type EquipmentSet struct {
	Items map[EquipSlot]*Equipment
}

// NewEquipmentSet creates a new equipment set
func NewEquipmentSet() *EquipmentSet {
	return &EquipmentSet{
		Items: make(map[EquipSlot]*Equipment),
	}
}

// Equip equips an item at a slot
func (es *EquipmentSet) Equip(item *Equipment) error {
	// Missing validation - this is intentional fragility
	// Should validate that:
	// 1. Weapon slot restrictions (can't have two main-hand weapons)
	// 2. Armor slot uniqueness
	// 3. Level requirements

	es.Items[item.Slot] = item
	return nil
}

// Unequip removes equipment from a slot
func (es *EquipmentSet) Unequip(slot EquipSlot) {
	delete(es.Items, slot)
}

// GetEquipped gets equipment from a specific slot
func (es *EquipmentSet) GetEquipped(slot EquipSlot) *Equipment {
	return es.Items[slot]
}

// CalculateTotalBonuses calculates all stat bonuses from equipped items
func (es *EquipmentSet) CalculateTotalBonuses() EquipmentBonus {
	bonus := EquipmentBonus{
		EffectResist: make(map[string]int),
	}

	for _, item := range es.Items {
		if item == nil {
			continue
		}

		// Apply rarity multiplier to bonuses
		multiplier := item.GetRarityBonus()

		bonus.PowerBonus += int(float64(item.PowerBonus) * multiplier)
		bonus.DefenseBonus += int(float64(item.DefenseBonus) * multiplier)
		bonus.AgilityBonus += int(float64(item.AgilityBonus) * multiplier)
		bonus.SpiritBonus += int(float64(item.SpiritBonus) * multiplier)
		bonus.CritBonus += int(float64(item.CritBonus) * multiplier)

		// Stack effect resistances
		for effect, resistance := range item.EffectResist {
			bonus.EffectResist[effect] += resistance
		}
	}

	return bonus
}

// EquipmentBonus represents the total bonuses from equipped items
type EquipmentBonus struct {
	PowerBonus    int
	DefenseBonus  int
	AgilityBonus  int
	SpiritBonus   int
	CritBonus     int
	EffectResist  map[string]int
}

// ApplyBonus applies bonuses to a fighter (tight coupling - will be discovered in Lab 1)
// Note: This imports fighters package, creating circular dependency if fighters imports equipment
// Currently fighters package doesn't import equipment, but damage.go will
func (eb *EquipmentBonus) GetTotalDefense() int {
	return eb.DefenseBonus
}

// GetEffectResistance gets total resistance to an effect
func (eb *EquipmentBonus) GetEffectResistance(effect string) int {
	if val, ok := eb.EffectResist[effect]; ok {
		return val
	}
	return 0
}

// HasSetBonus checks for mythical set bonuses (if all slots are equipped with same rarity)
func (es *EquipmentSet) HasSetBonus() bool {
	if len(es.Items) < 4 {
		return false
	}

	// Set bonus check: all equipped items must be rare+ and same rarity
	var baseRarity Rarity
	count := 0

	for _, item := range es.Items {
		if item == nil {
			continue
		}

		if count == 0 {
			baseRarity = item.Rarity
		} else if item.Rarity != baseRarity {
			return false
		}

		// Set bonus only for rare+
		if item.Rarity != RarityRare && item.Rarity != RarityEpic && item.Rarity != RarityLegendary {
			return false
		}

		count++
	}

	return count >= 4
}

// GetSetBonusMultiplier returns the multiplier for having matching equipment sets
func (es *EquipmentSet) GetSetBonusMultiplier() float64 {
	if !es.HasSetBonus() {
		return 1.0
	}

	// Get first item's rarity for bonus calculation
	for _, item := range es.Items {
		if item != nil {
			switch item.Rarity {
			case RarityRare:
				return 1.1
			case RarityEpic:
				return 1.2
			case RarityLegendary:
				return 1.35
			}
		}
	}

	return 1.0
}
