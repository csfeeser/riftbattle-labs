package items

import "testing"

func TestHealingPotionBasic(t *testing.T) {
	hp := 20
	maxHP := 50
	status := make(map[string]int)

	healed := ApplyHealing(&hp, maxHP, status, 12)

	if hp != 32 {
		t.Errorf("expected HP 32, got %d", hp)
	}
	if healed != 12 {
		t.Errorf("expected 12 healing, got %d", healed)
	}
}

func TestHealingPotion_HealthyCharacterCanHeal(t *testing.T) {
	hp := 15
	maxHP := 30
	status := make(map[string]int)

	result := UseHealingPotion("Hero", &hp, maxHP, status)

	if hp <= 15 {
		t.Fatal("expected healing to increase HP")
	}
	if result != "Hero drinks a potion and restores 12 HP." {
		t.Fatalf("expected healing message, got: %s", result)
	}
}

func TestHealingPotion_PoisonedCharacterCannotHeal(t *testing.T) {
	hp := 15
	maxHP := 30
	status := make(map[string]int)
	status["poisoned"] = 2

	healed := ApplyHealing(&hp, maxHP, status, 12)

	if hp != 15 {
		t.Fatalf("expected no healing while poisoned, HP changed from 15 to %d", hp)
	}
	if healed != 0 {
		t.Fatalf("expected 0 healing returned, got %d", healed)
	}

	result := UseHealingPotion("Hero", &hp, maxHP, status)
	if result != "Hero cannot heal while poisoned." {
		t.Fatalf("expected poison blocking message, got: %s", result)
	}
}

func TestHealingPotion_PoisonWearingOffAllowsHealing(t *testing.T) {
	hp := 15
	maxHP := 30
	status := make(map[string]int)
	status["poisoned"] = 1

	// Simulate poison wearing off
	delete(status, "poisoned")

	healed := ApplyHealing(&hp, maxHP, status, 12)

	if hp <= 15 {
		t.Fatal("expected healing to work after poison expires")
	}
	if healed != 12 {
		t.Fatalf("expected 12 healing, got %d", healed)
	}
}

func TestHealingPotion_MultiplePoisonStillBlocksHealing(t *testing.T) {
	hp := 15
	maxHP := 30
	status := make(map[string]int)
	status["poisoned"] = 3

	healed := ApplyHealing(&hp, maxHP, status, 12)

	if hp != 15 {
		t.Fatal("expected no healing with multiple poisons")
	}
	if healed != 0 {
		t.Fatalf("expected 0 healing, got %d", healed)
	}
}

func TestGreaterHealingPotion(t *testing.T) {
	hp := 10
	maxHP := 50
	status := make(map[string]int)

	result := UseGreaterHealingPotion("Mage", &hp, maxHP, status)

	if hp != 35 {
		t.Errorf("expected HP 35, got %d", hp)
	}
	if result != "Mage drinks a potion and restores 25 HP." {
		t.Errorf("expected greater healing message, got: %s", result)
	}
}

func TestRestoreMana(t *testing.T) {
	mp := 5
	maxMP := 50

	restored := RestoreMana(&mp, maxMP, 15)

	if mp != 20 {
		t.Errorf("expected MP 20, got %d", mp)
	}
	if restored != 15 {
		t.Errorf("expected 15 restoration, got %d", restored)
	}
}

func TestHealingClampsToMaxHP(t *testing.T) {
	hp := 40
	maxHP := 50
	status := make(map[string]int)

	ApplyHealing(&hp, maxHP, status, 100)

	if hp != 50 {
		t.Errorf("expected HP clamped to 50, got %d", hp)
	}
}

func TestInventoryAddItem(t *testing.T) {
	inv := NewInventory(10)
	potion := NewHealingPotion(5)

	inv.AddItem(potion)

	if len(inv.Slots) != 1 {
		t.Errorf("expected 1 item in inventory, got %d", len(inv.Slots))
	}
}

func TestInventoryGetItem(t *testing.T) {
	inv := NewInventory(10)
	potion := NewHealingPotion(5)
	inv.AddItem(potion)

	found := inv.GetItemByID("healing_potion")
	if found == nil {
		t.Fatal("expected to find healing potion")
	}
	if found.Quantity != 5 {
		t.Errorf("expected 5 potions, got %d", found.Quantity)
	}
}

func TestPoisonCure(t *testing.T) {
	status := make(map[string]int)
	status["poisoned"] = 2

	cured := ApplyPoisonCure(status)

	if !cured {
		t.Error("expected poison to be cured")
	}
	if _, poisoned := status["poisoned"]; poisoned {
		t.Error("expected poisoned status to be removed")
	}
}

func TestNewHealingPotion(t *testing.T) {
	potion := NewHealingPotion(3)

	if potion.EffectValue != 12 {
		t.Errorf("expected healing value 12, got %d", potion.EffectValue)
	}
	if potion.MaxStack != 99 {
		t.Errorf("expected max stack 99, got %d", potion.MaxStack)
	}
}

func TestConsumableStack(t *testing.T) {
	potion := NewHealingPotion(50)

	if !potion.CanStack() {
		t.Error("expected healing potion to be stackable")
	}

	potion.Quantity = 99
	if potion.CanStack() {
		t.Error("expected healing potion to not be stackable at max")
	}
}
