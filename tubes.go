package main

import "fmt"

const NMAXMahasiswa int = 100
const NMAXJurusan int = 10

type daftarMahasiswa [NMAXMahasiswa]mahasiswa

type mahasiswa struct {
	nama, jurusan string
	nilai         int
	diterima      bool
}

type daftarJurusan [NMAXJurusan]jurusan

type jurusan struct {
	namaJurusan string
}

//FLOW APLIKASI

func main() {
	var n int
	var data daftarMahasiswa

	header()
	inputDataMahasiswa(&data, &n)
	fmt.Print(data[0].nama, " ", data[0].jurusan, " ", data[0].nilai)
}

func header() {
	fmt.Println("	Selamat datang di	")
	fmt.Println(" Aplikasi Pendaftaran Mahasiswa Baru")
	fmt.Println("=====================================")
}

func menuAdmin() {
	headerAdmin()
}

func headerAdmin() {
	fmt.Println("=====================================")
	fmt.Println("	Menu Admin	")
	fmt.Println("=====================================")
}

func opsiAdmin(data *daftarMahasiswa) {
	var opsi int
	fmt.Println("1. Tambah data mahasiswa")
	fmt.Println("2. Ubah data mahasiswa")
	fmt.Println("3. Hapus data mahasiswa")
	fmt.Println("4. Tambah data jurusan")
	fmt.Println("5. Ubah data jurusan")
	fmt.Println("6. Hapus data jurusan")
	fmt.Println("7. Kembali")
	fmt.Println("=====================================")

	if opsi == 1 {
		inputDataMahasiswa()
	} else if opsi == 2 {
		ubahDataMahasiswa()
	} else if opsi == 3 {
		hapusDataMahasiswa()
	} else if opsi == 4 {
		inputJurusan()
	} else if opsi == 5 {
		ubahJurusan()
	} else if opsi == 6 {
		hapusJurusan()
	} else if opsi == 7 {
		menuAdmin()
	} else {
		fmt.Println("Pilihan tidak tersedia")
		opsiAdmin()
	}
}

func menuMahasiswa() {
	headerMahasiswa()
}

func headerMahasiswa() {
	fmt.Println("=====================================")
	fmt.Println("	Menu Mahasiswa	")
	fmt.Println("=====================================")
	fmt.Println("1. Daftar")
	fmt.Println("2. Masuk")
	fmt.Println("3. Keluar")
	fmt.Println("=====================================")
}

func pilih(pilihan *int) {
	fmt.Println("	Pilih tipe pengguna: ")
	fmt.Println("	1. Admin")
	fmt.Println("	2. Mahasiswa")
	fmt.Scan(&pilihan)
	pilihannya(pilihan)
}

func pilihannya(pilihan *int) {
	if *pilihan == 1 {
		menuAdmin()
	} else if *pilihan == 2 {
		menuMahasiswa()
	} else {
		fmt.Println("Pilihan tidak tersedia")
		pilih(pilihan)
	}
}

//ADMIN

func isMahasiswaDiterimaDitolak(data *daftarMahasiswa, jumlahMahasiswa *int) {
	var nilaiUntukDiterima int
	fmt.Print("Masukkan nilai untuk diterima: ")
	fmt.Scanln(&nilaiUntukDiterima)

	for i := 0; i < *jumlahMahasiswa; i++ {
		if data[i].nilai >= nilaiUntukDiterima {
			data[i].diterima = true
		} else {
			data[i].diterima = false
		}
	}
}

//MAHASISWA

func inputDataMahasiswa(data *daftarMahasiswa, i *int) {
	var n int
	fmt.Print("Masukkan jumlah data: ")
	fmt.Scanln(&i)
	fmt.Print("Masukkan Nama: ")
	fmt.Scanln(&data[n].nama)
	fmt.Print("Masukkan Jurusan: ")
	fmt.Scanln(&data[n].jurusan)
	fmt.Print("Masukkan Nilai: ")
	fmt.Scanln(&data[n].nilai)

	for n < *i {
		n++
		fmt.Print("Masukkan Nama: ")
		fmt.Scanln(&data[n].nama)
		fmt.Print("Masukkan Jurusan: ")
		fmt.Scanln(&data[n].jurusan)
		fmt.Print("Masukkan Nilai: ")
		fmt.Scanln(&data[n].nilai)
	}
}

