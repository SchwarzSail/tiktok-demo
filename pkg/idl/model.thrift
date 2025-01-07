namespace go model
# 定义所有服务之间要交流的业务实体

struct BaseResp {
    1: i64 code,
    2: string msg,
}


struct User {
    1: optional string id
    2: required string username
    3: required string avatar_url
    4: required string email
    5: required string phone
}

struct Video {
    1: optional string id
    2: optional string user_id
    3: required string video_url
    4: required string cover_url
    5: required string title
    6: required string description
    7: required i64 visit_count
    8: required i64 like_count
    9: required i64 comment_count
    10: required string created_at
    11: required string updated_at
    12: required string deleted_at
}

