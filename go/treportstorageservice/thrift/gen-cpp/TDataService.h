/**
 * Autogenerated by Thrift Compiler (0.13.0)
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
#ifndef TDataService_H
#define TDataService_H

#include <thrift/TDispatchProcessor.h>
#include <thrift/async/TConcurrentClientSyncInfo.h>
#include <memory>
#include "treportstorage_types.h"

namespace OpenStars { namespace Common { namespace TReportStorageService {

#ifdef _MSC_VER
  #pragma warning( push )
  #pragma warning (disable : 4250 ) //inheriting methods via dominance 
#endif

class TDataServiceIf {
 public:
  virtual ~TDataServiceIf() {}
  virtual void getReport(TDataResult& _return, const int64_t reportId) = 0;
  virtual TErrorCode::type putReport(const int64_t reportId, const TReportItem& data) = 0;
  virtual bool RemoveReport(const int64_t reportId) = 0;
  virtual void getListReports(TListDataResult& _return, const std::vector<int64_t> & lsReportIds) = 0;
};

class TDataServiceIfFactory {
 public:
  typedef TDataServiceIf Handler;

  virtual ~TDataServiceIfFactory() {}

  virtual TDataServiceIf* getHandler(const ::apache::thrift::TConnectionInfo& connInfo) = 0;
  virtual void releaseHandler(TDataServiceIf* /* handler */) = 0;
};

class TDataServiceIfSingletonFactory : virtual public TDataServiceIfFactory {
 public:
  TDataServiceIfSingletonFactory(const ::std::shared_ptr<TDataServiceIf>& iface) : iface_(iface) {}
  virtual ~TDataServiceIfSingletonFactory() {}

  virtual TDataServiceIf* getHandler(const ::apache::thrift::TConnectionInfo&) {
    return iface_.get();
  }
  virtual void releaseHandler(TDataServiceIf* /* handler */) {}

 protected:
  ::std::shared_ptr<TDataServiceIf> iface_;
};

class TDataServiceNull : virtual public TDataServiceIf {
 public:
  virtual ~TDataServiceNull() {}
  void getReport(TDataResult& /* _return */, const int64_t /* reportId */) {
    return;
  }
  TErrorCode::type putReport(const int64_t /* reportId */, const TReportItem& /* data */) {
    TErrorCode::type _return = (TErrorCode::type)0;
    return _return;
  }
  bool RemoveReport(const int64_t /* reportId */) {
    bool _return = false;
    return _return;
  }
  void getListReports(TListDataResult& /* _return */, const std::vector<int64_t> & /* lsReportIds */) {
    return;
  }
};

typedef struct _TDataService_getReport_args__isset {
  _TDataService_getReport_args__isset() : reportId(false) {}
  bool reportId :1;
} _TDataService_getReport_args__isset;

class TDataService_getReport_args {
 public:

  TDataService_getReport_args(const TDataService_getReport_args&);
  TDataService_getReport_args& operator=(const TDataService_getReport_args&);
  TDataService_getReport_args() : reportId(0) {
  }

  virtual ~TDataService_getReport_args() noexcept;
  int64_t reportId;

  _TDataService_getReport_args__isset __isset;

  void __set_reportId(const int64_t val);

  bool operator == (const TDataService_getReport_args & rhs) const
  {
    if (!(reportId == rhs.reportId))
      return false;
    return true;
  }
  bool operator != (const TDataService_getReport_args &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const TDataService_getReport_args & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

};


class TDataService_getReport_pargs {
 public:


  virtual ~TDataService_getReport_pargs() noexcept;
  const int64_t* reportId;

  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

};

typedef struct _TDataService_getReport_result__isset {
  _TDataService_getReport_result__isset() : success(false) {}
  bool success :1;
} _TDataService_getReport_result__isset;

class TDataService_getReport_result {
 public:

  TDataService_getReport_result(const TDataService_getReport_result&);
  TDataService_getReport_result& operator=(const TDataService_getReport_result&);
  TDataService_getReport_result() {
  }

  virtual ~TDataService_getReport_result() noexcept;
  TDataResult success;

  _TDataService_getReport_result__isset __isset;

  void __set_success(const TDataResult& val);

