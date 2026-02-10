package mathSlice

type Element int

type Slice []Element

// SumSlice returns sum of slice elements
func SumSlice(slice Slice) (result Element) {
	for _, val := range slice {
		result += val
	}
	return
}

// MapSlice applies the op function to each element
func MapSlice(slice Slice, op func(Element) Element) {
	for i, val := range slice {
		slice[i] = op(val)
	}
}

// FoldSlice - folds the slice
func FoldSlice(slice Slice, op func(Element, Element) Element, init Element) (result Element) {
	result = op(init, slice[0])
	for i := 1; i < len(slice); i++ {
		result = op(result, slice[i])
	}
	return
}
