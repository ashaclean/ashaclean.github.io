package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pesanan struct {
	NomorTelepon string
	Layanan      string
	Lokasi       string
}

var daftarLayanan = map[int]string{
	1: "Pembersihan Rumah Reguler",
	2: "Pembersihan Sofa & Karpet",
	3: "Pembersihan Pasca Renovasi",
	4: "Disinfeksi & Sterilisasi",
}

func main() {
	fmt.Println("=== SISTEM PEMESANAN ASHACLEAN.ID ===")
	fmt.Println("Silakan isi data pemesanan Anda:")

	// Input nomor telepon
	nomorTelepon := inputData("Nomor Telepon (contoh: 081234567890): ", func(input string) bool {
		return len(input) >= 10 && len(input) <= 15
	})

	// Pilih layanan
	fmt.Println("\nPilih Layanan:")
	for kode, layanan := range daftarLayanan {
		fmt.Printf("%d. %s\n", kode, layanan)
	}

	layanan := inputData("Masukkan nomor layanan (1-4): ", func(input string) bool {
		kode := 0
		_, err := fmt.Sscan(input, &kode)
		return err == nil && kode >= 1 && kode <= 4
	})
	namaLayanan := daftarLayanan[toInt(layanan)]

	// Input lokasi
	lokasi := inputData("Lokasi (contoh: Jl. Sudirman No. 10, Jakarta Selatan): ", func(input string) bool {
		return len(input) > 5
	})

	// Simpan pesanan
	pesanan := Pesanan{
		NomorTelepon: formatNomorTelepon(nomorTelepon),
		Layanan:      namaLayanan,
		Lokasi:       lokasi,
	}

	// Konfirmasi
	fmt.Println("\n=== RINCIAN PESANAN ===")
	fmt.Printf("No. Telepon: %s\n", pesanan.NomorTelepon)
	fmt.Printf("Layanan: %s\n", pesanan.Layanan)
	fmt.Printf("Lokasi: %s\n", pesanan.Lokasi)
	fmt.Println("Terima kasih! Pesanan Anda akan segera diproses.")
}

func inputData(prompt string, validator func(string) bool) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if validator(input) {
			return input
		}
		fmt.Println("Input tidak valid, silakan coba lagi")
	}
}

func toInt(s string) int {
	var i int
	fmt.Sscan(s, &i)
	return i
}

func formatNomorTelepon(nomor string) string {
	if strings.HasPrefix(nomor, "0") {
		return "62" + nomor[1:]
	}
	if strings.HasPrefix(nomor, "+62") {
		return nomor[1:]
	}
	if strings.HasPrefix(nomor, "62") {
		return nomor
	}
	return "62" + nomor
}