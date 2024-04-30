package uzeng

import (
	"database/sql"
	"fmt"
	"strings"

	md "github.com/Abduazim0811/dictionary/integral/models"
)

func Uzeng(db *sql.DB) {

	var (
		sentence string
	)
	mp := md.MP
	fmt.Println("Barcha lug'atlar:")
	for key, value := range mp {
		fmt.Println(value, "-", key)
	}
	i := 1
	for i > 0 {
		son:=0
		fmt.Println("Qidiruv: ")
		fmt.Scanln(&sentence)

		for key, value := range mp {
			if value == strings.TrimSpace(sentence) {
				fmt.Println(value, "-", key)
			}
		}

		fmt.Println("Bunday so'z mavjud emas!!")
		fmt.Println("yana malumot qidirasizmi?\n[1]HA\t[2]YOQ")
		fmt.Scanln(&son)
		if son==2{
			break
		}
	}

}
