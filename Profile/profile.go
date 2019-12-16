package profile

import "fmt"

type User struct {
	FirstName string
	LastName  string
	Address   string
}

func (u User) GetProfile() {
	u.FirstName = "Aditia"
	u.LastName = "Pratama"
	u.Address = "Jln jogja"

	fmt.Println(u)
}

func (u User) SetProfile(fname string, lname string, address string) {
	u.FirstName = fname
	u.LastName = lname
	u.Address = address

	fmt.Println(u)
}
