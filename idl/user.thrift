namespace go user

struct User{
1:i64 id
2:string name
3:i64 balance
4:i64 cost
5:string address
6:string avatar
}
struct Message{
1:string content
2:string create_time
3:i64 user_id
}
struct FriendInfo{
1:string name
2:string avatar
3:i64 user_id
}
struct UserInfo{
1:string name
2:i64 id
3:string avatar
}
struct FriendApplication{
1:string name
2:string avatar
3:i64 userid
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
struct GetUserInfoByUserIDRequest{
    1:i64 user_id
}
struct GetUserInfoByUserIDResponse{
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
struct SendMessageRequest{
1:string content
2:string token
3:i64 touser_id
}
struct SendMessageResponse{
1:i32 status_code
2:string status_msg
3:bool succed
}
struct GetFriendListRequest{
1:string token
}
struct GetFriendListResponse{
1:i32 status_code
2:string status_msg
3:list<FriendInfo> friend
}
struct AddFriendRequest{
1:string token
2:i64 touser_id
}
struct AddFriendResponse{
1:i32 status_code
2:string status_msg
3:bool succed
}
struct DeleteFriendRequest{
1:string token
2:i64 touser_id
}
struct DeleteFriendResponse{
1:i32 status_code
2:string status_msg
3:bool succed
}
struct GetMessageListRequest{
1:string token
2:i64 touser_id
}
struct GetMessageListResponse{
1:i32 status_code
2:string status_msg
3:list<Message> message
}
struct GetUserListByContentRequest{
1:string content
}
struct GetUserListByContentResponse{
1:i32 status_code
2:string status_msg
3:list<UserInfo> userinfolist
}
struct SendFriendApplicationRequest{
1:string token
2:i64 touserid
}
struct SendFriendApplicationResponse{
1:i32 status_code
2:string status_msg
3:bool succed
}
struct GetFriendApplicationListRequest{
1:string token
}
struct GetFriendApplicationListResponse{
1:i32 status_code
2:string status_msg
3:list<FriendApplication> friendapplicaiton
}
struct RejectFriendApplicationRequest{
1:string token
2:i64 touserid
}
struct RejectFriendApplicationResponse{
1:i32 status_code
2:string status_msg
3:bool succed
}
//kitex -module Eshop idl/user.thrift
//kitex -module Eshop -service Eshop.item -use Eshop/kitex_gen ../../idl/user.thrift
service UserService {
    UserLoginResponse UserLogin(1:UserLoginRequest req)
    UserRegisterResponse UserRegiter(1:UserRegisterRequest req)
    GetUserInfoResponse GetUserInfo(1:GetUserInfoRequest req)
    GetUserInfoByUserIDResponse GetUserInfoByUserID(1:GetUserInfoByUserIDRequest req)
    UpdateNameResponse UpdateName(1:UpdateNameRequest req)
    UpdatePasswordResponse UpdatePassword(1:UpdatePasswordRequest req)
    UpdateCostResponse UpdateCost(1:UpdateCostRequest req)
    UpdateBalanceResponse UpdateBalance(1:UpdateBalanceRequest req)
    UpdateBalanceAndCostResponse UpdateBalanceAndCost(1:UpdateBalanceAndCostRequest req)
    UpdateAddressResponse UpdateAddress(1:UpdateAddressRequest req)
    GetFriendListResponse GetFriendList(1:GetFriendListRequest req)
    AddFriendResponse AddFriend(1:AddFriendRequest req)
    DeleteFriendResponse DeleteFriend(1:DeleteFriendRequest req)
    GetMessageListResponse GetMessageList(1: GetMessageListRequest req)
    SendMessageResponse SendMessage(1: SendMessageRequest req)
    GetUserListByContentResponse GetUserListByContent(1: GetUserListByContentRequest req)
    SendFriendApplicationResponse SendFriendApplication(1: SendFriendApplicationRequest req)
    GetFriendApplicationListResponse GetFriendApplicationList(1: GetFriendApplicationListRequest req)
    RejectFriendApplicationResponse RejectFriendApplication(1: RejectFriendApplicationRequest req)
}
