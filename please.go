package please

import "os"
import "fmt"
import "os/exec"

func main() {
    brew := append([]string{"brew"}, os.Args[1:]...)
    fmt.Println(brew)
    out, err := exec.Command("brew", os.Args[1:]...).CombinedOutput()
    if err != nil {
        fmt.Println("error occured")
        fmt.Printf("%s", err)
    }
    fmt.Printf("%s", out)
}
