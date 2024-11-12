namespace go order
struct order{
1:i64 order_id
2:i64 user_id
3:string product_name
4:i64 number
5:i64 cost
6:string create_time
}
struct AddOrderRequest{
1:order ol
2:string token
}
struct AddOrderResponse{
1: i32    status_code
2: string status_msg
3: i64 order_id
}
struct DelOrderRequest{
1:i64 order_id
2:string token
}
struct DelOrderResponse{
1: i32    status_code
2: string status_msg
3: bool succed
}
struct GetOrderListByUserIDRequest{
1:i64 user_id
2:string token
}
struct GetOrderListByUserIDResponse{
1: i32    status_code
2: string status_msg
3:list<order> orderlist
}
struct GetOrderListByProductNameRequest{
1:string product_name
2:string token
}
struct GetOrderListByProductNameResponse{
1: i32    status_code
2: string status_msg
3:list<order> orderlist
}
service OrderListService{
AddOrderResponse AddOrder(1:AddOrderRequest req)
DelOrderResponse DelOrder(1:DelOrderRequest req)
GetOrderListByUserIDResponse GetOrderListByUserID(1:GetOrderListByUserIDRequest req)
GetOrderListByProductNameResponse GetOrderListByProductNameID(1:GetOrderListByProductNameRequest req)
}