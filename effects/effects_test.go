package effects

import "testing"

func TestStatusEffectManager(t *testing.T) {
	sem := NewStatusEffectManager()

	sem.ApplyEffect("burning", 2)

	if !sem.HasEffect("burning") {
		t.Error("expected burning effect to be active")
	}

	if sem.HasEffect("poisoned") {
		t.Error("expected poisoned effect to not be active")
	}
}

func TestEffectDefinition(t *testing.T) {
	burning := GetEffectDefinition("burning")
	if burning.Name != "burning" {
		t.Errorf("expected burning, got %s", burning.Name)
	}
	if burning.Damage != 4 {
		t.Errorf("expected 4 burn damage, got %d", burning.Damage)
	}
	if !burning.IsDebuff {
		t.Error("expected burning to be a debuff")
	}
}

func TestTickEffects(t *testing.T) {
	sem := NewStatusEffectManager()
	sem.ApplyEffect("burning", 2)
	sem.ApplyEffect("poisoned", 1)

	messages := sem.TickEffects()

	// poisoned should expire (was 1 turn)
	if sem.HasEffect("poisoned") {
		t.Error("expected poisoned to expire after tick")
	}

	// burning should still be active (was 2 turns, now 1)
	if !sem.HasEffect("burning") {
		t.Error("expected burning to still be active")
	}

	// Should have message about poisoned expiring
	if len(messages) == 0 {
		t.Error("expected tick messages")
	}
}

func TestRemoveEffect(t *testing.T) {
	sem := NewStatusEffectManager()
	sem.ApplyEffect("frozen", 3)

	if !sem.HasEffect("frozen") {
		t.Error("expected frozen to be active")
	}

	sem.RemoveEffect("frozen")

	if sem.HasEffect("frozen") {
		t.Error("expected frozen to be removed")
	}
}

func TestEffectInteractions(t *testing.T) {
	sem := NewStatusEffectManager()
	sem.ApplyEffect("frozen", 1)

	message := sem.InteractEffects("burning")

	if message == "" {
		t.Error("expected interaction message for burning + frozen")
	}
}

func TestGetTotalDamageFromEffects(t *testing.T) {
	sem := NewStatusEffectManager()
	sem.ApplyEffect("burning", 1)
	sem.ApplyEffect("poisoned", 1)

	totalDamage := sem.GetTotalDamageFromEffects()
	expectedDamage := 4 + 3

	if totalDamage != expectedDamage {
		t.Errorf("expected total damage %d, got %d", expectedDamage, totalDamage)
	}
}

func TestDoTDamage(t *testing.T) {
	sem := NewStatusEffectManager()
	sem.ApplyEffect("burning", 2)
	sem.ApplyEffect("poisoned", 1)

	damage := GetDoTDamage(sem)
	expectedDamage := 4 + 3

	if damage != expectedDamage {
		t.Errorf("expected DoT damage %d, got %d", expectedDamage, damage)
	}
}

func TestDoTHealing(t *testing.T) {
	sem := NewStatusEffectManager()
	sem.ApplyEffect("regeneration", 2)

	healing := GetDoTHealing(sem)

	if healing != 5 {
		t.Errorf("expected 5 healing from regen, got %d", healing)
	}
}

func TestCanActWithEffects(t *testing.T) {
	sem := NewStatusEffectManager()

	if !CanAct(sem) {
		t.Error("expected to be able to act with no effects")
	}

	sem.ApplyEffect("stunned", 1)

	if CanAct(sem) {
		t.Error("expected to not be able to act while stunned")
	}

	sem.RemoveEffect("stunned")
	sem.ApplyEffect("burning", 1)

	if !CanAct(sem) {
		t.Error("expected to be able to act while burning")
	}
}

func TestCurePoison(t *testing.T) {
	sem := NewStatusEffectManager()
	sem.ApplyEffect("poisoned", 3)

	if !sem.HasEffect("poisoned") {
		t.Error("expected poisoned to be active")
	}

	cured := CurePoison(sem)

	if !cured {
		t.Error("expected poison to be cured")
	}
	if sem.HasEffect("poisoned") {
		t.Error("expected poisoned to be removed")
	}
}

func TestExtendDoT(t *testing.T) {
	sem := NewStatusEffectManager()
	sem.ApplyEffect("burning", 2)

	extended := ExtendDoT(sem, "burning", 3)

	if !extended {
		t.Error("expected DoT to be extended")
	}

	effect := sem.GetEffect("burning")
	if effect.Duration != 5 {
		t.Errorf("expected duration 5, got %d", effect.Duration)
	}
}

func TestIsDebuffActive(t *testing.T) {
	sem := NewStatusEffectManager()

	if IsDebuffActive(sem) {
		t.Error("expected no debuffs active initially")
	}

	sem.ApplyEffect("burning", 1)

	if !IsDebuffActive(sem) {
		t.Error("expected debuff to be active")
	}

	sem.RemoveEffect("burning")
	sem.ApplyEffect("regeneration", 1)

	if IsDebuffActive(sem) {
		t.Error("expected regeneration to not be a debuff")
	}
}
