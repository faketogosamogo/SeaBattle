package titanikAPI
//возвращает true если вокруг точки есть ship or hit
func (f Field) shipAroundPoint(x, y int) bool {
	//d1, d2 := -1, -1
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if j == 0 && i == 0 {
				continue
			}
			if !f.isInside(x+j, y+i) {
				continue
			}
			if f.getValueFromPoint(x+j, i+y) == ship || f.getValueFromPoint(x+j, i+y) == hit {
				return true
			}
		}
	}
	return false
}
//-1 если у корабля ещё есть оставшиеся точки, в остальных случаях направление
func (f Field) getShipDirection(x, y int) int {

	if X, Y := f.toUp(x, y); X != -1 && Y != -1 {
		if f.isShip(f.toUp(x, y)) {
			return -1
		}
		if f.isHit(f.toUp(x, y)) {
			return up
		}
	}
	if X, Y := f.toDown(x, y); X != -1 && Y != -1 {
		if f.isShip(f.toDown(x, y)) {
			return -1
		}
		if f.isHit(f.toDown(x, y)) {
			return down
		}
	}
	if X, Y := f.toLeft(x, y); X != -1 && Y != -1 {
		if f.isShip(f.toLeft(x, y)) {
			return -1
		}
		if f.isHit(f.toLeft(x, y)) {
			return left
		}
	}
	if X, Y := f.toRight(x, y); X != -1 && Y != -1 {
		if f.isShip(f.toRight(x, y)) {
			return -1
		}
		if f.isHit(f.toRight(x, y)) {
			return right
		}
	}
	return -1
}
func (f Field) reverseDirection(direction int) int {
	if direction == up {
		return down
	}
	if direction == down {
		return up
	}
	if direction == left {
		return right
	}
	if direction == right {
		return left
	}
	return direction
}
func (f *Field) deletePoint(x, y int) {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !f.isInside(x+j, y+i) {
				continue
			}
			if !f.isHit(x+j, y+i) {
				f.setValueToPoint(x+j, y+i, miss)
			}
		}
	}
}
func (f *Field) deleteShip(startX, startY, finishX, finishY, direction int) {
	//fmt.Println(startX, startY, finishX, finishY, direction)
	if direction == up {
		for i := startY; i >= finishY; i-- {
			f.deletePoint(startX, i)
		}
	}
	if direction == down {
		for i := startY; i <= finishY; i++ {
			f.deletePoint(startX, i)
		}
	}
	if direction == left {
		for i := startX; i >= finishX; i-- {
			f.deletePoint(i, startY)
		}
	}
	if direction == right {
		for i := startX; i <= finishX; i++ {
			f.deletePoint(i, startY)
		}
	}

}

