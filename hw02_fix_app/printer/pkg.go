package printer

import (
	"fmt"

	"github.com/mankinddived/home_work_basic/hw02_fix_app/types"
)

func PrintStaff(staff []types.Employee) {
	var str string
	for i := 0; i < len(staff); i++ {
		// str = fmt.Sprintf("Name: %s; Age: %d; ", staff[i].Name, staff[i].Age)
		fmt.Println(staff[i])
	}

	fmt.Println(str)
}
