// Autogenerated by Thrift Compiler (0.9.1)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package TPostStorageService

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
)

// (needed to ensure safety because of naive import list construction.)
var _ = math.MinInt32
var _ = thrift.ZERO
var _ = fmt.Printf

type TDataService interface {
	// Parameters:
	//  - Key
	GetData(key TKey) (r *TDataResult, err error)
	// Parameters:
	//  - Key
	//  - Data
	PutData(key TKey, data *TPostItem) (r TErrorCode, err error)
	// Parameters:
	//  - Key
	RemoveData(key TKey) (r TErrorCode, err error)
	// Parameters:
	//  - Listkey
	GetListDatas(listkey []TKey) (r *TListDataResult, err error)
}

type TDataServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewTDataServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *TDataServiceClient {
	return &TDataServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewTDataServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *TDataServiceClient {
	return &TDataServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - Key
func (p *TDataServiceClient) GetData(key TKey) (r *TDataResult, err error) {
	if err = p.sendGetData(key); err != nil {
		return
	}
	return p.recvGetData()
}

func (p *TDataServiceClient) sendGetData(key TKey) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("getData", thrift.CALL, p.SeqId)
	args13 := NewGetDataArgs()
	args13.Key = key
	err = args13.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return
}

func (p *TDataServiceClient) recvGetData() (value *TDataResult, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error15 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error16 error
		error16, err = error15.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error16
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result14 := NewGetDataResult()
	err = result14.Read(iprot)
	iprot.ReadMessageEnd()
	value = result14.Success
	return
}

// Parameters:
//  - Key
//  - Data
func (p *TDataServiceClient) PutData(key TKey, data *TPostItem) (r TErrorCode, err error) {
	if err = p.sendPutData(key, data); err != nil {
		return
	}
	return p.recvPutData()
}

func (p *TDataServiceClient) sendPutData(key TKey, data *TPostItem) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("putData", thrift.CALL, p.SeqId)
	args17 := NewPutDataArgs()
	args17.Key = key
	args17.Data = data
	err = args17.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return
}

func (p *TDataServiceClient) recvPutData() (value TErrorCode, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error19 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error20 error
		error20, err = error19.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error20
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result18 := NewPutDataResult()
	err = result18.Read(iprot)
	iprot.ReadMessageEnd()
	value = result18.Success
	return
}

// Parameters:
//  - Key
func (p *TDataServiceClient) RemoveData(key TKey) (r TErrorCode, err error) {
	if err = p.sendRemoveData(key); err != nil {
		return
	}
	return p.recvRemoveData()
}

func (p *TDataServiceClient) sendRemoveData(key TKey) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("removeData", thrift.CALL, p.SeqId)
	args21 := NewRemoveDataArgs()
	args21.Key = key
	err = args21.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return
}

func (p *TDataServiceClient) recvRemoveData() (value TErrorCode, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error23 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error24 error
		error24, err = error23.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error24
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result22 := NewRemoveDataResult()
	err = result22.Read(iprot)
	iprot.ReadMessageEnd()
	value = result22.Success
	return
}

// Parameters:
//  - Listkey
func (p *TDataServiceClient) GetListDatas(listkey []TKey) (r *TListDataResult, err error) {
	if err = p.sendGetListDatas(listkey); err != nil {
		return
	}
	return p.recvGetListDatas()
}

func (p *TDataServiceClient) sendGetListDatas(listkey []TKey) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("getListDatas", thrift.CALL, p.SeqId)
	args25 := NewGetListDatasArgs()
	args25.Listkey = listkey
	err = args25.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return
}

func (p *TDataServiceClient) recvGetListDatas() (value *TListDataResult, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error27 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error28 error
		error28, err = error27.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error28
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result26 := NewGetListDatasResult()
	err = result26.Read(iprot)
	iprot.ReadMessageEnd()
	value = result26.Success
	return
}

type TDataServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      TDataService
}

