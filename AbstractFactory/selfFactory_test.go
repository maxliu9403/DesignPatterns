package AbstractFactory

import (
	"fmt"
	"testing"
)

func TestRdsFactories_NewRdsFactories(t *testing.T) {
	var rdsFactories RdsFactories
	rds := rdsFactories.NewRdsFactories()
	err := rds.Set("rds", "value")
	if err != nil {
		return
	}

	err, value := rds.Get("rds")
	fmt.Println(value)
}
