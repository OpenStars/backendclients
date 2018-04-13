// Autogenerated by Thrift Compiler (0.11.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
        "context"
        "flag"
        "fmt"
        "math"
        "net"
        "net/url"
        "os"
        "strconv"
        "strings"
        "git.apache.org/thrift.git/lib/go/thrift"
        "openstars/core/bigset/generic"
)


func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  TBigSetInfoResult createStringBigSet(TStringKey bsName)")
  fmt.Fprintln(os.Stderr, "  TBigSetInfoResult getBigSetInfoByName(TStringKey bsName)")
  fmt.Fprintln(os.Stderr, "  TBigSetInfoResult assignBigSetName(TStringKey bsName, TContainerKey bigsetID)")
  fmt.Fprintln(os.Stderr, "  TPutItemResult bsPutItem(TStringKey bsName, TItem item)")
  fmt.Fprintln(os.Stderr, "  bool bsRemoveItem(TStringKey bsName, TItemKey itemKey)")
  fmt.Fprintln(os.Stderr, "  TExistedResult bsExisted(TStringKey bsName, TItemKey itemKey)")
  fmt.Fprintln(os.Stderr, "  TItemResult bsGetItem(TStringKey bsName, TItemKey itemKey)")
  fmt.Fprintln(os.Stderr, "  TItemSetResult bsGetSlice(TStringKey bsName, i32 fromPos, i32 count)")
  fmt.Fprintln(os.Stderr, "  TItemSetResult bsGetSliceFromItem(TStringKey bsName, TItemKey fromKey, i32 count)")
  fmt.Fprintln(os.Stderr, "  TItemSetResult bsGetSliceR(TStringKey bsName, i32 fromPos, i32 count)")
  fmt.Fprintln(os.Stderr, "  TItemSetResult bsGetSliceFromItemR(TStringKey bsName, TItemKey fromKey, i32 count)")
  fmt.Fprintln(os.Stderr, "  TItemSetResult bsRangeQuery(TStringKey bsName, TItemKey startKey, TItemKey endKey)")
  fmt.Fprintln(os.Stderr, "  bool bsBulkLoad(TStringKey bsName, TItemSet setData)")
  fmt.Fprintln(os.Stderr, "  TMultiPutItemResult bsMultiPut(TStringKey bsName, TItemSet setData, bool getAddedItems, bool getReplacedItems)")
  fmt.Fprintln(os.Stderr, "  i64 getTotalCount(TStringKey bsName)")
  fmt.Fprintln(os.Stderr, "  i64 removeAll(TStringKey bsName)")
  fmt.Fprintln(os.Stderr, "  i64 totalStringKeyCount()")
  fmt.Fprintln(os.Stderr, "   getListKey(i64 fromIndex, i32 count)")
  fmt.Fprintln(os.Stderr, "   getListKeyFrom(TStringKey keyFrom, i32 count)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  var parsedUrl *url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Parse()
  
  if len(urlString) > 0 {
    var err error
    parsedUrl, err = url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  iprot := protocolFactory.GetProtocol(trans)
  oprot := protocolFactory.GetProtocol(trans)
  client := generic.NewTStringBigSetKVServiceClient(thrift.NewTStandardClient(iprot, oprot))
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "createStringBigSet":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CreateStringBigSet requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := generic.TStringKey(argvalue0)
    fmt.Print(client.CreateStringBigSet(context.Background(), value0))
    fmt.Print("\n")
    break
  case "getBigSetInfoByName":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetBigSetInfoByName requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := generic.TStringKey(argvalue0)
    fmt.Print(client.GetBigSetInfoByName(context.Background(), value0))
    fmt.Print("\n")
    break
  case "assignBigSetName":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "AssignBigSetName requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := generic.TStringKey(argvalue0)
    argvalue1, err152 := (strconv.ParseInt(flag.Arg(2), 10, 64))
    if err152 != nil {
      Usage()
      return
    }
    value1 := generic.TContainerKey(argvalue1)
    fmt.Print(client.AssignBigSetName(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "bsPutItem":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "BsPutItem requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := generic.TStringKey(argvalue0)
    arg154 := flag.Arg(2)
    mbTrans155 := thrift.NewTMemoryBufferLen(len(arg154))
    defer mbTrans155.Close()
    _, err156 := mbTrans155.WriteString(arg154)
    if err156 != nil {
      Usage()
      return
    }
    factory157 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt158 := factory157.GetProtocol(mbTrans155)
    argvalue1 := generic.NewTItem()
    err159 := argvalue1.Read(jsProt158)
    if err159 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.BsPutItem(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "bsRemoveItem":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "BsRemoveItem requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := generic.TStringKey(argvalue0)
    argvalue1 := []byte(flag.Arg(2))
    value1 := generic.TItemKey(argvalue1)
    fmt.Print(client.BsRemoveItem(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "bsExisted":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "BsExisted requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := generic.TStringKey(argvalue0)
    argvalue1 := []byte(flag.Arg(2))
    value1 := generic.TItemKey(argvalue1)
    fmt.Print(client.BsExisted(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "bsGetItem":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "BsGetItem requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := generic.TStringKey(argvalue0)
    argvalue1 := []byte(flag.Arg(2))
    value1 := generic.TItemKey(argvalue1)
    fmt.Print(client.BsGetItem(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "bsGetSlice":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "BsGetSlice requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := generic.TStringKey(argvalue0)
    tmp1, err167 := (strconv.Atoi(flag.Arg(2)))
    if err167 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    tmp2, err168 := (strconv.Atoi(flag.Arg(3)))
    if err168 != nil {
      Usage()
      return
    }
    argvalue2 := int32(tmp2)
    value2 := argvalue2
    fmt.Print(client.BsGetSlice(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "bsGetSliceFromItem":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "BsGetSliceFromItem requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := generic.TStringKey(argvalue0)
    argvalue1 := []byte(flag.Arg(2))
    value1 := generic.TItemKey(argvalue1)
    tmp2, err171 := (strconv.Atoi(flag.Arg(3)))
    if err171 != nil {
      Usage()
      return
    }
    argvalue2 := int32(tmp2)
    value2 := argvalue2
    fmt.Print(client.BsGetSliceFromItem(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "bsGetSliceR":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "BsGetSliceR requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := generic.TStringKey(argvalue0)
    tmp1, err173 := (strconv.Atoi(flag.Arg(2)))
    if err173 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    tmp2, err174 := (strconv.Atoi(flag.Arg(3)))
    if err174 != nil {
      Usage()
      return
    }
    argvalue2 := int32(tmp2)
    value2 := argvalue2
    fmt.Print(client.BsGetSliceR(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "bsGetSliceFromItemR":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "BsGetSliceFromItemR requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := generic.TStringKey(argvalue0)
    argvalue1 := []byte(flag.Arg(2))
    value1 := generic.TItemKey(argvalue1)
    tmp2, err177 := (strconv.Atoi(flag.Arg(3)))
    if err177 != nil {
      Usage()
      return
    }
    argvalue2 := int32(tmp2)
    value2 := argvalue2
    fmt.Print(client.BsGetSliceFromItemR(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "bsRangeQuery":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "BsRangeQuery requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := generic.TStringKey(argvalue0)
    argvalue1 := []byte(flag.Arg(2))
    value1 := generic.TItemKey(argvalue1)
    argvalue2 := []byte(flag.Arg(3))
    value2 := generic.TItemKey(argvalue2)
    fmt.Print(client.BsRangeQuery(context.Background(), value0, value1, value2))
    fmt.Print("\n")
    break
  case "bsBulkLoad":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "BsBulkLoad requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := generic.TStringKey(argvalue0)
    arg182 := flag.Arg(2)
    mbTrans183 := thrift.NewTMemoryBufferLen(len(arg182))
    defer mbTrans183.Close()
    _, err184 := mbTrans183.WriteString(arg182)
    if err184 != nil {
      Usage()
      return
    }
    factory185 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt186 := factory185.GetProtocol(mbTrans183)
    argvalue1 := generic.NewTItemSet()
    err187 := argvalue1.Read(jsProt186)
    if err187 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.BsBulkLoad(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "bsMultiPut":
    if flag.NArg() - 1 != 4 {
      fmt.Fprintln(os.Stderr, "BsMultiPut requires 4 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := generic.TStringKey(argvalue0)
    arg189 := flag.Arg(2)
    mbTrans190 := thrift.NewTMemoryBufferLen(len(arg189))
    defer mbTrans190.Close()
    _, err191 := mbTrans190.WriteString(arg189)
    if err191 != nil {
      Usage()
      return
    }
    factory192 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt193 := factory192.GetProtocol(mbTrans190)
    argvalue1 := generic.NewTItemSet()
    err194 := argvalue1.Read(jsProt193)
    if err194 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    argvalue2 := flag.Arg(3) == "true"
    value2 := argvalue2
    argvalue3 := flag.Arg(4) == "true"
    value3 := argvalue3
    fmt.Print(client.BsMultiPut(context.Background(), value0, value1, value2, value3))
    fmt.Print("\n")
    break
  case "getTotalCount":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetTotalCount requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := generic.TStringKey(argvalue0)
    fmt.Print(client.GetTotalCount(context.Background(), value0))
    fmt.Print("\n")
    break
  case "removeAll":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "RemoveAll requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := generic.TStringKey(argvalue0)
    fmt.Print(client.RemoveAll(context.Background(), value0))
    fmt.Print("\n")
    break
  case "totalStringKeyCount":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "TotalStringKeyCount requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.TotalStringKeyCount(context.Background()))
    fmt.Print("\n")
    break
  case "getListKey":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "GetListKey requires 2 args")
      flag.Usage()
    }
    argvalue0, err199 := (strconv.ParseInt(flag.Arg(1), 10, 64))
    if err199 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    tmp1, err200 := (strconv.Atoi(flag.Arg(2)))
    if err200 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    fmt.Print(client.GetListKey(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "getListKeyFrom":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "GetListKeyFrom requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := generic.TStringKey(argvalue0)
    tmp1, err202 := (strconv.Atoi(flag.Arg(2)))
    if err202 != nil {
      Usage()
      return
    }
    argvalue1 := int32(tmp1)
    value1 := argvalue1
    fmt.Print(client.GetListKeyFrom(context.Background(), value0, value1))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
