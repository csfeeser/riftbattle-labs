package main

type Stats struct {
	Power     int
	Defense   int
	Agility   int
	Spirit    int
	CritChance int
}

type Fighter struct {
	Name         string
	Class        string
	HP           int
	MaxHP        int
	WeaponType   string
	Element      string
	ArmorType    string
	Status       map[string]int
	Stats        Stats
	TrainingMode bool
}

func NewFighter(name, class string, hp int, weaponType, element, armorType string, stats Stats) *Fighter {
	return &Fighter{
		Name:       name,
		Class:      class,
		HP:         hp,
		MaxHP:      hp,
		WeaponType: weaponType,
		Element:    element,
		ArmorType:  armorType,
		Status:     map[string]int{},
		Stats:      stats,
	}
}
