package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tagihan struct {
	Nama string
	Nominal int
	Tanggal int
	Kategori string
	Status bool
}

var daftarTagihan = []Tagihan{
    { Nama: "Listrik PLN", Nominal: 150000, Tanggal: 20260630, Kategori: "Pokok", Status: false },
	{ Nama: "Air PDAM", Nominal: 80000, Tanggal: 20260616, Kategori: "Pokok", Status: true },
    { Nama: "Internet Indihome", Nominal: 350000, Tanggal: 20260615, Kategori: "Pokok", Status: true },
    { Nama: "Netflix", Nominal: 54000, Tanggal: 20260620, Kategori: "Sekunder", Status: false },
	{ Nama: "Youtube Premium", Nominal: 36000, Tanggal: 20260620, Kategori: "Sekunder", Status: true },
}

var scanner = bufio.NewScanner(os.Stdin)

/*
===========================
    1. TAMBAH TAGIHAN
===========================
*/
func tambahTagihan() {
	var t Tagihan
	var statusInput string

	fmt.Println()
	fmt.Println("--- Tambah Tagihan ---")

	fmt.Print("Nama Tagihan: ")
	if scanner.Scan() {
		t.Nama = scanner.Text()
	}

	fmt.Print("Nominal Biaya: ")
	fmt.Scan(&t.Nominal)
	scanner.Scan()

	fmt.Print("Tanggal Jatuh Tempo (YYYYMMDD): ")
	fmt.Scan(&t.Tanggal)
	scanner.Scan() 

	fmt.Print("Kategori: ")
	if scanner.Scan() {
		t.Kategori = scanner.Text()
	}

	fmt.Print("Status Pembayaran (lunas/belum): ")
	if scanner.Scan() {
		statusInput = scanner.Text()
	}

	if statusInput == "lunas" {
		t.Status = true
	} else {
		t.Status = false
	}

	daftarTagihan = append(daftarTagihan, t)
	fmt.Println()
	fmt.Println("Data berhasil ditambahkan!")
}

/*
=========================
    2. UBAH TAGIHAN
=========================
*/
func ubahTagihan() {
	tampilTagihan()
	if len(daftarTagihan) == 0 {
		return
	}

	var index int
	fmt.Println()
	fmt.Print("Pilih nomor tagihan yang ingin diubah: ")
	fmt.Scan(&index)
	scanner.Scan()

	if index > 0 && index <= len(daftarTagihan) {
		i := index - 1
		var statusInput string
		fmt.Println()
		fmt.Print("Nama Tagihan Baru: ")
		if scanner.Scan() {
			daftarTagihan[i].Nama = scanner.Text()
		}

		fmt.Print("Nominal Biaya Baru: ")
		fmt.Scan(&daftarTagihan[i].Nominal)
		scanner.Scan()

		fmt.Print("Tanggal Jatuh Tempo Baru (YYYYMMDD): ")
		fmt.Scan(&daftarTagihan[i].Tanggal)
		scanner.Scan()

		fmt.Print("Kategori Baru: ")
		if scanner.Scan() {
			daftarTagihan[i].Kategori = scanner.Text()
		}

		fmt.Print("Status Pembayaran Baru (lunas/belum): ")
		if scanner.Scan() {
			statusInput = scanner.Text()
		}

		if statusInput == "lunas" {
			daftarTagihan[i].Status = true
		} else {
			daftarTagihan[i].Status = false
		}

		fmt.Println()
		fmt.Println("Data berhasil diubah!")
	} else {
		fmt.Println()
		fmt.Println("Nomor tidak ditemukan.")
	}
}

/*
========================
    3. HAPUS TAGIHAN
========================
*/
func hapusTagihan() {
	tampilTagihan()
	if len(daftarTagihan) == 0 {
		return
	}

	var index int
	fmt.Println()
	fmt.Println("--- Hapus Tagihan ---")
	fmt.Print("Pilih nomor tagihan yang ingin dihapus: ")
	fmt.Scan(&index)
	scanner.Scan()

	if index > 0 && index <= len(daftarTagihan) {
		i := index - 1
		daftarTagihan = append(daftarTagihan[:i], daftarTagihan[i+1:]...)
		fmt.Println()
		fmt.Println("Data berhasil dihapus!")
	} else {
		fmt.Println()
		fmt.Println("Nomor tidak ditemukan.")
	}
}

/*
====================================
    4. TAMPILKAN SEMUA TAGIHAN
====================================
*/
func tampilTagihan() {
	fmt.Println()
	fmt.Println("--- Daftar Tagihan ---")
	if len(daftarTagihan) == 0 {
		fmt.Println("Data masih kosong.")
		return
	}
	for i, t := range daftarTagihan {
		statusStr := "Belum Lunas"
		if t.Status {
			statusStr = "Lunas"
		}
		fmt.Printf("%d. %s | Rp%d | Tempo: %d | Kategori: %s | Status: %s\n", i+1, t.Nama, t.Nominal, t.Tanggal, t.Kategori, statusStr)
	}
}