  bool operator == (const TDataService_getReport_result & rhs) const
  {
    if (!(success == rhs.success))
      return false;
    return true;
  }
  bool operator != (const TDataService_getReport_result &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const TDataService_getReport_result & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

};

typedef struct _TDataService_getReport_presult__isset {
  _TDataService_getReport_presult__isset() : success(false) {}
  bool success :1;
} _TDataService_getReport_presult__isset;

class TDataService_getReport_presult {
 public:


  virtual ~TDataService_getReport_presult() noexcept;
  TDataResult* success;

  _TDataService_getReport_presult__isset __isset;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);

};

typedef struct _TDataService_putReport_args__isset {
  _TDataService_putReport_args__isset() : reportId(false), data(false) {}
  bool reportId :1;
  bool data :1;
} _TDataService_putReport_args__isset;

class TDataService_putReport_args {
 public:

  TDataService_putReport_args(const TDataService_putReport_args&);
  TDataService_putReport_args& operator=(const TDataService_putReport_args&);
  TDataService_putReport_args() : reportId(0) {
  }

  virtual ~TDataService_putReport_args() noexcept;
  int64_t reportId;
  TReportItem data;

  _TDataService_putReport_args__isset __isset;

  void __set_reportId(const int64_t val);

  void __set_data(const TReportItem& val);

  bool operator == (const TDataService_putReport_args & rhs) const
  {
    if (!(reportId == rhs.reportId))
      return false;
    if (!(data == rhs.data))
      return false;
    return true;
  }
  bool operator != (const TDataService_putReport_args &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const TDataService_putReport_args & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

};


class TDataService_putReport_pargs {
 public:


  virtual ~TDataService_putReport_pargs() noexcept;
  const int64_t* reportId;
  const TReportItem* data;

  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

};

typedef struct _TDataService_putReport_result__isset {
  _TDataService_putReport_result__isset() : success(false) {}
  bool success :1;
} _TDataService_putReport_result__isset;

class TDataService_putReport_result {
 public:

  TDataService_putReport_result(const TDataService_putReport_result&);
  TDataService_putReport_result& operator=(const TDataService_putReport_result&);
  TDataService_putReport_result() : success((TErrorCode::type)0) {
  }

  virtual ~TDataService_putReport_result() noexcept;
  TErrorCode::type success;

  _TDataService_putReport_result__isset __isset;

  void __set_success(const TErrorCode::type val);

  bool operator == (const TDataService_putReport_result & rhs) const
  {
    if (!(success == rhs.success))
      return false;
    return true;
  }
  bool operator != (const TDataService_putReport_result &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const TDataService_putReport_result & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

};

typedef struct _TDataService_putReport_presult__isset {
  _TDataService_putReport_presult__isset() : success(false) {}
  bool success :1;
} _TDataService_putReport_presult__isset;

class TDataService_putReport_presult {
 public:


  virtual ~TDataService_putReport_presult() noexcept;
  TErrorCode::type* success;

  _TDataService_putReport_presult__isset __isset;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);

};

typedef struct _TDataService_RemoveReport_args__isset {
  _TDataService_RemoveReport_args__isset() : reportId(false) {}
  bool reportId :1;
} _TDataService_RemoveReport_args__isset;

class TDataService_RemoveReport_args {
 public:

  TDataService_RemoveReport_args(const TDataService_RemoveReport_args&);
  TDataService_RemoveReport_args& operator=(const TDataService_RemoveReport_args&);
  TDataService_RemoveReport_args() : reportId(0) {
  }

  virtual ~TDataService_RemoveReport_args() noexcept;
  int64_t reportId;

  _TDataService_RemoveReport_args__isset __isset;

  void __set_reportId(const int64_t val);

  bool operator == (const TDataService_RemoveReport_args & rhs) const
  {
    if (!(reportId == rhs.reportId))
      return false;
    return true;
  }
  bool operator != (const TDataService_RemoveReport_args &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const TDataService_RemoveReport_args & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

};


class TDataService_RemoveReport_pargs {
 public:


  virtual ~TDataService_RemoveReport_pargs() noexcept;
  const int64_t* reportId;

  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

};

typedef struct _TDataService_RemoveReport_result__isset {
  _TDataService_RemoveReport_result__isset() : success(false) {}
  bool success :1;
} _TDataService_RemoveReport_result__isset;

class TDataService_RemoveReport_result {
 public:

