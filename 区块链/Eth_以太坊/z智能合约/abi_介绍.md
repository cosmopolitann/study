# 深入以太坊智能合约Abi

智能合约 ABI

开发 DApp 时要调用在区块链上的 Ethereum 智能合约，就需要智能合约的 ABI。本文希望更多了解 ABI，如为什么需要 ABI？如何解读 Ethereum 的智能合约 ABI？以及如何取得智能的 ABI？

### ABI（Application Binary Interface）应用二进制接口

如果理解 API 就很容易了解 ABI。简单来说，API 是程序与程序间互动的接口。这个接口包含程序提供外界存取所需的 functions、variables 等。ABI 也是程序间互动的接口，但程序是被编译后的 binary code。所以同样的接口，但传递的是 binary 格式的信息。所以 ABI 就要描述如何 decode/encode 程序间传递的 binary 信息。编译和部署智能合约在 Ethereum 智能合约可以被大家使用前，必须先被部署到区块链上。

从智能合约的代码到使用智能合约，大概包含几个步骤：

1. 编写智能合约的代码（一般是用 Solidity 写）Remix 网址：http://remix.ethereum.org
2. 编译智能合约的代码变成可在 EVM 上执行的 bytecode（binary code）。同时可以通过编译取得智能合约的 ABI
3. 部署智能合约，实际上是把 bytecode 存储在链上（通过一个transaction），并取得一个专属于这个合约的地址
4. 如果要写个程序调用这个智能合约，就要把信息发送到这个合约的地址（一样的也是通过一个 transaction）。Ethereum 节点会根据输入的信息，选择要执行合约中的哪一个 function 和要输入的参数

而要如何知道這这个智能合约提供哪些 function 以及应该要传入什么样的参数呢？这些信息就是记录在智能合约的 ABI！

### Ethereum 智能合约 ABI

Ethereum 智能合约 ABI 用一个 array 表示，其中会包含数个用 JSON 格式表示的 Function 或 Event。根据最新的 Solidity 文件：

#### Function

共有 7 个参数：

1. `name`：a string，function 名称
2. `type`：a string，"function", "constructor", or "fallback"
3. `inputs`：an array，function 输入的参数，包含：
   - `name`：a string，参数名
   - `type`：a string，参数的 data type(e.g. uint256)
   - `components`：an array，如果输入的参数是 tuple(struct) type 才会有这个参数。描述 struct 中包含的参数类型
4. `outputs`：an array，function 的返回值，和 `inputs` 使用相同表示方式。如果沒有返回值可忽略，值为 `[]`
5. `payable`：`true`，function 是否可收 Ether，预设为 `false`
6. `constant`：`true`，function 是否会改写区块链状态，反之为 `false`
7. `stateMutability`：a string，其值可能为以下其中之一："pure"（不会读写区块链状态）、"view"（只读不写区块链状态）、"payable" and "nonpayable"（会改区块链状态，且如可收 Ether 为 "payable"，反之为 "nonpayable"）

仔细看会发现 `payable` 和 `constant` 这两个参数所描述的內容，似乎已包含在 `stateMutability` 中。

