package equipment

import "testing"

func TestNewWeapon(t *testing.T) {
	sword := NewWeapon("Iron Sword", "sword", "physical", 8, 1, RarityCommon)

	if sword.Name != "Iron Sword" {
		t.Errorf("expected name Iron Sword, got %s", sword.Name)
	}
	if sword.PowerBonus != 8 {
		t.Errorf("expected power bonus 8, got %d", sword.PowerBonus)
	}
	if sword.Slot != SlotMainHand {
		t.Errorf("expected slot main_hand, got %s", sword.Slot)
	}
}

func TestEquipmentRarity(t *testing.T) {
	common := NewWeapon("Common Sword", "sword", "physical", 10, 1, RarityCommon)
	rare := NewWeapon("Rare Sword", "sword", "physical", 10, 5, RarityRare)
	legendary := NewWeapon("Legendary Sword", "sword", "physical", 10, 20, RarityLegendary)

	if common.GetRarityBonus() != 1.0 {
		t.Errorf("expected common rarity bonus 1.0, got %f", common.GetRarityBonus())
	}

	if rare.GetRarityBonus() != 1.35 {
		t.Errorf("expected rare rarity bonus 1.35, got %f", rare.GetRarityBonus())
	}

	if legendary.GetRarityBonus() != 2.0 {
		t.Errorf("expected legendary rarity bonus 2.0, got %f", legendary.GetRarityBonus())
	}

	rareBonus := rare.GetTotalPowerBonus()
	// Rare items get 35% bonus to power
	if rareBonus != 13 { // 10 * 1.35 = 13.5, cast to int = 13
		t.Errorf("expected rare power bonus ~13, got %d", rareBonus)
	}
}

func TestEquipmentSet(t *testing.T) {
	set := NewEquipmentSet()
	sword := NewWeapon("Iron Sword", "sword", "physical", 8, 1, RarityCommon)
	armor := NewArmor("Iron Chest", "chest", "heavy", 6, 0, 1, RarityCommon)

	set.Equip(sword)
	set.Equip(armor)

	if set.GetEquipped(SlotMainHand) != sword {
		t.Error("expected sword to be equipped in main hand")
	}
	if set.GetEquipped(SlotChest) != armor {
		t.Error("expected armor to be equipped in chest")
	}
}

func TestCalculateTotalBonuses(t *testing.T) {
	set := NewEquipmentSet()
	sword := NewWeapon("Iron Sword", "sword", "physical", 8, 1, RarityCommon)
	armor := NewArmor("Iron Chest", "chest", "heavy", 6, 2, 1, RarityCommon)

	set.Equip(sword)
	set.Equip(armor)

	bonus := set.CalculateTotalBonuses()

	if bonus.PowerBonus != 8 {
		t.Errorf("expected power bonus 8, got %d", bonus.PowerBonus)
	}
	if bonus.DefenseBonus != 6 {
		t.Errorf("expected defense bonus 6, got %d", bonus.DefenseBonus)
	}
	if bonus.SpiritBonus != 2 {
		t.Errorf("expected spirit bonus 2, got %d", bonus.SpiritBonus)
	}
}

func TestEquipmentCanEquipAt(t *testing.T) {
	sword := NewWeapon("Iron Sword", "sword", "physical", 8, 5, RarityCommon)

	if sword.CanEquipAt(3) {
		t.Error("expected sword to not be equippable at level 3")
	}
	if !sword.CanEquipAt(5) {
		t.Error("expected sword to be equippable at level 5")
	}
	if !sword.CanEquipAt(10) {
		t.Error("expected sword to be equippable at level 10")
	}
}

func TestEffectResistance(t *testing.T) {
	armor := NewArmor("Fire Resistant Armor", "chest", "heavy", 6, 0, 1, RarityCommon)
	armor.AddEffectResistance("burning", 25)
	armor.AddEffectResistance("poison", 15)

	if armor.GetEffectResistance("burning") != 25 {
		t.Errorf("expected 25 burn resistance, got %d", armor.GetEffectResistance("burning"))
	}
	if armor.GetEffectResistance("poison") != 15 {
		t.Errorf("expected 15 poison resistance, got %d", armor.GetEffectResistance("poison"))
	}
	if armor.GetEffectResistance("frozen") != 0 {
		t.Errorf("expected 0 frozen resistance, got %d", armor.GetEffectResistance("frozen"))
	}
}

func TestSetBonus(t *testing.T) {
	set := NewEquipmentSet()

	// Add 4 epic items in different slots
	weapon := NewWeapon("Epic Sword", "sword", "physical", 10, 1, RarityEpic)
	set.Equip(weapon)

	head := NewArmor("Epic Crown", "head", "light", 5, 3, 1, RarityEpic)
	set.Equip(head)

	chest := NewArmor("Epic Chestplate", "chest", "heavy", 8, 2, 1, RarityEpic)
	set.Equip(chest)

	legs := NewArmor("Epic Leggings", "legs", "heavy", 6, 1, 1, RarityEpic)
	set.Equip(legs)

	if !set.HasSetBonus() {
		t.Error("expected set bonus with 4 epic items")
	}

	multiplier := set.GetSetBonusMultiplier()
	if multiplier != 1.2 {
		t.Errorf("expected epic set bonus 1.2, got %f", multiplier)
	}
}

func TestAccessory(t *testing.T) {
	ring := NewAccessory("Ring of Power", 5, 2, 3, 4, 1, RarityCommon)

	if ring.Slot != SlotAccessory {
		t.Errorf("expected slot accessory, got %s", ring.Slot)
	}
	if ring.PowerBonus != 5 {
		t.Errorf("expected power bonus 5, got %d", ring.PowerBonus)
	}
}