  TDataService_RemoveReport_result(const TDataService_RemoveReport_result&);
  TDataService_RemoveReport_result& operator=(const TDataService_RemoveReport_result&);
  TDataService_RemoveReport_result() : success(0) {
  }

  virtual ~TDataService_RemoveReport_result() noexcept;
  bool success;

  _TDataService_RemoveReport_result__isset __isset;

  void __set_success(const bool val);

  bool operator == (const TDataService_RemoveReport_result & rhs) const
  {
    if (!(success == rhs.success))
      return false;
    return true;
  }
  bool operator != (const TDataService_RemoveReport_result &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const TDataService_RemoveReport_result & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

};

typedef struct _TDataService_RemoveReport_presult__isset {
  _TDataService_RemoveReport_presult__isset() : success(false) {}
  bool success :1;
} _TDataService_RemoveReport_presult__isset;

class TDataService_RemoveReport_presult {
 public:


  virtual ~TDataService_RemoveReport_presult() noexcept;
  bool* success;

  _TDataService_RemoveReport_presult__isset __isset;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);

};

typedef struct _TDataService_getListReports_args__isset {
  _TDataService_getListReports_args__isset() : lsReportIds(false) {}
  bool lsReportIds :1;
} _TDataService_getListReports_args__isset;

class TDataService_getListReports_args {
 public:

  TDataService_getListReports_args(const TDataService_getListReports_args&);
  TDataService_getListReports_args& operator=(const TDataService_getListReports_args&);
  TDataService_getListReports_args() {
  }

  virtual ~TDataService_getListReports_args() noexcept;
  std::vector<int64_t>  lsReportIds;

  _TDataService_getListReports_args__isset __isset;

  void __set_lsReportIds(const std::vector<int64_t> & val);

  bool operator == (const TDataService_getListReports_args & rhs) const
  {
    if (!(lsReportIds == rhs.lsReportIds))
      return false;
    return true;
  }
  bool operator != (const TDataService_getListReports_args &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const TDataService_getListReports_args & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

};


class TDataService_getListReports_pargs {
 public:


  virtual ~TDataService_getListReports_pargs() noexcept;
  const std::vector<int64_t> * lsReportIds;

  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

};

typedef struct _TDataService_getListReports_result__isset {
  _TDataService_getListReports_result__isset() : success(false) {}
  bool success :1;
} _TDataService_getListReports_result__isset;

class TDataService_getListReports_result {
 public:

  TDataService_getListReports_result(const TDataService_getListReports_result&);
  TDataService_getListReports_result& operator=(const TDataService_getListReports_result&);
  TDataService_getListReports_result() {
  }

  virtual ~TDataService_getListReports_result() noexcept;
  TListDataResult success;

  _TDataService_getListReports_result__isset __isset;

  void __set_success(const TListDataResult& val);

  bool operator == (const TDataService_getListReports_result & rhs) const
  {
    if (!(success == rhs.success))
      return false;
    return true;
  }
  bool operator != (const TDataService_getListReports_result &rhs) const {
    return !(*this == rhs);
  }

  bool operator < (const TDataService_getListReports_result & ) const;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);
  uint32_t write(::apache::thrift::protocol::TProtocol* oprot) const;

};

typedef struct _TDataService_getListReports_presult__isset {
  _TDataService_getListReports_presult__isset() : success(false) {}
  bool success :1;
} _TDataService_getListReports_presult__isset;

class TDataService_getListReports_presult {
 public:


  virtual ~TDataService_getListReports_presult() noexcept;
  TListDataResult* success;

  _TDataService_getListReports_presult__isset __isset;

