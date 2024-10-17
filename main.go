package main

import (
	"day4/finance" // Path relatif dari folder proyek
	"fmt"
)

func main() {
	// Membuat akun baru
	account1, err := finance.NewAccount("Luno", "lunos@email.com")
	if err != nil {
		fmt.Println("Gagal membuat akun:", err)
		return
	}

	account2, err := finance.NewAccount("Lumo", "lumos@email.com")
	if err != nil {
		fmt.Println("Gagal membuat akun:", err)
		return
	}

	fmt.Printf("Akun berhasil ditambahkan: [%+v] [%+v]\n", *account1, *account2)

	// Menambahkan saldo ke akun 1
	err = account1.Saldo.AddSaldo(1000)
	if err != nil {
		fmt.Println("Gagal menambahkan saldo:", err)
		return
	}

	// Menambahkan saldo ke akun 2
	err = account2.Saldo.AddSaldo(2000)
	if err != nil {
		fmt.Println("Gagal menambahkan saldo:", err)
		return
	}

	fmt.Printf("Saldo berhasil ditambahkan: [%+v] [%+v]\n", *account1, *account2)

	// Mengurangi saldo dari akun 1
	err = account1.Saldo.DeductSaldo(500)
	if err != nil {
		fmt.Println("Gagal mengurangi saldo:", err)
		return
	}

	// Mengurangi saldo dari akun 2
	err = account2.Saldo.DeductSaldo(1000)
	if err != nil {
		fmt.Println("Gagal mengurangi saldo:", err)
		return
	}

	fmt.Printf("Saldo berhasil dikurangi: [%+v] [%+v]\n", *account1, *account2)
}
