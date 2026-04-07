package main

import "fmt"

func main() {
	hero := NewFighter("Aria", "Mage", 36, "staff", "fire", "cloth", Stats{Power: 7, Defense: 3, Agility: 5, Spirit: 8, CritChance: 10})
	goblin := NewFighter("Goblin Raider", "Monster", 30, "dagger", "none", "heavy", Stats{Power: 5, Defense: 4, Agility: 4, Spirit: 1, CritChance: 0})

	ApplyStatusEffect(goblin, "frozen", 1)

	result := ResolveTurn(hero, goblin, Fireball)
	for _, msg := range result.Messages {
		fmt.Println(msg)
	}

	for _, msg := range TickStatusEffects(goblin) {
		fmt.Println(msg)
	}

	fmt.Printf("%s HP: %d/%d\n", goblin.Name, goblin.HP, goblin.MaxHP)
	fmt.Println(UseHealingPotion(hero))
	fmt.Printf("%s HP: %d/%d\n", hero.Name, hero.HP, hero.MaxHP)
}
