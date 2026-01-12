package stringset

import (
	"fmt"
	"strings"
)

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
	set := New()
	for _, str := range l {
		if _, present := set[str]; !present {
			set[str] = struct{}{}
		}
	}
	return set
}

func formatStringSliceIntoString(strSlice []string) string {
	for i, element := range strSlice {
		strSlice[i] = fmt.Sprintf("\"%s\"", element)
	}
	return fmt.Sprintf("{%s}", strings.Join(strSlice, ", "))
}

func (s Set) String() string {
	var strs []string
	for element, _ := range s {
		strs = append(strs, element)
	}
	return formatStringSliceIntoString(strs)
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
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
	if s1.IsEmpty() {
		return true
	}
	for element, _ := range s1 {
		if !s2.Has(element) {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	return Intersection(s1, s2).IsEmpty()
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
	return true
}

func Intersection(s1, s2 Set) Set {
	inter := New()
	for element, _ := range s1 {
		if _, exists := s2[element]; exists {
			inter[element] = struct{}{}
		}
	}
	return inter
}

func Difference(s1, s2 Set) Set {
	inter := New()
	for element, _ := range s1 {
		if _, exists := s2[element]; !exists {
			inter[element] = struct{}{}
		}
	}
	return inter
}

func cloneSet(s Set) Set {
	clonedSet := New()
	for element, _ := range s {
		clonedSet[element] = struct{}{}
	}
	return clonedSet
}

func Union(s1, s2 Set) Set {
	union := cloneSet(s2)
	if s1.IsEmpty() {
		return union
	}
	diff := Difference(s1, s2)
	for element, _ := range diff {
		union[element] = struct{}{}
	}
	return union
}
