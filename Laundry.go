package main

import "fmt"

const NMAX = 1000

type Transaksi struct {
	IDTrx         int
	NamaPelanggan string
	Tanggal       int
	Berat         float64
	JenisLayanan  string
	TotalHarga    float64
}

type ArrTransaksi [NMAX]Transaksi

var dataTransaksi ArrTransaksi
var nTransaksi int = 0

func formatRupiah(nominal float64) string {
	angka := int(nominal)
	if angka == 0 {
		return "Rp.0"
	}

	hasil := ""
	for angka > 0 {
		ratusan := angka % 1000
		angka = angka / 1000
		if angka > 0 {
			hasil = fmt.Sprintf(".%03d", ratusan) + hasil
		} else {
			hasil = fmt.Sprintf("%d", ratusan) + hasil
		}
	}
	return "Rp." + hasil
}

func validasiMenu() int {
	var pilihan int
	var valid bool = false

	for !valid {
		fmt.Print("Pilih menu (0-6): ")
		fmt.Scan(&pilihan)

		if pilihan >= 0 && pilihan <= 6 {
			valid = true
		} else {
			fmt.Println(">> ERROR: Pilihan tidak valid! Silakan masukkan angka antara 0 sampai 6.")
		}
	}
	return pilihan
}

// =====================================================================

func hitungBiaya(berat float64, jenis string) float64 {
	var hargaPerKg float64 = 0

	// Menggunakan switch case agar lebih rapi dari if-else berantai
	switch jenis {
	case "komplit":
		hargaPerKg = 8000
	case "cucikering":
		hargaPerKg = 5000
	case "setrika":
		hargaPerKg = 4000
	default:
		hargaPerKg = 0
	}

	return berat * hargaPerKg
}

func tambahTransaksi(idTrx int, nama string, tgl int, berat float64, jenis string) {
	if nTransaksi < NMAX {
		dataTransaksi[nTransaksi].IDTrx = idTrx
		dataTransaksi[nTransaksi].NamaPelanggan = nama
		dataTransaksi[nTransaksi].Tanggal = tgl
		dataTransaksi[nTransaksi].Berat = berat
		dataTransaksi[nTransaksi].JenisLayanan = jenis
		dataTransaksi[nTransaksi].TotalHarga = hitungBiaya(berat, jenis)
		nTransaksi++
		fmt.Println(">> Data transaksi berhasil dicatat!")
	} else {
		fmt.Println(">> Kapasitas data transaksi penuh!")
	}
}

func seqSearchNama(nama string) int {
	idx := -1
	i := 0
	for i < nTransaksi && idx == -1 {
		if dataTransaksi[i].NamaPelanggan == nama {
			idx = i
		}
		i++
	}
	return idx
}

