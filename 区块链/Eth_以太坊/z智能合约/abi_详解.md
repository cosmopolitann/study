ABI[1](https://me.tryblockchain.org/Solidity-abi-abstraction.html#fn1)是以太坊的一种合约间调用时的一个消息格式。类似Webservice里的SOAP协议一样；也就是定义操作函数签名，参数编码，返回结果编码等。



## 函数

### 基本设计思想

使用ABI协议时必须要求在编译时知道类型，也就是说不支持动态语言那样的声明的变量还会变类型的情况。由于协议假设合约在编译期间知道另一个合约的接口定义，所以协议内没有明确定义存的内容类型（协议非类型内省）。

所以这个协议并不支持合约接口是动态的，或者是仅在运行时才知道类型的情况。如果这些情况很重要，可以使用以太坊生态系统建立自己的基础设施来解决这个问题。

### 函数选择器

一个函数调用的前四个字节数据指定了要调用的函数签名。计算方式是使用函数签名的`keccak256`的哈希，取4个字节。

```text
bytes4(keccak256("foo(uint32,bool)"))
```

函数签名使用基本类型的典型格式（canonical expression）定义，如果有多个参数使用`,`隔开，要去掉表达式中的所有空格。

### 参数编码

由于前面的函数签名使用了四个字节，参数的数据将从第五个字节开始。参数的编码方式与返回值，事件的编码方式一致，后面一起介绍。

### 支持的类型

支持的类型可参考原文[2](https://me.tryblockchain.org/Solidity-abi-abstraction.html#fn2)。支持的类型里面有一些比较特殊的是动态内容的类型，比如`string`，需要存储的空间是不固定的。

### 编码方式

针对数组参数中的嵌套数组的优化：

- 访问一个参数属性需要的读取次数，在一个数组结构中最多是数组的深度，比如a_i[k][l][r]，最多4次。在之前的ABI协议版本中，在最差情况下读取次数会随着总的动态类型的参数量线性增长。
- 变量的值或数组的元素间不应该是隔开存储的，可支持重定位，比如使用相对地址来定位。

区分了`动态内容类型`和`固定大小的类型`。固定大小的类型按原位置存储到当前块。动态类型的数据则独立存储在其它数据块。

#### `动态内容类型`的定义

- bytes
- string
- T[] 某个类型的不定长数组
- T[k] 某个类型的定长数组

所有其它类型则称为`固定大小的类型`。

#### 长度函数的定义

`len(a)`是二进制字符串`a`的中的字节数。`len(a)`的结果类型是`uint256`。

我们定义`enc`，编码函数，是一个ABI类型值到二进制串的映射函数，也就是ABI类型到二进制字符串的映射函数。由此`len(enc(X))`的结果也将因为`X`是不是`动态内容类型`而有所不同（也就是说`动态内容类型`的编码方式稍有不同）。

#### 进一步定义

对于任何ABI的值，根据`X`的类型不同，我们递归定义enc(X)，如下：

- 对于`X`是任意类型的`T`和长度值`k`，`T[k]`。

```text
enc(X) = head(X[0]) ... head(x[k-1] tail(X[0]) ... tail(X[k-1])
```

对于`X[i]`，如果其为`固定大小的类型`，`head`函数定义为，`head(X[i]) = enc(X[i])`。`tail`函数定义为`tail(X[i]) = ""`。

对于`动态内容类型`:

```text
head(X[i]) = enc(len(head(X[0]) ... head(X[k-1]) tail(X[0]) ... tail(X[i-1]))) tail(X[i]) = enc(X[i])
```

而对于是动态长度的类型的`X[i]`，虽然其长度不确定，但`head(X[i])`所存值其实是非常明确的，头部中只是存的一个偏移值（offset），这是偏移是实际存储内容处相对`enc(X)`整个编码内容开始位置来定义的。

> 上面这个表达式看得有点云里雾里的，但如果没有理解错的话，`固定大小的类型`在head里就依次编码了，`动态内容类型`只在head里放了一个从开始到真正内容开始的偏移，在偏移处才真正放内容，内容如果是变长的，就用len(enc(X))函数计算一个值放在前面，标识这个值有多大的内容。

- T[] 其中`X`有`k`个元素。其中`k`被认为是`uint256`，所以`enc(k)`实际是编码一个`uint256`。

```text
enc(X) = enc(k) enc(X[1], ..., X[k])
```

它被以一个静态长度的数据来编码，但将数组所含元素的个长度作为前缀。

#### 具体类型的编码方式

具体编码方式由于细节太多，不完全保证翻译正确，如果要自己实现这样的细节，建议再仔细研究原文文档，下面翻译仅做参考。

- `bytes`，长度`k`，长度值`k`是`uint256`。

`enc(X) = enc(k) pad_right(X)`，先将长度按`uint256`编码，然后紧跟字节序列格式的`X`，再用零补足，来保证`len(enc(X))`是32字节（256位）的倍数。

- `string`

`enc(X) = enc(enc_UTF8(X))`，这里的utf-8编码被按字节解释及编码；所以后续涉及的长度都是指按字节算的，不是按字符计算。

- `uint<M>`：enc(X)是按大端序编码`X`，并在左侧高位补足0，使之为32字节的倍数。
- `address`：按`uint160`编码。
- `int<M>`: enc(X) 是`X`的按大端序值2的补码表示，如果是负数左侧用1补足，正数左侧用0被足，直到是32的倍数。
- `bool`：按`uint8`编码。1代表`true`，0代表`false`。
- fixedx: enc(X) is enc(X * 2**N) where X \* 2**N is interpreted as a int256.
- fixed: as in the fixed128x128 case
- ufixedx: enc(X) is enc(X * 2**N) where X \* 2**N is interpreted as a uint256.
- ufixed: as in the ufixed128x128 case
- `bytes<M>`: enc(X) 是将字节序列用0补足为32位。

所以对于任意的`X`，`len(enc(X))`都是32的倍数。

### 函数选择器和参数编码

总的来说，对函数的`f`的参数a_1， ...， a_n按以下方式编码：

```text
function_selector(f) enc([a_1，...，a_n])
```

`f`函数的对应的返回值v_1，...，v_k编码如下：

```text
enc([v_1, ..., v_k])
```

这里的`[a_1, ..., a_n]`和`[v_1, ..., v_k]`，是定长数组，长度分别是`n`和`k`。严格说来，`[a_1, ..., a_n]`是一个含不同类型元素的数组。但即便如此，编码仍然是明确的，因为实际上我们并没有使用这样一种类型`T`。

### 例子

```text
contract Foo {
  function bar(fixed[2] xy) {}
  function baz(uint32 x, bool y) returns (bool r) { r = x > 32 || y; }
  function sam(bytes name, bool z, uint[] data) {}
}
```

如果要调用`baz(69, true)`，要传的字节拆解如下：

- 0xcdcd77c0: 使用函数选择器确定的函数ID。通过`bytes4(keccak256("baz(uint32,bool)"))`。
- 0x0000000000000000000000000000000000000000000000000000000000000045。第一个参数，uint32位的值`69`，并补位到32字节。
- 

- 0x0000000000000000000000000000000000000000000000000000000000000001。第二值`boolean`类型值`true`。补位到32字节。

所以最终的串值为：

```text
0xcdcd77c000000000000000000000000000000000000000000000000000000000000000450000000000000000000000000000000000000000000000000000000000000001
```

返回结果是一个`bool`值，在这里，返回的是`false`。所以输出就是一个`bool`。

```text
0x0000000000000000000000000000000000000000000000000000000000000000
```

#### 动态类型的使用例子

如果我们要值用`(0x123, [0x456, 0x789], "1234567890", "Hello, world!")`调用函数`f(uint,uint32[],bytes10,bytes)`，编码拆解如下：

`bytes4(sha3("f(uint256,uint32[],bytes10,bytes)"))`计算MethodID值。对于`固定大小的类型`值`uint256`和`bytes10`，直接编码值。而对于`动态内容类型`值`uint32[]`和`bytes`，我们先编码偏移值，偏移值是整个值编码的开始到真正存这个数据的偏移值（这里不计算头四个用于表示函数签名的字节）。所以依次为：



- 0x0000000000000000000000000000000000000000000000000000000000000123，32字节的`0x123`。
- 0x0000000000000000000000000000000000000000000000000000000000000080 （第二个参数的由于是动态内容类型，所以这里存储偏移值，4*32 字节，刚好是头部部分的大小）
- 0x3132333435363738393000000000000000000000000000000000000000000000 （"1234567890" 在右侧补0到32字节大小）
- 0x00000000000000000000000000000000000000000000000000000000000000e0 （第四个参数的偏移 = 第一个动态参数的偏移值 + 第一个动态参数的大小 = 4*32 + 3*32 动态长度的计算见后）

尾部部分的第一个动态参数，`[0x456, 0x789]`编码拆解如下：

- 0x0000000000000000000000000000000000000000000000000000000000000002 （整个数组的长度，2）。
- 0x0000000000000000000000000000000000000000000000000000000000000456 （第一个元素）
- 

- 0x0000000000000000000000000000000000000000000000000000000000000789（第二个元素）

最后我们来看看第二个动态参数的的编码，`Hello, world!`。

- 0x000000000000000000000000000000000000000000000000000000000000000d (元素的字节长度，13)
- 0x48656c6c6f2c20776f726c642100000000000000000000000000000000000000 ("Hello, world!" 补位到32字节，里面是按ascii编码的，可以查查对应的编码。)

最终我们得到了下述的编码，为了清晰在函数签名的四个字节后，加了一个换行。

```text
0x8be65246
0000000000000000000000000000000000000000000000000000000000000123
0000000000000000000000000000000000000000000000000000000000000080
3132333435363738393000000000000000000000000000000000000000000000
00000000000000000000000000000000000000000000000000000000000000e0
0000000000000000000000000000000000000000000000000000000000000002
0000000000000000000000000000000000000000000000000000000000000456
0000000000000000000000000000000000000000000000000000000000000789
000000000000000000000000000000000000000000000000000000000000000d
48656c6c6f2c20776f726c642100000000000000000000000000000000000000
```

## Events 事件

`Events`是抽象出来的以太坊的日志，事件监听协议。`日志实体`包含合约的地址，一系列的最多可以达到4个的`Topic`，和任意长度的二进制数据内容。`Events`依赖ABI函数来解释，`日志实体`被当成为一个自定义数据结构。

事件有一个给定的事件名称，一系列的事件参数，我们将他们分为两个类别：需要索引的和不需要索引的。需要索引的，可以最多允许有三个，包含使用`Keccak hash`算法哈希过的事件签名，来组成现在`日志实体`的`Topic`。那些不需要索引的组成了`Events`的字节数组。

一个`日志实体`使用ABI描述如下：

- `address`: 合约的地址。（由以太坊内部提供）
- `topics[0]`: `keccak(EVENT_NAME+"("+EVENT_ARGS.map(canonical_type_of).join(",")+")")`，其中的`canonical_type_of`是返回函数的规范型(Canonical form)，如，`uint indexed foo`，返回的应该是`uint256`。如果事件本身是匿名定义的，那么`Topic[0]`将不会自动生成。
- `Topics[n]`: `EVENT_INDEXED_ARGS[n-1]`，其中的`EVENT_INDEXED_ARGS`表示指定成要索引的事件参数。
- 

- `data`: `abi_serialise(EVENT_NON_INDEXED_ARGS)`使用ABI协议序列化的没有指定为索引的其它的参数。`abi_serialise()`是ABI序列函数，用来返回一系列的函数定义的类型值。

## JSON格式

合约接口的JSON格式。包含一系列的函数和或事件的描述。一个描述函数的JSON包含下述的字段：

- `type`: 可取值有`function`，`constructor`，`fallback`（无名称的默认函数）

- ```
  inputs
  ```

  : 一系列的对象，每个对象包含下述属性：

  - `name`: 参数名称
  - `type`: 参数的`规范型(Canonical Type)`。

- `outputs`： 一系列的类似`inputs`的对象，如果无返回值时，可以省略。

- `constant`: `true`表示函数声明自己[不会改变区块链的状态](http://solidity.readthedocs.io/en/develop/contracts.html#constant-functions)。

- `payable`: `true`表示函数可以接收`ether`，否则表示不能。

其中`type`字段可以省略，默认是`function`类型。构造器函数和回退函数没有`name`或`outputs`。回退函数也没有`inputs`。

向不支持`payable`发送`ether`将会引发异常，禁止这么做。

事件用JSON描述时几乎是一样的：

- `type`: 总是`event`

- `name`: 事件的名称

- ```
  inputs
  ```

  : 一系列的对象，每个对象包含下述属性：

  - `name`: 参数名称
  - `type`: 参数的`规范型(Canonical Type)`。
  - `indexed`: `true`代表这个这段是日志主题的一部分，`false`代表是日志数据的一部分。

- `anonymous`: `true`代表事件是匿名声明的。

示例：

```text
contract Test {
    function Test() {
        b = 0x12345678901234567890123456789012;
    }
    event Event(uint indexed a, bytes32 b) event Event2(uint indexed a, bytes32 b) function foo(uint a) {
        Event(a, b);
    }
    bytes32 b;
}
```

上述代码的JSON格式如下：

```text
[
    {
        "type": "event",
        "inputs": [
            {
                "name": "a",
                "type": "uint256",
                "indexed": true
            },
            {
                "name": "b",
                "type": "bytes32",
                "indexed": false
            }
        ],
        "name": "Event"
    },
    {
        "type": "event",
        "inputs": [
            {
                "name": "a",
                "type": "uint256",
                "indexed": true
            },
            {
                "name": "b",
                "type": "bytes32",
                "indexed": false
            }
        ],
        "name": "Event2"
    },
    {
        "type": "event",
        "inputs": [
            {
                "name": "a",
                "type": "uint256",
                "indexed": true
            },
            {
                "name": "b",
                "type": "bytes32",
                "indexed": false
            }
        ],
        "name": "Event2"
    },
    {
        "type": "function",
        "inputs": [
            {
                "name": "a",
                "type": "uint256"
            }
        ],
        "name": "foo",
        "outputs": []
    }
]
```

在Javascript中的使用示例：

```text
var Test = eth.contract(
[
    {
        "type": "event",
        "inputs": [
            {
                "name": "a",
                "type": "uint256",
                "indexed": true
            },
            {
                "name": "b",
                "type": "bytes32",
                "indexed": false
            }
        ],
        "name": "Event"
    },
    {
        "type": "event",
        "inputs": [
            {
                "name": "a",
                "type": "uint256",
                "indexed": true
            },
            {
                "name": "b",
                "type": "bytes32",
                "indexed": false
            }
        ],
        "name": "Event2"
    },
    {
        "type": "function",
        "inputs": [
            {
                "name": "a",
                "type": "uint256"
            }
        ],
        "name": "foo",
        "outputs": []
    }
]);
var theTest = new Test(addrTest);

// examples of usage:
// every log entry ("event") coming from theTest (i.e. Event & Event2):
var f0 = eth.filter(theTest);
// just log entries ("events") of type "Event" coming from theTest:
var f1 = eth.filter(theTest.Event);
// also written as
var f1 = theTest.Event();
// just log entries ("events") of type "Event" and "Event2" coming from theTest:
var f2 = eth.filter([theTest.Event, theTest.Event2]);
// just log entries ("events") of type "Event" coming from theTest with indexed parameter 'a' equal to 69:
var f3 = eth.filter(theTest.Event, {'a': 69});
// also written as
var f3 = theTest.Event({'a': 69});
// just log entries ("events") of type "Event" coming from theTest with indexed parameter 'a' equal to 69 or 42:
var f4 = eth.filter(theTest.Event, {'a': [69, 42]});
// also written as
var f4 = theTest.Event({'a': [69, 42]});

// options may also be supplied as a second parameter with `earliest`, `latest`, `offset` and `max`, as defined for `eth.filter`.
var options = { 'max': 100 };
var f4 = theTest.Event({'a': [69, 42]}, options);

var trigger;
f4.watch(trigger);

// call foo to make an Event:
theTest.foo(69);

// would call trigger like:
//trigger(theTest.Event, {'a': 69, 'b': '0x12345678901234567890123456789012'}, n);
// where n is the block number that the event triggered in.
```

实现：

```text
// e.g. f4 would be similar to:
web3.eth.filter({'max': 100, 'address': theTest.address, 'topics': [ [69, 42] ]});
// except that the resultant data would need to be converted from the basic log entry format like:
{
  'address': theTest.address,
  'topics': [web3.sha3("Event(uint256,bytes32)"), 0x00...0045 /* 69 in hex format */],
  'data': '0x12345678901234567890123456789012',
  'number': n
}
// into data good for the trigger, specifically the three fields:
  Test.Event // derivable from the first topic
  {'a': 69, 'b': '0x12345678901234567890123456789012'} // derivable from the 'indexed' bool in the interface, the later 'topics' and the 'data'
  n // from the 'number'
```

事件结果：

```text
[ {
  'event': Test.Event,
  'args': {'a': 69, 'b': '0x12345678901234567890123456789012'},
  'number': n
  },
  { ...
  } ...
]
```