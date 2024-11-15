

// jawaban pertanyaan a dan b

// a. Jika nam diberikan adalah 80.1, apa keluaran dari program tersebut? Apakah eksekusi program tersebut sesuai spesifikasi soal?
//jika nam diberikan 80.1 program akan mencetak BC karena setelah pernyataan if nam > 80
//kondisi selanjutnya adalah nam >= 72.5, yang akan dievaluasi sebagai true dan variabel mk akan diatur ke BC

// b. Apa saja kesalahan dari program tersebut? Mengapa demikian? Jelaskan alur program seharusnya!
//urutan pernyataan if dalam program tidak tersusun dari terbesar ke terkecil, jadi beberapa kondisi dievaluasi secara salah,
//program mengevaluasi nam >= 72.5 setelah nam > 80 jadi menyebabkan program memilih BC meskipun nam seharusnya dikategorikan sebagai AB


package main
import "fmt"

func main() {
    var nam float64
    var mk string

    fmt.Print("Nilai akhir mata kuliah: ")
    fmt.Scan(&nam)

    if nam >= 89 {
        mk = "A"
    } else if nam >= 72.5 {
        mk = "AB"
    } else if nam >= 65 {
        mk = "B"
    } else if nam >= 57.5 {
        mk = "BC"
    } else if nam >= 50 {
        mk = "C"
    } else if nam >= 40 {
        mk = "D"
    } else {
        mk = "E"
    }

    fmt.Println("Nilai mata kuliah: ", mk)
}