namespace go user

struct User{
1:i64 id
2:string name
3:i64 balance
4:i64 cost
5:string address
6:string avatar
}
struct UserLoginRequest {
    1: string username;
    2: string password;
    3: string captcha;
}
struct UserLoginResponse {
    1: i32    status_code
    2: string status_msg
    3: i64 user_id
    4: string token
}
struct UserRegisterRequest {
    1: string username;
    2: string password;
    3: string address;
    4: string captcha;
    5: string avatar;
}
struct UserRegisterResponse {
    1: i32    status_code
    2: string status_msg
    3: i64 user_id
    4: string token
}
struct GetUserInfoRequest {
    1: string token
}
struct GetUserInfoResponse {
    1: i32    status_code
    2: string status_msg
    3: User user
}
struct UpdateNameRequest{
1:string token
2:string newname
}
struct UpdateNameResponse{
1: i32    status_code
2: string status_msg
3: bool succed
}
struct UpdatePasswordRequest{
1:string token
2:string oldpassword
3:string newpassword
}
struct UpdatePasswordResponse{
1: i32    status_code
2: string status_msg
3: bool succed
}
struct UpdateCostRequest{
1:string token
2:i64 addcost
}
struct UpdateCostResponse{
1: i32    status_code
2: string status_msg
3: bool succed
}
struct UpdateBalanceRequest{
1:string token
2:i64 addbalance
}
struct UpdateBalanceResponse{
1: i32    status_code
2: string status_msg
3: bool succed
}
struct UpdateBalanceAndCostRequest{
1:string token
2:i64 number
}
struct UpdateBalanceAndCostResponse{
1: i32    status_code
2: string status_msg
3: bool succed
}
struct UpdateAddressRequest{
1:string token
2:string address
}
struct UpdateAddressResponse{
1:i32 status_code
2:string status_msg
3:bool succed
}
service UserService {
    UserLoginResponse UserLogin(1:UserLoginRequest req)
    UserRegisterResponse UserRegiter(1:UserRegisterRequest req)
    GetUserInfoResponse GetUserInfo(1:GetUserInfoRequest req)
    UpdateNameResponse UpdateName(1:UpdateNameRequest req)
    UpdatePasswordResponse UpdatePassword(1:UpdatePasswordRequest req)
    UpdateCostResponse UpdateCost(1:UpdateCostRequest req)
    UpdateBalanceResponse UpdateBalance(1:UpdateBalanceRequest req)
    UpdateBalanceAndCostResponse UpdateBalanceAndCost(1:UpdateBalanceAndCostRequest req)
    UpdateAddressResponse UpdateAddress(1:UpdateAddressRequest req)
}
