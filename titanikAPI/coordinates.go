package titanikAPI

import "encoding/json"
const(
	empty     = 0  //если на клетке ничего нет и не было
	hit       = -1 //если здесь произошло попадание по кораблю
	ship      = 1  //если здесь находится корабль
	miss      = 2	//если здесь ничего не находилось
	killed    = 3//
	notInside = -2

	size = 10

	up    = 1
	down  = 2
	left  = 3
	right = 4
)


type Point struct{
	X int `json:"x"`
	Y int `json:"y"`
}
type pointValue struct{
	X int `json:"x"`
	Y int `json:"y"`
	Value int `json:"value"`
}


type Coordinates map[Point]int


func(c *Coordinates)GetCoordinatesJSON() ([]byte, error){
	pointsValue:= make([]pointValue, 0)
	for key, value:= range *c{
		pointsValue = append(pointsValue, pointValue{key.X,key.Y,value})
	}
	return json.Marshal(pointsValue)
}
func (c* Coordinates) SetCoordinatesJSON(data []byte) error{
	*c = Coordinates{}
	pointsValue:= make([]pointValue, 0)
	err:= json.Unmarshal(data,&pointsValue)
	if err!=nil{
		return err
	}
	*c = Coordinates{}
	for _, value := range pointsValue{
		(*c)[Point{value.X,value.Y}] = value.Value
	}
	return nil
}

func (c* Coordinates)setValuePoint(x,y, value int){
	(*c)[Point{x,y}] = value
}
func (c Coordinates)getValuePoint(x,y int) int{
	return c[Point{x,y}]
}
//true если точка внутри поля!
func (c Coordinates)isInside(x, y int) bool{
	if (x<=9 &&x >=0) && (y <=9 && y >=0){
		return true
	}
	return false
}




///запись и чтение из файла записать