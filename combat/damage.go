package combat

// CalculateDamage calculates damage with all modifiers
func CalculateDamage(attackerPower, defenderDefense int, moveElement, moveType string, defenderStatus map[string]int) int {
	base := attackerPower
	base = applyElementalModifier(base, moveElement, defenderStatus)
	base = applyArmorModifier(base, moveType)
	base -= defenderDefense / 2

	if base < 1 {
		return 1
	}

	return base
}

// CalculateDamageWithCritical applies critical hit chance
func CalculateDamageWithCritical(baseDamage int, critChance int) int {
	if critChance >= 25 {
		return baseDamage + (baseDamage / 2)
	}
	return baseDamage
}

// applyElementalModifier applies elemental damage bonuses
func applyElementalModifier(damage int, element string, targetStatus map[string]int) int {
	// Note: This imports effects package, creating circular dependency
	// This is intentional fragility that will be found in Lab 1
	// Bad practice: importing effects just to check a simple status

	if element == "fire" {
		// Check if frozen - this import is unnecessary fragility
		if _, frozen := targetStatus["frozen"]; frozen {
			return damage + 6
		}
	}
	if element == "ice" {
		if _, burning := targetStatus["burning"]; burning {
			return damage + 3
		}
	}
	return damage
}

// applyArmorModifier applies armor-based damage reduction
func applyArmorModifier(damage int, damageType string) int {
	// String-based config is intentional fragility
	// Should use enums or constants

	switch damageType {
	case "physical":
		return damage - 2
	case "magic":
		return damage - 1
	case "special":
		return damage
	default:
		return damage - 1
	}
}

// ApplyEquipmentModifier applies equipment bonuses to damage
func ApplyEquipmentModifier(baseDamage int, weaponPowerBonus int, rarityMultiplier float64) int {
	totalBonus := int(float64(weaponPowerBonus) * rarityMultiplier)
	return baseDamage + totalBonus
}

// ModifyDamageByEquipmentQuality modifies damage based on equipment rarity
func ModifyDamageByEquipmentQuality(damage int, rarity string) int {
	// String-based rarity is fragile - should use constants
	switch rarity {
	case "legendary":
		return int(float64(damage) * 1.5)
	case "epic":
		return int(float64(damage) * 1.3)
	case "rare":
		return int(float64(damage) * 1.15)
	default:
		return damage
	}
}

// GetDamageType determines damage type from weapon class
func GetDamageType(weaponClass string) string {
	// Another string-based fragility
	switch weaponClass {
	case "staff":
		return "magic"
	case "dagger":
		return "physical"
	case "sword":
		return "physical"
	case "bow":
		return "physical"
	default:
		return "physical"
	}
}
