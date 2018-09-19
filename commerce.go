package faker

import (
	"fmt"
)

// Commercer Interface
type Commercer interface {
	Color() string
	Department() string
	ProductName() string
	Price(min, max int, s string) string
	ProductAdjective() string
	ProductMaterial() string
	Product() string
}

// Commerce struct
type Commerce struct {
	*Fake
}

// Color returns a commerce color
func (c *Commerce) Color() string {
	return c.pick(commercePrefix + "/color")
}

// Department returns a commerce department
func (c *Commerce) Department() string {
	return c.pick(commercePrefix + "/department")
}

// ProductName returns a commerce product name
func (c *Commerce) ProductName() string {
	return fmt.Sprintf("%s %s %s", c.ProductAdjective(), c.ProductMaterial(), c.Product())
}

// ProductAdjective returns a commerce product adjective
func (c *Commerce) ProductAdjective() string {
	return c.pick(commercePrefix + "/product_name/adjective")
}

// ProductMaterial returns a commerce product material
func (c *Commerce) ProductMaterial() string {
	return c.pick(commercePrefix + "/product_name/material")
}

// Product returns a commerce product
func (c *Commerce) Product() string {
	return c.pick(commercePrefix + "/product_name/product")
}

// Price returns a commerce price
func (c *Commerce) Price(min, max int, symbol string) string {
	if min < 0 || max < 0 {
		return fmt.Sprintf("%v%v", symbol, 0.0)
	}
	r := randomFloatRange(min, max)
	return fmt.Sprintf("%v%.2f", symbol, r)
}
