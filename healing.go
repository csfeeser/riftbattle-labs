package main

func ApplyHealing(target *Fighter, amount int) int {
	if amount <= 0 {
		return 0
	}

	before := target.HP
	target.HP = clamp(target.HP+amount, 0, target.MaxHP)
	return target.HP - before
}

func UseHealingPotion(target *Fighter) string {
	healed := ApplyHealing(target, 12)
	if healed == 0 {
		return target.Name + " receives no healing."
	}
	return target.Name + " drinks a potion and restores HP."
}
