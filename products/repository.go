package products

import "github.com/Cepheix/example-gin-webapi/common"

func AddProduct(product *ProductEntity) error {
	database := common.GetDB()

	database.Create(&product)
	err := database.Save(&product).Error
	return err
}
