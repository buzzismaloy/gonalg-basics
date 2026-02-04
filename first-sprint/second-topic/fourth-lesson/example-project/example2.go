package main

import (
	"fmt"

	"example-project/contacts"
)

func main() {
	contacts.SetSupport("Support service")
	fmt.Println(contacts.GetContact())
	fmt.Println("Email:", contacts.Email)
}
