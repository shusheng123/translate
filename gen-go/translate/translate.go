// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package translate

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

// Attributes:
//  - Code
//  - Msg
type ServerError struct {
  Code string `thrift:"code,1" db:"code" json:"code"`
  Msg string `thrift:"msg,2" db:"msg" json:"msg"`
}

func NewServerError() *ServerError {
  return &ServerError{}
}


func (p *ServerError) GetCode() string {
  return p.Code
}

func (p *ServerError) GetMsg() string {
  return p.Msg
}
func (p *ServerError) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
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

func (p *ServerError)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Code = v
}
  return nil
}

func (p *ServerError)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Msg = v
}
  return nil
}

func (p *ServerError) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("ServerError"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *ServerError) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("code", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:code: ", p), err) }
  if err := oprot.WriteString(string(p.Code)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.code (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:code: ", p), err) }
  return err
}

func (p *ServerError) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("msg", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:msg: ", p), err) }
  if err := oprot.WriteString(string(p.Msg)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.msg (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:msg: ", p), err) }
  return err
}

func (p *ServerError) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ServerError(%+v)", *p)
}

func (p *ServerError) Error() string {
  return p.String()
}

type Translate interface {
  // Parameters:
  //  - SrcString
  Translate(src_string string) (r string, err error)
}

type TranslateClient struct {
  Transport thrift.TTransport
  ProtocolFactory thrift.TProtocolFactory
  InputProtocol thrift.TProtocol
  OutputProtocol thrift.TProtocol
  SeqId int32
}

func NewTranslateClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *TranslateClient {
  return &TranslateClient{Transport: t,
    ProtocolFactory: f,
    InputProtocol: f.GetProtocol(t),
    OutputProtocol: f.GetProtocol(t),
    SeqId: 0,
  }
}

func NewTranslateClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *TranslateClient {
  return &TranslateClient{Transport: t,
    ProtocolFactory: nil,
    InputProtocol: iprot,
    OutputProtocol: oprot,
    SeqId: 0,
  }
}

// Parameters:
//  - SrcString
func (p *TranslateClient) Translate(src_string string) (r string, err error) {
  if err = p.sendTranslate(src_string); err != nil { return }
  return p.recvTranslate()
}

func (p *TranslateClient) sendTranslate(src_string string)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("translate", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := TranslateTranslateArgs{
  SrcString : src_string,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *TranslateClient) recvTranslate() (value string, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "translate" {
    err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "translate failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "translate failed: out of sequence response")
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
    err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "translate failed: invalid message type")
    return
  }
  result := TranslateTranslateResult{}
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


type TranslateProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler Translate
}

func (p *TranslateProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *TranslateProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *TranslateProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewTranslateProcessor(handler Translate) *TranslateProcessor {

  self2 := &TranslateProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self2.processorMap["translate"] = &translateProcessorTranslate{handler:handler}
return self2
}

func (p *TranslateProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err := iprot.ReadMessageBegin()
  if err != nil { return false, err }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(seqId, iprot, oprot)
  }
  iprot.Skip(thrift.STRUCT)
  iprot.ReadMessageEnd()
  x3 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
  x3.Write(oprot)
  oprot.WriteMessageEnd()
  oprot.Flush()
  return false, x3

}

type translateProcessorTranslate struct {
  handler Translate
}

func (p *translateProcessorTranslate) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := TranslateTranslateArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("translate", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return false, err
  }

  iprot.ReadMessageEnd()
  result := TranslateTranslateResult{}
var retval string
  var err2 error
  if retval, err2 = p.handler.Translate(args.SrcString); err2 != nil {
  switch v := err2.(type) {
    case *ServerError:
  result.E = v
    default:
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing translate: " + err2.Error())
    oprot.WriteMessageBegin("translate", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return true, err2
  }
  } else {
    result.Success = &retval
}
  if err2 = oprot.WriteMessageBegin("translate", thrift.REPLY, seqId); err2 != nil {
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
//  - SrcString
type TranslateTranslateArgs struct {
  SrcString string `thrift:"src_string,1" db:"src_string" json:"src_string"`
}

func NewTranslateTranslateArgs() *TranslateTranslateArgs {
  return &TranslateTranslateArgs{}
}


func (p *TranslateTranslateArgs) GetSrcString() string {
  return p.SrcString
}
func (p *TranslateTranslateArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
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

func (p *TranslateTranslateArgs)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.SrcString = v
}
  return nil
}

func (p *TranslateTranslateArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("translate_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TranslateTranslateArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("src_string", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:src_string: ", p), err) }
  if err := oprot.WriteString(string(p.SrcString)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.src_string (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:src_string: ", p), err) }
  return err
}

func (p *TranslateTranslateArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TranslateTranslateArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type TranslateTranslateResult struct {
  Success *string `thrift:"success,0" db:"success" json:"success,omitempty"`
  E *ServerError `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewTranslateTranslateResult() *TranslateTranslateResult {
  return &TranslateTranslateResult{}
}

var TranslateTranslateResult_Success_DEFAULT string
func (p *TranslateTranslateResult) GetSuccess() string {
  if !p.IsSetSuccess() {
    return TranslateTranslateResult_Success_DEFAULT
  }
return *p.Success
}
var TranslateTranslateResult_E_DEFAULT *ServerError
func (p *TranslateTranslateResult) GetE() *ServerError {
  if !p.IsSetE() {
    return TranslateTranslateResult_E_DEFAULT
  }
return p.E
}
func (p *TranslateTranslateResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *TranslateTranslateResult) IsSetE() bool {
  return p.E != nil
}

func (p *TranslateTranslateResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
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

func (p *TranslateTranslateResult)  ReadField0(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  p.Success = &v
}
  return nil
}

func (p *TranslateTranslateResult)  ReadField1(iprot thrift.TProtocol) error {
  p.E = &ServerError{}
  if err := p.E.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
  }
  return nil
}

func (p *TranslateTranslateResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("translate_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TranslateTranslateResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteString(string(*p.Success)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *TranslateTranslateResult) writeField1(oprot thrift.TProtocol) (err error) {
  if p.IsSetE() {
    if err := oprot.WriteFieldBegin("e", thrift.STRUCT, 1); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err) }
    if err := p.E.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err) }
  }
  return err
}

func (p *TranslateTranslateResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TranslateTranslateResult(%+v)", *p)
}

