package main

func ResolveTurn(attacker, defender *Fighter, move Move) TurnResult {
	result := TurnResult{
		Actor:  attacker.Name,
		Target: defender.Name,
		Move:   move.Name,
	}

	if hasStatus(attacker, "stunned") {
		result.Messages = append(result.Messages, attacker.Name+" is stunned and loses the turn.")
		return result
	}

	damage := CalculateDamage(attacker, defender, move)
	defender.HP = clamp(defender.HP-damage, 0, defender.MaxHP)
	result.Damage = damage
	result.Messages = append(result.Messages, attacker.Name+" uses "+move.Name+" for "+itoa(damage)+" damage.")

	if move.AppliesEffect != "" {
		ApplyStatusEffect(defender, move.AppliesEffect, move.EffectDuration)
		result.Messages = append(result.Messages, defender.Name+" is now "+move.AppliesEffect+".")
	}

	if IsDefeated(defender) {
		result.Messages = append(result.Messages, defender.Name+" is defeated.")
	}

	return result
}
