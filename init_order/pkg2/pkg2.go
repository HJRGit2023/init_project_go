package pkg2

import "fmt"

const PkgName = "pkg2"

var PkgVar = getPkgName()

func init() {
	fmt.Println("pkg2 init method invoked")
}

func getPkgName() string {
	fmt.Println("pkg2.PkgNameVar has been initialized")
	return PkgName
}


