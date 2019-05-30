package triangle

import "testing"

func TestGetTriangleType(t *testing.T) {
	tests := []struct {
		a, b, c  int
		expected triangleType
	}{
		{1001, 5, 6, UnknownTriangle},
		{1,5000,5,UnknownTriangle},
		{1,1,5000,UnknownTriangle},
		{-1,1,5000,UnknownTriangle},
		{1,-1,5000,UnknownTriangle},
		{0,3,-5,UnknownTriangle},
		{2,4,0,UnknownTriangle},
		{1,0,6,UnknownTriangle},
		{1,10,12,InvalidTriangle},
		{3,4,5,RightTriangle},
		{4,4,6,AcuteTriangle},
		{6,6,8,ObtuseTriangle},

	}

	for _, test := range tests {
		actual := getTriangleType(test.a, test.b, test.c)
		if actual != test.expected {
			t.Errorf("getTriangleType(%d, %d, %d)=%v; want %v", test.a, test.b, test.c, actual, test.expected)
		}
	}
}
