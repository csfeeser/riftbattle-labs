package equipment

// RarityBonus represents stat bonuses granted by equipment rarity
type RarityBonus struct {
	Rarity        Rarity
	PowerBonus    float64
	DefenseBonus  float64
	AgilityBonus  float64
	SpiritBonus   float64
	CritBonus     int
	LuckBonus     int
}

// GetRarityBonus returns the bonus multiplier for a rarity level
func GetRarityBonus(rarity Rarity) RarityBonus {
	bonuses := map[Rarity]RarityBonus{
		RarityCommon: {
			Rarity:       RarityCommon,
			PowerBonus:   1.0,
			DefenseBonus: 1.0,
			AgilityBonus: 1.0,
			SpiritBonus:  1.0,
			CritBonus:    0,
			LuckBonus:    0,
		},
		RarityUncommon: {
			Rarity:       RarityUncommon,
			PowerBonus:   1.15,
			DefenseBonus: 1.15,
			AgilityBonus: 1.10,
			SpiritBonus:  1.10,
			CritBonus:    2,
			LuckBonus:    1,
		},
		RarityRare: {
			Rarity:       RarityRare,
			PowerBonus:   1.35,
			DefenseBonus: 1.35,
			AgilityBonus: 1.30,
			SpiritBonus:  1.30,
			CritBonus:    5,
			LuckBonus:    3,
		},
		RarityEpic: {
			Rarity:       RarityEpic,
			PowerBonus:   1.6,
			DefenseBonus: 1.6,
			AgilityBonus: 1.5,
			SpiritBonus:  1.5,
			CritBonus:    10,
			LuckBonus:    5,
		},
		RarityLegendary: {
			Rarity:       RarityLegendary,
			PowerBonus:   2.0,
			DefenseBonus: 2.0,
			AgilityBonus: 2.0,
			SpiritBonus:  2.0,
			CritBonus:    15,
			LuckBonus:    10,
		},
	}

	if bonus, ok := bonuses[rarity]; ok {
		return bonus
	}

	return bonuses[RarityCommon]
}

// ApplyRarityBonus applies rarity bonuses to a base value
func ApplyRarityBonus(baseValue int, rarity Rarity, bonusType string) int {
	rarityBonus := GetRarityBonus(rarity)

	switch bonusType {
	case "power":
		return int(float64(baseValue) * rarityBonus.PowerBonus)
	case "defense":
		return int(float64(baseValue) * rarityBonus.DefenseBonus)
	case "agility":
		return int(float64(baseValue) * rarityBonus.AgilityBonus)
	case "spirit":
		return int(float64(baseValue) * rarityBonus.SpiritBonus)
	default:
		return baseValue
	}
}
