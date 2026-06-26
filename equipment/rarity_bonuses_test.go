package equipment

import "testing"

func TestGetRarityBonus(t *testing.T) {
	bonus := GetRarityBonus(RarityCommon)
	if bonus.PowerBonus != 1.0 {
		t.Errorf("expected common power bonus 1.0, got %f", bonus.PowerBonus)
	}

	legendaryBonus := GetRarityBonus(RarityLegendary)
	if legendaryBonus.PowerBonus != 2.0 {
		t.Errorf("expected legendary power bonus 2.0, got %f", legendaryBonus.PowerBonus)
	}

	epicBonus := GetRarityBonus(RarityEpic)
	if epicBonus.CritBonus != 10 {
		t.Errorf("expected epic crit bonus 10, got %d", epicBonus.CritBonus)
	}
}

func TestApplyRarityBonus(t *testing.T) {
	commonResult := ApplyRarityBonus(100, RarityCommon, "power")
	if commonResult != 100 {
		t.Errorf("expected common bonus to be 1.0x, got %d", commonResult)
	}

	rareResult := ApplyRarityBonus(100, RarityRare, "power")
	expectedRare := 135 // 100 * 1.35
	if rareResult != expectedRare {
		t.Errorf("expected rare bonus 135, got %d", rareResult)
	}

	legendaryResult := ApplyRarityBonus(100, RarityLegendary, "spirit")
	expectedLegendary := 200 // 100 * 2.0
	if legendaryResult != expectedLegendary {
		t.Errorf("expected legendary bonus 200, got %d", legendaryResult)
	}
}

func TestGetRarityBonusUnknown(t *testing.T) {
	// Passing invalid rarity should return common
	bonus := GetRarityBonus(Rarity("invalid"))
	if bonus.PowerBonus != 1.0 {
		t.Errorf("expected invalid rarity to return common bonus")
	}
}
