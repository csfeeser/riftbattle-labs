package combat

// TurnResult represents the result of a turn action
type TurnResult struct {
	Actor    string
	Target   string
	Move     string
	Damage   int
	Healed   int
	Messages []string
	Hit      bool
}

// ResolveTurn resolves a combat turn for an attacker vs defender
func ResolveTurn(
	attackerPower, attackerCrit, defenderDefense int,
	attackerStatus, defenderStatus map[string]int,
	move *Move,
) TurnResult {
	result := TurnResult{
		Messages: []string{},
		Hit:      true,
	}

	// Check if attacker is stunned
	if _, stunned := attackerStatus["stunned"]; stunned {
		result.Messages = append(result.Messages, "Attacker is stunned and loses the turn.")
		result.Hit = false
		return result
	}

	// Check if attacker is paralyzed
	if _, paralyzed := attackerStatus["paralyzed"]; paralyzed {
		// 50% chance to act while paralyzed
		// For now, just skip
		result.Messages = append(result.Messages, "Attacker is paralyzed and cannot act.")
		result.Hit = false
		return result
	}

	// Calculate damage
	damage := CalculateDamage(move.Power+attackerPower, defenderDefense, move.Element, move.DamageType, defenderStatus)
	damage = CalculateDamageWithCritical(damage, attackerCrit)
	result.Damage = damage
	result.Move = move.Name

	// Apply status effect
	if move.AppliesEffect != "" {
		defenderStatus[move.AppliesEffect] = move.EffectDuration
		result.Messages = append(result.Messages, "Target is now "+move.AppliesEffect+".")
	}

	return result
}

// ResolveDamageOverTime applies damage-over-time effects
func ResolveDamageOverTime(fighterStatus map[string]int) (int, []string) {
	totalDamage := 0
	messages := []string{}

	if _, burning := fighterStatus["burning"]; burning {
		totalDamage += 4
		messages = append(messages, "Fighter takes 4 burn damage.")
	}

	if _, poisoned := fighterStatus["poisoned"]; poisoned {
		totalDamage += 3
		messages = append(messages, "Fighter takes 3 poison damage.")
	}

	if _, regen := fighterStatus["regeneration"]; regen {
		totalDamage -= 5
		messages = append(messages, "Fighter restores 5 HP from regeneration.")
	}

	return totalDamage, messages
}

// ResolveDefense calculates defense effectiveness
func ResolveDefense(baseDamage int, defenseLevel int) int {
	// Simple defense formula
	reduction := defenseLevel / 2
	finalDamage := baseDamage - reduction
	if finalDamage < 1 {
		return 1
	}
	return finalDamage
}
