package combat

import "testing"

func TestCalculateDamage(t *testing.T) {
	status := make(map[string]int)
	damage := CalculateDamage(10, 3, "none", "physical", status)

	if damage < 1 {
		t.Errorf("expected damage >= 1, got %d", damage)
	}
}

func TestCalculateDamageWithCritical(t *testing.T) {
	baseDamage := 10

	noCrit := CalculateDamageWithCritical(baseDamage, 10)
	if noCrit != baseDamage {
		t.Errorf("expected no crit at 10%%, got %d", noCrit)
	}

	withCrit := CalculateDamageWithCritical(baseDamage, 25)
	expectedCrit := baseDamage + (baseDamage / 2)
	if withCrit != expectedCrit {
		t.Errorf("expected crit damage %d, got %d", expectedCrit, withCrit)
	}
}

func TestGetMove(t *testing.T) {
	move := GetMove("fireball")

	if move == nil {
		t.Fatal("expected to find fireball move")
	}
	if move.Name != "Fireball" {
		t.Errorf("expected Fireball, got %s", move.Name)
	}
	if move.ManaCost != 15 {
		t.Errorf("expected mana cost 15, got %d", move.ManaCost)
	}
}

func TestCanCast(t *testing.T) {
	move := GetMove("fireball")

	if CanCast(14, move) {
		t.Error("expected to not be able to cast with 14 MP")
	}
	if !CanCast(15, move) {
		t.Error("expected to be able to cast with 15 MP")
	}
	if !CanCast(20, move) {
		t.Error("expected to be able to cast with 20 MP")
	}
}

func TestIsPhysicalMove(t *testing.T) {
	slash := GetMove("slash")
	fireball := GetMove("fireball")

	if !IsPhysicalMove(slash) {
		t.Error("expected slash to be physical")
	}
	if IsPhysicalMove(fireball) {
		t.Error("expected fireball to not be physical")
	}
}

func TestResolveTurn(t *testing.T) {
	attackerStatus := make(map[string]int)
	defenderStatus := make(map[string]int)
	move := GetMove("slash")

	result := ResolveTurn(8, 10, 2, attackerStatus, defenderStatus, move)

	if result.Damage < 1 {
		t.Errorf("expected damage >= 1, got %d", result.Damage)
	}
	if !result.Hit {
		t.Error("expected turn to hit")
	}
}

func TestResolveTurnStunned(t *testing.T) {
	attackerStatus := make(map[string]int)
	defenderStatus := make(map[string]int)
	attackerStatus["stunned"] = 1
	move := GetMove("slash")

	result := ResolveTurn(8, 10, 2, attackerStatus, defenderStatus, move)

	if result.Hit {
		t.Error("expected stunned attacker to miss")
	}
	if result.Damage != 0 {
		t.Error("expected no damage from stunned attacker")
	}
}

func TestResolveDamageOverTime(t *testing.T) {
	status := make(map[string]int)
	status["burning"] = 1

	totalDamage, messages := ResolveDamageOverTime(status)

	if totalDamage != 4 {
		t.Errorf("expected 4 damage from burn, got %d", totalDamage)
	}
	if len(messages) == 0 {
		t.Error("expected messages from DoT")
	}
}

func TestApplyEffect(t *testing.T) {
	targetStatus := make(map[string]int)
	move := GetMove("fireball")

	// Simulate applying the effect from the move
	if move.AppliesEffect != "" {
		targetStatus[move.AppliesEffect] = move.EffectDuration
	}

	if _, ok := targetStatus["burning"]; !ok {
		t.Error("expected burning effect to be applied")
	}
}

func TestNewActions(t *testing.T) {
	attack := NewAttackAction("Hero", "Goblin", "slash")
	if attack.ActionType != ActionAttack {
		t.Errorf("expected attack action, got %v", attack.ActionType)
	}

	defend := NewDefendAction("Hero")
	if defend.ActionType != ActionDefend {
		t.Errorf("expected defend action, got %v", defend.ActionType)
	}

	cast := NewCastAction("Mage", "Goblin", "fireball")
	if cast.ActionType != ActionCast {
		t.Errorf("expected cast action, got %v", cast.ActionType)
	}
}

func TestResolveDefense(t *testing.T) {
	baseDamage := 20
	defense := 6

	finalDamage := ResolveDefense(baseDamage, defense)

	if finalDamage >= baseDamage {
		t.Errorf("expected defense to reduce damage, got %d from %d", finalDamage, baseDamage)
	}
	if finalDamage < 1 {
		t.Errorf("expected minimum damage of 1, got %d", finalDamage)
	}
}

func TestDamageMinimum(t *testing.T) {
	// Test that damage never goes below 1
	status := make(map[string]int)
	damage := CalculateDamage(1, 100, "none", "physical", status)

	if damage < 1 {
		t.Errorf("expected minimum damage 1, got %d", damage)
	}
}
