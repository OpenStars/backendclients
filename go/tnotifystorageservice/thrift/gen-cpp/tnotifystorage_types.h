/**
 * Autogenerated by Thrift Compiler (0.11.0)
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
#ifndef tnotifystorage_TYPES_H
#define tnotifystorage_TYPES_H

#include <iosfwd>

#include <thrift/Thrift.h>
#include <thrift/TApplicationException.h>
#include <thrift/TBase.h>
#include <thrift/protocol/TProtocol.h>
#include <thrift/transport/TTransport.h>

#include <thrift/stdcxx.h>


namespace OpenStars { namespace Common { namespace TNotifyStorageService {

struct TErrorCode {
  enum type {
    EGood = 0,
    ENotFound = -1,
    EUnknown = -2,
    EDataExisted = -3
  };
};

extern const std::map<int, const char*> _TErrorCode_VALUES_TO_NAMES;

std::ostream& operator<<(std::ostream& out, const TErrorCode::type& val);

typedef int64_t TKey;

typedef class TNotifyItem TData;

class TNotifyItem;

class TDataResult;

class TListDataResult;

typedef struct _TNotifyItem__isset {
  _TNotifyItem__isset() : key(false), objectId(false), actionId(false), targetId(false), extendId(false), extendmapdata(false), extend(false), seen(false), timestamps(false) {}
  bool key :1;
  bool objectId :1;
  bool actionId :1;
  bool targetId :1;
  bool extendId :1;
  bool extendmapdata :1;
  bool extend :1;
  bool seen :1;
  bool timestamps :1;
} _TNotifyItem__isset;

class TNotifyItem : public virtual ::apache::thrift::TBase {
 public:

  TNotifyItem(const TNotifyItem&);
  TNotifyItem& operator=(const TNotifyItem&);
  TNotifyItem() : key(0), objectId(0), actionId(0), targetId(0), extend(), seen(0), timestamps(0) {
  }

  virtual ~TNotifyItem() throw();
  int64_t key;
  int64_t objectId;
  int64_t actionId;
  int64_t targetId;
  std::vector<int64_t>  extendId;
  std::map<std::string, std::string>  extendmapdata;
  std::string extend;
  bool seen;
  int64_t timestamps;

  _TNotifyItem__isset __isset;

  void __set_key(const int64_t val);

  void __set_objectId(const int64_t val);

  void __set_actionId(const int64_t val);

  void __set_targetId(const int64_t val);

  void __set_extendId(const std::vector<int64_t> & val);

  void __set_extendmapdata(const std::map<std::string, std::string> & val);

  void __set_extend(const std::string& val);

  void __set_seen(const bool val);

  void __set_timestamps(const int64_t val);

  bool operator == (const TNotifyItem & rhs) const
  {
    if (!(key == rhs.key))
      return false;
    if (!(objectId == rhs.objectId))
      return false;
    if (!(actionId == rhs.actionId))
      return false;
    if (__isset.targetId != rhs.__isset.targetId)
      return false;
    else if (__isset.targetId && !(targetId == rhs.targetId))
      return false;
    if (__isset.extendId != rhs.__isset.extendId)
      return false;
    else if (__isset.extendId && !(extendId == rhs.extendId))
      return false;
    if (__isset.extendmapdata != rhs.__isset.extendmapdata)
      return false;
    else if (__isset.extendmapdata && !(extendmapdata == rhs.extendmapdata))
      return false;
    if (__isset.extend != rhs.__isset.extend)
      return false;
    else if (__isset.extend && !(extend == rhs.extend))
      return false;
    if (__isset.seen != rhs.__isset.seen)
      return false;
    else if (__isset.seen && !(seen == rhs.seen))
      return false;
    if (!(timestamps == rhs.timestamps))
      return false;
    return true;
  }
  bool operator != (const TNotifyItem &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const TNotifyItem & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(TNotifyItem &a, TNotifyItem &b);

std::ostream& operator<<(std::ostream& out, const TNotifyItem& obj);

typedef struct _TDataResult__isset {
  _TDataResult__isset() : errorCode(false), data(false) {}
  bool errorCode :1;
  bool data :1;
} _TDataResult__isset;

class TDataResult : public virtual ::apache::thrift::TBase {
 public:

  TDataResult(const TDataResult&);
  TDataResult& operator=(const TDataResult&);
  TDataResult() : errorCode((TErrorCode::type)0) {
  }

  virtual ~TDataResult() throw();
  TErrorCode::type errorCode;
  TNotifyItem data;

  _TDataResult__isset __isset;

  void __set_errorCode(const TErrorCode::type val);

  void __set_data(const TNotifyItem& val);

  bool operator == (const TDataResult & rhs) const
  {
    if (!(errorCode == rhs.errorCode))
      return false;
    if (__isset.data != rhs.__isset.data)
      return false;
    else if (__isset.data && !(data == rhs.data))
      return false;
    return true;
  }
  bool operator != (const TDataResult &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const TDataResult & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(TDataResult &a, TDataResult &b);

std::ostream& operator<<(std::ostream& out, const TDataResult& obj);

typedef struct _TListDataResult__isset {
  _TListDataResult__isset() : errorCode(false), datass(false) {}
  bool errorCode :1;
  bool datass :1;
} _TListDataResult__isset;

class TListDataResult : public virtual ::apache::thrift::TBase {
 public:

  TListDataResult(const TListDataResult&);
  TListDataResult& operator=(const TListDataResult&);
  TListDataResult() : errorCode((TErrorCode::type)0) {
  }

  virtual ~TListDataResult() throw();
  TErrorCode::type errorCode;
  std::vector<TNotifyItem>  datass;

  _TListDataResult__isset __isset;

  void __set_errorCode(const TErrorCode::type val);

  void __set_datass(const std::vector<TNotifyItem> & val);

  bool operator == (const TListDataResult & rhs) const
  {
    if (!(errorCode == rhs.errorCode))
      return false;
    if (__isset.datass != rhs.__isset.datass)
      return false;
    else if (__isset.datass && !(datass == rhs.datass))
      return false;
    return true;
  }
  bool operator != (const TListDataResult &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const TListDataResult & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(TListDataResult &a, TListDataResult &b);

std::ostream& operator<<(std::ostream& out, const TListDataResult& obj);

}}} // namespace

#endif