package main

import (
	"Eshop/dal/db"
	"Eshop/dal/rs"
	product "Eshop/kitex_gen/product"
	product2 "Eshop/pkg/kafka/producer/product"
	"Eshop/pkg/middlerware"
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"time"
)

const cacheKey string = "productlistinfo"

// ProductServiceImpl implements the last service interface defined in the IDL.
type ProductServiceImpl struct{}

func BadAddProductResponse(s string) *product.AddProductResponse {
	return &product.AddProductResponse{StatusCode: -1, StatusMsg: s}
}
func GoodAddProductResponse(s string, ID int64) *product.AddProductResponse {
	return &product.AddProductResponse{StatusCode: 200, StatusMsg: s, ProductId: ID}
}
func BadGetProductInfoResponse(s string) *product.GetProductInfoResponse {
	return &product.GetProductInfoResponse{StatusCode: -1, StatusMsg: s}
}
func GoodGetProductInfoResponse(s string, pro *product.Product) *product.GetProductInfoResponse {
	return &product.GetProductInfoResponse{StatusCode: 200, StatusMsg: s, Product: pro}
}
func BadGetProductListInfoResponse(s string) *product.GetProductListInfoResponse {
	return &product.GetProductListInfoResponse{StatusCode: -1, StatusMsg: s}
}
func GoodGetProductListInfoResponse(s string, prolist []*product.Product) *product.GetProductListInfoResponse {
	return &product.GetProductListInfoResponse{StatusCode: 200, StatusMsg: s, Productlist: prolist}
}
func BadDelProductResponse(s string) *product.DelProductResponse {
	return &product.DelProductResponse{StatusCode: -1, StatusMsg: s}
}
func GoodDelProductResponse(s string, flag bool) *product.DelProductResponse {
	return &product.DelProductResponse{StatusCode: 200, StatusMsg: s, Succed: flag}
}
func BadUpdatePriceResponse(s string) *product.UpdatePriceResponse {
	return &product.UpdatePriceResponse{StatusCode: -1, StatusMsg: s}
}
func GoodUpdatePriceResponse(s string, flag bool) *product.UpdatePriceResponse {
	return &product.UpdatePriceResponse{StatusCode: 200, StatusMsg: s, Succed: flag}
}
func BadUpdateStockResponse(s string) *product.UpdateStockResponse {
	return &product.UpdateStockResponse{StatusCode: -1, StatusMsg: s}
}
func GoodUpdateStockResponse(s string, flag bool) *product.UpdateStockResponse {
	return &product.UpdateStockResponse{StatusCode: 200, StatusMsg: s, Succed: flag}
}
func BadGetProductListInfoByUserResponse(s string) *product.GetProductListInfoByUserResponse {
	return &product.GetProductListInfoByUserResponse{StatusCode: -1, StatusMsg: s}
}
func GoodGetProductListInfoByUserResponse(s string, prolist []*product.Product) *product.GetProductListInfoByUserResponse {
	return &product.GetProductListInfoByUserResponse{StatusCode: 200, StatusMsg: s, Productlist: prolist}
}
func BadUpdateStockAndSalesResponse(s string) *product.UpdateStockAndSalesResponse {
	return &product.UpdateStockAndSalesResponse{StatusCode: -1, StatusMsg: s}
}
func GoodUpdateStockAndSalesResponse(s string, flag bool) *product.UpdateStockAndSalesResponse {
	return &product.UpdateStockAndSalesResponse{StatusCode: 200, StatusMsg: s, Succed: flag}
}

// AddProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) AddProduct(ctx context.Context, req *product.AddProductRequest) (resp *product.AddProductResponse, err error) {
	mc, err := middlerware.ParseToken(req.Token)
	if err != nil || mc == nil {
		logger.Error("token解析失败：", zap.Error(err))
		return BadAddProductResponse("token解析失败"), err
	}
	usr, err := db.GetUserByID(ctx, mc.UserId)
	if err != nil {
		logger.Error("添加失败，服务器内部错误：", zap.Error(err))
		return BadAddProductResponse("添加失败，服务器内部错误"), err
	}
	pro := &db.Product{
		ProductName:  req.Name,
		Price:        req.Price,
		Stock:        req.Stock,
		ProductImage: req.Productimage,
		UserID:       int64(usr.ID),
	}
	err = db.CreateProduct(ctx, pro)
	if err != nil {
		logger.Error("添加商品失败：", zap.Error(err))
		return BadAddProductResponse("添加商品失败"), err
	}
	err = rs.DelKey(ctx, cacheKey) // 删除缓存
	if err != nil {

		logger.Error("删除缓存失败：", zap.Error(err))
	}
	return GoodAddProductResponse("添加商品成功", int64(pro.ID)), nil
}

