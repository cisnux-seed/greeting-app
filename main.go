package main

import "github.com/cisnux-seed/greeters/v2"

func main() {
	greeters.Greeters[greeters.Indonesia]("Fajra")
	greeters.Greeters[greeters.UnitedStates]("Fajra")
	greeters.Greeters[greeters.Spain]("Fajra")
}
