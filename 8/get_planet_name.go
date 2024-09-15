package kata8

/*
#Get Planet Name By ID

The function is not returning the correct values. Can you figure out why?

Example (Input --> Output ):

3 --> "Earth"

1: "Mercury"
2: "Venus"
3: "Earth"
4: "Mars"
5: "Jupiter"
6: "Saturn"
7: "Uranus"
8: "Neptune"
9: "Pluto"
*/
func GetPlanetName(ID int) string {
	switch ID {
	case 1:
		return "Mercury"
	case 2:
		return "Venus"
	case 3:
		return "Earth"
	case 4:
		return "Mars"
	case 5:
		return "Jupiter"
	case 6:
		return "Saturn"
	case 7:
		return "Uranus"
	case 8:
		return "Neptune"
	default:
		return "Pluto"
	}
}

var galaxy = []string{"Mercury","Venus","Earth","Mars","Jupiter","Saturn","Uranus","Neptune","Pluto"}
func GetPlanetName1(ID int) string {
  return galaxy[ID-1]
}