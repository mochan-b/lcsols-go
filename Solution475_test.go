package lcsols

import "testing"

func TestSoln475_1(t *testing.T) {
	var houses = []int{1, 2, 3}
	var heaters = []int{2}

	result := findRadius(houses, heaters)
	if result != 1 {
		t.Errorf("[1,2,3], [2] => %d; want 1", result)
	}
}

func TestSoln475_2(t *testing.T) {
	var houses = []int{1, 3, 2}
	var heaters = []int{2}

	result := findRadius(houses, heaters)
	if result != 1 {
		t.Errorf("[1,3,2], [2] => %d; want 1", result)
	}
}

func TestSoln475_3(t *testing.T) {
	var houses = []int{1, 2, 3, 4}
	var heaters = []int{1, 4}

	result := findRadius(houses, heaters)
	if result != 1 {
		t.Errorf("[1,2,3,4], [1,4] => %d; want 1", result)
	}
}

func TestSoln475_4(t *testing.T) {
	var houses = []int{1, 5}
	var heaters = []int{2}

	result := findRadius(houses, heaters)
	if result != 3 {
		t.Errorf("[1,5], [2] => %d; want 3", result)
	}
}

func TestSoln475_5(t *testing.T) {
	var houses = []int{1, 5}
	var heaters = []int{10}

	result := findRadius(houses, heaters)
	if result != 9 {
		t.Errorf("[1,5], [10] => %d; want 9", result)
	}
}

func TestSoln475_6(t *testing.T) {
	var houses = []int{282475249, 622650073, 984943658, 144108930, 470211272, 101027544, 457850878, 458777923}
	var heaters = []int{823564440, 115438165, 784484492, 74243042, 114807987, 137522503, 441282327, 16531729, 823378840, 143542612}

	result := findRadius(houses, heaters)
	if result != 161834419 {
		t.Errorf("unorderd vals => %d; want 161834419", result)
	}
}

func TestSoln475_7(t *testing.T) {
	var houses = []int{474833169, 264817709, 998097157, 817129560}
	var heaters = []int{197493099, 404280278, 893351816, 505795335}

	result := findRadius(houses, heaters)
	if result != 104745341 {
		t.Errorf("[474833169,264817709,998097157,817129560],[197493099,404280278,893351816,505795335] => %d; want 104745341", result)
	}
}
