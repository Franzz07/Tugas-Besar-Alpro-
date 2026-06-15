# Aplikasi Pengelolaan Laundry - Tugas Besar Algoritma & Pemrograman

Aplikasi Pengelolaan Laundry berbasis *Command Line Interface* (CLI) ini dikembangkan menggunakan bahasa pemrograman **Go (Golang)**. Program ini dirancang untuk mendokumentasikan, mengolah, dan merangkum transaksi *laundry* pelanggan secara terstruktur dengan menerapkan konsep-konsep dasar Algoritma dan Pemrograman (Alpro).

## 📌 Deskripsi & Spesifikasi Proyek

Aplikasi ini mencatat transaksi *laundry* pelanggan dengan spesifikasi sebagai berikut:
1. **Pengelolaan Data:** Mencatat entitas data pelanggan dan detail transaksi (*ID Transaksi, Nama, Tanggal, Berat, Jenis Layanan, dan Total Harga*).
2. **Operasi CRUD Dasar:** Menambah, mengubah, dan menghapus data pelanggan serta transaksi.
3. **Kalkulasi Biaya Otomatis:** Menghitung biaya *laundry* secara otomatis berdasarkan berat (Kg) dan jenis layanan yang dipilih.
4. **Pencarian Data (Searching):** Mencari transaksi berdasarkan Nama Pelanggan atau Tanggal tertentu.
5. **Pengurutan Laporan (Sorting):** Menampilkan daftar transaksi yang terurut berdasarkan Tanggal atau Total Harga secara *ascending* maupun *descending*.
6. **Rekapitulasi Keuangan:** Menampilkan total pendapatan bersih toko *laundry* dalam periode tanggal tertentu.

---

## ⚠️ Batasan Aturan Implementasi (Constraints)

Sesuai dengan ketentuan instruksi tugas akademik, program ini wajib mematuhi aturan struktur kode berikut:
* **Modular:** Kode dibuat menggunakan subprogram berupa fungsi dan prosedur yang dilengkapi parameter dan spesifikasi yang jelas.
* **Array Statis:** Penyimpanan data wajib menggunakan array statis dengan batas maksimum elemen (`NMAX = 1000`). Penggunaan *slice* dinamis atau fungsi `append` dilarang.
* **Tanpa Statement Lompat:** Dilarang menggunakan instruksi `break` (selain untuk `repeat-until`) maupun `continue` di dalam seluruh struktur perulangan. Pengendalian alur *looping* wajib memanfaatkan *Boolean Flag*.
* **Variabel Global Terbatas:** Penggunaan variabel global murni hanya diperbolehkan untuk array utama penampung data yang akan diolah.
* **Error Handling Mandiri:** Input menu utama, kategori layanan, dan opsi pengurutan dikunci dengan fungsi validasi khusus agar program tidak *crash* atau keluar secara tidak sengaja ketika menerima input yang salah.

---

## 🛠️ Algoritma yang Diimplementasikan

### 1. Searching (Pencarian Data)
* **Sequential Search:** Digunakan pada fitur pencarian berdasarkan **Nama Pelanggan** (`seqSearchNama`). Digunakan karena data string nama tersimpan secara acak di dalam memori.
* **Binary Search:** Digunakan pada fitur pencarian berdasarkan **Tanggal** (`binSearchTanggal`). Karena Binary Search hanya mengembalikan satu titik indeks, program ditambahkan logika pencarian batas kiri (`kiri--`) dan kanan (`kanan++`) untuk menyapu bersih semua transaksi kembar yang terjadi pada tanggal yang sama.

### 2. Sorting (Pengurutan Data)
* **Selection Sort:** Diterapkan untuk mengurutkan transaksi berdasarkan **Tanggal** (`selectionSortTanggal`). Algoritma ini wajib dijalankan secara urut naik (*ascending*) sebagai syarat mutlak sebelum melakukan *Binary Search*.
* **Insertion Sort:** Diterapkan untuk mengurutkan transaksi berdasarkan **Total Harga** (`insertionSortHarga`). Berfungsi menyajikan laporan transaksi dari yang termurah ke termahal (*ascending*) atau sebaliknya (*descending*).

---

## 🚀 Cara Menjalankan Aplikasi

1. Pastikan Anda sudah menginstal **Go Compiler** di perangkat Anda.
2. Unduh atau klon *repository* ini.
3. Buka Terminal / Command Prompt pada direktori proyek.
4. Jalankan perintah berikut:
   ```bash
   go run file.go
