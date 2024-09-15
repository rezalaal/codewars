package kata7

func SequenceSum(start, end, step int) int {
	sum := 0
	for i:=start;i<=end;i+=step {
		sum += i
	}	
	return sum
}

func SequenceSum1(start, end, step int) int {
	if end < start { return 0 }
	n := (end-start)/step
	return (n + 1) * (2*start + n * step)/2
}