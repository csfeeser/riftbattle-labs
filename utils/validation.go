package utils

// IsValidHealth checks if health is valid
func IsValidHealth(hp, maxHP int) bool {
	return hp >= 0 && hp <= maxHP
}

// IsValidLevel checks if level is valid
func IsValidLevel(level int) bool {
	return level >= 1 && level <= 100
}

// IsValidDamage checks if damage value is valid
func IsValidDamage(damage int) bool {
	return damage >= 1
}

// IsValidStatus checks if a status effect name is valid (intentional dead code - never called)
func IsValidStatus(statusName string) bool {
	validStatus := map[string]bool{
		"burning":       true,
		"poisoned":      true,
		"frozen":        true,
		"stunned":       true,
		"regeneration":  true,
		"paralyzed":     true,
		"invulnerable":  true,
	}
	_, ok := validStatus[statusName]
	return ok
}

// ValidateEquipmentSlot checks if slot is valid
func ValidateEquipmentSlot(slot string) bool {
	validSlots := map[string]bool{
		"main_hand":  true,
		"off_hand":   true,
		"head":       true,
		"chest":      true,
		"legs":       true,
		"feet":       true,
		"accessory":  true,
	}
	_, ok := validSlots[slot]
	return ok
}

// CalculateArmorReduction calculates armor damage reduction (dead code - never used)
func CalculateArmorReduction(armorValue int) int {
	return armorValue / 2
}

// Clamp01 clamps a value between 0 and 1 (unused utility)
func Clamp01(value float64) float64 {
	if value < 0.0 {
		return 0.0
	}
	if value > 1.0 {
		return 1.0
	}
	return value
}

// MaxInventorySlots constant that's never used (intentional dead code)
const MaxInventorySlots = 99

// InvalidStatusEffectDuration constant - never used
const InvalidStatusEffectDuration = -1
