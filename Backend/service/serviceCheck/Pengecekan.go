package servicecheck

import "strings"

func CheckGender(jenisKelamin string) int {
	// Mengubah string jenis kelamin menjadi huruf kecil untuk memudahkan perbandingan
	jenisKelamin = strings.ToLower(jenisKelamin)

	// Mengecek jenis kelamin
	if jenisKelamin == "laki-laki" {
		return 1
	}

	return 2

}

func CheckNikah(Nikah string) int {
	// Mengubah string jenis kelamin menjadi huruf kecil untuk memudahkan perbandingan
	Nikah = strings.ToLower(Nikah)

	// Mengecek jenis kelamin
	if Nikah == "belum menikah" {
		return 1
	}

	if Nikah == "menikah" {
		return 2
	}

	if Nikah == "cerai" {
		return 3
	}

	if Nikah == "janda/duda" {
		return 4
	}

	return 5

}

func CheckFinancial(Financial string) int {
	// Mengubah string jenis kelamin menjadi huruf kecil untuk memudahkan perbandingan
	Financial = strings.ToLower(Financial)

	// Mengecek jenis kelamin
	if Financial == "sangat mampu" {
		return 1
	}

	if Financial == "mampu" {
		return 2
	}

	if Financial == "tidak mampu" {
		return 3
	}

	return 4

}

func CheckReligil(Financial string) int {
	// Mengubah string jenis kelamin menjadi huruf kecil untuk memudahkan perbandingan
	Financial = strings.ToLower(Financial)

	// Mengecek jenis kelamin
	if Financial == "islam" {
		return 1
	}

	if Financial == "kristen protestan" {
		return 2
	}

	if Financial == "katolik" {
		return 3
	}

	if Financial == "hindu" {
		return 4
	}

	return 5

}
