package main

import "testing"

func TestInputNull(t *testing.T) {
	list := Init()
	list.AddNewPair(1, 1)

	data := []Pair{Pair{1, 1}}
	if !compare(list.GetList(), data) {
		t.Errorf("List was incorrect, got: %v", list.GetList())
	}
}

func TestInputAfter1(t *testing.T) {
	list := Init()
	list.AddNewPair(1, 1)
	list.AddNewPair(2, 2)
	list.AddNewPair(4, 2)

	data := []Pair{Pair{1, 1}, Pair{2, 2}}
	if !compare(list.GetList(), data) {
		t.Errorf("List was incorrect, got: %v", list.GetList())
	}
}

func TestInputMedian1(t *testing.T) {
	list := Init()
	list.AddNewPair(1, 1)
	list.AddNewPair(2, 2)
	list.AddNewPair(4, 2)
	list.AddNewPair(5, 5)
	list.AddNewPair(1, 6)
	list.AddNewPair(4, 4)
	list.AddNewPair(3, 3)

	data := []Pair{Pair{1, 1}, Pair{2, 2}, Pair{3, 3}, Pair{4, 4}, Pair{5, 5}}
	if !compare(list.GetList(), data) {
		t.Errorf("List was incorrect, got: %v", list.GetList())
	}
}

func TestInputBefore1(t *testing.T) {
	list := Init()
	list.AddNewPair(1, 1)
	list.AddNewPair(2, 2)
	list.AddNewPair(4, 2)
	list.AddNewPair(5, 5)
	list.AddNewPair(1, 6)
	list.AddNewPair(4, 4)
	list.AddNewPair(3, 3)
	list.AddNewPair(0, 0)

	data := []Pair{Pair{0, 0}, Pair{1, 1}, Pair{2, 2}, Pair{3, 3}, Pair{4, 4}, Pair{5, 5}}
	if !compare(list.GetList(), data) {
		t.Errorf("List was incorrect, got: %v", list.GetList())
	}
}

func TestInputBefore2(t *testing.T) {
	list := Init()
	list.AddNewPair(1, 1)
	list.AddNewPair(2, 2)
	list.AddNewPair(4, 2)
	list.AddNewPair(5, 5)
	list.AddNewPair(1, 6)
	list.AddNewPair(4, 4)
	list.AddNewPair(3, 3)
	list.AddNewPair(0, 0)
	list.AddNewPair(6, 7)
	list.AddNewPair(7, 6)
	list.AddNewPair(7, 8)

	data := []Pair{Pair{0, 0}, Pair{1, 1}, Pair{2, 2}, Pair{3, 3}, Pair{4, 4}, Pair{5, 5}, Pair{6, 7}, Pair{7, 8}}
	if !compare(list.GetList(), data) {
		t.Errorf("List was incorrect, got: %v", list.GetList())
	}
}

func TestNearestHeightBefore1(t *testing.T) {
	list := Init()
	list.AddNewPair(1, 1)
	list.AddNewPair(2, 2)
	list.AddNewPair(4, 2)
	list.AddNewPair(5, 5)
	list.AddNewPair(1, 6)
	list.AddNewPair(4, 4)
	list.AddNewPair(3, 3)
	list.AddNewPair(0, 0)
	list.AddNewPair(6, 7)
	list.AddNewPair(7, 6)
	list.AddNewPair(7, 8)

	if list.nearestHeightBefore(-1) != -1 {
		t.Errorf("List was incorrect, got: %v", list.GetList())
	}
}

func TestNearestHeightBefore2(t *testing.T) {
	list := Init()
	list.AddNewPair(1, 1)
	list.AddNewPair(2, 2)
	list.AddNewPair(4, 2)
	list.AddNewPair(5, 5)
	list.AddNewPair(1, 6)
	list.AddNewPair(4, 4)
	list.AddNewPair(3, 3)
	list.AddNewPair(0, 0)
	list.AddNewPair(6, 7)
	list.AddNewPair(7, 6)
	list.AddNewPair(7, 8)

	if list.nearestHeightBefore(-0) != -1 {
		t.Errorf("List was incorrect, got: %v", list.GetList())
	}
}

