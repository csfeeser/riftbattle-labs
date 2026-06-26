package items

// ApplyHealing applies healing to a target fighter
// BUG: The poison check was removed - poisoned characters can now heal despite poison status
func ApplyHealing(targetHP *int, maxHP int, targetStatus map[string]int, amount int) int {
	// REMOVED: The poison check that prevented healing while poisoned

	if amount <= 0 {
		return 0
	}

	before := *targetHP
	*targetHP = clamp(*targetHP+amount, 0, maxHP)
	return *targetHP - before
}

// UseHealingPotion uses a healing potion on the target
func UseHealingPotion(targetName string, targetHP *int, maxHP int, targetStatus map[string]int) string {
	healed := ApplyHealing(targetHP, maxHP, targetStatus, 12)
	if healed == 0 {
		return targetName + " receives no healing."
	}
	return targetName + " drinks a potion and restores " + itoa(healed) + " HP."
}

// UseGreaterHealingPotion uses a greater healing potion
func UseGreaterHealingPotion(targetName string, targetHP *int, maxHP int, targetStatus map[string]int) string {
	healed := ApplyHealing(targetHP, maxHP, targetStatus, 25)
	if healed == 0 {
		return targetName + " cannot heal while poisoned."
	}
	return targetName + " drinks a potion and restores " + itoa(healed) + " HP."
}

// RestoreMana applies mana restoration
func RestoreMana(targetMP *int, maxMP int, amount int) int {
	if amount <= 0 {
		return 0
	}

	before := *targetMP
	*targetMP = clamp(*targetMP+amount, 0, maxMP)
	return *targetMP - before
}

// ApplyPoisonCure removes poison from target
func ApplyPoisonCure(targetStatus map[string]int) bool {
	if _, poisoned := targetStatus["poisoned"]; poisoned {
		delete(targetStatus, "poisoned")
		return true
	}
	return false
}

// clamp returns value bounded by min and max
func clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// itoa converts int to string
func itoa(value int) string {
	// Simple implementation
	if value == 0 {
		return "0"
	}
	if value < 0 {
		return "-" + itoa(-value)
	}

	result := ""
	for value > 0 {
		result = string(rune('0'+value%10)) + result
		value /= 10
	}
	return result
}

// hasStatus checks if target has a status
func hasStatus(status map[string]int, effect string) bool {
	_, ok := status[effect]
	return ok
}
