/**
 * Autogenerated by Thrift Compiler (0.11.0)
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
#ifndef tpoststorage_TYPES_H
#define tpoststorage_TYPES_H

#include <iosfwd>

#include <thrift/Thrift.h>
#include <thrift/TApplicationException.h>
#include <thrift/TBase.h>
#include <thrift/protocol/TProtocol.h>
#include <thrift/transport/TTransport.h>

#include <thrift/stdcxx.h>


namespace OpenStars { namespace Common { namespace TPostStorageService {

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

struct TPrivacy {
  enum type {
    Public = 0,
    Friends = 1,
    FriendsExcept = 2,
    Onlyme = 3,
    SpecificFriends = 4
  };
};

extern const std::map<int, const char*> _TPrivacy_VALUES_TO_NAMES;

std::ostream& operator<<(std::ostream& out, const TPrivacy::type& val);

typedef int64_t TKey;

typedef class TPostItem TData;

class ActionLink;

class MediaItem;

class TPostItem;

class TDataResult;

class TListDataResult;

typedef struct _ActionLink__isset {
  _ActionLink__isset() : text(false), href(false) {}
  bool text :1;
  bool href :1;
} _ActionLink__isset;

class ActionLink : public virtual ::apache::thrift::TBase {
 public:

  ActionLink(const ActionLink&);
  ActionLink& operator=(const ActionLink&);
  ActionLink() : text(), href() {
  }

  virtual ~ActionLink() throw();
  std::string text;
  std::string href;

  _ActionLink__isset __isset;

  void __set_text(const std::string& val);

  void __set_href(const std::string& val);

  bool operator == (const ActionLink & rhs) const
  {
    if (__isset.text != rhs.__isset.text)
      return false;
    else if (__isset.text && !(text == rhs.text))
      return false;
    if (__isset.href != rhs.__isset.href)
      return false;
    else if (__isset.href && !(href == rhs.href))
      return false;
    return true;
  }
  bool operator != (const ActionLink &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const ActionLink & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(ActionLink &a, ActionLink &b);

std::ostream& operator<<(std::ostream& out, const ActionLink& obj);

typedef struct _MediaItem__isset {
  _MediaItem__isset() : name(false), mediaType(false), url(false) {}
  bool name :1;
  bool mediaType :1;
  bool url :1;
} _MediaItem__isset;

class MediaItem : public virtual ::apache::thrift::TBase {
 public:

  MediaItem(const MediaItem&);
  MediaItem& operator=(const MediaItem&);
  MediaItem() : name(), mediaType(0), url() {
  }

  virtual ~MediaItem() throw();
  std::string name;
  int64_t mediaType;
  std::string url;

  _MediaItem__isset __isset;

  void __set_name(const std::string& val);

  void __set_mediaType(const int64_t val);

  void __set_url(const std::string& val);

  bool operator == (const MediaItem & rhs) const
  {
    if (!(name == rhs.name))
      return false;
    if (!(mediaType == rhs.mediaType))
      return false;
    if (!(url == rhs.url))
      return false;
    return true;
  }
  bool operator != (const MediaItem &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const MediaItem & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(MediaItem &a, MediaItem &b);

std::ostream& operator<<(std::ostream& out, const MediaItem& obj);

typedef struct _TPostItem__isset {
  _TPostItem__isset() : idpost(false), uid(false), content(false), listMediaItems(false), idbackground(false), idfeeling(false), privacy(false), friendsexcept(false), specificfriends(false), tagusers(false), locationId(false), timestamps(false), pubkey(false), touid(false), togroupid(false), actionLinks(false), poolid(false), pageid(false), extend(false), totalReaction(false), totalComment(false), totalShare(false), originPostID(false) {}
  bool idpost :1;
  bool uid :1;
  bool content :1;
  bool listMediaItems :1;
  bool idbackground :1;
  bool idfeeling :1;
  bool privacy :1;
  bool friendsexcept :1;
  bool specificfriends :1;
  bool tagusers :1;
  bool locationId :1;
  bool timestamps :1;
  bool pubkey :1;
  bool touid :1;
  bool togroupid :1;
  bool actionLinks :1;
  bool poolid :1;
  bool pageid :1;
  bool extend :1;
  bool totalReaction :1;
  bool totalComment :1;
  bool totalShare :1;
  bool originPostID :1;
} _TPostItem__isset;

class TPostItem : public virtual ::apache::thrift::TBase {
 public:

  TPostItem(const TPostItem&);
  TPostItem& operator=(const TPostItem&);
  TPostItem() : idpost(0), uid(0), content(), idbackground(), idfeeling(), privacy(0), locationId(), timestamps(0), pubkey(), touid(0), togroupid(0), poolid(0), pageid(0), extend(), totalReaction(0), totalComment(0), totalShare(0), originPostID(0) {
  }

  virtual ~TPostItem() throw();
  TKey idpost;
  int64_t uid;
  std::string content;
  std::vector<MediaItem>  listMediaItems;
  std::string idbackground;
  std::string idfeeling;
  int64_t privacy;
  std::vector<int64_t>  friendsexcept;
  std::vector<int64_t>  specificfriends;
  std::vector<int64_t>  tagusers;
  std::string locationId;
  int64_t timestamps;
  std::string pubkey;
  int64_t touid;
  int64_t togroupid;
  std::vector<ActionLink>  actionLinks;
  int64_t poolid;
  int64_t pageid;
  std::string extend;
  int64_t totalReaction;
  int64_t totalComment;
  int64_t totalShare;
  int64_t originPostID;

  _TPostItem__isset __isset;

  void __set_idpost(const TKey val);

  void __set_uid(const int64_t val);

  void __set_content(const std::string& val);

  void __set_listMediaItems(const std::vector<MediaItem> & val);

  void __set_idbackground(const std::string& val);

  void __set_idfeeling(const std::string& val);

  void __set_privacy(const int64_t val);

  void __set_friendsexcept(const std::vector<int64_t> & val);

  void __set_specificfriends(const std::vector<int64_t> & val);

  void __set_tagusers(const std::vector<int64_t> & val);

  void __set_locationId(const std::string& val);

  void __set_timestamps(const int64_t val);

  void __set_pubkey(const std::string& val);

  void __set_touid(const int64_t val);

  void __set_togroupid(const int64_t val);

  void __set_actionLinks(const std::vector<ActionLink> & val);

  void __set_poolid(const int64_t val);

  void __set_pageid(const int64_t val);

  void __set_extend(const std::string& val);

  void __set_totalReaction(const int64_t val);

  void __set_totalComment(const int64_t val);

  void __set_totalShare(const int64_t val);

  void __set_originPostID(const int64_t val);

  bool operator == (const TPostItem & rhs) const
  {
    if (!(idpost == rhs.idpost))
      return false;
    if (!(uid == rhs.uid))
      return false;
    if (!(content == rhs.content))
      return false;
    if (__isset.listMediaItems != rhs.__isset.listMediaItems)
      return false;
    else if (__isset.listMediaItems && !(listMediaItems == rhs.listMediaItems))
      return false;
    if (__isset.idbackground != rhs.__isset.idbackground)
      return false;
    else if (__isset.idbackground && !(idbackground == rhs.idbackground))
      return false;
    if (__isset.idfeeling != rhs.__isset.idfeeling)
      return false;
    else if (__isset.idfeeling && !(idfeeling == rhs.idfeeling))
      return false;
    if (!(privacy == rhs.privacy))
      return false;
    if (__isset.friendsexcept != rhs.__isset.friendsexcept)
      return false;
    else if (__isset.friendsexcept && !(friendsexcept == rhs.friendsexcept))
      return false;
    if (__isset.specificfriends != rhs.__isset.specificfriends)
      return false;
    else if (__isset.specificfriends && !(specificfriends == rhs.specificfriends))
      return false;
    if (__isset.tagusers != rhs.__isset.tagusers)
      return false;
    else if (__isset.tagusers && !(tagusers == rhs.tagusers))
      return false;
    if (__isset.locationId != rhs.__isset.locationId)
      return false;
    else if (__isset.locationId && !(locationId == rhs.locationId))
      return false;
    if (!(timestamps == rhs.timestamps))
      return false;
    if (!(pubkey == rhs.pubkey))
      return false;
    if (__isset.touid != rhs.__isset.touid)
      return false;
    else if (__isset.touid && !(touid == rhs.touid))
      return false;
    if (__isset.togroupid != rhs.__isset.togroupid)
      return false;
    else if (__isset.togroupid && !(togroupid == rhs.togroupid))
      return false;
    if (__isset.actionLinks != rhs.__isset.actionLinks)
      return false;
    else if (__isset.actionLinks && !(actionLinks == rhs.actionLinks))
      return false;
    if (__isset.poolid != rhs.__isset.poolid)
      return false;
    else if (__isset.poolid && !(poolid == rhs.poolid))
      return false;
    if (__isset.pageid != rhs.__isset.pageid)
      return false;
    else if (__isset.pageid && !(pageid == rhs.pageid))
      return false;
    if (__isset.extend != rhs.__isset.extend)
      return false;
    else if (__isset.extend && !(extend == rhs.extend))
      return false;
    if (!(totalReaction == rhs.totalReaction))
      return false;
    if (!(totalComment == rhs.totalComment))
      return false;
    if (!(totalShare == rhs.totalShare))
      return false;
    if (!(originPostID == rhs.originPostID))
      return false;
    return true;
  }
  bool operator != (const TPostItem &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const TPostItem & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

  virtual void printTo(std::ostream& out) const;
};

void swap(TPostItem &a, TPostItem &b);

std::ostream& operator<<(std::ostream& out, const TPostItem& obj);

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
  TPostItem data;

  _TDataResult__isset __isset;

  void __set_errorCode(const TErrorCode::type val);

  void __set_data(const TPostItem& val);

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
  _TListDataResult__isset() : errorCode(false), listDatas(false) {}
  bool errorCode :1;
  bool listDatas :1;
} _TListDataResult__isset;

class TListDataResult : public virtual ::apache::thrift::TBase {
 public:

  TListDataResult(const TListDataResult&);
  TListDataResult& operator=(const TListDataResult&);
  TListDataResult() : errorCode((TErrorCode::type)0) {
  }

  virtual ~TListDataResult() throw();
  TErrorCode::type errorCode;
  std::vector<TPostItem>  listDatas;

  _TListDataResult__isset __isset;

  void __set_errorCode(const TErrorCode::type val);

  void __set_listDatas(const std::vector<TPostItem> & val);

  bool operator == (const TListDataResult & rhs) const
  {
    if (!(errorCode == rhs.errorCode))
      return false;
    if (!(listDatas == rhs.listDatas))
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
