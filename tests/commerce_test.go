package faker

import (
	"strings"
	"testing"
)

func init() {
	if err != nil {
		panic(err)
	}
}
func TestCommerceColor(t *testing.T) {
	color := f.Commerce().Color()
	if color != "" {
		return
	}
	t.Error("Could not generate color")
}

func TestDepartment(t *testing.T) {
	dept := f.Commerce().Department()
	if dept != "" {
		return
	}
	t.Error("Could not generate department")
}

func TestProductName(t *testing.T) {
	productName := f.Commerce().ProductName()
	splitName := strings.Split(productName, " ")

	if len(splitName) >= 3 {
		return
	}
	t.Errorf("Expected product name of length %v, got %v", 3, len(splitName))
}
