namespace cpp OpenStars.Common.TPostStorageService
namespace go OpenStars.Common.TPostStorageService
namespace java OpenStars.Common.TPostStorageService

typedef i64 TKey

enum TErrorCode{
    EGood = 0,
    ENotFound = -1,
    EUnknown = -2 ,
    EDataExisted = -3
}

// which user can see your post
enum TPrivacy {
    Public = 0, // all user
    Friends = 1,
    FriendsExcept = 2,
    Onlyme = 3,
    SpecificFriends = 4,
}

struct ActionLink {
    1:optional string text,
    2:optional string href,
}

struct MediaItem {
    1:string name,
    2:i64 mediaType, // 1 = image ; 2 = video; 3 = gif; 
    3:string url,
}

struct TPostItem{
    1: TKey idpost, // id of post gen by kvcounter
    2: i64 uid, // id of user author of this post
    3: string content, // content of this post
    4:optional list<MediaItem> listMediaItems, // list mediaUrls such as photo,video,audio,gift
    5:optional string idbackground, // background of post like facebook
    6:optional string idfeeling, // feeling or activity of user like facebook
    7:i64 privacy,
    8:optional list<i64> friendsexcept , // list friends can not see this post
    9:optional list<i64> specificfriends, // list friends only see this post
    10:optional list<i64> tagusers, // list friend tag in this post
    11:optional string locationId, // check in location
    12:i64 timestamps;
    13:string pubkey, //pubkey of user author of this post
    14:optional string topubkey,
    15:optional i64 touid,
    16:optional i64 togroupid,
    17:optional list<ActionLink> actionLinks,
    18:optional list<string> mediaUrls,
}

typedef TPostItem TData


struct TDataResult{ 
    1: TErrorCode errorCode,
    2: optional TPostItem data
    
}

struct TListDataResult {
    1: TErrorCode errorCode,
    2: list<TPostItem> listDatas;
}

service TDataServiceR{
    TDataResult getData(1: TKey key),
}

service TDataService{
    TDataResult getData(1: TKey key), 
    TErrorCode putData(1: TKey key, 2: TPostItem data)
    TErrorCode removeData(1:TKey key),
    TListDataResult getListDatas(1:list<TKey> listkey),
}

service TPostStorageService extends TDataService{
    
}


