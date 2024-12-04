namespace go product
struct Product{//商品有id,商品名，有单价，库存
1:i64 id
2:string name
3:i64 price
4:i64 stock
5:string productimage
}
struct AddProductRequest{
1:string name
2:i64 stock
3:i64 price
}
struct AddProductResponse{
1: i32    status_code
2: string status_msg
3: i64 product_id
}
struct GetProductInfoRequest{
1:i64 productid
}
struct GetProductInfoResponse{
1: i32    status_code
2: string status_msg
3: Product product
}
struct GetProductListInfoResponse{
1: i32    status_code
2: string status_msg
3: list<Product> productlist
}
struct DelProductRequest{
1:i64 productid
}
struct DelProductResponse{
1: i32    status_code
2: string status_msg
3:bool succed
}
struct UpdatePriceRequest{
1:i64 product_id
2:i64 newprice
}
struct UpdatePriceResponse{
1: i32    status_code
2: string status_msg
3:bool succed
}
struct UpdateStockRequest{
1:i64 product_id
2:i64 addstock
}
struct UpdateStockResponse{
1: i32    status_code
2: string status_msg
3:bool succed
}
service ProductService{
AddProductResponse AddProduct(1:AddProductRequest req)
GetProductInfoResponse GetProductInfo(1:GetProductInfoRequest req)
GetProductListInfoResponse GetProductListInfo()
DelProductResponse DelProduct(1:DelProductRequest req)
UpdatePriceResponse UpdatePrice(1:UpdatePriceRequest req)
UpdateStockResponse UpdateStock(1:UpdateStockRequest req)
}