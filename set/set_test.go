package set

import (
	"testing"
)

func TestAdd(t *testing.T) {
	s := NewHashSet[int](5)
	s.Add(1)
	s.Add(2)
	if s.Len() != 2 {
		t.Errorf("Expected length to be 2 but got %d", s.Len())
	}
}

func TestRemove(t *testing.T) {
	s := NewHashSet[int](5)
	s.Add(1)
	s.Add(2)
	s.Remove(1)
	if s.Len() != 1 {
		t.Errorf("Expected length to be 1 but got %d", s.Len())
	}
	if s.Contains(1) {
		t.Errorf("Expected set to not contain 1")
	}
}

func TestContains(t *testing.T) {
	s := NewHashSet[int](5)
	s.Add(1)
	if !s.Contains(1) {
		t.Errorf("Expected set to contain 1")
	}
	if s.Contains(2) {
		t.Errorf("Expected set to not contain 2")
	}
}

func TestLen(t *testing.T) {
	s := NewHashSet[int](5)
	s.Add(1)
	s.Add(2)
	s.Add(3)
	if s.Len() != 3 {
		t.Errorf("Expected length to be 3 but got %d", s.Len())
	}
}

func TestClear(t *testing.T) {
	s := NewHashSet[int](5)
	s.Add(1)
	s.Add(2)
	s.Clear()
	if s.Len() != 0 {
		t.Errorf("Expected length to be 0 but got %d", s.Len())
	}
}

func TestKeys(t *testing.T) {
	s := NewHashSet[int](5)
	s.Add(1)
	s.Add(2)
	keys := s.Keys()
	if len(keys) != 2 {
		t.Errorf("Expected length to be 2 but got %d", len(keys))
	}
	contains1 := false
	contains2 := false
	for _, k := range keys {
		if k == 1 {
			contains1 = true
		}
		if k == 2 {
			contains2 = true
		}
	}
	if !contains1 || !contains2 {
		t.Errorf("Keys do not contain the expected elements")
	}
}
