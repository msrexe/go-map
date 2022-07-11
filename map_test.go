package gomap

import "testing"

func TestNewGoMap(t *testing.T) {
	gm := New[string, int](10)
	if gm.Length() != 0 {
		t.Errorf("Expected length of 0, got %d", gm.Length())
	}
}

func TestGoMap_Get(t *testing.T) {
	gm := New[string, int](10)
	gm.Set("key", 1)
	if gm.Get("key") != 1 {
		t.Errorf("Expected value of 1, got %d", gm.Get("key"))
	}
}

func TestGoMap_Set(t *testing.T) {
	gm := New[string, int](10)
	gm.Set("key", 1)
	if gm.Get("key") != 1 {
		t.Errorf("Expected value of 1, got %d", gm.Get("key"))
	}
}

func TestGoMap_Delete(t *testing.T) {
	gm := New[string, int](10)
	gm.Set("key", 1)
	gm.Delete("key")
	if gm.Length() != 0 {
		t.Errorf("Expected length of 0, got %d", gm.Length())
	}
}

func TestGoMap_Length(t *testing.T) {
	gm := New[string, int](10)
	gm.Set("key", 1)
	if gm.Length() != 1 {
		t.Errorf("Expected length of 1, got %d", gm.Length())
	}
}

func TestGoMap_Keys(t *testing.T) {
	gm := New[string, int](10)
	gm.Set("key", 1)
	keys := gm.Keys()
	if len(keys) != 1 {
		t.Errorf("Expected length of 1, got %d", len(keys))
	}
}

func TestGoMap_Values(t *testing.T) {
	gm := New[string, int](10)
	gm.Set("key", 1)
	values := gm.Values()
	if len(values) != 1 {
		t.Errorf("Expected length of 1, got %d", len(values))
	}
}

func TestGoMap_Clear(t *testing.T) {
	gm := New[string, int](10)
	gm.Set("key", 1)
	if gm.Length() != 1 {
		t.Errorf("Expected length of 1, got %d", gm.Length())
	}

	gm.Clear()
	if gm.Length() != 0 {
		t.Errorf("Expected length of 0, got %d", gm.Length())
	}
}

func TestGoMap_IsExists(t *testing.T) {
	gm := New[string, int](10)
	gm.Set("key", 1)
	if !gm.IsExists("key") {
		t.Errorf("Expected key to exist")
	}
}
