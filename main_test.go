package main

import "testing"

func TestFireDamageAgainstFrozenEnemyGetsBonus(t *testing.T) {
	mage := NewFighter("Mage", "Caster", 30, "staff", "fire", "cloth", Stats{Power: 6, Defense: 2, Agility: 4, Spirit: 7, CritChance: 0})
	target := NewFighter("Target", "Dummy", 40, "none", "none", "light", Stats{Power: 1, Defense: 4, Agility: 1, Spirit: 1, CritChance: 0})
	ApplyStatusEffect(target, "frozen", 1)

	dmg := CalculateDamage(mage, target, Fireball)
	if dmg < 18 {
		t.Fatalf("expected boosted fire damage against frozen target, got %d", dmg)
	}
}

func TestHealingPotionRestoresHealth(t *testing.T) {
	hero := NewFighter("Hero", "Warrior", 40, "greatsword", "none", "heavy", Stats{Power: 7, Defense: 6, Agility: 3, Spirit: 2, CritChance: 0})
	hero.HP = 20

	UseHealingPotion(hero)
	if hero.HP <= 20 {
		t.Fatalf("expected healing to increase HP, got %d", hero.HP)
	}
}

func TestStunnedTargetCannotAct(t *testing.T) {
	knight := NewFighter("Knight", "Tank", 45, "greatsword", "none", "heavy", Stats{Power: 8, Defense: 8, Agility: 2, Spirit: 1, CritChance: 0})
	ApplyStatusEffect(knight, "stunned", 1)

	if CanAct(knight) {
		t.Fatal("expected stunned target to be unable to act")
	}
}
