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
    2: i64 objectId // chu the phat ra hanh dong
    3: i64 actionId, // hanh dong
    4: optional i64 targetId, // doi tuong chiu tac dong cua chu the
    5: optional list<i64> extendId,
    6: optional map<string,string> extendmapdata,
    7: optional string extend,
    8: optional bool seen,
    9: i64 timestamps,
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