func (p *TDataServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *TDataServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *TDataServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewTDataServiceProcessor(handler TDataService) *TDataServiceProcessor {

	self29 := &TDataServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self29.processorMap["getData"] = &tDataServiceProcessorGetData{handler: handler}
	self29.processorMap["putData"] = &tDataServiceProcessorPutData{handler: handler}
	self29.processorMap["removeData"] = &tDataServiceProcessorRemoveData{handler: handler}
	self29.processorMap["getListDatas"] = &tDataServiceProcessorGetListDatas{handler: handler}
	return self29
}

func (p *TDataServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x30 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x30.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x30

}

type tDataServiceProcessorGetData struct {
	handler TDataService
}

func (p *tDataServiceProcessorGetData) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewGetDataArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getData", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewGetDataResult()
	if result.Success, err = p.handler.GetData(args.Key); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getData: "+err.Error())
		oprot.WriteMessageBegin("getData", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("getData", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type tDataServiceProcessorPutData struct {
	handler TDataService
}

func (p *tDataServiceProcessorPutData) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewPutDataArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("putData", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewPutDataResult()
	if result.Success, err = p.handler.PutData(args.Key, args.Data); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing putData: "+err.Error())
		oprot.WriteMessageBegin("putData", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("putData", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type tDataServiceProcessorRemoveData struct {
	handler TDataService
}

func (p *tDataServiceProcessorRemoveData) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewRemoveDataArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("removeData", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewRemoveDataResult()
	if result.Success, err = p.handler.RemoveData(args.Key); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing removeData: "+err.Error())
		oprot.WriteMessageBegin("removeData", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("removeData", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type tDataServiceProcessorGetListDatas struct {
	handler TDataService
}

func (p *tDataServiceProcessorGetListDatas) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewGetListDatasArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("getListDatas", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewGetListDatasResult()
	if result.Success, err = p.handler.GetListDatas(args.Listkey); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getListDatas: "+err.Error())
		oprot.WriteMessageBegin("getListDatas", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("getListDatas", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

type GetDataArgs struct {
	Key TKey `thrift:"key,1"`
}

func NewGetDataArgs() *GetDataArgs {
	return &GetDataArgs{}
}

func (p *GetDataArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *GetDataArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 1: %s")
	} else {
		p.Key = TKey(v)
	}
	return nil
}

func (p *GetDataArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getData_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *GetDataArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("key", thrift.I64, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:key: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.Key)); err != nil {
		return fmt.Errorf("%T.key (1) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:key: %s", p, err)
	}
	return err
}

func (p *GetDataArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetDataArgs(%+v)", *p)
}

type GetDataResult struct {
	Success *TDataResult `thrift:"success,0"`
}

func NewGetDataResult() *GetDataResult {
	return &GetDataResult{}
}

func (p *GetDataResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *GetDataResult) readField0(iprot thrift.TProtocol) error {
	p.Success = NewTDataResult()
	if err := p.Success.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Success)
	}
	return nil
}

func (p *GetDataResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getData_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	switch {
	default:
		if err := p.writeField0(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *GetDataResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.Success != nil {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Success)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 0:success: %s", p, err)
		}
	}
	return err
}

func (p *GetDataResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetDataResult(%+v)", *p)
}

type PutDataArgs struct {
	Key  TKey       `thrift:"key,1"`
	Data *TPostItem `thrift:"data,2"`
}

func NewPutDataArgs() *PutDataArgs {
	return &PutDataArgs{}
}

func (p *PutDataArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *PutDataArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 1: %s")
	} else {
		p.Key = TKey(v)
	}
	return nil
}

func (p *PutDataArgs) readField2(iprot thrift.TProtocol) error {
	p.Data = NewTPostItem()
	if err := p.Data.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Data)
	}
	return nil
}

func (p *PutDataArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("putData_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *PutDataArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("key", thrift.I64, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:key: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.Key)); err != nil {
		return fmt.Errorf("%T.key (1) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:key: %s", p, err)
	}
	return err
}

func (p *PutDataArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if p.Data != nil {
		if err := oprot.WriteFieldBegin("data", thrift.STRUCT, 2); err != nil {
			return fmt.Errorf("%T write field begin error 2:data: %s", p, err)
		}
		if err := p.Data.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Data)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 2:data: %s", p, err)
		}
	}
	return err
}

func (p *PutDataArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PutDataArgs(%+v)", *p)
}

type PutDataResult struct {
	Success TErrorCode `thrift:"success,0"`
}

func NewPutDataResult() *PutDataResult {
	return &PutDataResult{
		Success: math.MinInt32 - 1, // unset sentinal value
	}
}

func (p *PutDataResult) IsSetSuccess() bool {
	return int64(p.Success) != math.MinInt32-1
}

func (p *PutDataResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *PutDataResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 0: %s")
	} else {
		p.Success = TErrorCode(v)
	}
	return nil
}

func (p *PutDataResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("putData_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	switch {
	default:
		if err := p.writeField0(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *PutDataResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.I32, 0); err != nil {
			return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
		}
		if err := oprot.WriteI32(int32(p.Success)); err != nil {
			return fmt.Errorf("%T.success (0) field write error: %s", p)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 0:success: %s", p, err)
		}
	}
	return err
}

func (p *PutDataResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PutDataResult(%+v)", *p)
}

type RemoveDataArgs struct {
	Key TKey `thrift:"key,1"`
}

func NewRemoveDataArgs() *RemoveDataArgs {
	return &RemoveDataArgs{}
}

func (p *RemoveDataArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *RemoveDataArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 1: %s")
	} else {
		p.Key = TKey(v)
	}
	return nil
}

func (p *RemoveDataArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("removeData_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *RemoveDataArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("key", thrift.I64, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:key: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.Key)); err != nil {
		return fmt.Errorf("%T.key (1) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:key: %s", p, err)
	}
	return err
}

func (p *RemoveDataArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RemoveDataArgs(%+v)", *p)
}

type RemoveDataResult struct {
	Success TErrorCode `thrift:"success,0"`
}

func NewRemoveDataResult() *RemoveDataResult {
	return &RemoveDataResult{
		Success: math.MinInt32 - 1, // unset sentinal value
	}
}

func (p *RemoveDataResult) IsSetSuccess() bool {
	return int64(p.Success) != math.MinInt32-1
}

func (p *RemoveDataResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *RemoveDataResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 0: %s")
	} else {
		p.Success = TErrorCode(v)
	}
	return nil
}

