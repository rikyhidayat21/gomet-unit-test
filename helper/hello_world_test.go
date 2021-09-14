package helper

import (
	"fmt"
	"testing"
)

/*
Selain aturan nama file, di Golang juga sudah diatur untuk nama function unit test
Nama function untuk unit test harus diawali dengan nama Test
Misal jika kita ingin mengetest function HelloWorld, maka kita akan membuat function unit test dnegan nama TestHelloWorld
Tak ada aturan untuk nama belakang function unit test harus sama dengan nama function yang akan di test, yang penting adalah harus diawalin dengan kata Test
Selanjutnya harus memiliki parameter (t *testing.T) dan tidak mengembalikan return value

go test -v => running test di package tertentu
go test -v -run=FUNCTIONTEST => running test di function tertentu
go test ./... => running test di root project

jangan menggunakan panic untuk menggagalkan test, tapi menggunakan Fail(), FailNow(), Error(), dan Fatal()

Fail() => akan menggagalkan unit test, namun tetap melanjutkan eksekusi unit test. Namun diakhir ketika selsai, maka unit test tersebut dianggap gagal
FailNow() => akan menggagalkan unit test saat itu juga, tanpa melanjutkan eksekusi unit test
Lebih baik menggunakan function dibawah:
Error() => akan melakukan log error, kemudian menjalankan Fail()
Fatal() => akan melakukan log error, kemudian menjalankan FailNow() sehingga unit test berhenti
*/

func TestHelloWorldRiky(t *testing.T) {
	result := HelloWorld("Riky")
	if result != "Hello Riky" {
		// unit test failed
		t.Error("Result should be Hello Riky")
	}

	fmt.Println("TestHelloWorldRiky Done")
}
func TestHelloWorldHidayat(t *testing.T) {
	result := HelloWorld("Hidayat")
	if result != "Hello Hidayat" {
		// unit test failed
		t.Fatal("Result should be Hello Riky")
	}

	fmt.Println("TestHelloWorldHidayat Done")
}