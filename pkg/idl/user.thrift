namespace go user
include "model.thrift"

//注册
struct RegisterRequest {
    1: required string username
    2: required string password
}

struct RegisterResponse {
    1: required model.BaseResp resp
}

//登录
struct LoginRequest {
    1: required string username
    2: required string password
}

struct LoginResponse {
    1: required model.BaseResp resp
    2: optional model.User user
}

service UserService {
    RegisterResponse Register(1: RegisterRequest req)
    LoginResponse Login(1: LoginRequest req)
}
