// Autogenerated by Thrift Compiler (1.0.0-dev)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package users

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/brettshollenberger/thrift-example/gen-go/types"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var _ = types.GoUnusedProtection__

type UsersService interface {
	// Parameters:
	//  - Id
	Find(id int64) (r *User, err error)
	// Parameters:
	//  - User
	Create(user *User) (r *User, err error)
}

type UsersServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewUsersServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *UsersServiceClient {
	return &UsersServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewUsersServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *UsersServiceClient {
	return &UsersServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - Id
func (p *UsersServiceClient) Find(id int64) (r *User, err error) {
	if err = p.sendFind(id); err != nil {
		return
	}
	return p.recvFind()
}

func (p *UsersServiceClient) sendFind(id int64) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("find", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := FindArgs{
		Id: id,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *UsersServiceClient) recvFind() (value *User, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "find" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "find failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "find failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error0 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error1 error
		error1, err = error0.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error1
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "find failed: invalid message type")
		return
	}
	result := FindResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.E != nil {
		err = result.E
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - User
func (p *UsersServiceClient) Create(user *User) (r *User, err error) {
	if err = p.sendCreate(user); err != nil {
		return
	}
	return p.recvCreate()
}

func (p *UsersServiceClient) sendCreate(user *User) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("create", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := CreateArgs{
		User: user,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *UsersServiceClient) recvCreate() (value *User, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "create" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "create failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "create failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error2 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error3 error
		error3, err = error2.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error3
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "create failed: invalid message type")
		return
	}
	result := CreateResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.E != nil {
		err = result.E
		return
	}
	value = result.GetSuccess()
	return
}

type UsersServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      UsersService
}

func (p *UsersServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *UsersServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *UsersServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewUsersServiceProcessor(handler UsersService) *UsersServiceProcessor {

	self4 := &UsersServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self4.processorMap["find"] = &usersServiceProcessorFind{handler: handler}
	self4.processorMap["create"] = &usersServiceProcessorCreate{handler: handler}
	return self4
}

func (p *UsersServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x5 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x5.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x5

}

type usersServiceProcessorFind struct {
	handler UsersService
}

func (p *usersServiceProcessorFind) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := FindArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("find", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := FindResult{}
	var retval *User
	var err2 error
	if retval, err2 = p.handler.Find(args.Id); err2 != nil {
		switch v := err2.(type) {
		case *types.ServiceException:
			result.E = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing find: "+err2.Error())
			oprot.WriteMessageBegin("find", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("find", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type usersServiceProcessorCreate struct {
	handler UsersService
}

func (p *usersServiceProcessorCreate) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := CreateArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("create", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := CreateResult{}
	var retval *User
	var err2 error
	if retval, err2 = p.handler.Create(args.User); err2 != nil {
		switch v := err2.(type) {
		case *types.ServiceException:
			result.E = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing create: "+err2.Error())
			oprot.WriteMessageBegin("create", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("create", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Id
type FindArgs struct {
	Id int64 `thrift:"id,1" json:"id"`
}

func NewFindArgs() *FindArgs {
	return &FindArgs{}
}

func (p *FindArgs) GetId() int64 {
	return p.Id
}
func (p *FindArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
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
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *FindArgs) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Id = v
	}
	return nil
}

func (p *FindArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("find_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *FindArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("id", thrift.I64, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:id: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.Id)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.id (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:id: ", p), err)
	}
	return err
}

func (p *FindArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FindArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type FindResult struct {
	Success *User                   `thrift:"success,0" json:"success,omitempty"`
	E       *types.ServiceException `thrift:"e,1" json:"e,omitempty"`
}

func NewFindResult() *FindResult {
	return &FindResult{}
}

var FindResult_Success_DEFAULT *User

func (p *FindResult) GetSuccess() *User {
	if !p.IsSetSuccess() {
		return FindResult_Success_DEFAULT
	}
	return p.Success
}

var FindResult_E_DEFAULT *types.ServiceException

func (p *FindResult) GetE() *types.ServiceException {
	if !p.IsSetE() {
		return FindResult_E_DEFAULT
	}
	return p.E
}
func (p *FindResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FindResult) IsSetE() bool {
	return p.E != nil
}

func (p *FindResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.ReadField0(iprot); err != nil {
				return err
			}
		case 1:
			if err := p.ReadField1(iprot); err != nil {
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
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *FindResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &User{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *FindResult) ReadField1(iprot thrift.TProtocol) error {
	p.E = &types.ServiceException{}
	if err := p.E.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *FindResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("find_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *FindResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *FindResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin("e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *FindResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FindResult(%+v)", *p)
}

// Attributes:
//  - User
type CreateArgs struct {
	User *User `thrift:"user,1" json:"user"`
}

func NewCreateArgs() *CreateArgs {
	return &CreateArgs{}
}

var CreateArgs_User_DEFAULT *User

func (p *CreateArgs) GetUser() *User {
	if !p.IsSetUser() {
		return CreateArgs_User_DEFAULT
	}
	return p.User
}
func (p *CreateArgs) IsSetUser() bool {
	return p.User != nil
}

func (p *CreateArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
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
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *CreateArgs) ReadField1(iprot thrift.TProtocol) error {
	p.User = &User{}
	if err := p.User.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.User), err)
	}
	return nil
}

func (p *CreateArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("create_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *CreateArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("user", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:user: ", p), err)
	}
	if err := p.User.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.User), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:user: ", p), err)
	}
	return err
}

func (p *CreateArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CreateArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type CreateResult struct {
	Success *User                   `thrift:"success,0" json:"success,omitempty"`
	E       *types.ServiceException `thrift:"e,1" json:"e,omitempty"`
}

func NewCreateResult() *CreateResult {
	return &CreateResult{}
}

var CreateResult_Success_DEFAULT *User

func (p *CreateResult) GetSuccess() *User {
	if !p.IsSetSuccess() {
		return CreateResult_Success_DEFAULT
	}
	return p.Success
}

var CreateResult_E_DEFAULT *types.ServiceException

func (p *CreateResult) GetE() *types.ServiceException {
	if !p.IsSetE() {
		return CreateResult_E_DEFAULT
	}
	return p.E
}
func (p *CreateResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CreateResult) IsSetE() bool {
	return p.E != nil
}

func (p *CreateResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.ReadField0(iprot); err != nil {
				return err
			}
		case 1:
			if err := p.ReadField1(iprot); err != nil {
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
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *CreateResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &User{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *CreateResult) ReadField1(iprot thrift.TProtocol) error {
	p.E = &types.ServiceException{}
	if err := p.E.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *CreateResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("create_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *CreateResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *CreateResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin("e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *CreateResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CreateResult(%+v)", *p)
}
