package titanikAPI

func (f Field) isInside(x, y int) bool {
	if !f.coordinates.isInside(x,y) {
		return false
	}
	return true
}
func (f Field) isEmpty(x, y int) bool {
	if f.coordinates.getValuePoint(x, y) == empty {
		return true
	}
	return false
}
func (f Field) isHit(x, y int) bool {
	if f.coordinates.getValuePoint(x, y) == hit {
		return true
	}
	return false
}
func (f Field) isMiss(x, y int) bool {
	if f.coordinates.getValuePoint(x, y) == miss {
		return true
	}
	return false
}
func (f Field) isShip(x, y int) bool {
	if f.coordinates.getValuePoint(x, y) == ship {
		return true
	}
	return false
}
func (f *Field) isKilled(x, y int) bool {
	if !f.shipAroundPoint(x, y) { //работает только для корабля длины==1
		f.deletePoint(x, y)
		return true
	}

	direction := f.getShipDirection(x, y)
	if direction == -1 {
		return false
	}
	startX := x
	startY := y
	//fmt.Println(x, y, direction)
	for {
		if t1, t2 := f.move(startX, startY, direction); t1 == -1 && t2 == -1 {
			break
		}
		if f.isShip(f.move(startX, startY, direction)) {
			return false
		}
		if f.isEmpty(f.move(startX, startY, direction)) || f.isMiss(f.move(startX, startY, direction)) {
			break
		}
		if f.isHit(f.move(startX, startY, direction)) {
			startX, startY = f.move(startX, startY, direction)
		}
	}

	finishX := startX
	finishY := startY

	//fmt.Println(startX, startY)
	direction = f.reverseDirection(direction)

	for {
		if t1, t2 := f.move(finishX, finishY, direction); t1 == -1 && t2 == -1 {
			break
		}
		if f.isShip(f.move(finishX, finishY, direction)) {
			return false
		}
		if f.isEmpty(f.move(finishX, finishY, direction)) || f.isMiss(f.move(finishX, finishY, direction)) {
			break
		}
		if f.isHit(f.move(finishX, finishY, direction)) {
			finishX, finishY = f.move(finishX, finishY, direction)
		}
	}

	f.deleteShip(startX, startY, finishX, finishY, direction)
	f.countOfShips--
	return true
}
func (f Field) checkPoint(x, y int) bool {
	if !f.isInside(x, y) {
		return false
	}
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !f.isInside(x+j, y+i) {
				continue
			}
			if !f.isEmpty(x+j, y+i) && !f.isMiss(x+j, y+i) {
				return false
			}
		}
	}
	return true
}
func (f Field) checkArea(lenght, x, y, direction int) bool {
	if direction == up {
		if y-lenght < 0 {
			return false
		}
		for i := y; i > y-lenght; i-- {
			if !f.checkPoint(x, i) {
				return false
			}
		}
		return true
	}
	if direction == down {
		if y+lenght > 9 {
			return false
		}
		for i := y; i < y+lenght; i++ {
			if !f.checkPoint(x, i) {
				return false
			}
		}
		return true
	}
	if direction == left {
		if x-lenght < 0 {
			return false
		}
		for i := x; i > x-lenght; i-- {
			if !f.checkPoint(i, y) {
				return false
			}
		}
		return true
	}
	if direction == right {
		if x+lenght > 9 {
			return false
		}
		for i := x; i < x+lenght; i++ {
			if !f.checkPoint(i, y) {
				return false
			}
		}
		return true
	}
	return false
}



