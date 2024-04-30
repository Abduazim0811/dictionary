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
	fmt.Println("Barcha lug'atlar:\n")
	for key, value := range mp {
		fmt.Println(value, "-", key)
	}

	fmt.Println("Qidiruv: ")
	fmt.Scanln(&sentence)

	for key, value := range mp {
		if value == strings.TrimSpace(sentence) {
			fmt.Println(value, "-", key)
		}
	}

	fmt.Println("Bunday so'z mavjud emas!!")
}
