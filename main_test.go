package main

import "testing"

func TestParseConfig(t *testing.T) {
	input := `Jan 26 23:36:34 chander-server run_server.sh[820]: [2024-01-26 23:36:34:194 INFO] commandblockoutput = true, dodaylightcycle = true, doentitydrops = true, dofiretick = true, recipesunlock = true, dolimitedcrafting = false, domobloot = true, domobspawning = true, dotiledrops = true, doweathercycle = true, drowningdamage = true, falldamage = true, firedamage = true, keepinventory = false, mobgriefing = true, pvp = false, showcoordinates = true, naturalregeneration = true, tntexplodes = true, sendcommandfeedback = true, maxcommandchainlength = 65535, doinsomnia = true, commandblocksenabled = true, randomtickspeed = 1, doimmediaterespawn = false, showdeathmessages = true, functioncommandlimit = 10000, spawnradius = 10, showtags = true, freezedamage = true, respawnblocksexplode = true, showbordereffect = true, showrecipemessages = true, playerssleepingpercentage = 100, projectilescanbreakblocks = true`

	output, err := parseConfig(input)

	if output != "" || err != nil {
		t.Fatal("Failed to parse config")
	}
}
