package main

import (
	"log"
	"os"
	"text/template"
)

type item struct {
	Name, Descrip string
	Price         float64
}

type meal struct {
	Meal  string
	Items []item
}

type menu []meal

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	m := menu{
		meal{
			Meal: "Breakfast",
			Items: []item{
				item{
					Name:    "Oatmeal",
					Descrip: "yum yum",
					Price:   4.95,
				},
				item{
					Name:    "Cheerios",
					Descrip: "American eating food traditional now",
					Price:   3.95,
				},
				item{
					Name:    "Juice Orange",
					Descrip: "Delicious drinking in throat squeezed fresh",
					Price:   2.95,
				},
			},
		},
		meal{
			Meal: "Lunch",
			Items: []item{
				item{
					Name:    "Hamburger",
					Descrip: "Delicous good eating for you",
					Price:   6.95,
				},
				item{
					Name:    "Cheese Melted Sandwich",
					Descrip: "Make cheese bread melt grease hot",
					Price:   3.95,
				},
				item{
					Name:    "French Fries",
					Descrip: "French eat potatoe fingers",
					Price:   2.95,
				},
			},
		},
		meal{
			Meal: "Dinner",
			Items: []item{
				item{
					Name:    "Pasta Bolognese",
					Descrip: "From Italy delicious eating",
					Price:   7.95,
				},
				item{
					Name:    "Steak",
					Descrip: "Dead cow grilled bloody",
					Price:   13.95,
				},
				item{
					Name:    "Bistro Potatoe",
					Descrip: "Bistro bar wood American bacon",
					Price:   6.95,
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, m)
	if err != nil {
		log.Fatalln(err)
	}

}