func searchMahasiswa(data *daftarMahasiswa, jumlahMahasiswa *int, nama string) mahasiswa { //sequential/linear search
	var ketemu mahasiswa
	var i int
	for i < *jumlahMahasiswa {
		if data[i].nama == nama {
			ketemu = data[i]
			return ketemu
		}
		i++
	}
	return ketemu
}

func ubahDataMahasiswa(data *daftarMahasiswa, jumlahMahasiswa *int, nama string) {
	var ketemu mahasiswa = searchMahasiswa(data, jumlahMahasiswa, nama)
	if ketemu.nama == nama {
		fmt.Println("Nama: ")
		fmt.Scanln(&ketemu.nama)
		fmt.Print("Jurusan: ", ketemu.jurusan)
		fmt.Scanln(&ketemu.jurusan)
		fmt.Print("Nilai: ", ketemu.nilai)
		fmt.Scanln(&ketemu.nilai)
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func hapusDataMahasiswa(data *daftarMahasiswa, jumlahMahasiswa *int, nama string) {
	var ketemu mahasiswa = searchMahasiswa(data, jumlahMahasiswa, nama)
	if ketemu.nama == nama {
		for i := 0; i < *jumlahMahasiswa; i++ {
			if data[i].nama == nama {
				for j := i; j < *jumlahMahasiswa; j++ {
					data[j] = data[j+1]
				}
				*jumlahMahasiswa--
			}
		}
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

//JURUSAN

func inputJurusan(data *daftarJurusan, i *int) {
	var n int
	fmt.Print("Masukkan jumlah data: ")
	fmt.Scanln(&i)
	fmt.Print("Masukkan Nama Jurusan: ")
	fmt.Scanln(&data[n].namaJurusan)

	for n < *i {
		n++
		fmt.Print("Masukkan Nama Jurusan: ")
		fmt.Scanln(&data[n].namaJurusan)
	}
}

func searchJurusan(data *daftarJurusan, jumlahJurusan *int, nama string) jurusan { //binary search
	var ketemu jurusan
	terendah := 0
	tertinggi := *jumlahJurusan - 1

	for terendah <= tertinggi {
		mid := (terendah + tertinggi) / 2

		if data[mid].namaJurusan == nama {
			ketemu = data[mid]
			return ketemu
		} else if data[mid].namaJurusan < nama {
			terendah = mid + 1
		} else {
			tertinggi = mid - 1
		}
	}
	return ketemu
}

func ubahJurusan(data *daftarJurusan, jumlahJurusan *int, nama string) {
	var ketemu jurusan = searchJurusan(data, jumlahJurusan, nama)
	if ketemu.namaJurusan == nama {
		fmt.Println("Nama Jurusan: ")
		fmt.Scanln(&ketemu.namaJurusan)
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func hapusJurusan(data *daftarJurusan, jumlahJurusan *int, nama string) {
	var ketemu jurusan = searchJurusan(data, jumlahJurusan, nama)
	if ketemu.namaJurusan == nama {
		for i := 0; i < *jumlahJurusan; i++ {
			if data[i].namaJurusan == nama {
				for j := i; j < *jumlahJurusan; j++ {
					data[j] = data[j+1]
				}
				*jumlahJurusan--
			}
		}
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

//TAMPIL DATA

func tampilkanMahasiswaBerdasarkanJurusan(data *daftarMahasiswa, jumlahMahasiswa *int, jurusan string) {
	fmt.Println("Data Mahasiswa Jurusan ", jurusan)
	fmt.Println("============================")
	var i int
	for i < *jumlahMahasiswa {
		if data[i].jurusan == jurusan {
			fmt.Println("Nama: ", data[i].nama)
			fmt.Println("Jurusan: ", data[i].jurusan)
			fmt.Println("Nilai: ", data[i].nilai)
		}
		i++
	}
}

func tampilkanMahasiswaYangDitolakBerdasarkanJurusan(data *daftarMahasiswa, jumlahMahasiswa *int, jurusan string) {
	fmt.Println("Data Mahasiswa Jurusan ", jurusan)
	fmt.Println("============================")
	var i int
	for i < *jumlahMahasiswa {
		if data[i].jurusan == jurusan && data[i].diterima == false {
			fmt.Println("Nama: ", data[i].nama)
			fmt.Println("Jurusan: ", data[i].jurusan)
			fmt.Println("Nilai: ", data[i].nilai)
		}
		i++
	}
}

func tampilkanMahasiswaYangDiterimaBerdasarkanJurusan(data *daftarMahasiswa, jumlahMahasiswa *int, jurusan string) {
	fmt.Println("Data Mahasiswa Jurusan ", jurusan)
	fmt.Println("============================")
	var i int
	for i < *jumlahMahasiswa {
		if data[i].jurusan == jurusan && data[i].diterima {
			fmt.Println("Nama: ", data[i].nama)
			fmt.Println("Jurusan: ", data[i].jurusan)
			fmt.Println("Nilai: ", data[i].nilai)
		}
		i++
	}
}

func tampilkanDaftarMahasiswa(data *daftarMahasiswa, jumlahMahasiswa int) {
	fmt.Println("Daftar Mahasiswa:")
	fmt.Println("============================")
	for i := 0; i < jumlahMahasiswa; i++ {
		fmt.Println("Nama: ", data[i].nama)
		fmt.Println("Jurusan: ", data[i].jurusan)
		fmt.Println("Nilai: ", data[i].nilai)
		fmt.Println("============================")
	}
}

func mahasiswaTerurutBerdasarkanNilaiAscending(data *daftarMahasiswa, jumlahMahasiswa int) {
	for i := 1; i < jumlahMahasiswa; i++ {
		key := data[i]
		j := i - 1

		for j >= 0 && data[j].nilai > key.nilai {
			data[j+1] = data[j]
			j--
		}

		data[j+1] = key
	}
}

func mahasiswaTerurutBerdasarkanNilaiDescending(data *daftarMahasiswa, jumlahMahasiswa int) {
	for i := 1; i < jumlahMahasiswa; i++ {
		key := data[i]
		j := i - 1

		for j >= 0 && data[j].nilai < key.nilai {
			data[j+1] = data[j]
			j--
		}

		data[j+1] = key
	}
}

func mahasiswaTerurutBerdasarkanJurusanAscending(data *daftarMahasiswa, jumlahMahasiswa int) {
	for i := 0; i < jumlahMahasiswa-1; i++ {
		minIndex := i

		for j := i + 1; j < jumlahMahasiswa; j++ {
			if data[j].jurusan < data[minIndex].jurusan {
				minIndex = j
			}
		}

		data[i], data[minIndex] = data[minIndex], data[i]
	}
}

func mahasiswaTerurutBerdasarkanJurusanDescending(data *daftarMahasiswa, jumlahMahasiswa int) {
	for i := 0; i < jumlahMahasiswa-1; i++ {
		maxIndex := i

		for j := i + 1; j < jumlahMahasiswa; j++ {
			if data[j].jurusan > data[maxIndex].jurusan {
				maxIndex = j
			}
		}

		data[i], data[maxIndex] = data[maxIndex], data[i]
	}
}

func mahasiswaTerurutBerdasarkanNamaAscending(data *daftarMahasiswa, jumlahMahasiswa int) {
	for i := 0; i < jumlahMahasiswa-1; i++ {
		for j := 0; j < jumlahMahasiswa-i-1; j++ {
			if data[j].nama > data[j+1].nama {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

func mahasiswaTerurutBerdasarkanNamaDescending(data *daftarMahasiswa, jumlahMahasiswa int) {
	for i := 0; i < jumlahMahasiswa-1; i++ {
		for j := 0; j < jumlahMahasiswa-i-1; j++ {
			if data[j].nama < data[j+1].nama {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}
