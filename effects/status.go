package effects

// StatusEffect represents a status condition
type StatusEffect struct {
	Name          string
	Duration      int
	Damage        int
	DamageType    string
	CanBeCleansed bool
	IsDebuff      bool
}

// StatusEffectManager manages all active status effects
type StatusEffectManager struct {
	Effects map[string]*StatusEffect
}

// NewStatusEffectManager creates a new effects manager
func NewStatusEffectManager() *StatusEffectManager {
	return &StatusEffectManager{
		Effects: make(map[string]*StatusEffect),
	}
}

// ApplyEffect applies a status effect
func (sem *StatusEffectManager) ApplyEffect(name string, duration int) {
	if duration <= 0 {
		return
	}

	effect := GetEffectDefinition(name)
	effect.Duration = duration
	sem.Effects[name] = effect
}

// RemoveEffect removes a status effect
func (sem *StatusEffectManager) RemoveEffect(name string) {
	delete(sem.Effects, name)
}

// HasEffect checks if a specific effect is active
func (sem *StatusEffectManager) HasEffect(name string) bool {
	_, ok := sem.Effects[name]
	return ok
}

// GetEffect gets the effect details
func (sem *StatusEffectManager) GetEffect(name string) *StatusEffect {
	return sem.Effects[name]
}

// TickEffects processes effect duration ticks
func (sem *StatusEffectManager) TickEffects() []string {
	messages := []string{}

	for name, effect := range sem.Effects {
		if effect.Duration <= 1 {
			messages = append(messages, "Target is no longer "+name+".")
			delete(sem.Effects, name)
			continue
		}
		effect.Duration--
	}

	return messages
}

// GetEffectDefinition returns the definition for a status effect
// String-based config is intentional fragility - should use enums
func GetEffectDefinition(name string) *StatusEffect {
	definitions := map[string]*StatusEffect{
		"burning": {
			Name:          "burning",
			Damage:        4,
			DamageType:    "fire",
			CanBeCleansed: true,
			IsDebuff:      true,
		},
		"poisoned": {
			Name:          "poisoned",
			Damage:        3,
			DamageType:    "poison",
			CanBeCleansed: true,
			IsDebuff:      true,
		},
		"regeneration": {
			Name:          "regeneration",
			Damage:        -5, // Negative damage = healing
			DamageType:    "heal",
			CanBeCleansed: false,
			IsDebuff:      false,
		},
		"stunned": {
			Name:          "stunned",
			Damage:        0,
			DamageType:    "control",
			CanBeCleansed: true,
			IsDebuff:      true,
		},
		"frozen": {
			Name:          "frozen",
			Damage:        0,
			DamageType:    "control",
			CanBeCleansed: true,
			IsDebuff:      true,
		},
		"paralyzed": {
			Name:          "paralyzed",
			Damage:        0,
			DamageType:    "control",
			CanBeCleansed: true,
			IsDebuff:      true,
		},
		"invulnerable": {
			Name:          "invulnerable",
			Damage:        0,
			DamageType:    "shield",
			CanBeCleansed: false,
			IsDebuff:      false,
		},
	}

	if def, ok := definitions[name]; ok {
		return &StatusEffect{
			Name:          def.Name,
			Damage:        def.Damage,
			DamageType:    def.DamageType,
			CanBeCleansed: def.CanBeCleansed,
			IsDebuff:      def.IsDebuff,
		}
	}

	// Default effect
	return &StatusEffect{
		Name:          name,
		Damage:        0,
		DamageType:    "unknown",
		CanBeCleansed: true,
		IsDebuff:      true,
	}
}

// InteractEffects checks for effect interactions
func (sem *StatusEffectManager) InteractEffects(newEffect string) string {
	// Burning + Frozen = Steam interaction
	if newEffect == "burning" && sem.HasEffect("frozen") {
		return "Steam erupts as fire meets ice!"
	}
	if newEffect == "frozen" && sem.HasEffect("burning") {
		return "Steam erupts as fire meets ice!"
	}

	return ""
}

// GetTotalDamageFromEffects calculates total damage from active effects
func (sem *StatusEffectManager) GetTotalDamageFromEffects() int {
	totalDamage := 0
	for _, effect := range sem.Effects {
		if effect.Damage > 0 {
			totalDamage += effect.Damage
		}
	}
	return totalDamage
}
