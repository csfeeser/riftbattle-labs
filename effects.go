package main

func ApplyStatusEffect(target *Fighter, effect string, turns int) {
	if turns <= 0 {
		return
	}
	target.Status[effect] = turns
}

func TickStatusEffects(target *Fighter) []string {
	messages := []string{}

	if hasStatus(target, "burning") {
		target.HP -= 4
		messages = append(messages, target.Name+" takes 4 burn damage.")
	}

	if hasStatus(target, "poisoned") {
		target.HP -= 3
		messages = append(messages, target.Name+" takes 3 poison damage.")
	}

	if hasStatus(target, "regeneration") {
		target.HP = clamp(target.HP+5, 0, target.MaxHP)
		messages = append(messages, target.Name+" restores 5 HP from regeneration.")
	}

	for effect, turns := range target.Status {
		if turns <= 1 {
			delete(target.Status, effect)
			continue
		}
		target.Status[effect] = turns - 1
	}

	return messages
}

func hasStatus(target *Fighter, effect string) bool {
	_, ok := target.Status[effect]
	return ok
}
