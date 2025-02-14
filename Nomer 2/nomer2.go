package main

import "fmt"

func main() {
	const maxDaerah = 100
	const maxRumah = 100

	var n int
	fmt.Print("Masukkan jumlah daerah kerabat Hercules: ")
	fmt.Scan(&n)

	var jumlahRumah [maxDaerah]int
	var masukan [maxDaerah][maxRumah]int

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan jumlah rumah di daerah %d: ", i+1)
		fmt.Scan(&jumlahRumah[i])

		fmt.Printf("Masukkan nomor rumah di daerah %d (pisahkan dengan spasi): ", i+1)
		for j := 0; j < jumlahRumah[i]; j++ {
			fmt.Scan(&masukan[i][j])
		}
	}

	fmt.Println("\nKeluarannya:")
	fmt.Printf("%-5s %-15s %s\n", "No", "Masukan", "Keluaran")
	for i := 0; i < n; i++ {
		ganjil, genap, panjangGanjil, panjangGenap := pisahkanGanjilGenap(masukan[i][:jumlahRumah[i]])
		sortAscending(ganjil[:panjangGanjil])
		sortDescending(genap[:panjangGenap])
		fmt.Printf("%-5d %-15s %s\n", i+1, formatMasukan(masukan[i][:jumlahRumah[i]]), formatKeluaran(ganjil[:panjangGanjil], genap[:panjangGenap]))
	}
}

func pisahkanGanjilGenap(arr []int) ([100]int, [100]int, int, int) {
	var ganjil [100]int
	var genap [100]int
	var panjangGanjil, panjangGenap int

	for i := 0; i < lenArray(arr); i++ {
		if arr[i]%2 == 0 {
			genap[panjangGenap] = arr[i]
			panjangGenap++
		} else {
			ganjil[panjangGanjil] = arr[i]
			panjangGanjil++
		}
	}
	return ganjil, genap, panjangGanjil, panjangGenap
}

func sortAscending(arr []int) {
	for i := 0; i < lenArray(arr)-1; i++ {
		minIdx := i
		for j := i + 1; j < lenArray(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func sortDescending(arr []int) {
	for i := 0; i < lenArray(arr)-1; i++ {
		maxIdx := i
		for j := i + 1; j < lenArray(arr); j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
	}
}

func formatMasukan(arr []int) string {
	var hasil string
	for i, v := range arr {
		if i > 0 {
			hasil += " "
		}
		hasil += fmt.Sprint(v)
	}
	return hasil
}

func formatKeluaran(ganjil, genap []int) string {
	var hasil string
	for i, v := range ganjil {
		if i > 0 {
			hasil += " "
		}
		hasil += fmt.Sprint(v)
	}
	for i, v := range genap {
		if len(ganjil) > 0 || i > 0 {
			hasil += " "
		}
		hasil += fmt.Sprint(v)
	}
	return hasil
}

func lenArray(arr []int) int {
	count := 0
	for range arr {
		count++
	}
	return count
}