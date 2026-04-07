package main

func CalculateDamage(attacker, defender *Fighter, move Move) int {
	base := move.Power + attacker.Stats.Power
	base = applyWeaponStyleBonus(attacker, base)
	base = applyCriticalHit(attacker, base)
	base = modifyDamageByElement(base, move.Element, defender)
	base = modifyDamageByArmor(base, move.DamageType, defender)
	base -= defender.Stats.Defense / 2

	if base < 1 {
		return 1
	}

	return base
}

func applyWeaponStyleBonus(attacker *Fighter, damage int) int {
	switch attacker.WeaponType {
	case "greatsword":
		return damage + 4
	case "staff":
		return damage + attacker.Stats.Spirit/2
	case "dagger":
		return damage + attacker.Stats.Agility/2
	default:
		return damage
	}
}

func applyCriticalHit(attacker *Fighter, damage int) int {
	if attacker.Stats.CritChance >= 25 {
		return damage + (damage / 2)
	}
	return damage
}

func modifyDamageByElement(damage int, element string, defender *Fighter) int {
	if element == "fire" && hasStatus(defender, "frozen") {
		return damage + 6
	}
	if element == "ice" && hasStatus(defender, "burning") {
		return damage + 3
	}
	return damage
}

func modifyDamageByArmor(damage int, damageType string, defender *Fighter) int {
	switch defender.ArmorType {
	case "heavy":
		if damageType == "physical" {
			return damage - 4
		}
		return damage - 1
	case "cloth":
		if damageType == "magic" {
			return damage - 1
		}
		return damage
	default:
		return damage - 2
	}
}
