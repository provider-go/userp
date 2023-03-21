# userp

```
package main

import (
	"fmt"
	"github.com/provider-go/userp"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	DBConn, err := gorm.Open(mysql.Open("root:pawword@tcp(ip:3306)/db?charset=utf8"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	provider := userp.GetUserProvider("default")
	provider.SetDB(DBConn)
	res, err := provider.View("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
```