  uint32_t read(::apache::thrift::protocol::TProtocol* iprot);

};

class TDataServiceClient : virtual public TDataServiceIf {
 public:
  TDataServiceClient(std::shared_ptr< ::apache::thrift::protocol::TProtocol> prot) {
    setProtocol(prot);
  }
  TDataServiceClient(std::shared_ptr< ::apache::thrift::protocol::TProtocol> iprot, std::shared_ptr< ::apache::thrift::protocol::TProtocol> oprot) {
    setProtocol(iprot,oprot);
  }
 private:
  void setProtocol(std::shared_ptr< ::apache::thrift::protocol::TProtocol> prot) {
  setProtocol(prot,prot);
  }
  void setProtocol(std::shared_ptr< ::apache::thrift::protocol::TProtocol> iprot, std::shared_ptr< ::apache::thrift::protocol::TProtocol> oprot) {
    piprot_=iprot;
    poprot_=oprot;
    iprot_ = iprot.get();
    oprot_ = oprot.get();
  }
 public:
  std::shared_ptr< ::apache::thrift::protocol::TProtocol> getInputProtocol() {
    return piprot_;
  }
  std::shared_ptr< ::apache::thrift::protocol::TProtocol> getOutputProtocol() {
    return poprot_;
  }
  void getReport(TDataResult& _return, const int64_t reportId);
  void send_getReport(const int64_t reportId);
  void recv_getReport(TDataResult& _return);
  TErrorCode::type putReport(const int64_t reportId, const TReportItem& data);
  void send_putReport(const int64_t reportId, const TReportItem& data);
  TErrorCode::type recv_putReport();
  bool RemoveReport(const int64_t reportId);
  void send_RemoveReport(const int64_t reportId);
  bool recv_RemoveReport();
  void getListReports(TListDataResult& _return, const std::vector<int64_t> & lsReportIds);
  void send_getListReports(const std::vector<int64_t> & lsReportIds);
  void recv_getListReports(TListDataResult& _return);
 protected:
  std::shared_ptr< ::apache::thrift::protocol::TProtocol> piprot_;
  std::shared_ptr< ::apache::thrift::protocol::TProtocol> poprot_;
  ::apache::thrift::protocol::TProtocol* iprot_;
  ::apache::thrift::protocol::TProtocol* oprot_;
};

class TDataServiceProcessor : public ::apache::thrift::TDispatchProcessor {
 protected:
  ::std::shared_ptr<TDataServiceIf> iface_;
  virtual bool dispatchCall(::apache::thrift::protocol::TProtocol* iprot, ::apache::thrift::protocol::TProtocol* oprot, const std::string& fname, int32_t seqid, void* callContext);
 private:
  typedef  void (TDataServiceProcessor::*ProcessFunction)(int32_t, ::apache::thrift::protocol::TProtocol*, ::apache::thrift::protocol::TProtocol*, void*);
  typedef std::map<std::string, ProcessFunction> ProcessMap;
  ProcessMap processMap_;
  void process_getReport(int32_t seqid, ::apache::thrift::protocol::TProtocol* iprot, ::apache::thrift::protocol::TProtocol* oprot, void* callContext);
  void process_putReport(int32_t seqid, ::apache::thrift::protocol::TProtocol* iprot, ::apache::thrift::protocol::TProtocol* oprot, void* callContext);
  void process_RemoveReport(int32_t seqid, ::apache::thrift::protocol::TProtocol* iprot, ::apache::thrift::protocol::TProtocol* oprot, void* callContext);
  void process_getListReports(int32_t seqid, ::apache::thrift::protocol::TProtocol* iprot, ::apache::thrift::protocol::TProtocol* oprot, void* callContext);
 public:
  TDataServiceProcessor(::std::shared_ptr<TDataServiceIf> iface) :
    iface_(iface) {
    processMap_["getReport"] = &TDataServiceProcessor::process_getReport;
    processMap_["putReport"] = &TDataServiceProcessor::process_putReport;
    processMap_["RemoveReport"] = &TDataServiceProcessor::process_RemoveReport;
    processMap_["getListReports"] = &TDataServiceProcessor::process_getListReports;
  }

  virtual ~TDataServiceProcessor() {}
};

class TDataServiceProcessorFactory : public ::apache::thrift::TProcessorFactory {
 public:
  TDataServiceProcessorFactory(const ::std::shared_ptr< TDataServiceIfFactory >& handlerFactory) :
      handlerFactory_(handlerFactory) {}

  ::std::shared_ptr< ::apache::thrift::TProcessor > getProcessor(const ::apache::thrift::TConnectionInfo& connInfo);

 protected:
  ::std::shared_ptr< TDataServiceIfFactory > handlerFactory_;
};

class TDataServiceMultiface : virtual public TDataServiceIf {
 public:
  TDataServiceMultiface(std::vector<std::shared_ptr<TDataServiceIf> >& ifaces) : ifaces_(ifaces) {
  }
  virtual ~TDataServiceMultiface() {}
 protected:
  std::vector<std::shared_ptr<TDataServiceIf> > ifaces_;
  TDataServiceMultiface() {}
  void add(::std::shared_ptr<TDataServiceIf> iface) {
    ifaces_.push_back(iface);
  }
 public:
  void getReport(TDataResult& _return, const int64_t reportId) {
    size_t sz = ifaces_.size();
    size_t i = 0;
    for (; i < (sz - 1); ++i) {
      ifaces_[i]->getReport(_return, reportId);
    }
    ifaces_[i]->getReport(_return, reportId);
    return;
  }

