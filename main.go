package main

import (
	"oop/blog"
	"oop/product"
	"oop/profile"
	"oop/store"
)

func main() {
	profile := profile.User{}
	profile.GetProfile()
	profile.SetProfile("siapa", "kamu", "dimana")

	product := product.New("Aqua", 100)
	product.GetProduct()

	blog := blog.New("Tutorial golang", "oop in golang", "aditia")
	blog.GetBlog()

	branchStore := store.BranchStore{StoreName: "cabang toko bakpia", OwnerBranch: "Aditia"}
	store := store.Store{BranchStore: branchStore, StoreName: "Toko bakpia", Owner: "Rizky"}

	store.GetBranchStore()
}
