package main

import "fmt"

/**
변수선언 및 초기화(Variable declaration & initialization)
var a int
var f float32
var f16 float64

a = 10
f = 12.0

var i,j,k int = 1,2,3

const c int = 10
const s string = "Hi"

const c = 10
const s = "Hi"
const (
	Apple = iota	// 0
	Grape 			// 1
	Orange			// 2
)

부울린 타입
 bool
문자열 타입
 string: string은 한번 생성되면 수정될 수 없는 Immutable 타입임
정수형 타입
 int int8 int16 int32 int64
 uint uint8 uint16 uint32 uint64 uintptr
Float 및 복소수 타입
 float32 float64 complex64 complex128
기타 타입
 byte: uint8과 동일하며 바이트 코드에 사용
 rune: int32과 동일하며 유니코드 코드포인트에 사용한다
**/

func main() {
	// Raw String Literal. 복수라인.
	rawLiteral := `아리랑\n
	아리랑\n
	  아라리요`
 
    // Interpreted String Literal
    interLiteral := "아리랑아리랑\n아리리요"
    // 아래와 같이 +를 사용하여 두 라인에 걸쳐 사용할 수도 있다.
    // interLiteral := "아리랑아리랑\n" + 
    //                 "아리리요"   
 
    fmt.Println(rawLiteral)
    fmt.Println()
    fmt.Println(interLiteral)

    var i int = -100
    var u uint = uint(i)
    var f float32 = float32(i)  
    println(f, u)
 
    str := "ABC"
    bytes := []byte(str)
    str2 := string(bytes)
    println(bytes, str2)
}