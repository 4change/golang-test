package main

import "fmt"

type observer interface {
	update(string)
	getID() string
}

// =====================================================================================================================

type customer struct {
	id string
}

func (c *customer) update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.id, itemName)
}

func (c *customer) getID() string {
	return c.id
}
