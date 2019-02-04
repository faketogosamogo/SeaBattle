package titanikAPI

import (
	"fmt"
	"net/http"
	"strconv"
)
////здесь будут логика ходов
type GameManagement struct{
	bot Bot
	userField Field
}

func (g *GameManagement)StartServerGame(){
	g.bot.initBot()
	g.userField.coordinates = Coordinates{}

	http.HandleFunc("/startGame",g.startGame)
	http.HandleFunc("/printUserField", g.printUserField)
	http.HandleFunc("/printBotField", g.printBotField)
	http.HandleFunc("/makeShotUser",g.makeShotUser)
	http.HandleFunc("/makeShotBot", g.makeShotBot)
	http.HandleFunc("/getUserJSONMap",g. getUserJSONMap)
	http.HandleFunc("/getBotJSONMap", g. getBotJSONMap)

	http.ListenAndServe("localhost:8000",nil)
}
//startGame: choise:1- ручное заполнение, 2 - автозаполнение!

func (g *GameManagement) startGame(w http.ResponseWriter, r* http.Request){
	g.userField.generateField()
	g.bot.initBot()
	g.bot.loadData()
	fmt.Fprintln(w,"Совершайте ход!")
}


func (g *GameManagement) printUserField(w http.ResponseWriter, r* http.Request){
	switch r.Method {
	case "GET":
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			value := g.userField.coordinates.getValuePoint(j, i)
			if value == empty {
				fmt.Fprint(w, "*  ")
			} else if value == ship {
				fmt.Fprint(w, "+  ")
			} else if value == hit {
				fmt.Fprint(w, "-  ")
			} else if value == miss {
				fmt.Fprint(w, "/  ")
			}
		}
		fmt.Fprintln(w)
	}
	//default:
		//fmt.Fprintln(w, "Только GET поддерживается")
	}
}
func(g *GameManagement) printBotField(w http.ResponseWriter, r* http.Request){
	switch r.Method {
	case "GET":
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			value := g.bot.field.coordinates.getValuePoint(j, i)
			if value == empty {
				fmt.Fprint(w, "*  ")
			} else if value == ship {
				fmt.Fprint(w, "+  ")
			} else if value == hit {
				fmt.Fprint(w, "-  ")
			} else if value == miss {
				fmt.Fprint(w, "/  ")
			}
		}
		fmt.Fprintln(w)
	}
	//default:
	//	fmt.Fprintln(w, "Только GET поддерживается")
	}
}

/*
func (g *GameManagement) manualFilling(w http.ResponseWriter, r* http.Request) {
	k := 1
	for i := 4; i > 0; i-- {
		for j := 0; j < k; j++ {
			var x, y, direction int
			fmt.Println("Введите x,y и направление(up = 1, down = 2, left  = 3,	right = 4 для корабля с длинной: ", i)
			fmt.Scan(&x, &y, &direction)
			if !g.userField.addShip(i, x, y, direction) {
				fmt.Println("Неверные координаты!")
				j--
			} else {
				g.printUserField(w,r)
			}
		}
		k++
	}

	file, err := os.Create("manualFilling.txt")
	if err != nil {
		fmt.Println("Ошибка файла!", err)
		os.Exit(1)
	}
	jsonCoordinates, _:= g.userField.coordinates.GetCoordinatesJSON()
	_,err=file.Write(jsonCoordinates)

	if err != nil {
		fmt.Println("Ошибка файла!", err)
		os.Exit(1)
	}


}
*/
func (g *GameManagement) makeShotUser(w http.ResponseWriter, r* http.Request){
	switch r.Method {
	case "PUT":
		X := r.URL.Query().Get("x")
		Y := r.URL.Query().Get("y")
		if X == "" || Y == "" {
			fmt.Fprintln(w, "Вы не ввели параметры")
			//g.printBotField(w,r)
			return
		}
		x, err := strconv.Atoi(X)
		if err != nil {
			fmt.Fprintln(w, "Походите заново!(Неверный ввод данных)")
			//g.printBotField(w,r)
			return
		}
		y, err := strconv.Atoi(Y)
		if err != nil {
			fmt.Fprintln(w, "Походите заново!(Неверный ввод данных)")
			//g.printBotField(w,r)
			return
		}
		if (x > 9 || x < 0) || (y > 9 || y < 0) {
			fmt.Fprintln(w, "Походите заново! Выход за пределы поля? !")
			//g.printBotField(w,r)
			return
		}
		result:=g.bot.getShot(x,y)
		fmt.Fprintln(w,x,y,result)
		if result!=empty{
			fmt.Fprintln(w,"Совершите ход Пользователя!")
		}else{
			if g.bot.field.countOfShips==0{
				fmt.Fprintln(w, "Вы выиграли!")
			}else {
				fmt.Fprintln(w, "Ход бота!")
			}
		}

	//default:
		//fmt.Fprintln(w, "Только PUT поддерживается")
	}
}
func (g *GameManagement) makeShotBot(w http.ResponseWriter, r* http.Request){
	switch r.Method {
	case "PUT":
	x,y:=g.bot.makeShot()
	result:=g.userField.getShot(x,y)
	g.bot.setResult(result)
	fmt.Fprintln(w,x,y,result)
		if result!=empty{
			fmt.Fprintln(w,"Совершите ход Бота!")
		}else{
			if g.bot.field.countOfShips==0{
				fmt.Fprintln(w, "Бот выиграл!")
			}else {
				fmt.Fprintln(w, "Ход Пользователя!")
			}
		}
	//default:
		//fmt.Fprintln(w, "Только PUT поддерживается")
	}
}

func (g *GameManagement) getUserJSONMap(w http.ResponseWriter, r* http.Request){
	switch r.Method {
	case "GET":
		JSON, _ := g.userField.coordinates.GetCoordinatesJSON()
		fmt.Fprintln(w, string(JSON))
	}
}
func (g *GameManagement) getBotJSONMap(w http.ResponseWriter, r* http.Request){
	switch r.Method {
	case "GET":
	JSON, _:= g.bot.field.coordinates.GetCoordinatesJSON()
	fmt.Fprintln(w,string(JSON))
	}
}