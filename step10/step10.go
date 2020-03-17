package main

import ( // <- Importacion de paquetes

	"hotaka-formacion-go/step8/restaurante"
)

func createOrders(orderlist chan *restaurante.Order) {

	var order *restaurante.Order = new(restaurante.Order)

	order.AddDish("Patatas Fritas", 3)
	order.AddDish("Ensalada Cesar", 2)

	orderlist <- order

	newOrder := new(restaurante.Order)
	newOrder.AddDish("Ensalada Cesar", 2)

	orderlist <- newOrder

	close(orderlist)
}

func main() {
	restaurante.CreateProductList()
	var orderlist chan *restaurante.Order

	orderlist = make(chan *restaurante.Order)

	//Esto levanta un hilo donde ejecuta la funcion
	go createOrders(orderlist)

	// El for se espera hasta se hace un close del channel
	for ord := range orderlist {
		ord.Show()
	}

}