func (p *RemoveDataResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("removeData_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	switch {
	default:
		if err := p.writeField0(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *RemoveDataResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.I32, 0); err != nil {
			return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
		}
		if err := oprot.WriteI32(int32(p.Success)); err != nil {
			return fmt.Errorf("%T.success (0) field write error: %s", p)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 0:success: %s", p, err)
		}
	}
	return err
}

func (p *RemoveDataResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RemoveDataResult(%+v)", *p)
}

type GetListDatasArgs struct {
	Listkey []TKey `thrift:"listkey,1"`
}

func NewGetListDatasArgs() *GetListDatasArgs {
	return &GetListDatasArgs{}
}

func (p *GetListDatasArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *GetListDatasArgs) readField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list being: %s")
	}
	p.Listkey = make([]TKey, 0, size)
	for i := 0; i < size; i++ {
		var _elem31 TKey
		if v, err := iprot.ReadI64(); err != nil {
			return fmt.Errorf("error reading field 0: %s")
		} else {
			_elem31 = TKey(v)
		}
		p.Listkey = append(p.Listkey, _elem31)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s")
	}
	return nil
}

func (p *GetListDatasArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getListDatas_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *GetListDatasArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if p.Listkey != nil {
		if err := oprot.WriteFieldBegin("listkey", thrift.LIST, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:listkey: %s", p, err)
		}
		if err := oprot.WriteListBegin(thrift.I64, len(p.Listkey)); err != nil {
			return fmt.Errorf("error writing list begin: %s")
		}
		for _, v := range p.Listkey {
			if err := oprot.WriteI64(int64(v)); err != nil {
				return fmt.Errorf("%T. (0) field write error: %s", p)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return fmt.Errorf("error writing list end: %s")
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:listkey: %s", p, err)
		}
	}
	return err
}

func (p *GetListDatasArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetListDatasArgs(%+v)", *p)
}

type GetListDatasResult struct {
	Success *TListDataResult `thrift:"success,0"`
}

func NewGetListDatasResult() *GetListDatasResult {
	return &GetListDatasResult{}
}

func (p *GetListDatasResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error", p)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *GetListDatasResult) readField0(iprot thrift.TProtocol) error {
	p.Success = NewTListDataResult()
	if err := p.Success.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Success)
	}
	return nil
}

func (p *GetListDatasResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getListDatas_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	switch {
	default:
		if err := p.writeField0(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *GetListDatasResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.Success != nil {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Success)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 0:success: %s", p, err)
		}
	}
	return err
}

func (p *GetListDatasResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetListDatasResult(%+v)", *p)
}