package titanikAPI


func (f Field) toUp(x, y int) (int, int) {
	if !f.isInside(x, y-1) {
		return -1, -1
	}
	return x, y - 1
}
func (f Field) toDown(x, y int) (int, int) {
	if !f.isInside(x, y+1) {
		return -1, -1
	}
	return x, y + 1
}
func (f Field) toLeft(x, y int) (int, int) {
	if !f.isInside(x-1, y) {
		return -1, -1
	}
	return x - 1, y
}
func (f Field) toRight(x, y int) (int, int) {
	if !f.isInside(x+1, y) {
		return -1, -1
	}
	return x + 1, y
}
func (f Field) move(x, y, direction int) (int, int) {
	if direction == up {
		return f.toUp(x, y)
	}
	if direction == down {
		return f.toDown(x, y)
	}
	if direction == left {
		return f.toLeft(x, y)
	}
	if direction == right {
		return f.toRight(x, y)
	}
	return -1, -1
}
func (f Field) incrDirection(x, y, direction int) int {

	for {
		if direction == right {
			direction = up
		} else {
			direction += 1
		}
		a, b := f.move(x, y, direction)
		if a != -1 && b != -1 {
			break
		}
	}
	return direction

}

