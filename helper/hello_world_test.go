package helper

import "testing"

/*
Selain aturan nama file, di Golang juga sudah diatur untuk nama function unit test
Nama function untuk unit test harus diawali dengan nama Test
Misal jika kita ingin mengetest function HelloWorld, maka kita akan membuat function unit test dnegan nama TestHelloWorld
Tak ada aturan untuk nama belakang function unit test harus sama dengan nama function yang akan di test, yang penting adalah harus diawalin dengan kata Test
Selanjutnya harus memiliki parameter (t *testing.T) dan tidak mengembalikan return value

go test -v => running test di package tertentu
go test -v -run=FUNCTIONTEST => running test di function tertentu
go test ./... => running test di root project
*/

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Riky")
	if result != "Hello Riky" {
		// unit test failed
		panic("Result is not Hello Riky")
	}
}