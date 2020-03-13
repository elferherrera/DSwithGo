package set

// IntSet map to store bool options
type IntSet map[int]bool

// NewIntSet function creates a new map
func NewIntSet() IntSet {
	set := IntSet{}
	return set
}

// Add function adds new value to the map
func (set IntSet) Add(val int) IntSet {
	set[val] = true
	return set
}

// NewIntSetFromSlice creates a new map from a slice
func NewIntSetFromSlice(slice []int) IntSet {
	set := IntSet{}
	for _, val := range slice {
		set[val] = true
	}
	return set
}

// Contains checks if a map has an specific value
func (set IntSet) Contains(val int) bool {
	if _, ok := set[val]; ok {
		return true
	}
	return false
}

// IsSubsetOf checks if a map is a subset of another map
func (set IntSet) IsSubsetOf(set2 IntSet) bool {
	for key1 := range set {
		if !set2.Contains(key1) {
			return false
		}
	}
	return true
}

// IsSupersetOf checks if map is a superset of another map
func (set IntSet) IsSupersetOf(set2 IntSet) bool {
	for key2 := range set2 {
		if !set.Contains(key2) {
			return false
		}
	}
	return true
}

// Intersect creates a new map considering the intersection
// of two maps
func (set IntSet) Intersect(set2 IntSet) IntSet {
	newSet := NewIntSet()

	for key1 := range set {
		for key2 := range set2 {
			if key1 == key2 {
				newSet.Add(key1)
			}
		}
	}
	return newSet
}
