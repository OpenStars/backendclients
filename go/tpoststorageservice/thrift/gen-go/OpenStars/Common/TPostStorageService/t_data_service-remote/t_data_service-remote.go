// Autogenerated by Thrift Compiler (0.9.1)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"OpenStars/Common/TPostStorageService"
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  TDataResult getData(TKey key)")
	fmt.Fprintln(os.Stderr, "  TErrorCode putData(TKey key, TPostItem data)")
	fmt.Fprintln(os.Stderr, "  TErrorCode removeData(TKey key)")
	fmt.Fprintln(os.Stderr, "  TListDataResult getListDatas( listkey)")
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
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = math.MinInt32 // will become unneeded eventually
	_ = strconv.Atoi
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
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
	client := TPostStorageService.NewTDataServiceClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "getData":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetData requires 1 args")
			flag.Usage()
		}
		argvalue0, err32 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err32 != nil {
			Usage()
			return
		}
		value0 := TPostStorageService.TKey(argvalue0)
		fmt.Print(client.GetData(value0))
		fmt.Print("\n")
		break
	case "putData":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "PutData requires 2 args")
			flag.Usage()
		}
		argvalue0, err33 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err33 != nil {
			Usage()
			return
		}
		value0 := TPostStorageService.TKey(argvalue0)
		arg34 := flag.Arg(2)
		mbTrans35 := thrift.NewTMemoryBufferLen(len(arg34))
		defer mbTrans35.Close()
		_, err36 := mbTrans35.WriteString(arg34)
		if err36 != nil {
			Usage()
			return
		}
		factory37 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt38 := factory37.GetProtocol(mbTrans35)
		argvalue1 := TPostStorageService.NewTPostItem()
		err39 := argvalue1.Read(jsProt38)
		if err39 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.PutData(value0, value1))
		fmt.Print("\n")
		break
	case "removeData":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "RemoveData requires 1 args")
			flag.Usage()
		}
		argvalue0, err40 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err40 != nil {
			Usage()
			return
		}
		value0 := TPostStorageService.TKey(argvalue0)
		fmt.Print(client.RemoveData(value0))
		fmt.Print("\n")
		break
	case "getListDatas":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetListDatas requires 1 args")
			flag.Usage()
		}
		arg41 := flag.Arg(1)
		mbTrans42 := thrift.NewTMemoryBufferLen(len(arg41))
		defer mbTrans42.Close()
		_, err43 := mbTrans42.WriteString(arg41)
		if err43 != nil {
			Usage()
			return
		}
		factory44 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt45 := factory44.GetProtocol(mbTrans42)
		containerStruct0 := TPostStorageService.NewGetListDatasArgs()
		err46 := containerStruct0.ReadField1(jsProt45)
		if err46 != nil {
			Usage()
			return
		}
		argvalue0 := containerStruct0.Listkey
		value0 := argvalue0
		fmt.Print(client.GetListDatas(value0))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
