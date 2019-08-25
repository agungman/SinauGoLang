package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var items = make(map[string]item, 5)

//type shippingaddress struct {
//	streetName, city, postcode string
//}

//type defaultaddress struct {
//	streetName, city, postCode string
//}

type customer struct {
	name, email, telephone, shippingaddress, defaultaddress string
}

type item struct{
	id string
	name string
	stock int
	price int
}

type cart struct {
	customer customer
	orderedQty int
	item item
}

func showMenu() {
	fmt.Printf("\nAdmin Menu:\n")
	fmt.Println("1.1 Product Entry")
	fmt.Println("1.2 Edit Product")
	fmt.Println()
	fmt.Println("\nCustomer Menu:\n")
	fmt.Println("2.1 User Registration")
	fmt.Println("2.2 Browse Product")
	fmt.Println("2.3 Add to Cart")
	fmt.Println("2.4 Show Cart Items")
	fmt.Println("2.5 Checkout")
	fmt.Println()
	fmt.Println("3. Exit")
	fmt.Println()
	fmt.Println("Please enter your selected menu number [e.g.: '1.1', '2.1', etc], or '3' to end the program.")
	fmt.Print("> ")
}

func productEntry(scanner *bufio.Scanner) {
	var item item
	fmt.Print("\nItem code:")
	scanner.Scan()
	item.id = scanner.Text()
	fmt.Print("\nItem name: ")
	scanner.Scan()
	item.name = scanner.Text()
	fmt.Print("\nItem stock: ")
	scanner.Scan()
	item.stock, _ = strconv.Atoi(scanner.Text())
	fmt.Print("\nItem price: ")
	scanner.Scan()
	item.price, _ = strconv.Atoi(scanner.Text())
	items[item.id] = item
}

var custRegist = customer{
	name:            "Apalah Saya",
	email:           "apalahsaya@ini.com",
	telephone:       "0983223232",
	shippingaddress: "Jakarta",
	defaultaddress:  "Jakarta",
}

func showProducts(scanner bufio.Scanner) {
	    for _, v := range items {
		fmt.Println("Item Code:", v.id)
		fmt.Println("Item Name:", v.name)
		fmt.Println()
//	shprod := items
//	fmt.Printf("%+v\n", shprod)
	}

func custRegist(scanner *bufio.Scanner) {
	var customer customer
	fmt.Print("\nNama Anda: ")
	scanner.Scan()
	customer.name = scanner.Text()
	fmt.Print("\nEmail Anda: ")
	scanner.Scan()
	customer.email = scanner.Text()
	fmt.Print("\nNo Telephone Anda: ")
	scanner.Scan()
	customer.telephone = scanner.Text()
	fmt.Print("\nAlamat Utama Anda: ")
	scanner.Scan()
	customer.defaultaddress = scanner.Text()
	fmt.Print("\nAlamat Pengiriman Anda: ")
	scanner.Scan()
	customer.shippingaddress = scanner.Text()
}

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	var selectedOption string
menu:
	for {
		showMenu()
		scanner.Scan()
		selectedOption = scanner.Text()
		switch selectedOption {
		case "1.1":
			productEntry(scanner)
		case "1.2":
			showProducts()
		case "2.1":
			custRegister(scanner)
		case "3":
			fmt.Println("Thanks for your visit!")
			break menu
		}
	}
}