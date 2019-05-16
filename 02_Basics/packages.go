// goではmainパッケージから開始される
package main

// このプログラムでは"fmt"と"math/rand"パッケージをインポートしている。
// 
// 規約で、パッケージ名はインポートパスの最後の要素と同じ名前になる。
// 例えば、インポートパスが "math/rand" のパッケージは、pakcage randステートメントで始まるファイル群で構成される。
import (
    "fmt"
    "math/rand"
)

func main() {
    fmt.Println("My favorite number is", rand.Intn(10))
}

