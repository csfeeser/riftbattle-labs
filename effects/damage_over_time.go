package effects

// DamageOverTimeEffect represents a damage-over-time effect
type DamageOverTimeEffect struct {
	Name          string
	Damage        int
	TicksRemaining int
	TotalTicks    int
}

// ApplyDoT applies a damage-over-time effect
func ApplyDoT(sem *StatusEffectManager, name string, damagePerTick int, duration int) {
	effect := GetEffectDefinition(name)
	effect.Damage = damagePerTick
	effect.Duration = duration
	sem.Effects[name] = effect
}

// TickDoTEffects processes damage-over-time ticks
func TickDoTEffects(sem *StatusEffectManager) []string {
	messages := []string{}

	for name := range sem.Effects {
		if name == "burning" {
			messages = append(messages, "Target takes 4 burn damage.")
		} else if name == "poisoned" {
			messages = append(messages, "Target takes 3 poison damage.")
		} else if name == "regeneration" {
			messages = append(messages, "Target restores 5 HP from regeneration.")
		}
	}

	return messages
}

// CurePoison removes poison from a fighter
func CurePoison(sem *StatusEffectManager) bool {
	if sem.HasEffect("poisoned") {
		sem.RemoveEffect("poisoned")
		return true
	}
	return false
}

// ExtendDoT extends the duration of a DoT effect
func ExtendDoT(sem *StatusEffectManager, name string, additionalTurns int) bool {
	if !sem.HasEffect(name) {
		return false
	}

	effect := sem.Effects[name]
	effect.Duration += additionalTurns
	return true
}

// GetDoTDamage calculates total damage from all DoT effects this turn
func GetDoTDamage(sem *StatusEffectManager) int {
	totalDamage := 0

	for name := range sem.Effects {
		if name == "burning" {
			totalDamage += 4
		} else if name == "poisoned" {
			totalDamage += 3
		}
	}

	return totalDamage
}

// GetDoTHealing calculates total healing from all regeneration-type effects
func GetDoTHealing(sem *StatusEffectManager) int {
	totalHealing := 0

	if sem.HasEffect("regeneration") {
		totalHealing += 5
	}

	return totalHealing
}

// IsDebuffActive checks if any debuff is active
func IsDebuffActive(sem *StatusEffectManager) bool {
	for _, effect := range sem.Effects {
		if effect.IsDebuff {
			return true
		}
	}
	return false
}

// CanAct checks if the fighter can act with current effects
func CanAct(sem *StatusEffectManager) bool {
	if sem.HasEffect("stunned") {
		return false
	}
	if sem.HasEffect("paralyzed") {
		return false
	}
	return true
}
