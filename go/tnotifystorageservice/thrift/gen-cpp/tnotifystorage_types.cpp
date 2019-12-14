/**
 * Autogenerated by Thrift Compiler (0.11.0)
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
#include "tnotifystorage_types.h"

#include <algorithm>
#include <ostream>

#include <thrift/TToString.h>

namespace OpenStars { namespace Common { namespace TNotifyStorageService {

int _kTErrorCodeValues[] = {
  TErrorCode::EGood,
  TErrorCode::ENotFound,
  TErrorCode::EUnknown,
  TErrorCode::EDataExisted
};
const char* _kTErrorCodeNames[] = {
  "EGood",
  "ENotFound",
  "EUnknown",
  "EDataExisted"
};
const std::map<int, const char*> _TErrorCode_VALUES_TO_NAMES(::apache::thrift::TEnumIterator(4, _kTErrorCodeValues, _kTErrorCodeNames), ::apache::thrift::TEnumIterator(-1, NULL, NULL));

std::ostream& operator<<(std::ostream& out, const TErrorCode::type& val) {
  std::map<int, const char*>::const_iterator it = _TErrorCode_VALUES_TO_NAMES.find(val);
  if (it != _TErrorCode_VALUES_TO_NAMES.end()) {
    out << it->second;
  } else {
    out << static_cast<int>(val);
  }
  return out;
}


TNotifyItem::~TNotifyItem() throw() {
}


void TNotifyItem::__set_key(const int64_t val) {
  this->key = val;
}

void TNotifyItem::__set_subjectId(const int64_t val) {
  this->subjectId = val;
}

void TNotifyItem::__set_actionId(const int64_t val) {
  this->actionId = val;
}

void TNotifyItem::__set_objectId(const int64_t val) {
  this->objectId = val;
}

void TNotifyItem::__set_subjectType(const int64_t val) {
  this->subjectType = val;
}

void TNotifyItem::__set_objectType(const int64_t val) {
  this->objectType = val;
}

void TNotifyItem::__set_extendSubjectId(const std::vector<int64_t> & val) {
  this->extendSubjectId = val;
__isset.extendSubjectId = true;
}

void TNotifyItem::__set_extendObjectId(const std::vector<int64_t> & val) {
  this->extendObjectId = val;
__isset.extendObjectId = true;
}

void TNotifyItem::__set_message(const std::string& val) {
  this->message = val;
__isset.message = true;
}

void TNotifyItem::__set_extend(const std::string& val) {
  this->extend = val;
__isset.extend = true;
}

void TNotifyItem::__set_seen(const bool val) {
  this->seen = val;
__isset.seen = true;
}

void TNotifyItem::__set_timestamps(const int64_t val) {
  this->timestamps = val;
}
std::ostream& operator<<(std::ostream& out, const TNotifyItem& obj)
{
  obj.printTo(out);
  return out;
}


uint32_t TNotifyItem::read(::apache::thrift::protocol::TProtocol* iprot) {

  ::apache::thrift::protocol::TInputRecursionTracker tracker(*iprot);
  uint32_t xfer = 0;
  std::string fname;
  ::apache::thrift::protocol::TType ftype;
  int16_t fid;

  xfer += iprot->readStructBegin(fname);

  using ::apache::thrift::protocol::TProtocolException;


  while (true)
  {
    xfer += iprot->readFieldBegin(fname, ftype, fid);
    if (ftype == ::apache::thrift::protocol::T_STOP) {
      break;
    }
    switch (fid)
    {
      case 1:
        if (ftype == ::apache::thrift::protocol::T_I64) {
          xfer += iprot->readI64(this->key);
          this->__isset.key = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      case 2:
        if (ftype == ::apache::thrift::protocol::T_I64) {
          xfer += iprot->readI64(this->subjectId);
          this->__isset.subjectId = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      case 3:
        if (ftype == ::apache::thrift::protocol::T_I64) {
          xfer += iprot->readI64(this->actionId);
          this->__isset.actionId = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      case 4:
        if (ftype == ::apache::thrift::protocol::T_I64) {
          xfer += iprot->readI64(this->objectId);
          this->__isset.objectId = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      case 5:
        if (ftype == ::apache::thrift::protocol::T_I64) {
          xfer += iprot->readI64(this->subjectType);
          this->__isset.subjectType = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      case 6:
        if (ftype == ::apache::thrift::protocol::T_I64) {
          xfer += iprot->readI64(this->objectType);
          this->__isset.objectType = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      case 7:
        if (ftype == ::apache::thrift::protocol::T_LIST) {
          {
            this->extendSubjectId.clear();
            uint32_t _size0;
            ::apache::thrift::protocol::TType _etype3;
            xfer += iprot->readListBegin(_etype3, _size0);
            this->extendSubjectId.resize(_size0);
            uint32_t _i4;
            for (_i4 = 0; _i4 < _size0; ++_i4)
            {
              xfer += iprot->readI64(this->extendSubjectId[_i4]);
            }
            xfer += iprot->readListEnd();
          }
          this->__isset.extendSubjectId = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      case 8:
        if (ftype == ::apache::thrift::protocol::T_LIST) {
          {
            this->extendObjectId.clear();
            uint32_t _size5;
            ::apache::thrift::protocol::TType _etype8;
            xfer += iprot->readListBegin(_etype8, _size5);
            this->extendObjectId.resize(_size5);
            uint32_t _i9;
            for (_i9 = 0; _i9 < _size5; ++_i9)
            {
              xfer += iprot->readI64(this->extendObjectId[_i9]);
            }
            xfer += iprot->readListEnd();
          }
          this->__isset.extendObjectId = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      case 9:
        if (ftype == ::apache::thrift::protocol::T_STRING) {
          xfer += iprot->readString(this->message);
          this->__isset.message = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      case 10:
        if (ftype == ::apache::thrift::protocol::T_STRING) {
          xfer += iprot->readString(this->extend);
          this->__isset.extend = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      case 11:
        if (ftype == ::apache::thrift::protocol::T_BOOL) {
          xfer += iprot->readBool(this->seen);
          this->__isset.seen = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      case 12:
        if (ftype == ::apache::thrift::protocol::T_I64) {
          xfer += iprot->readI64(this->timestamps);
          this->__isset.timestamps = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      default:
        xfer += iprot->skip(ftype);
        break;
    }
    xfer += iprot->readFieldEnd();
  }

  xfer += iprot->readStructEnd();

  return xfer;
}

uint32_t TNotifyItem::write(::apache::thrift::protocol::TProtocol* oprot) const {
  uint32_t xfer = 0;
  ::apache::thrift::protocol::TOutputRecursionTracker tracker(*oprot);
  xfer += oprot->writeStructBegin("TNotifyItem");

  xfer += oprot->writeFieldBegin("key", ::apache::thrift::protocol::T_I64, 1);
  xfer += oprot->writeI64(this->key);
  xfer += oprot->writeFieldEnd();

  xfer += oprot->writeFieldBegin("subjectId", ::apache::thrift::protocol::T_I64, 2);
  xfer += oprot->writeI64(this->subjectId);
  xfer += oprot->writeFieldEnd();

  xfer += oprot->writeFieldBegin("actionId", ::apache::thrift::protocol::T_I64, 3);
  xfer += oprot->writeI64(this->actionId);
  xfer += oprot->writeFieldEnd();

  xfer += oprot->writeFieldBegin("objectId", ::apache::thrift::protocol::T_I64, 4);
  xfer += oprot->writeI64(this->objectId);
  xfer += oprot->writeFieldEnd();

  xfer += oprot->writeFieldBegin("subjectType", ::apache::thrift::protocol::T_I64, 5);
  xfer += oprot->writeI64(this->subjectType);
  xfer += oprot->writeFieldEnd();

  xfer += oprot->writeFieldBegin("objectType", ::apache::thrift::protocol::T_I64, 6);
  xfer += oprot->writeI64(this->objectType);
  xfer += oprot->writeFieldEnd();

  if (this->__isset.extendSubjectId) {
    xfer += oprot->writeFieldBegin("extendSubjectId", ::apache::thrift::protocol::T_LIST, 7);
    {
      xfer += oprot->writeListBegin(::apache::thrift::protocol::T_I64, static_cast<uint32_t>(this->extendSubjectId.size()));
      std::vector<int64_t> ::const_iterator _iter10;
      for (_iter10 = this->extendSubjectId.begin(); _iter10 != this->extendSubjectId.end(); ++_iter10)
      {
        xfer += oprot->writeI64((*_iter10));
      }
      xfer += oprot->writeListEnd();
    }
    xfer += oprot->writeFieldEnd();
  }
  if (this->__isset.extendObjectId) {
    xfer += oprot->writeFieldBegin("extendObjectId", ::apache::thrift::protocol::T_LIST, 8);
    {
      xfer += oprot->writeListBegin(::apache::thrift::protocol::T_I64, static_cast<uint32_t>(this->extendObjectId.size()));
      std::vector<int64_t> ::const_iterator _iter11;
      for (_iter11 = this->extendObjectId.begin(); _iter11 != this->extendObjectId.end(); ++_iter11)
      {
        xfer += oprot->writeI64((*_iter11));
      }
      xfer += oprot->writeListEnd();
    }
    xfer += oprot->writeFieldEnd();
  }
  if (this->__isset.message) {
    xfer += oprot->writeFieldBegin("message", ::apache::thrift::protocol::T_STRING, 9);
    xfer += oprot->writeString(this->message);
    xfer += oprot->writeFieldEnd();
  }
  if (this->__isset.extend) {
    xfer += oprot->writeFieldBegin("extend", ::apache::thrift::protocol::T_STRING, 10);
    xfer += oprot->writeString(this->extend);
    xfer += oprot->writeFieldEnd();
  }
  if (this->__isset.seen) {
    xfer += oprot->writeFieldBegin("seen", ::apache::thrift::protocol::T_BOOL, 11);
    xfer += oprot->writeBool(this->seen);
    xfer += oprot->writeFieldEnd();
  }
  xfer += oprot->writeFieldBegin("timestamps", ::apache::thrift::protocol::T_I64, 12);
  xfer += oprot->writeI64(this->timestamps);
  xfer += oprot->writeFieldEnd();

  xfer += oprot->writeFieldStop();
  xfer += oprot->writeStructEnd();
  return xfer;
}

void swap(TNotifyItem &a, TNotifyItem &b) {
  using ::std::swap;
  swap(a.key, b.key);
  swap(a.subjectId, b.subjectId);
  swap(a.actionId, b.actionId);
  swap(a.objectId, b.objectId);
  swap(a.subjectType, b.subjectType);
  swap(a.objectType, b.objectType);
  swap(a.extendSubjectId, b.extendSubjectId);
  swap(a.extendObjectId, b.extendObjectId);
  swap(a.message, b.message);
  swap(a.extend, b.extend);
  swap(a.seen, b.seen);
  swap(a.timestamps, b.timestamps);
  swap(a.__isset, b.__isset);
}

TNotifyItem::TNotifyItem(const TNotifyItem& other12) {
  key = other12.key;
  subjectId = other12.subjectId;
  actionId = other12.actionId;
  objectId = other12.objectId;
  subjectType = other12.subjectType;
  objectType = other12.objectType;
  extendSubjectId = other12.extendSubjectId;
  extendObjectId = other12.extendObjectId;
  message = other12.message;
  extend = other12.extend;
  seen = other12.seen;
  timestamps = other12.timestamps;
  __isset = other12.__isset;
}
TNotifyItem& TNotifyItem::operator=(const TNotifyItem& other13) {
  key = other13.key;
  subjectId = other13.subjectId;
  actionId = other13.actionId;
  objectId = other13.objectId;
  subjectType = other13.subjectType;
  objectType = other13.objectType;
  extendSubjectId = other13.extendSubjectId;
  extendObjectId = other13.extendObjectId;
  message = other13.message;
  extend = other13.extend;
  seen = other13.seen;
  timestamps = other13.timestamps;
  __isset = other13.__isset;
  return *this;
}
void TNotifyItem::printTo(std::ostream& out) const {
  using ::apache::thrift::to_string;
  out << "TNotifyItem(";
  out << "key=" << to_string(key);
  out << ", " << "subjectId=" << to_string(subjectId);
  out << ", " << "actionId=" << to_string(actionId);
  out << ", " << "objectId=" << to_string(objectId);
  out << ", " << "subjectType=" << to_string(subjectType);
  out << ", " << "objectType=" << to_string(objectType);
  out << ", " << "extendSubjectId="; (__isset.extendSubjectId ? (out << to_string(extendSubjectId)) : (out << "<null>"));
  out << ", " << "extendObjectId="; (__isset.extendObjectId ? (out << to_string(extendObjectId)) : (out << "<null>"));
  out << ", " << "message="; (__isset.message ? (out << to_string(message)) : (out << "<null>"));
  out << ", " << "extend="; (__isset.extend ? (out << to_string(extend)) : (out << "<null>"));
  out << ", " << "seen="; (__isset.seen ? (out << to_string(seen)) : (out << "<null>"));
  out << ", " << "timestamps=" << to_string(timestamps);
  out << ")";
}


TDataResult::~TDataResult() throw() {
}


void TDataResult::__set_errorCode(const TErrorCode::type val) {
  this->errorCode = val;
}

void TDataResult::__set_data(const TNotifyItem& val) {
  this->data = val;
__isset.data = true;
}
std::ostream& operator<<(std::ostream& out, const TDataResult& obj)
{
  obj.printTo(out);
  return out;
}


uint32_t TDataResult::read(::apache::thrift::protocol::TProtocol* iprot) {

  ::apache::thrift::protocol::TInputRecursionTracker tracker(*iprot);
  uint32_t xfer = 0;
  std::string fname;
  ::apache::thrift::protocol::TType ftype;
  int16_t fid;

  xfer += iprot->readStructBegin(fname);

  using ::apache::thrift::protocol::TProtocolException;


  while (true)
  {
    xfer += iprot->readFieldBegin(fname, ftype, fid);
    if (ftype == ::apache::thrift::protocol::T_STOP) {
      break;
    }
    switch (fid)
    {
      case 1:
        if (ftype == ::apache::thrift::protocol::T_I32) {
          int32_t ecast14;
          xfer += iprot->readI32(ecast14);
          this->errorCode = (TErrorCode::type)ecast14;
          this->__isset.errorCode = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      case 2:
        if (ftype == ::apache::thrift::protocol::T_STRUCT) {
          xfer += this->data.read(iprot);
          this->__isset.data = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      default:
        xfer += iprot->skip(ftype);
        break;
    }
    xfer += iprot->readFieldEnd();
  }

  xfer += iprot->readStructEnd();

  return xfer;
}

uint32_t TDataResult::write(::apache::thrift::protocol::TProtocol* oprot) const {
  uint32_t xfer = 0;
  ::apache::thrift::protocol::TOutputRecursionTracker tracker(*oprot);
  xfer += oprot->writeStructBegin("TDataResult");

  xfer += oprot->writeFieldBegin("errorCode", ::apache::thrift::protocol::T_I32, 1);
  xfer += oprot->writeI32((int32_t)this->errorCode);
  xfer += oprot->writeFieldEnd();

  if (this->__isset.data) {
    xfer += oprot->writeFieldBegin("data", ::apache::thrift::protocol::T_STRUCT, 2);
    xfer += this->data.write(oprot);
    xfer += oprot->writeFieldEnd();
  }
  xfer += oprot->writeFieldStop();
  xfer += oprot->writeStructEnd();
  return xfer;
}

void swap(TDataResult &a, TDataResult &b) {
  using ::std::swap;
  swap(a.errorCode, b.errorCode);
  swap(a.data, b.data);
  swap(a.__isset, b.__isset);
}

TDataResult::TDataResult(const TDataResult& other15) {
  errorCode = other15.errorCode;
  data = other15.data;
  __isset = other15.__isset;
}
TDataResult& TDataResult::operator=(const TDataResult& other16) {
  errorCode = other16.errorCode;
  data = other16.data;
  __isset = other16.__isset;
  return *this;
}
void TDataResult::printTo(std::ostream& out) const {
  using ::apache::thrift::to_string;
  out << "TDataResult(";
  out << "errorCode=" << to_string(errorCode);
  out << ", " << "data="; (__isset.data ? (out << to_string(data)) : (out << "<null>"));
  out << ")";
}


TListDataResult::~TListDataResult() throw() {
}


void TListDataResult::__set_errorCode(const TErrorCode::type val) {
  this->errorCode = val;
}

void TListDataResult::__set_datass(const std::vector<TNotifyItem> & val) {
  this->datass = val;
__isset.datass = true;
}
std::ostream& operator<<(std::ostream& out, const TListDataResult& obj)
{
  obj.printTo(out);
  return out;
}


uint32_t TListDataResult::read(::apache::thrift::protocol::TProtocol* iprot) {

  ::apache::thrift::protocol::TInputRecursionTracker tracker(*iprot);
  uint32_t xfer = 0;
  std::string fname;
  ::apache::thrift::protocol::TType ftype;
  int16_t fid;

  xfer += iprot->readStructBegin(fname);

  using ::apache::thrift::protocol::TProtocolException;


  while (true)
  {
    xfer += iprot->readFieldBegin(fname, ftype, fid);
    if (ftype == ::apache::thrift::protocol::T_STOP) {
      break;
    }
    switch (fid)
    {
      case 1:
        if (ftype == ::apache::thrift::protocol::T_I32) {
          int32_t ecast17;
          xfer += iprot->readI32(ecast17);
          this->errorCode = (TErrorCode::type)ecast17;
          this->__isset.errorCode = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      case 2:
        if (ftype == ::apache::thrift::protocol::T_LIST) {
          {
            this->datass.clear();
            uint32_t _size18;
            ::apache::thrift::protocol::TType _etype21;
            xfer += iprot->readListBegin(_etype21, _size18);
            this->datass.resize(_size18);
            uint32_t _i22;
            for (_i22 = 0; _i22 < _size18; ++_i22)
            {
              xfer += this->datass[_i22].read(iprot);
            }
            xfer += iprot->readListEnd();
          }
          this->__isset.datass = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      default:
        xfer += iprot->skip(ftype);
        break;
    }
    xfer += iprot->readFieldEnd();
  }

  xfer += iprot->readStructEnd();

  return xfer;
}

uint32_t TListDataResult::write(::apache::thrift::protocol::TProtocol* oprot) const {
  uint32_t xfer = 0;
  ::apache::thrift::protocol::TOutputRecursionTracker tracker(*oprot);
  xfer += oprot->writeStructBegin("TListDataResult");

  xfer += oprot->writeFieldBegin("errorCode", ::apache::thrift::protocol::T_I32, 1);
  xfer += oprot->writeI32((int32_t)this->errorCode);
  xfer += oprot->writeFieldEnd();

  if (this->__isset.datass) {
    xfer += oprot->writeFieldBegin("datass", ::apache::thrift::protocol::T_LIST, 2);
    {
      xfer += oprot->writeListBegin(::apache::thrift::protocol::T_STRUCT, static_cast<uint32_t>(this->datass.size()));
      std::vector<TNotifyItem> ::const_iterator _iter23;
      for (_iter23 = this->datass.begin(); _iter23 != this->datass.end(); ++_iter23)
      {
        xfer += (*_iter23).write(oprot);
      }
      xfer += oprot->writeListEnd();
    }
    xfer += oprot->writeFieldEnd();
  }
  xfer += oprot->writeFieldStop();
  xfer += oprot->writeStructEnd();
  return xfer;
}

void swap(TListDataResult &a, TListDataResult &b) {
  using ::std::swap;
  swap(a.errorCode, b.errorCode);
  swap(a.datass, b.datass);
  swap(a.__isset, b.__isset);
}

TListDataResult::TListDataResult(const TListDataResult& other24) {
  errorCode = other24.errorCode;
  datass = other24.datass;
  __isset = other24.__isset;
}
TListDataResult& TListDataResult::operator=(const TListDataResult& other25) {
  errorCode = other25.errorCode;
  datass = other25.datass;
  __isset = other25.__isset;
  return *this;
}
void TListDataResult::printTo(std::ostream& out) const {
  using ::apache::thrift::to_string;
  out << "TListDataResult(";
  out << "errorCode=" << to_string(errorCode);
  out << ", " << "datass="; (__isset.datass ? (out << to_string(datass)) : (out << "<null>"));
  out << ")";
}

}}} // namespace
