package enguz

import (
	"database/sql"
	"fmt"

	md "github.com/Abduazim0811/dictionary/integral/models"
)

func Enguz(db *sql.DB) {
	var (
		sentence string
	)
	mp := md.MP

	fmt.Println("All dictinaries:")
	for key, value := range mp {
		fmt.Println(key, "-", value)
	}
	i := 0
	for i > 0 {
		son:=0
		fmt.Println("Search: ")
		fmt.Scanln(&sentence)
		value, found := mp[sentence]
		if found {
			fmt.Println(sentence, "-", value)
		} else {
			fmt.Println("Bunday so'zni tarjimasi yoq!!!")
		}
		fmt.Println("Bunday so'z mavjud emas!!")
		fmt.Println("yana malumot qidirasizmi?\n[1]HA\t[2]YOQ")
		fmt.Scanln(&son)
		if son==2{
			break
		}
	}

}
