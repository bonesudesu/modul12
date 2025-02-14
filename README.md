nomer 1 :
Fungsi Program:
Input Data:

Pengguna diminta memasukkan jumlah daerah yang akan dikelola.
Untuk setiap daerah, pengguna memasukkan jumlah rumah dan daftar nomor rumah di daerah tersebut.
Pengurutan Data:

Daftar nomor rumah dalam setiap daerah diurutkan menggunakan algoritma Selection Sort dalam urutan menaik.
Output Data:

Program menampilkan tabel yang mencantumkan nomor daerah, masukan asli (sebelum pengurutan), dan keluaran (setelah pengurutan).
Alur Eksekusi Program:
Pengguna memasukkan jumlah daerah.
Untuk setiap daerah:
Memasukkan jumlah rumah.
Memasukkan daftar nomor rumah.
Daftar nomor rumah diurutkan menggunakan Selection Sort.
Program menampilkan tabel dengan nomor daerah, daftar masukan, dan daftar keluaran (yang sudah diurutkan).
Penjelasan Fungsi:
selectionSort(arr []int): Mengurutkan array menggunakan algoritma Selection Sort.
formatMasukan(angka []int): Mengubah array angka menjadi string dengan format terpisah spasi.
lenArray(arr []int): Menghitung panjang array menggunakan loop manual (alternatif dari fungsi built-in len).

nomer 2 : 
Fungsi Program:
Input Data:

Pengguna diminta untuk memasukkan jumlah daerah.
Untuk setiap daerah, pengguna memasukkan jumlah rumah dan daftar nomor rumahnya.
Pemrosesan Data:

Pisahkan Ganjil dan Genap: Nomor rumah dipisahkan menjadi dua kelompok: ganjil dan genap.
Pengurutan:
Nomor rumah ganjil diurutkan dalam urutan menaik.
Nomor rumah genap diurutkan dalam urutan menurun.
Output Data:

Program menampilkan tabel dengan nomor daerah, daftar masukan asli, dan daftar keluaran (gabungan nomor rumah ganjil dan genap yang telah diurutkan sesuai aturan).
Penjelasan Fungsi Utama:
pisahkanGanjilGenap:
Memisahkan nomor rumah menjadi dua array: ganjil dan genap.
Mengembalikan array ganjil, genap, serta panjang masing-masing array.
sortAscending:
Mengurutkan array dalam urutan menaik menggunakan algoritma Selection Sort.
sortDescending:
Mengurutkan array dalam urutan menurun menggunakan algoritma Selection Sort.
formatMasukan dan formatKeluaran:
Membuat string terformat untuk menampilkan daftar masukan dan keluaran.
Alur Eksekusi Program:
Pengguna memasukkan jumlah daerah.
Untuk setiap daerah:
Memasukkan jumlah rumah dan nomor rumah.
Program memisahkan nomor rumah menjadi ganjil dan genap, mengurutkan keduanya, dan menggabungkannya kembali untuk ditampilkan.
Program menampilkan tabel dengan nomor daerah, daftar masukan, dan daftar keluaran.

nomer 3 : 
Fungsi Program:
Input Data:

Pengguna memasukkan angka secara berulang.
Jika angka 0 dimasukkan, program mengurutkan semua angka yang telah dimasukkan dan menghitung median dari angka-angka tersebut.
Program berhenti jika angka -5313 dimasukkan.
Pengurutan dan Median:

Angka-angka diurutkan menggunakan algoritma Insertion Sort.
Median dihitung sebagai:
Jika jumlah data ganjil, median adalah nilai tengah dari data yang diurutkan.
Jika jumlah data genap, median adalah rata-rata dari dua nilai tengah.
Output Data:

Median dari data yang telah diurutkan ditampilkan setiap kali angka 0 dimasukkan.
Fungsi Utama:
findMedian(arr []int):
Menghitung median dari array yang sudah diurutkan.
insertionSort(arr []int):
Mengurutkan array menggunakan algoritma Insertion Sort.
lenArray(arr []int):
Menghitung jumlah elemen yang valid dalam array (hingga elemen dengan nilai 0).
Alur Eksekusi:
Pengguna memasukkan angka satu per satu.
Jika angka 0 dimasukkan:
Angka-angka yang telah dimasukkan diurutkan.
Median dari angka-angka tersebut dihitung dan ditampilkan.
Program terus menerima input hingga pengguna memasukkan -5313, yang mengakhiri program.
