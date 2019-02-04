package titanikAPI

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"sort"
)

type Bot struct {
	field            Field
	checkedCoordinates Coordinates
	//repetitivePoints   //частоповторяющиеся точки
	repetitivePoints []pointValue
	x                int
	y                int
	direction        int
	trueAttempts     int //количество верных попыток
	iterPoints       int //чтобы двигаться по repetitivePoints
}

//регулирование данных
func (b *Bot) regulationData() {
	repetitivePointsMF:= Coordinates{}

	file, err:= ioutil.ReadFile("manualFilling.txt")
	if err!=nil{
		return
	}
	repetitivePointsMF.SetCoordinatesJSON(file)

	repetitivePointsB:= Coordinates{}
	file, err = ioutil.ReadFile("botData.txt")
	if err!=nil{
		fmt.Println(err)
	}
	repetitivePointsB.SetCoordinatesJSON(file)

	var max *Coordinates
	var min *Coordinates
	if len(repetitivePointsB) >len(repetitivePointsMF){
		max = &repetitivePointsB
		min = & repetitivePointsMF
	}else{
		min = &repetitivePointsB
		 max= & repetitivePointsMF
	}
	for index, _ := range *max{
		(*min)[index]++
	}

	for index, value:= range *min{
		b.repetitivePoints = append(b.repetitivePoints,pointValue{index.X,index.Y,value})
	}

	wFile,err := os.Create("botData.txt")
	if err!= nil{
		return
	}
	jsData, _ := min.GetCoordinatesJSON()
	wFile.Write(jsData)
}


func (b *Bot) sortRepetitivePoints() {
	sort.Slice(b.repetitivePoints, func(i, j int) bool {
		return b.repetitivePoints[i].Value > b.repetitivePoints[j].Value
	})
	fmt.Println(b.repetitivePoints)
}

func (b *Bot) loadData(){
	repetitivePointsB:= Coordinates{}
	file, err := ioutil.ReadFile("botData.txt")
	if err!=nil{
		fmt.Println(err)
	}
	repetitivePointsB.SetCoordinatesJSON(file)
	for index, value:= range repetitivePointsB{
		b.repetitivePoints = append(b.repetitivePoints,pointValue{index.X,index.Y,value})
	}
}

func (b *Bot) initBot() {
	b.field.coordinates = Coordinates{}
	b.checkedCoordinates = Coordinates{}
	b.field.generateField()
	b.repetitivePoints = make([]pointValue, 0)
	b.x = -1
	b.y = -1
	b.direction = 1



}

func (b *Bot) getShot(x, y int) int {
	return b.field.getShot(x, y)
}

func (b *Bot) setResult(result int) {
	fmt.Println("result", result, b.trueAttempts)
	b.checkedCoordinates[Point{b.x, b.y}] += 1
	if result == killed {
		b.x, b.y = -1, -1
		b.trueAttempts = 0
		return
	}

	if result != ship && b.trueAttempts == 0 {
		b.x, b.y = -1, -1
		return
	}

	if result == ship && b.trueAttempts == 0 {
		b.trueAttempts += 1
		b.direction = b.field.incrDirection(b.x, b.y, b.direction)
		b.x, b.y = b.field.move(b.x, b.y, b.direction)
		return
	}

	if result != ship && b.trueAttempts == 1 {
		b.x, b.y = b.field.move(b.x, b.y, b.field.reverseDirection(b.direction)) //двигаемся в обратном направлении

		b.direction = b.field.incrDirection(b.x, b.y, b.direction)
		b.x, b.y = b.field.move(b.x, b.y, b.direction)
		return
	}
	if result == ship && b.trueAttempts > 0 {
		b.trueAttempts += 1
		b.x, b.y = b.field.move(b.x, b.y, b.direction)
		return
	}
	if result != ship && b.trueAttempts > 1 {
		b.direction = b.field.reverseDirection(b.direction) //разворачиваемся
		for i := 0; i <= b.trueAttempts; i++ {
			b.x, b.y = b.field.move(b.x, b.y, b.direction)
		}
		return
	}

}
func (b *Bot) makeShot() (int, int) {
	var x, y int

	if b.x == -1 && b.y == -1 {
		for {
			restart := false
			if b.iterPoints < len(b.repetitivePoints) {
				x = b.repetitivePoints[b.iterPoints].X
				y = b.repetitivePoints[b.iterPoints].Y
				b.checkedCoordinates[Point{x,y}]++
				b.iterPoints += 1
			} else {

				x = rand.Intn(size)
				y = rand.Intn(size)
			}
			if b.checkedCoordinates[Point{x, y}] > 0 {
				restart = true
			} else {
				b.checkedCoordinates[Point{x, y}]++
			}
			if !restart {
				break
			}
		}
		b.x, b.y = x, y
		return x, y
	} else {
		return b.x, b.y
	}

}