// GetProductInfo implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) GetProductInfo(ctx context.Context, req *product.GetProductInfoRequest) (resp *product.GetProductInfoResponse, err error) {
	pro, err := db.GetProductByID(ctx, req.Productid)
	if err != nil {
		logger.Error("获取商品信息失败，服务器内部错误：", zap.Error(err))
		return BadGetProductInfoResponse("获取商品信息失败，服务器内部错误"), err
	}
	if pro == nil {
		return BadGetProductInfoResponse("商品不存在"), nil
	}
	p := &product.Product{
		Name:         pro.ProductName,
		Id:           int64(pro.ID),
		Price:        pro.Price,
		Stock:        pro.Stock,
		Productimage: pro.ProductImage,
	}
	return GoodGetProductInfoResponse("获取商品信息成功", p), nil
}

// DelProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) DelProduct(ctx context.Context, req *product.DelProductRequest) (resp *product.DelProductResponse, err error) {
	pro, err := db.GetProductByID(ctx, req.Productid)
	if err != nil {
		logger.Error("删除商品失败，服务器内部错误：", zap.Error(err))
		return BadDelProductResponse("删除商品失败，服务器内部错误"), err
	}
	if pro == nil {
		return BadDelProductResponse("商品不存在"), nil
	}
	err = db.DeleteProduct(ctx, pro)
	if err != nil {
		logger.Error("删除商品失败：", zap.Error(err))
		return BadDelProductResponse("删除商品失败"), err
	}
	kafkaProducer, err := product2.NewKafkaProducer([]string{kafkaAddr}) //初始化kafka生产者
	if err != nil {
		logger.Error("kafka生产者创建失败：", zap.Error(err))
		return BadDelProductResponse("Kafka生产者创建失败"), err
	}
	err = kafkaProducer.SendDeleteProductImageEvent(pro.ProductName)
	if err != nil {
		logger.Error("删除商品图片消息创建成功，但更新消息发送失败：", zap.Error(err))
		return BadDelProductResponse("删除商品图片消息创建成功，但更新消息发送失败"), err
	}
	err = rs.DelKey(ctx, cacheKey) // 删除缓存
	if err != nil {
		logger.Error("删除缓存失败：", zap.Error(err))
	}
	return GoodDelProductResponse("删除商品成功", true), nil
}

// UpdateStock implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) UpdatePrice(ctx context.Context, req *product.UpdatePriceRequest) (resp *product.UpdatePriceResponse, err error) {
	lockKey := rs.RedisLockKeyPrefix + string(req.ProductId) + "price"
	locked, err := rs.AcquireLock(ctx, lockKey) //获取分布式锁
	if err != nil {
		logger.Error("获取Redis锁失败：", zap.Error(err)) //锁已存在
		return BadUpdatePriceResponse("获取价格锁失败，稍后重试"), err
	}
	if !locked {
		return BadUpdatePriceResponse("价格操作正在进行中，请稍后重试"), nil
	}
	defer rs.ReleaseLock(ctx, lockKey) // 确保操作完成后释放锁
	pro, err := db.GetProductByID(ctx, req.ProductId)
	if err != nil {
		logger.Error("修改商品价格失败，服务器内部错误：", zap.Error(err))
		return BadUpdatePriceResponse("修改商品价格失败，服务器内部错误"), err
	}
	if pro == nil {
		return BadUpdatePriceResponse("商品不存在"), nil
	}
	err = db.UpdatePrice(ctx, pro, req.Newprice_)
	if err != nil {
		logger.Error("修改商品价格失败：", zap.Error(err))
		return BadUpdatePriceResponse("修改商品价格失败"), err
	}
	err = rs.DelKey(ctx, cacheKey) // 删除缓存
	if err != nil {
		logger.Error("删除缓存失败：", zap.Error(err))
	}
	return GoodUpdatePriceResponse("修改商品价格成功", true), nil
}

// UpdateStock implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) UpdateStock(ctx context.Context, req *product.UpdateStockRequest) (resp *product.UpdateStockResponse, err error) {
	lockKey := rs.RedisLockKeyPrefix + string(req.ProductId) + "stock"
	locked, err := rs.AcquireLock(ctx, lockKey) //获取分布式锁
	if err != nil {
		logger.Error("获取Redis锁失败：", zap.Error(err)) //锁已存在
		return BadUpdateStockResponse("获取库存锁失败，稍后重试"), err
	}
	if !locked {
		return BadUpdateStockResponse("库存操作正在进行中，请稍后重试"), nil
	}
	defer rs.ReleaseLock(ctx, lockKey) // 确保操作完成后释放锁
	pro, err := db.GetProductByID(ctx, req.ProductId)
	if err != nil {
		logger.Error("修改商品库存失败，服务器内部错误：", zap.Error(err))
		return BadUpdateStockResponse("修改商品库存失败，服务器内部错误"), err
	}
	if pro == nil {
		return BadUpdateStockResponse("商品不存在"), nil
	}
	err = db.UpdateStock(ctx, pro, req.Addstock)
	if err != nil {
		logger.Error("修改商品价格失败：", zap.Error(err))
		return BadUpdateStockResponse("修改商品库存失败"), err
	}
	cacheKey := "productlistinfo"
	err = rs.DelKey(ctx, cacheKey) // 删除缓存
	if err != nil {
		logger.Error("删除缓存失败：", zap.Error(err))
	}
	return GoodUpdateStockResponse("修改商品库存成功", true), nil
}

