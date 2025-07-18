package factorial_sums

import "testing"

func TestBelowRange(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{	
		{0, -1},   
		{-1, -1},      
		{-25303, -1},   
		{-2147483648, -1},
	}

	for _, tc := range testCases {
		result, err := GetFactorialSums(tc.input)
		if err != nil {
			t.Errorf("Unexpected error for input %d: %v", tc.input, err)
		}
		if result != tc.expected {
			t.Errorf("Expected %d for input %d, got %d", tc.expected, tc.input, result)
		}
	}
}

func TestBottomBoundary(t *testing.T) {
	result, err := GetFactorialSums(1)
	if err != nil {
		t.Errorf("Unexpected error for input 1: %v", err)
	}
	if result != 1 {
		t.Errorf("Expected 1 for input 1, got %d", result)
	}
}

func TestSmallFactorials(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{	
		{2, 2},   
		{5, 3},      
		{7, 9},   
		{8, 9},   
		{10, 27}, 
		{15, 45},
	}

	for _, tc := range testCases {
		result, err := GetFactorialSums(tc.input)
		if err != nil {
			t.Errorf("Unexpected error for input %d: %v", tc.input, err)
		}
		if result != tc.expected {
			t.Errorf("Expected %d for input %d, got %d", tc.expected, tc.input, result)
		}
	}
}

func TestMediumFactorials(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{22, 72},   
		{70, 459},
	}

	for _, tc := range testCases {
		result, err := GetFactorialSums(tc.input)
		if err != nil {
			t.Errorf("Unexpected error for input %d: %v", tc.input, err)
		}
		if result != tc.expected {
			t.Errorf("Expected %d for input %d, got %d", tc.expected, tc.input, result)
		}
	}
}

func TestLargeFactorials(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{500, 4599},  
		{1501, 16650}, 
		{9000, 132777},
		{25204, 421668},
	}

	for _, tc := range testCases {
		result, err := GetFactorialSums(tc.input)
		if err != nil {
			t.Errorf("Unexpected error for input %d: %v", tc.input, err)
		}
		if result != tc.expected {
			t.Errorf("Expected %d for input %d, got %d", tc.expected, tc.input, result)
		}
	}
}


func TestTopBoundary(t *testing.T) {
	result, err := GetFactorialSums(25205)
	if err != nil {
		t.Errorf("Unexpected error for input 25205: %v", err)
	}
	if result != 421443 {
		t.Errorf("Expected 421,443 for input 25,205, instead got %d", result)
	}
}

func TestAboveRange(t *testing.T) {
	testCases := []int{25206, 25302}

	for _, input := range testCases {
		_, err := GetFactorialSums(input)
		if err == nil {
			t.Errorf("Expected error for input %d (too large), but no error was returned", input)
		}
	}
}