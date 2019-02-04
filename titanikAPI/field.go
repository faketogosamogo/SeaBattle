package titanikAPI

type Field struct{
	coordinates Coordinates
	countOfShips int
}


func (f *Field) setValueToPoint(x, y, value int) {
	f.coordinates.setValuePoint(x,y,value)
}
func (f *Field) getValueFromPoint(x, y int) int {
	return f.coordinates.getValuePoint(x,y)
}
func (f *Field) getShot(x, y int) int {
	if !f.isInside(x, y) {
		return notInside
	}

	if f.isEmpty(x, y) {
		f.setValueToPoint(x, y, miss)
		return empty ////tut pomanyat
	}
	if f.isHit(x, y) {
		return hit
	}

	if f.isMiss(x, y) {
		return miss
	}
	if f.isShip(x, y) {
		f.setValueToPoint(x, y, hit)
		if f.isKilled(x, y) {
			return killed
		}
		return ship
	}
	return empty
}
//ручное заполнение
