namespace cpp OpenStars.Common.TMediaStorageService
namespace go OpenStars.Common.TMediaStorageService
namespace java OpenStars.Common.TMediaStorageService

typedef i64 TKey

enum TErrorCode{
    EGood = 0,
    ENotFound = -1,
    EUnknown = -2 ,
    EDataExisted = -3
}

struct TMediaItem{
    1: TKey uid
    2: i64 mediaType, // 1= image,2= video,3= audio,4 =gif,
    3: string name,
    4: binary rawsdata,
    5: string mediaurls,
}

typedef TMediaItem TData


struct TDataResult{
    1: TErrorCode errorCode,
    2: optional TMediaItem data
    
}

service TDataServiceR{
    TDataResult getData(1: TKey key), 
}

service TDataService{
    TDataResult getData(1: TKey key), 
    TErrorCode putData(1: TKey key, 2: TMediaItem data)
    TErrorCode removeData(1:TKey key)
}

service TMediaStorageService extends TDataService{
    
}


