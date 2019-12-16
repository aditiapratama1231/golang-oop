package main

import (
	"oop/blog"
	"oop/product"
	"oop/profile"
)

func main() {
	profile := profile.User{}
	profile.GetProfile()
	profile.SetProfile("siapa", "kamu", "dimana")

	product := product.New("Aqua", 100)
	product.GetProduct()

	blog := blog.New("Tutorial golang", "oop in golang", "aditia")
	blog.GetBlog()
}
