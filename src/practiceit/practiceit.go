package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Dimension struct {
	length int
	width  int
	height int
}

func (d *Dimension) Volume() int {
	d.length = 20
	return d.length * d.width * d.height
}

func (d Dimension) Area() int {
	d.length = 10
	return d.length * d.height
}

func dimensions(length, width, height int) (area int, volume int) {
	area = length * width
	volume = length * width * height

	return
}

type Product struct {
	id        int
	name      string
	inventory int
	price     int
}

func main() {
	// area, volume := dimensions(5, 5, 5)
	// fmt.Println("Area : ", area, " :: Volume : ", volume)

	// d := Dimension{10, 10, 15}
	// fmt.Println("Volume of d is : ", d.Volume())

	// x, y := 10, 20
	// n := &x
	// fmt.Println("The pointer points to the address: ", n)
	// fmt.Println("The pointer n contains the address of x which has the value: ", *n)

	// *n = 50
	// fmt.Println("Value of x after it is changed by accessing through n is : ", x, ", Value of y is : ", y)

	// t := &y
	// *t = *t + 30
	// fmt.Println("Value of changed y is : ", y)

	// dim := Dimension{10, 5, 7}
	// fmt.Println("Printing volume in a pass method : ", dim.Volume())
	// fmt.Println("The value of d after the value is changed in Volume method: ", dim)
	// fmt.Println("The value of Area after reverted back to legth 10 : ", dim.Area())
	// fmt.Println("The value of dim still intact after change in Volume as Area is a value method : ", dim)

	// Connecting to DB

	db, err := sql.Open("sqlite3", "../../practiceit.db")
	if err != nil {
		log.Fatal(err.Error())
	}

	rows, err := db.Query("SELECT id, name, inventory, price FROM products")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var p Product
		rows.Scan(&p.id, &p.name, &p.inventory, &p.price)
		fmt.Println("Product : ", p.id, " ", p.name, " ", p.inventory, " ", p.price)
	}
}
