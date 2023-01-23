# RRI Audiobook API

Repositori untuk server API RRI Audiobook yang menjadi Project Magang di RRI Suemenep.

# Cara Pakai

- Instal [go](https://go.dev) dan [cmake](https://cmake.org).
- Jalankan perintah `go get` untuk memasang semua dependensi yang dibutuhkan.
- Jalankan perintah `make run` atau `go run main.go` untuk menjalankan server.
- Buka alamat `127.0.0.1:8080` atau `localhost:8080` pada browser.

# Postman

Impor berkas `RRI Play.postman_collection.json` ke postman, lalu buat global variable pada postman dengan nama:

- HOST: isi alamat server http yang berjalan.
- BEARER: _copas_ token dengan nama `acces_token` yang didapat saat login disini.
