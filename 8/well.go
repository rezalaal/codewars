package kata8

/*
#Well of Ideas - Easy Version

For every good kata idea there seem to be quite a few bad ones!

In this kata you need to check the provided array (x) for good ideas 'good' and bad ideas 'bad'. 
If there are one or two good ideas, return 'Publish!', if there are more than 2 return 'I smell a series!'. 
If there are no good ideas, as is often the case, return 'Fail!'.
*/
func Well(x []string) string {
	nGood := 0
	for _, word :=range x {
		if word == "good" {
			nGood++
		}
	}
	switch nGood {
		case 0: return "Fail!"
		case 1,2: return "Publish!"
		default: return "I smell a series!"
	}
}