func binSearchTanggal(tgl int) int {
	kiri := 0
	kanan := nTransaksi - 1
	idx := -1

	for kiri <= kanan && idx == -1 {
		tengah := (kiri + kanan) / 2
		if dataTransaksi[tengah].Tanggal == tgl {
			idx = tengah
		} else if dataTransaksi[tengah].Tanggal < tgl {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return idx
}

// Fungsi Selection Sort Tanggal murni Ascending (tanpa parameter bool)
func selectionSortTanggal() {
	var i, j, idxTujuan int
	var temp Transaksi

	i = 0
	for i < nTransaksi-1 {
		idxTujuan = i
		j = i + 1
		for j < nTransaksi {
			// Langsung urutkan dari terkecil ke terbesar
			if dataTransaksi[j].Tanggal < dataTransaksi[idxTujuan].Tanggal {
				idxTujuan = j
			}
			j++
		}
		temp = dataTransaksi[i]
		dataTransaksi[i] = dataTransaksi[idxTujuan]
		dataTransaksi[idxTujuan] = temp

		i++
	}
}

func insertionSortHarga(isAscending bool) {
	var i, j int
	var temp Transaksi

	i = 1
	for i < nTransaksi {
		temp = dataTransaksi[i]
		j = i

		if isAscending {
			for j > 0 && temp.TotalHarga < dataTransaksi[j-1].TotalHarga {
				dataTransaksi[j] = dataTransaksi[j-1]
				j--
			}
		} else {
			for j > 0 && temp.TotalHarga > dataTransaksi[j-1].TotalHarga {
				dataTransaksi[j] = dataTransaksi[j-1]
				j--
			}
		}
		dataTransaksi[j] = temp
		i++
	}
}

func hapusTransaksi(nama string) {
	idx := seqSearchNama(nama)

	if idx != -1 {
		i := idx
		for i < nTransaksi-1 {
			dataTransaksi[i] = dataTransaksi[i+1]
			i++
		}
		nTransaksi--
		fmt.Println(">> Data transaksi berhasil dihapus.")
	} else {
		fmt.Println(">> Data tidak ditemukan, gagal menghapus.")
	}
}

func hitungPendapatanPeriode(awal int, akhir int) float64 {
	var total float64 = 0
	i := 0
	for i < nTransaksi {
		if dataTransaksi[i].Tanggal >= awal && dataTransaksi[i].Tanggal <= akhir {
			total += dataTransaksi[i].TotalHarga
		}
		i++
	}
	return total
}

func tampilDataTransaksi() {
	if nTransaksi == 0 {
		fmt.Println(">> Belum ada data transaksi.")
		return
	}

	i := 0
	fmt.Println("--- Daftar Transaksi ---")
	for i < nTransaksi {
		fmt.Printf("ID: %d | Tgl: %d | Nama: %s | Layanan: %s | Berat: %.2f Kg | Total: %s\n", dataTransaksi[i].IDTrx, dataTransaksi[i].Tanggal, dataTransaksi[i].NamaPelanggan, dataTransaksi[i].JenisLayanan, dataTransaksi[i].Berat, formatRupiah(dataTransaksi[i].TotalHarga))
		i++
	}
	fmt.Println("------------------------")
}

func main() {
	var pilihan int
	var selesai bool = false

	for !selesai {
		fmt.Println("\n=== APLIKASI PENGELOLAAN LAUNDRY ===")
		fmt.Println("1. Tambah Transaksi")
		fmt.Println("2. Tampilkan Data (insertion)")
		fmt.Println("3. Cari Transaksi (Berdasarkan Nama)")
		fmt.Println("4. Urutkan & Cari Tanggal (Binary Search)")
		fmt.Println("5. Hapus Transaksi (Berdasarkan Nama)")
		fmt.Println("6. Hitung Total Pendapatan Periode")
		fmt.Println("0. Keluar")

		pilihan = validasiMenu()

		// Menggunakan switch case untuk menu utama
		switch pilihan {
		case 1:
			var jumlah int
			fmt.Print("Masukkan jumlah transaksi yang ingin ditambahkan: ")
			fmt.Scan(&jumlah)

			var k int = 0
			for k < jumlah {
				var id, tgl int
				var nama, jenis string
				var berat float64
				var pilihanLayanan int

				fmt.Printf("\n--- Data Transaksi ke-%d ---\n", k+1)
				fmt.Print("Masukkan ID Transaksi: ")
				fmt.Scan(&id)
				fmt.Print("Masukkan Nama Pelanggan (1 kata): ")
				fmt.Scan(&nama)
				fmt.Print("Masukkan Tanggal (YYYYMMDD): ")
				fmt.Scan(&tgl)
				fmt.Print("Masukkan Berat (Kg): ")
				fmt.Scan(&berat)

				// Validasi Pilihan Jenis Layanan
				var validLayanan bool = false
				for !validLayanan {
					fmt.Println("Pilih Jenis Layanan:")
					fmt.Println("1. Komplit		(8000/KG)")
					fmt.Println("2. Cuci Kering		(5000/KG)")
					fmt.Println("3. Setrika		(4000/KG)")
					fmt.Print("Pilih (1/2/3): ")
					fmt.Scan(&pilihanLayanan)

					// Switch case untuk pilihan layanan
					switch pilihanLayanan {
					case 1:
						jenis = "komplit"
						validLayanan = true
					case 2:
						jenis = "cucikering"
						validLayanan = true
					case 3:
						jenis = "setrika"
						validLayanan = true
					default:
						fmt.Println(">> ERROR: Pilihan layanan tidak valid! Masukkan angka 1, 2, atau 3.")
					}
				}

				tambahTransaksi(id, nama, tgl, berat, jenis)
				k++
			}

		case 2:
			var urut int

			// Validasi Pilihan Tampil Data
			var validUrut bool = false
			for !validUrut {
				fmt.Println("\nPilihan Tampil Data:")
				fmt.Println("1. Sesuai Urutan Input")
				fmt.Println("2. Urutkan Harga Termurah ke Termahal")
				fmt.Println("3. Urutkan Harga Termahal ke Termurah")
				fmt.Print("Pilih (1/2/3): ")
				fmt.Scan(&urut)

				if urut >= 1 && urut <= 3 {
					validUrut = true
				} else {
					fmt.Println(">> ERROR: Pilihan pengurutan tidak valid! Masukkan angka 1, 2, atau 3.")
				}
			}

			// Switch case untuk jenis sorting
			switch urut {
			case 2:
				insertionSortHarga(true)
				fmt.Println(">> Data telah diurutkan dari Termurah ke Termahal")
			case 3:
				insertionSortHarga(false)
				fmt.Println(">> Data telah diurutkan dari Termahal ke Termurah")
			}
			tampilDataTransaksi()

		case 3:
			var nama string
			fmt.Print("Masukkan Nama Pelanggan yang dicari (sama persis saat input): ")
			fmt.Scan(&nama)

			idx := seqSearchNama(nama)
			if idx != -1 {
				fmt.Printf(">> Ditemukan! Transaksi ID %d atas nama %s senilai %s\n", dataTransaksi[idx].IDTrx, dataTransaksi[idx].NamaPelanggan, formatRupiah(dataTransaksi[idx].TotalHarga))
			} else {
				fmt.Println(">> Data tidak ditemukan.")
			}

		case 4:
			var tgl int
			selectionSortTanggal()

			fmt.Print("Masukkan Tanggal yang dicari (YYYYMMDD): ")
			fmt.Scan(&tgl)

			idx := binSearchTanggal(tgl)
			if idx != -1 {
				fmt.Println(">> Ditemukan transaksi pada tanggal tersebut:")

				kiri := idx
				for kiri > 0 && dataTransaksi[kiri-1].Tanggal == tgl {
					kiri--
				}

				kanan := idx
				for kanan < nTransaksi-1 && dataTransaksi[kanan+1].Tanggal == tgl {
					kanan++
				}

				i := kiri
				for i <= kanan {
					fmt.Printf("- Trx ID: %d | Nama: %s | Layanan: %s | Total: %s\n", dataTransaksi[i].IDTrx, dataTransaksi[i].NamaPelanggan, dataTransaksi[i].JenisLayanan, formatRupiah(dataTransaksi[i].TotalHarga))
					i++
				}
			} else {
				fmt.Println(">> Tidak ada transaksi pada tanggal tersebut.")
			}

		case 5:
			var nama string
			fmt.Print("Masukkan Nama Pelanggan yang ingin dihapus transaksinya: ")
			fmt.Scan(&nama)
			hapusTransaksi(nama)

		case 6:
			var awal, akhir int
			fmt.Print("Masukkan Tanggal Awal (YYYYMMDD): ")
			fmt.Scan(&awal)
			fmt.Print("Masukkan Tanggal Akhir (YYYYMMDD): ")
			fmt.Scan(&akhir)

			total := hitungPendapatanPeriode(awal, akhir)
			fmt.Printf(">> Total Pendapatan dari %d s/d %d adalah: %s\n", awal, akhir, formatRupiah(total))

		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini!")
			selesai = true
		}
	}
}
