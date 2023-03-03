# Test Tahap Pertama
Test tahap pertama dari PT Semesta Arus Teknologi


Cara menjalan program dari test ini
Pastikan sudah ter install Golang dan MongoDB

# Soal No. 1
Hanya bisa menyelesaikan di postman
Untuk hasil gambar bisa dilihat di folder Foto Result Soal No.1
Buka CMD didalam folder ini lalu ketik
"go run main.go"

Lalu buka postman
Pilih method yg sesuai berikut ini

Create Data
Method POST
http://localhost:9090/v1/user/create
Body raw: 
{
    "nama": "Kahfi",
    "umur": 22,
    "kelas": "SMK"
}

Get All Data
Method GET
http://localhost:9090/v1/user/getall

Get By Name
Method GET
http://localhost:9090/v1/user/get/Kahfi

Update Data by Name
Method PUT
http://localhost:9090/v1/user/update
Body raw: 
{
    "nama": "Kahfi",
    "umur": 22,
    "kelas": "Kuliah"
}

Delete Data by Name
Method DELETE
http://localhost:9090/v1/user/delete/Kahfi

# Soal No. 2
Buka CMD didalam folder ini lalu ketik

A. "go run prima.go"
B. "go run bintang.go"
C. "go run sort.go"
D. "go run palindrome.go"

