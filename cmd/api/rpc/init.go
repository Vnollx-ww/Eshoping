package rpc

import "Eshop/pkg/viper"

func init() {
	userConfig := viper.Init("user")
	InitUser(&userConfig)
	productConfig := viper.Init("product")
	InitProduct(&productConfig)
	orderConfig := viper.Init("order")
	InitOrder(&orderConfig)
}