  TErrorCode::type putReport(const int64_t reportId, const TReportItem& data) {
    size_t sz = ifaces_.size();
    size_t i = 0;
    for (; i < (sz - 1); ++i) {
      ifaces_[i]->putReport(reportId, data);
    }
    return ifaces_[i]->putReport(reportId, data);
  }

  bool RemoveReport(const int64_t reportId) {
    size_t sz = ifaces_.size();
    size_t i = 0;
    for (; i < (sz - 1); ++i) {
      ifaces_[i]->RemoveReport(reportId);
    }
    return ifaces_[i]->RemoveReport(reportId);
  }

  void getListReports(TListDataResult& _return, const std::vector<int64_t> & lsReportIds) {
    size_t sz = ifaces_.size();
    size_t i = 0;
    for (; i < (sz - 1); ++i) {
      ifaces_[i]->getListReports(_return, lsReportIds);
    }
    ifaces_[i]->getListReports(_return, lsReportIds);
    return;
  }

};

// The 'concurrent' client is a thread safe client that correctly handles
// out of order responses.  It is slower than the regular client, so should
// only be used when you need to share a connection among multiple threads
class TDataServiceConcurrentClient : virtual public TDataServiceIf {
 public:
  TDataServiceConcurrentClient(std::shared_ptr< ::apache::thrift::protocol::TProtocol> prot, std::shared_ptr<::apache::thrift::async::TConcurrentClientSyncInfo> sync) : sync_(sync)
{
    setProtocol(prot);
  }
  TDataServiceConcurrentClient(std::shared_ptr< ::apache::thrift::protocol::TProtocol> iprot, std::shared_ptr< ::apache::thrift::protocol::TProtocol> oprot, std::shared_ptr<::apache::thrift::async::TConcurrentClientSyncInfo> sync) : sync_(sync)
{
    setProtocol(iprot,oprot);
  }
 private:
  void setProtocol(std::shared_ptr< ::apache::thrift::protocol::TProtocol> prot) {
  setProtocol(prot,prot);
  }
  void setProtocol(std::shared_ptr< ::apache::thrift::protocol::TProtocol> iprot, std::shared_ptr< ::apache::thrift::protocol::TProtocol> oprot) {
    piprot_=iprot;
    poprot_=oprot;
    iprot_ = iprot.get();
    oprot_ = oprot.get();
  }
 public:
  std::shared_ptr< ::apache::thrift::protocol::TProtocol> getInputProtocol() {
    return piprot_;
  }
  std::shared_ptr< ::apache::thrift::protocol::TProtocol> getOutputProtocol() {
    return poprot_;
  }
  void getReport(TDataResult& _return, const int64_t reportId);
  int32_t send_getReport(const int64_t reportId);
  void recv_getReport(TDataResult& _return, const int32_t seqid);
  TErrorCode::type putReport(const int64_t reportId, const TReportItem& data);
  int32_t send_putReport(const int64_t reportId, const TReportItem& data);
  TErrorCode::type recv_putReport(const int32_t seqid);
  bool RemoveReport(const int64_t reportId);
  int32_t send_RemoveReport(const int64_t reportId);
  bool recv_RemoveReport(const int32_t seqid);
  void getListReports(TListDataResult& _return, const std::vector<int64_t> & lsReportIds);
  int32_t send_getListReports(const std::vector<int64_t> & lsReportIds);
  void recv_getListReports(TListDataResult& _return, const int32_t seqid);
 protected:
  std::shared_ptr< ::apache::thrift::protocol::TProtocol> piprot_;
  std::shared_ptr< ::apache::thrift::protocol::TProtocol> poprot_;
  ::apache::thrift::protocol::TProtocol* iprot_;
  ::apache::thrift::protocol::TProtocol* oprot_;
  std::shared_ptr<::apache::thrift::async::TConcurrentClientSyncInfo> sync_;
};

#ifdef _MSC_VER
  #pragma warning( pop )
#endif

}}} // namespace

#endif