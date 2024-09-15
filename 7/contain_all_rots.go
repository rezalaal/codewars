package kata7


/*
All Inclusive?

input:

a string strng
an array of strings arr
Output of function contain_all_rots(strng, arr) (or containAllRots or contain-all-rots):

a boolean true if all rotations of strng are included in arr
false otherwise
Examples:
contain_all_rots(

	"bsjq", ["bsjq", "qbsj", "sjqb", "twZNsslC", "jqbs"]) -> true

contain_all_rots(

	"Ajylvpy", ["Ajylvpy", "ylvpyAj", "jylvpyA", "lvpyAjy", "pyAjylv", "vpyAjyl", "ipywee"]) -> false)

Note:
Though not correct in a mathematical sense

we will consider that there are no rotations of strng == ""
and for any array arr: contain_all_rots("", arr) --> true
Ref: https://en.wikipedia.org/wiki/String_(computer_science)#Rotations
*/
func ContainAllRots(strng string, arr []string) bool {
	if strng == "" {
		return true
	}
	rotates := map[string]bool{}
	for i := 0; i <= len(strng); i++ {
		strng = rotate(strng)
		rotates[strng] = false
		for _, item := range arr {
			if strng == item {
				rotates[strng] = true
			}
		}
	}
	for _, item := range rotates {
		if !item {
			return false
		}
	}
	return true
}
func ContainAllRots2(strng string, arr []string) bool { 
    counter := 0
    master := []string{}
    for i,_ := range strng{
        master = append(master,strng[i:]+strng[:i])
    }
    for _,j := range master{
        for _,k := range arr{
            if j == k{
                counter ++
                break
            }
        }
    }
    return len(strng) == counter
}
func ContainAllRots3(strng string, arr []string) bool {
	rotations := make(map[string]struct{})
	for i := range strng {
	  rotations[strng[i:]+strng[:i]] = struct{}{}
	}
	
	for _, v := range arr {
	  delete(rotations, v)
	}
	return len(rotations) == 0
  }