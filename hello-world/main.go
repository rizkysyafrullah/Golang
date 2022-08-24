package main //File yg berpackage 'main' akan diexec pertama kali ketika program dijalankan

import "fmt" //fmt adalah package bawaan dr go, isinya berupa fungsi untuk keperluan I/O yg berhubungan dengan Text

func main() { /*dlm sebuah project hrs ada file prgrm yg di dlmnya berisi sebuah fungsi bernama main(),
	fgsi tsb hrs ada di file yg package nya bernama 'main'*/

	fmt.Println("Rizky Si Tampan") //PRINTLN fgsinya untuk menampilkan text ke layar, semaca 'cout' lah
} //fgsi main adlh yg dipanggil pertama kali pada saat eksekusi program