// GetProductListInfo implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) GetProductListInfo(ctx context.Context) (resp *product.GetProductListInfoResponse, err error) {
	plist, err := rs.GetProductListInfo(ctx, cacheKey)
	if err == nil && plist != nil {
		return GoodGetProductListInfoResponse("获取商品列表信息成功", plist), nil
	}
	prolist, err := db.GetProductListInfo(ctx)
	if err != nil {
		logger.Error("获取商品列表信息失败：", zap.Error(err))
		return BadGetProductListInfoResponse("获取商品列表信息失败"), err
	}
	var productlist []*product.Product
	for _, pro := range prolist {
		var p product.Product
		p.Id = int64(pro.ID)
		p.Name = pro.ProductName
		p.Price = pro.Price
		p.Stock = pro.Stock
		p.Productimage = pro.ProductImage
		p.Sales = pro.Sales
		p.UserId = pro.UserID
		productlist = append(productlist, &p)
	}
	cacheddatabytes, err := json.Marshal(productlist)
	if err != nil {
		logger.Error("商品列表数据序列化失败：", zap.Error(err))
	}
	err = rs.SetKey(ctx, cacheKey, cacheddatabytes, 10*time.Minute)
	if err != nil {
		logger.Error("缓存设置失败：", zap.Error(err))
	}
	return GoodGetProductListInfoResponse("获取商品列表信息成功", productlist), nil
}

// GetProductListInfoByUser implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) GetProductListInfoByUser(ctx context.Context, req *product.GetProductListInfoByUserRequest) (resp *product.GetProductListInfoByUserResponse, err error) {
	mc, err := middlerware.ParseToken(req.Token)
	if err != nil || mc == nil {
		logger.Error("token解析失败：", zap.Error(err))
		return BadGetProductListInfoByUserResponse("token解析失败"), err
	}
	usr, err := db.GetUserByID(ctx, mc.UserId)
	prolist, err := db.GetProductListInfoByUser(ctx, int64(usr.ID))
	if err != nil {
		logger.Error("获取商品列表信息失败：", zap.Error(err))
		return BadGetProductListInfoByUserResponse("获取商品列表信息失败"), err
	}
	var productlist []*product.Product
	for _, pro := range prolist {
		var p product.Product
		p.Id = int64(pro.ID)
		p.Name = pro.ProductName
		p.Price = pro.Price
		p.Stock = pro.Stock
		p.Productimage = pro.ProductImage
		p.Sales = pro.Sales
		p.UserId = pro.UserID
		productlist = append(productlist, &p)
	}
	return GoodGetProductListInfoByUserResponse("获取商品列表信息成功", productlist), nil
}

// UpdateStockAndSales implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) UpdateStockAndSales(ctx context.Context, req *product.UpdateStockAndSalesRequest) (resp *product.UpdateStockAndSalesResponse, err error) {
	lockKey := rs.RedisLockKeyPrefix + string(req.ProductId) + "stock"
	locked, err := rs.AcquireLock(ctx, lockKey) //获取分布式锁
	if err != nil {
		logger.Error("获取Redis锁失败：", zap.Error(err)) //锁已存在
		return BadUpdateStockAndSalesResponse("获取库存锁失败，稍后重试"), err
	}
	if !locked {
		return BadUpdateStockAndSalesResponse("库存操作正在进行中，请稍后重试"), nil
	}
	defer rs.ReleaseLock(ctx, lockKey) // 确保操作完成后释放锁
	pro, err := db.GetProductByID(ctx, req.ProductId)
	if err != nil {
		logger.Error("修改商品库存销量失败，服务器内部错误：", zap.Error(err))
		return BadUpdateStockAndSalesResponse("修改商品库存销量失败，服务器内部错误"), err
	}
	if pro == nil {
		return BadUpdateStockAndSalesResponse("商品不存在"), err
	}
	err = db.UpdateStockAndSales(ctx, pro, req.Number)
	if err != nil {
		logger.Error("修改商品库存和销量失败：", zap.Error(err))
		return BadUpdateStockAndSalesResponse("修改商品库存和销量失败"), err
	}
	cacheKey := "productlistinfo"
	err = rs.DelKey(ctx, cacheKey) // 删除缓存
	if err != nil {
		logger.Error("删除缓存失败：", zap.Error(err))
	}
	return GoodUpdateStockAndSalesResponse("修改商品库存和销量成功", true), nil
}
