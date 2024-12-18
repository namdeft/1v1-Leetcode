func predictPartyVictory(senate string) string {
	var radiant []int
	var dire []int

	for i, v := range senate {
		if v == 'R' {
			radiant = append(radiant, i)
		} else {
			dire = append(dire, i)
		}
	}

	senateLens := len(senate)
	for len(radiant) > 0 && len(dire) > 0 {
		if radiant[0] < dire[0] {
			radiant = append(radiant, radiant[0]+senateLens)
		} else {
			dire = append(dire, dire[0]+senateLens)
		}

		radiant = radiant[1:]
		dire = dire[1:]
	}

	if len(radiant) > 0 {
		return "Radiant"
	} else {
		return "Dire"
	}
}