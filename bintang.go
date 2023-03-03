package main
import "fmt"

func main() {

    var b int

    fmt.Printf("\n->Silahkan ketik jumlah bintang : \n")//mulai dari angka terkecil
    fmt.Scan(&b)
    fmt.Println("input : ", b)

    fmt.Printf("result :\n")
    for brs:=0; brs<b; brs++{
        //melakukan pengulangan bintang(*) sampai 30

        for spasi:=0 ; spasi<=brs; spasi++{
            fmt.Printf("* ");
        }

        fmt.Printf("\n");
    }
}