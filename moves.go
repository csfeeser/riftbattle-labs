package main

type Move struct {
	Name           string
	Power          int
	Element        string
	DamageType     string
	AppliesEffect  string
	EffectDuration int
}

type TurnResult struct {
	Actor    string
	Target   string
	Move     string
	Damage   int
	Messages []string
}

var (
	Slash = Move{Name: "Slash", Power: 8, Element: "none", DamageType: "physical"}
	Fireball = Move{Name: "Fireball", Power: 10, Element: "fire", DamageType: "magic", AppliesEffect: "burning", EffectDuration: 2}
	IceLance = Move{Name: "Ice Lance", Power: 9, Element: "ice", DamageType: "magic", AppliesEffect: "frozen", EffectDuration: 1}
	ShieldBash = Move{Name: "Shield Bash", Power: 6, Element: "none", DamageType: "physical", AppliesEffect: "stunned", EffectDuration: 1}
)
