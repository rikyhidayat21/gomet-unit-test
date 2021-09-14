package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

***Assertion***
downlaod 3rd party using testify

***Skip Test***
t.Skip("args")

**Before and After Test**
testing.M => digunakan untuk mengatur eksekusi unit test, hal ini bisa digunakan untuk before dan after di unit test

Untuk membuatnya harus function dengan nama TestMain, dengan parameter testing.M
Jika terdapat funcion TestMain tersebut, maka secara otomatis Golang akan mengeksekusi function setiap kali
TestMain hanya dieksekusi sekali per golang pacckage, bukan setiap kali function ada

**Table Test**
digunakan untuk membuat 1 function, kemudian memasukan request dan expectasi banyak data
*/

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name string
		request string
		expected string
	}{
		{
			name: "riky",
			request: "riky",
			expected: "Hello riky",
		},
		{
			name: "hidayat",
			request: "hidayat",
			expected: "Hello hidayat",
		},
		{
			name: "paten",
			request: "paten",
			expected: "Hello paten",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Riky", func(t *testing.T) {
		result := HelloWorld("Riky")
		require.Equal(t, "Hello Riky", result, "Result must be 'Hello Riky'")
	})

	t.Run("Hidayat", func(t *testing.T) {
		result := HelloWorld("Hidayat")
		require.Equal(t, "Hello Hidayat", result, "Result must be 'Hello Hidayat'")
	})
}

func TestMain(m *testing.M) {
	// before => disini bisa untuk koneksi ke database dan lain lain
	fmt.Println("BEFORE unit test")

	m.Run()

	// after
	fmt.Println("After unit test")

}

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

func TestHelloWorldAssert (t *testing.T) {
	result := HelloWorld("Riky")
	assert.Equal(t, "Hello Riky", result, "Result should be 'Hello Riky'")

	fmt.Println("Executed")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can't run on windows OS")
	}

	result := HelloWorld("Riky")
	require.Equal(t, "Hello Riky", result, "Result should be 'Hello Riky'")
}