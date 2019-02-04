package titanikAPI

import "math/rand"

func (f *Field) addShip(lenght, x, y, direction int) bool {
	if !f.checkPoint(x, y) {
		return false
	}
	if !f.checkArea(lenght, x, y, direction) {
		return false
	}

	if direction == up {
		for i := y; i > y-lenght; i-- {
			f.setValueToPoint(x, i, ship)
		}
	}
	if direction == down {
		for i := y; i < y+lenght; i++ {
			f.setValueToPoint(x, i, ship)
		}
	}
	if direction == left {
		for i := x; i > x-lenght; i-- {
			f.setValueToPoint(i, y, ship)

		}
	}
	if direction == right {
		for i := x; i < x+lenght; i++ {
			f.setValueToPoint(i, y, ship)
		}
	}
	f.countOfShips++
	return true
}

func (f *Field) generateField() {
	k := 1
	for i := 4; i > 0; i-- {
		for j := 0; j < k; j++ {
			x := rand.Intn(size)
			y := rand.Intn(size)
			direction := rand.Intn(4)
			if !f.addShip(i, x, y, direction) {
				j--
			}
		}
		k++
	}

}
