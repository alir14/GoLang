package main

// import (
// 	"fmt"
// 	"net/http"
// )

// // const x int = 5

// type car struct {
// 	gas_pedal       uint16
// 	break_pedal     uint16
// 	streering_wheel int16
// 	top_speed_kmh   float64
// }

// // value receiver and pointer reciever
// // const usixteenbitmax float64 = 65535
// // const kmh_multiple float64 = 1.60934

// func main() {
// 	// fmt.Println("Hello Go language")
// 	// foo()
// 	// fmt.Println("Random Number ", rand.Intn(150))
// 	// var num1 float64 = 2.4
// 	// var num2 float64 = 8.43
// 	// var num1, num2 float64 = 5.6, 9.4
// 	// num1, num2 := 5.6, 9.6
// 	// fmt.Println(add(num1, num2))
// 	// x := 15
// 	// a := &x
// 	// fmt.Println(a)
// 	// fmt.Println(*a)
// 	// http.HandleFunc("/", index_Handler)
// 	// http.HandleFunc("/about", about_Handler)
// 	// http.ListenAndServe(":8000", nil)

// 	// a_car := car{
// 	// 	gas_pedal:       144,
// 	// 	break_pedal:     500,
// 	// 	streering_wheel: 12,
// 	// 	top_speed_kmh:   500.00}
// 	// fmt.Println(a_car.gas_pedal)
// 	// fmt.Println(a_car.kmh())
// 	// fmt.Println(a_car.mph())
// 	// a_car.new_top_speed(200)
// 	// fmt.Println(a_car.kmh())
// 	// fmt.Println(a_car.mph())
// 	// a_car = update_new_tope_speed(a_car, 100)
// 	// fmt.Println(a_car.top_speed_kmh)
// 	http.HandleFunc("/", index_handler)
// 	http.ListenAndServe(":8000", nil)
// }

// // func (c car) kmh() float64 {
// // 	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax)
// // }

// // func (c car) mph() float64 {
// // 	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax / kmh_multiple)
// // }

// // func (c *car) new_top_speed(newspeed float64) {
// // 	c.top_speed_kmh = newspeed
// // }

// // func update_new_tope_speed(c car, speed float64) car {
// // 	c.top_speed_kmh = speed
// // 	return c
// // }

// // func index_Handler(w http.ResponseWriter, r *http.Request) {
// // 	fmt.Fprint(w, "Uhm..., what do you want ?")
// // }

// // func about_Handler(w http.ResponseWriter, r *http.Request) {
// // 	fmt.Fprint(w, "Uhm..., is it what you are looking for ?")
// // }

// func index_handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "<h1>Hey there</h1>")
// 	fmt.Fprint(w, "<h2>Go lang</h2>")
// 	fmt.Fprintf(w, `<p>OMG </p>
// 		<p>I do not like this lang </p>
// 		<p>Why on earth do people look for easy language like that</p>
// 		`)
// }

// // func foo() {
// // 	fmt.Println("The square root of 4 is ", math.Sqrt(4))
// // }
// // func add(x float64, y float64) float64 {
// // 	return x + y
// // }
// // func multiple(a, b string) (string, string) {
// // 	return a, b
// // }
