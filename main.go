package main

import (
	"fmt"
	"math"
)

func cekPrima(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func kelipatanTujuh(n int) bool {
	if n%7 == 0 {
		return true
	}else {
		return false
	}
}

func hitungLuasTrapesium(alas, tinggi, sisiMiring float64) float64 {
    luas := 0.5 * (alas + sisiMiring) * tinggi
    return luas
}

func main(){
	var num, prim, kelipatanT int
	var alas, tinggi, sisiM float64
	fmt.Println("Masukan Menu Pilihan : ")
	fmt.Println("1. Cek Prima ")
	fmt.Println("2. Cek Kelipatan 7 ")
	fmt.Println("3. Hitung Luas Trapesium ")

	fmt.Scanln(&num)
	switch num {
    case 1:
		fmt.Println("CEK PRIMA")
		fmt.Println("Masukkan bilangan :")
		fmt.Scanln(&prim)
		if cekPrima(prim){
			fmt.Printf("%d adalah bilangan prima\n", prim)
		}else {
			fmt.Printf("%d bukan bilangan prima\n", prim)
		}
    case 2:
		fmt.Println("CEK KELIPATAN TUJUH")
        fmt.Println("Masukkan bilangan :")
		fmt.Scanln(&kelipatanT)
		if kelipatanTujuh(kelipatanT) {
			fmt.Printf("%d adalah bilangan kelipatan tujuh\n", kelipatanT)
		}else {
			fmt.Printf("%d bukan bilangan kelipatan tujuh\n", kelipatanT)
		}
    case 3:
        fmt.Println("LUAS TRAPESIUM")
        fmt.Println("Masukkan alas :")
		fmt.Scanln(&alas)
		fmt.Println("Masukkan tinggi :")
		fmt.Scanln(&tinggi)
		fmt.Println("sisiMiring :")
		fmt.Scanln(&sisiM)
		luasTrapesium := hitungLuasTrapesium(alas, tinggi, sisiM)
    fmt.Println("Luas Trapesium:", luasTrapesium)
    default:
        fmt.Println("Input tidak valid")
    }
	
	
	
	
}