事实也确实是这样的，在 [Solidity v0.4.16](https://link.jianshu.com/?t=https%3A%2F%2Fgithub.com%2Fethereum%2Fsolidity%2Freleases) 中把 `constant` 这个修饰function 的 key words 分成： `view`（neither reads from nor writes to the state）和 `pure`（does not modify the state），并从 v0.4.17 开始 Type Checker 会强制检查。`constant` 改为只用来修饰不能被修改的 variable。并在 ABI 中加入 `stateMutability` 这个参数统一表示，`payable` 和 `constant` 目前保留是为了向后兼容。这个改动详细的內容和讨论可参考：
[https://github.com/ethereum/solidity/issues/992](https://link.jianshu.com/?t=https%3A%2F%2Fgithub.com%2Fethereum%2Fsolidity%2Fissues%2F992)

#### Event

共有 4 个参数：

1. `name`: a string，event 的名称
2. `type`: a string，always "event"
3. `inputs`: an array，输入参数，包含：
   - `name`: a string，参数名称
   - `type`: a string，参数的 data type(e.g. uint256)
   - `components`: an array，如果输入参数是 tuple(struct) type 才会有这个参数。描述 struct 中包含的信息类型
   - `indexed`: `true`，如果这个参数被定义为 indexed ，反之为 `false`
4. `anonymous`: `true`，如果 event 被定义为 anonymous

更新智能合约状态需要发送 transaction，transaction 需要等待验证，所以更新合约状态是非同步的，无法马上取得返回值。使用 Event 可以在状态更新成功后，将相关信息记录到 Log，并让监听这个 Event 的 DApp 或任何应用这个接口的程序收到通知。每笔 transaction 都有对应的 Log。

所以简单来说，Event 可用來：1. 取得 function 更新合约状态的返回值 2. 也可作为合约另外的存储空间。

Event 的参数分为：有 `indexed`，和其他没有 `indexed` 的。有 `indexed` 的参数可以使用 filter，例如同一个 Event，我可以选择只监听从特定 address 发出来的交易。每笔 Log 的信息同样分为两个部分：Topics（长度最多为 4 的 array） 和 Data。有 `indexed` 的参数会存储存在 Log 的 Topics，其他的存在 Data。如果定义为 `anonymous`，就不会产生以下示例中的 Topics[0]，其值为 Event signature 的 hash，作为這個 Event 的 ID。





Event

```
event Set(address indexed _from, uint value)
```

### 用一个简单的智能合约举个例子

这个智能合约包含：

- `data`：一个可修改的 state variable，会自动产生一个只能读取的 `data()` function
- `set()`：一个修改 `data` 值的 function
- `Set()`：一个在每次修写 `data` 时记录 Log 的 event

智能合约 Source Code：

```
pragma solidity ^0.4.20;



contract SimpleStorage {



    uint public data;



    event Set(address indexed _from, uint value);



    function set(uint x) public {



        data = x;



        Set(msg.sender, x);



    }



}
```

智能合约 ABI：

```
[{



        "constant": true,



        "inputs": [],



        "name": "data",



        "outputs": [{"name": "","type": "uint256"}],



        "payable": false,



        "stateMutabㄒility": "view",



        "type": "function"



    },



    {



        "anonymous": false,



        "inputs": [{"indexed": true,"name": "_from","type": "address"},{"indexed": false,"name": "value","type": "uint256"}],



        "name": "Set",



        "type": "event"



    },



    {



        "constant": false,



        "inputs": [{"name": "x","type": "uint256"}],



        "name": "set",



        "outputs": [],



        "payable": false,



        "stateMutability": "nonpayable",



        "type": "function"



}]
```

### 取得 Ethereum 智能合约 ABI

#### Solidity Compiler

可以用 Solidity Compiler 取得合约 ABI，我使用 JavaScript 版本的 Compiler 为例。

安装：

```
npm install solc -g
```

取得合约 ABI：

```
solcjs simpleStorage.sol --abi
```

会生成一个 simpleStorage_sol_SimpleStorage.abi 文件，里面就是合约ABI 內容。

也可以取得合约的 binary code：

```
solcjs your_contract.sol --bin
```

#### Remix

同样的使用 Solidity Compiler，也可以用 Remix。在合约的 Details 可以看到完整的 ABI。可以在 Settings 中指定 Compiler 版本。





Remix

#### Etherscan

许多知名合约会把合约 source code 放上 Etherscan 做验证，可以同时看到h 合约ABI。





Etherscan

另外 Etherscan 提供 [API](https://link.jianshu.com/?t=https%3A%2F%2Fetherscan.io%2Fapis%23contracts)，可用来取得经过验证的合约 ABI。