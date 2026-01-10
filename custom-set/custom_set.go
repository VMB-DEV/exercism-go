package stringset

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set map[string]struct{}

func New() Set {
	return map[string]struct{}{}
}

func NewFromSlice(l []string) Set {
	var set map[string]struct{}
	for _, str := range l {
		if _, present := set[str]; !present {
			set[str] = struct{}{}
		}
	}
	return set
}

func (s Set) String() string {
	var str string
	for element, _ := range s {
		str += element
	}
	return str
}

func (s Set) IsEmpty() bool {
	return s.IsEmpty()
}

func (s Set) Has(elem string) bool {
	_, has := s[elem]
	return has
}

func (s Set) Add(elem string) {
	if !s.Has(elem) {
		s[elem] = struct{}{}
	}
}

func Subset(s1, s2 Set) bool {
	for element, _ := range s2 {
		if s1.Has(element) {
			delete(s1, element)
		} else {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	panic("Please implement the Disjoint function")
}

func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}
	for element, _ := range s2 {
		if !s1.Has(element) {
			return false
		}
	}
}

func Intersection(s1, s2 Set) Set {
	inter := New()
	for element, _ := range s1 {

	}
}

func Difference(s1, s2 Set) Set {

}

func Union(s1, s2 Set) Set {
}