/*
=======================================================
    5. CARI TAGIHAN (SEQUENTIAL & BINARY SEARCH)
=======================================================
*/
func cariTagihan() {
	if len(daftarTagihan) == 0 {
		fmt.Println()
		fmt.Println("Data masih kosong.")
		return
	}

	var subMenu int
	fmt.Println()
	fmt.Println("--- Cari Tagihan ---")
	fmt.Println("1. Cari Berdasarkan Nama (Binary Search)")
	fmt.Println("2. Cari Berdasarkan Kategori (Sequential Search)")
	fmt.Print("Pilih metode pencarian: ")
	fmt.Scan(&subMenu)
	scanner.Scan()

	if subMenu == 1 {
		for i := 0; i < len(daftarTagihan)-1; i++ {
			for j := 0; j < len(daftarTagihan)-i-1; j++ {
				if daftarTagihan[j].Nama > daftarTagihan[j+1].Nama {
					daftarTagihan[j], daftarTagihan[j+1] = daftarTagihan[j+1], daftarTagihan[j]
				}
			}
		}

		var cari string
		fmt.Print("Masukkan Nama Tagihan yang dicari: ")
		if scanner.Scan() {
			cari = scanner.Text()
		}

		kiri := 0
		kanan := len(daftarTagihan) - 1
		ditemukan := false

		for kiri <= kanan {
			tengah := (kiri + kanan) / 2
			if daftarTagihan[tengah].Nama == cari {
				t := daftarTagihan[tengah]
				statusStr := "Belum Lunas"
				if t.Status {
					statusStr = "Lunas"
				}
				fmt.Printf("Ditemukan: %s | Rp%d | Tempo: %d | Kategori: %s | %s\n", t.Nama, t.Nominal, t.Tanggal, t.Kategori, statusStr)
				ditemukan = true
				break
			} else if daftarTagihan[tengah].Nama < cari {
				kiri = tengah + 1
			} else {
				kanan = tengah - 1
			}
		}

		if !ditemukan {
			fmt.Println()	
			fmt.Println("Nama tagihan tidak ditemukan.")
		}

	} else if subMenu == 2 {
		var cari string
		fmt.Println()
		fmt.Print("Masukkan Kategori Tagihan yang dicari: ")
		if scanner.Scan() {
			cari = scanner.Text()
		}

		ditemukan := false
		fmt.Println("Hasil Pencarian:")
		for _, t := range daftarTagihan {
			if t.Kategori == cari {
				statusStr := "Belum Lunas"
				if t.Status {
					statusStr = "Lunas"
				}
				fmt.Printf("- %s | Rp%d | Tempo: %d | %s\n", t.Nama, t.Nominal, t.Tanggal, statusStr)
				ditemukan = true
			}
		}
		if !ditemukan {
			fmt.Println()
			fmt.Println("Kategori tidak ditemukan.")
		}
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

/*
====================================
    6. URUTKAN TAGIHAN TERDEKAT 
====================================
*/
func urutkanTagihan() {
	if len(daftarTagihan) == 0 {
		fmt.Println()
		fmt.Println("Data masih kosong.")
		return
	}

	var subMenu int
	fmt.Println()
	fmt.Println("--- Urutkan Tagihan Terdekat ---")
	fmt.Println("1. Gunakan Selection Sort")
	fmt.Println("2. Gunakan Insertion Sort")
	fmt.Print("Pilih algoritma pengurutan: ")
	fmt.Scan(&subMenu)
	scanner.Scan()

	n := len(daftarTagihan)

	if subMenu == 1 {
		for i := 0; i < n-1; i++ {
			minIdx := i
			for j := i + 1; j < n; j++ {
				if daftarTagihan[j].Tanggal < daftarTagihan[minIdx].Tanggal {
					minIdx = j
				}
			}
			daftarTagihan[i], daftarTagihan[minIdx] = daftarTagihan[minIdx], daftarTagihan[i]
		}
		fmt.Println("Data berhasil diurutkan dengan Selection Sort!")
	} else if subMenu == 2 {
		for i := 1; i < n; i++ {
			key := daftarTagihan[i]
			j := i - 1
			for j >= 0 && daftarTagihan[j].Tanggal > key.Tanggal {
				daftarTagihan[j+1] = daftarTagihan[j]
				j--
			}
			daftarTagihan[j+1] = key
		}
		fmt.Println()
		fmt.Println("Data berhasil diurutkan dengan Insertion Sort!")
	} else {
		fmt.Println("Pilihan tidak valid.")
		return
	}
	tampilTagihan()
}

/*
====================================
    7. STATISTIK TOTAL BIAYA
====================================
*/
func tampilStatistik() {
	if len(daftarTagihan) == 0  {
		fmt.Println()
		fmt.Println("Belum ada data tagihan.")
		return
	}

	totalHarusDibayar := 0
	jumlahLunas := 0
	for _, t := range daftarTagihan {
		if t.Status {
			jumlahLunas++
		} else {
			totalHarusDibayar += t.Nominal
		}
	}
	persentaseLunas := float64(jumlahLunas) / float64(len(daftarTagihan)) * 100
	fmt.Println()
	fmt.Printf("--- Statistik Tagihan ---")
	fmt.Println()
	fmt.Printf("Jumlah Total Tagihan           : %d tagihan\n", len(daftarTagihan))
	fmt.Printf("Total Biaya                    : Rp%d\n", totalHarusDibayar)
	fmt.Printf("Persentase Tagihan Sudah Lunas : %.2f%%\n", persentaseLunas)
}

func main() {
	var pilihan int

	for {
		fmt.Println()
		fmt.Println("=== Aplikasi Manajemen Tagihan Bulanan (SIMTAB) ===")
		fmt.Println("1. Tambah Tagihan")
		fmt.Println("2. Ubah Tagihan")
		fmt.Println("3. Hapus Tagihan")
		fmt.Println("4. Tampilkan Semua Tagihan")
		fmt.Println("5. Cari Tagihan Berdasarkan Nama")
		fmt.Println("6. Urutkan Tagihan Terdekat")
		fmt.Println("7. Statistik Total Biaya")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		
		fmt.Scan(&pilihan)
		scanner.Scan() 

		switch pilihan {
		case 1:
			tambahTagihan()
		case 2:
			ubahTagihan()
		case 3:
			hapusTagihan()
		case 4:
			tampilTagihan()
		case 5:
			cariTagihan()
		case 6:
			urutkanTagihan()
		case 7:
			tampilStatistik()
		case 0:
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}