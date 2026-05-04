package main

import "fmt"

func main() {
	store := NewStore()
	store.Put("name", []byte("joe"))

	val, err := store.Get("name")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("name", string(val))
	store.Delete("name")

	_, err = store.Get("name")
	if err != nil {
		fmt.Println("Deleted", err)
	}
}
