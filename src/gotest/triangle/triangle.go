package polygon

type Polygon struct {
	sides []PolygonSide
}

type PolygonSide struct {
	length int
	innerAngle int
}

func CreateTriangle() *Polygon{
	var poly = new(Polygon)
	poly.sides = make([]PolygonSide, 3, 3)
	poly.sides[0] = PolygonSide{length:4, innerAngle:90}
	poly.sides[1] = PolygonSide{length:3, innerAngle:30}
	poly.sides[2] = PolygonSide{length:5, innerAngle:60}

	return poly
}