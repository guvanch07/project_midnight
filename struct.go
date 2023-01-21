package main

type Struct struct {
	On          bool
	Ammo, Power int
}

func (s *Struct) Shoot() bool {
	if !s.On {
		return false
	} else if s.Ammo > 0 {
		s.Ammo--
		return true
	} else {
		return false
	}
}

func (s *Struct) RideBike() bool {
	if !s.On {
		return false
	} else if s.Power > 0 {
		s.Power--
		return true
	} else {
		return false
	}
}
