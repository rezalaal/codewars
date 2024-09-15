package kata8

// Surface Area and Volume of a Box
// Write a function that returns the total surface area and volume of a box as an array: [area, volume]
func GetSize(w, h, d int) [2]int {
	result := [2]int{}
	volume := w * h * d
	area := 2 * (w*h + w*d + h*d)
	result[0] = area
	result[1] = volume
	return result
}