func TestNearestHeightBefore3(t *testing.T) {
	list := Init()
	list.AddNewPair(1, 1)
	list.AddNewPair(2, 2)
	list.AddNewPair(4, 2)
	list.AddNewPair(5, 5)
	list.AddNewPair(1, 6)
	list.AddNewPair(4, 4)
	list.AddNewPair(3, 3)
	list.AddNewPair(0, 0)
	list.AddNewPair(6, 7)
	list.AddNewPair(7, 6)
	list.AddNewPair(7, 8)

	if list.nearestHeightBefore(1) != 0 {
		t.Errorf("List was incorrect, got: %v", list.GetList())
	}
}

func TestNearestHeightBefore4(t *testing.T) {
	list := Init()
	list.AddNewPair(1, 1)
	list.AddNewPair(2, 2)
	list.AddNewPair(4, 2)
	list.AddNewPair(5, 5)
	list.AddNewPair(1, 6)
	list.AddNewPair(4, 4)
	list.AddNewPair(3, 3)
	list.AddNewPair(0, 0)
	list.AddNewPair(6, 7)
	list.AddNewPair(7, 6)
	list.AddNewPair(7, 8)

	if list.nearestHeightBefore(2) != 1 {
		t.Errorf("List was incorrect, got: %v", list.GetList())
	}
}
func TestNearestHeightBefore5(t *testing.T) {
	list := Init()
	list.AddNewPair(1, 1)
	list.AddNewPair(2, 2)
	list.AddNewPair(4, 2)
	list.AddNewPair(5, 5)
	list.AddNewPair(1, 6)
	list.AddNewPair(4, 4)
	list.AddNewPair(3, 3)
	list.AddNewPair(0, 0)
	list.AddNewPair(6, 7)
	list.AddNewPair(7, 6)
	list.AddNewPair(7, 8)

	if list.nearestHeightBefore(5) != 4 {
		t.Errorf("List was incorrect, got: %v", list.GetList())
	}
}
func TestNearestHeightBefore6(t *testing.T) {
	list := Init()
	list.AddNewPair(1, 1)
	list.AddNewPair(2, 2)
	list.AddNewPair(4, 2)
	list.AddNewPair(5, 5)
	list.AddNewPair(1, 6)
	list.AddNewPair(4, 4)
	list.AddNewPair(3, 3)
	list.AddNewPair(0, 0)
	list.AddNewPair(6, 7)
	list.AddNewPair(7, 6)
	list.AddNewPair(7, 8)

	if list.nearestHeightBefore(7) != 5 {
		t.Errorf("List was incorrect, got: %v", list.GetList())
	}
}
func TestNearestHeightBefore7(t *testing.T) {
	list := Init()
	list.AddNewPair(1, 1)
	list.AddNewPair(2, 2)
	list.AddNewPair(4, 2)
	list.AddNewPair(5, 5)
	list.AddNewPair(1, 6)
	list.AddNewPair(4, 4)
	list.AddNewPair(3, 3)
	list.AddNewPair(0, 0)
	list.AddNewPair(6, 7)
	list.AddNewPair(7, 6)
	list.AddNewPair(7, 8)

	if list.nearestHeightBefore(8) != 6 {
		t.Errorf("List was incorrect, got: %v", list.GetList())
	}
}
func TestNearestHeightBefore8(t *testing.T) {
	list := Init()
	list.AddNewPair(1, 1)
	list.AddNewPair(2, 2)
	list.AddNewPair(4, 2)
	list.AddNewPair(5, 5)
	list.AddNewPair(1, 6)
	list.AddNewPair(4, 4)
	list.AddNewPair(3, 3)
	list.AddNewPair(0, 0)
	list.AddNewPair(6, 7)
	list.AddNewPair(7, 6)
	list.AddNewPair(7, 8)

	if list.nearestHeightBefore(9) != 7 {
		t.Errorf("List was incorrect, got: %v", list.GetList())
	}
}
func TestNearestHeightBefore9(t *testing.T) {
	list := Init()
	list.AddNewPair(1, 1)
	list.AddNewPair(2, 2)
	list.AddNewPair(4, 2)
	list.AddNewPair(5, 5)
	list.AddNewPair(1, 6)
	list.AddNewPair(4, 4)
	list.AddNewPair(3, 3)
	list.AddNewPair(0, 0)
	list.AddNewPair(6, 7)
	list.AddNewPair(7, 6)
	list.AddNewPair(7, 8)

	if list.nearestHeightBefore(10) != 7 {
		t.Errorf("List was incorrect, got: %v", list.GetList())
	}
}
