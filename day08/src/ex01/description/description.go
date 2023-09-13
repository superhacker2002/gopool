package description

import (
	"fmt"
	"reflect"
)

func DescribePlant(plant any) {
	v := reflect.ValueOf(plant)
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Print(field.Name)
		if len(field.Tag) != 0 {
			fmt.Print("(", field.Tag, ")")
		}
		fmt.Printf(":%v\n", v.FieldByName(field.Name))
	}
}
