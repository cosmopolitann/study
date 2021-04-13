```go
//版权所有2015年的以太坊作者
//此文件是go-ethereum库的一部分。
//
// go-ethereum库是免费软件：您可以重新分发和/或修改
//它是根据由GNU发布的GNU通用通用公共许可证的条款
//自由软件基金会（许可的第3版）或
//（根据您的选择）任何更高版本。
//
//分发以太坊库，希望它会有用，
//但没有任何保证；甚至没有默示担保
//特定用途的适销性或适用性。看到
//有关更多详细信息，请参见GNU次通用公共许可证。
//
//您应该已经收到了GNU次通用公共许可证的副本
//以及go-ethereum库。如果不是，请参见<http://www.gnu.org/licenses/>。
package abi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/ethereum/go-ethereum/common"
)

// ABI保留有关合同上下文和可用信息的信息
//可调用的方法。它将允许您键入检查函数调用和
//相应地打包数据。
type ABI struct {
	Constructor Method
	Methods     map[string]Method
	Events      map[string]Event

	//固体v0.6.0中引入的其他“特殊”功能。
	//它与原始的默认后备广告分开。每份合约
	//只能定义一个后备和接收函数。
	//后备 方法 注意，它也用于表示v0.6.0之前的旧版后备
	//接收  方式
	Fallback Method // Note it's also used to represent legacy fallback before v0.6.0
	Receive  Method
}

// JSON返回已解析的ABI接口，如果失败则返回错误。
func JSON(reader io.Reader) (ABI, error) {
	dec := json.NewDecoder(reader)

	var abi ABI
	if err := dec.Decode(&abi); err != nil {
		return ABI{}, err
	}
	return abi, nil
}

//打包给定的方法名称以符合ABI。方法调用的数据
//将由method_id，args0，arg1，... argN组成。方法ID由
// 4个字节，参数全为32个字节。
//方法ID是从的哈希值的前4个字节创建的
//方法字符串签名。（签名= baz（uint32，string32））
func (abi ABI) Pack(name string, args ...interface{}) ([]byte, error) {
	// Fetch the ABI of the requested method
	if name == "" {
		// constructor
		arguments, err := abi.Constructor.Inputs.Pack(args...)
		if err != nil {
			return nil, err
		}
		return arguments, nil
	}
	method, exist := abi.Methods[name]
	if !exist {
		return nil, fmt.Errorf("method '%s' not found", name)
	}
	arguments, err := method.Inputs.Pack(args...)
	if err != nil {
		return nil, err
	}
	//如果不是构造函数，也要打包方法ID并返回
	return append(method.ID, arguments...), nil
}

//根据abi规范在v中解压缩输出
	//由于合同和事件之间无法发生命名冲突，
	//我们需要确定是要调用方法还是事件
func (abi ABI) Unpack(v interface{}, name string, data []byte) (err error) {
	if method, ok := abi.Methods[name]; ok {
		if len(data)%32 != 0 {
			return fmt.Errorf("abi: improperly formatted output: %s - Bytes: [%+v]", string(data), data)
		}
		return method.Outputs.Unpack(v, data)
	}
	if event, ok := abi.Events[name]; ok {
		return event.Inputs.Unpack(v, data)
	}
	return fmt.Errorf("abi: could not locate named method or event")
}

// UnpackIntoMap将日志解压缩到提供的map [string] interface {}中
func (abi ABI) UnpackIntoMap(v map[string]interface{}, name string, data []byte) (err error) {
	// since there can't be naming collisions with contracts and events,
	// we need to decide whether we're calling a method or an event
	if method, ok := abi.Methods[name]; ok {
		if len(data)%32 != 0 {
			return fmt.Errorf("abi: improperly formatted output")
		}
		return method.Outputs.UnpackIntoMap(v, data)
	}
	if event, ok := abi.Events[name]; ok {
		return event.Inputs.UnpackIntoMap(v, data)
	}
	return fmt.Errorf("abi: could not locate named method or event")
}

// UnmarshalJSON实现json.Unmarshaler接口
func (abi *ABI) UnmarshalJSON(data []byte) error {
	var fields []struct {
		Type    string
		Name    string
		Inputs  []Argument
		Outputs []Argument

		// Status indicator which can be: "pure", "view",
		// "nonpayable" or "payable".
		StateMutability string

		// Deprecated Status indicators, but removed in v0.6.0.
		Constant bool // True if function is either pure or view
		Payable  bool // True if function is payable

		// Event relevant indicator represents the event is
		// declared as anonymous.
		Anonymous bool
	}
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	abi.Methods = make(map[string]Method)
	abi.Events = make(map[string]Event)
	for _, field := range fields {
		switch field.Type {
		case "constructor":
			abi.Constructor = NewMethod("", "", Constructor, field.StateMutability, field.Constant, field.Payable, field.Inputs, nil)
		case "function":
			name := abi.overloadedMethodName(field.Name)
			abi.Methods[name] = NewMethod(name, field.Name, Function, field.StateMutability, field.Constant, field.Payable, field.Inputs, field.Outputs)
		case "fallback":
			// New introduced function type in v0.6.0, check more detail
			// here https://solidity.readthedocs.io/en/v0.6.0/contracts.html#fallback-function
			if abi.HasFallback() {
				return errors.New("only single fallback is allowed")
			}
			abi.Fallback = NewMethod("", "", Fallback, field.StateMutability, field.Constant, field.Payable, nil, nil)
		case "receive":
			// New introduced function type in v0.6.0, check more detail
			// here https://solidity.readthedocs.io/en/v0.6.0/contracts.html#fallback-function
			if abi.HasReceive() {
				return errors.New("only single receive is allowed")
			}
			if field.StateMutability != "payable" {
				return errors.New("the statemutability of receive can only be payable")
			}
			abi.Receive = NewMethod("", "", Receive, field.StateMutability, field.Constant, field.Payable, nil, nil)
		case "event":
			name := abi.overloadedEventName(field.Name)
			abi.Events[name] = NewEvent(name, field.Name, field.Anonymous, field.Inputs)
		default:
			return fmt.Errorf("abi: could not recognize type %v of field %v", field.Type, field.Name)
		}
	}
	return nil
}

//重载方法名称返回给定函数的下一个可用名称。
//因为坚固性允许函数重载，所以需要。
//
//例如，如果abi包含方法send，send1
//重载方法名称将为输入send返回send2。
func (abi *ABI) overloadedMethodName(rawName string) string {
	name := rawName
	_, ok := abi.Methods[name]
	for idx := 0; ok; idx++ {
		name = fmt.Sprintf("%s%d", rawName, idx)
		_, ok = abi.Methods[name]
	}
	return name
}

//重载的事件名称返回给定事件的下一个可用名称。
//因为坚固性允许事件重载，所以需要。
//
//例如，如果abi包含接收到的事件，则receive1
//重载的事件名称将为接收到的输入返回received2。
func (abi *ABI) overloadedEventName(rawName string) string {
	name := rawName
	_, ok := abi.Events[name]
	for idx := 0; ok; idx++ {
		name = fmt.Sprintf("%s%d", rawName, idx)
		_, ok = abi.Events[name]
	}
	return name
}

// MethodById通过4字节ID查找方法
//如果找不到则返回nil
func (abi *ABI) MethodById(sigdata []byte) (*Method, error) {
	if len(sigdata) < 4 {
		return nil, fmt.Errorf("data too short (%d bytes) for abi method lookup", len(sigdata))
	}
	for _, method := range abi.Methods {
		if bytes.Equal(method.ID, sigdata[:4]) {
			return &method, nil
		}
	}
	return nil, fmt.Errorf("no method with id: %#x", sigdata[:4])
}

// EventByID通过其主题哈希在
// ABI，如果找不到，则返回nil。
func (abi *ABI) EventByID(topic common.Hash) (*Event, error) {
	for _, event := range abi.Events {
		if bytes.Equal(event.ID.Bytes(), topic.Bytes()) {
			return &event, nil
		}
	}
	return nil, fmt.Errorf("no event with id: %#x", topic.Hex())
}

// HasFallback返回一个指示符，该指示符是否包含回退功能。
func (abi *ABI) HasFallback() bool {
	return abi.Fallback.Type == Fallback
}

// HasReceive返回一个指示符，该指示符是否包含接收函数。
func (abi *ABI) HasReceive() bool {
	return abi.Receive.Type == Receive
}
```

