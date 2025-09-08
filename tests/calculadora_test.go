package main

import "testing"

func TestShouldSumCorrectly(t *testing.T) {
	result := soma(1, 2, 3, 4, 5)
	if result != 15 {
		t.Errorf("ShouldSumCorrectly: Esperado 15, mas obteve %d", result)
	}
}

func TestShouldSumTwoElements(t *testing.T) {
	result := soma(1, 2)
	if result != 3 {
		t.Errorf("ShouldSumTwoElements: Esperado 3, mas obteve %d", result)
	}
}

func TestShouldSubtractTwoElements(t *testing.T) {
	result := subtrai(10, 10)
	if result != 0 {
		t.Errorf("ShouldSubtractTwoElements: Esperado 0, mas obteve %d", result)
	}
}

func TestShouldSubtractCorrectly(t *testing.T) {
	result := subtrai(10, 10)
	if result != 0 {
		t.Errorf("ShouldSubtractCorrectly: Esperado 0, mas obteve %d", result)
	}
}

func TestShouldMultiplyCorrectly(t *testing.T) {
	result := multiplica(10, 10)
	if result != 100 {
		t.Errorf("ShouldMultiplyCorrectly: Esperado 100, mas obteve %d", result)
	}
}

func TestShouldDivideCorrectly(t *testing.T) {
	result := divide(10, 10)
	if result != 1 {
		t.Errorf("ShouldDivideCorrectly: Esperado 1, mas obteve %d", result)
	}
}
