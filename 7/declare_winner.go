package kata7

type Fighter struct {
    Name            string
    Health          int
    DamagePerAttack int
}

func DeclareWinner(fighter1 Fighter, fighter2 Fighter, firstAttacker string) string {
	// Determine the order of attackers
	attacker := &fighter1
	defender := &fighter2

	if defender.Name == firstAttacker {
		attacker, defender = defender, attacker
	}

	// Simulate the fight
	for attacker.Health > 0 && defender.Health > 0 {
		defender.Health -= attacker.DamagePerAttack

		// Check if defender is dead
		if defender.Health <= 0 {
			return attacker.Name
		}

		// Swap roles
		attacker, defender = defender, attacker
	}
	return ""
}

func DeclareWinner1(fighter1 Fighter, fighter2 Fighter, firstAttacker string) string {
  for {
    if firstAttacker == fighter1.Name {
      fighter2.Health -= fighter1.DamagePerAttack
      firstAttacker = fighter2.Name
    } else {
      fighter1.Health -= fighter2.DamagePerAttack
      firstAttacker = fighter1.Name
    }
        
    if fighter1.Health <= 0 {
      return fighter2.Name
    }
    
    if fighter2.Health <= 0 {
      return fighter1.Name
    }
  }
}

func DeclareWinner2(a Fighter, b Fighter, f string) string {
	if b.Name == f {a,b = b, a}
	if (b.Health/a.DamagePerAttack)<(a.Health/b.DamagePerAttack)||(b.Health/a.DamagePerAttack)==(a.Health/b.DamagePerAttack)&& a.Health%b.DamagePerAttack != 0 {
	  return a.Name
	} else {
	  return b.Name
	}
}