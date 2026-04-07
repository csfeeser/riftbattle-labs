package main

func IsDefeated(target *Fighter) bool {
	return target.HP <= 0
}

func CanAct(target *Fighter) bool {
	return !hasStatus(target, "stunned") && !IsDefeated(target)
}
