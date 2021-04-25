package main

import (
    "fmt"

    "main_mod/test_pkg3"
    "common_mod/test_pkg1"
    "common_mod/test_pkg2"
)

func main() {
    fmt.Println(test_pkg3.MyPkgFunc())
    fmt.Println(test_pkg1.MyPkgFunc1())
    fmt.Println(test_pkg2.MyPkgFunc1st())
    fmt.Println(test_pkg2.MyPkgFunc2nd())
    //fmt.Println(test_pkg1_hahaha.MyPkgFunc1())
}



