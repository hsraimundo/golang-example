package polygon

import "testing"
import "math"

func TestTriangle_ThreeSides(t *testing.T){
	//Arrange
	poly := CreateTriangle()

	//Assert
	if len(poly.sides) != 3 {
		t.Fail()
	}
}

func TestTriangle_180InnerDegrees(t *testing.T){
	//Arrange
	poly := CreateTriangle()

	//Assert
	if (poly.sides[0].innerAngle + poly.sides[1].innerAngle + poly.sides[2].innerAngle) != 180 {
		t.Fail()
	}
}

func TestTriangle_SidesLength(t *testing.T){
	//Arrange
	poly := CreateTriangle()
	a := poly.sides[0].length
	b := poly.sides[1].length
	c := poly.sides[2].length

	//Assert
	if !(math.Abs(math.Abs(float64(b-c)))<float64(a) && float64(a)<float64(b+c)){
		t.Fail()
	}
}