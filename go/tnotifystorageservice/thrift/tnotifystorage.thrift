namespace cpp OpenStars.Common.TNotifyStorageService
namespace go OpenStars.Common.TNotifyStorageService
namespace java OpenStars.Common.TNotifyStorageService

typedef i64 TKey

enum TErrorCode{
    EGood = 0,
    ENotFound = -1,
    EUnknown = -2 ,
    EDataExisted = -3
}

struct TNotifyItem{
    1: i64 key, // notifyitem id
    2: i64 subjectId // chu the phat ra hanh dong
    3: i64 actionId, // hanh dong
    4: i64  objectId, // doi tuong chiu tac dong cua chu the
    5: i64 subjectType, // uid, pageid or groupid
    6: i64 objectType, // comment or post
    7: optional list<i64> extendSubjectId,
    8: optional list<i64> extendObjectId,
    9: optional string message, // render message
    10: optional string extend,
    11: optional bool seen,
    12: i64 timestamps,
}

typedef TNotifyItem TData


struct TDataResult{
    1: TErrorCode errorCode,
    2: optional TNotifyItem data  
}

struct TListDataResult{
    1:TErrorCode errorCode,
    2:optional list<TNotifyItem> datass,
}

service TDataServiceR{
    TDataResult getData(1:i64 key), 
}

service TDataService{
    TDataResult getData(1:i64 key), 
    TErrorCode putData(1:i64 key, 2:TNotifyItem data),
    TListDataResult getListData(1:list<i64> lskeys)
    bool removeData(1:i64 key)
}

service TNotifyStorageService extends TDataService{
    
}


