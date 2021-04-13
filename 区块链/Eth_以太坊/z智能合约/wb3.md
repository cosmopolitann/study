## evm/wasm虚拟机对比

### 1. 支持功能是否完善

- 1.1 对于wasm，以太坊2.0的虚拟机SSVM

| 序号 | 描述                                                         | 说明                                                         |
| :--: | :----------------------------------------------------------- | :----------------------------------------------------------- |
|  1   | Ewasm字节码                                                  | (1)要使用Ewasm字节码运行SSVM，需要提供有效的Ewasm字节码和calldata. (2)目前，SSVM不支持以太坊的ABI编码，因此用户需要自己编写calldata数据. |
|  2   | 配套合约编译工具[SOLL](https://github.com/second-state/soll) | (1) Solidity有80个测试合约，SOLL可通过27个，通过率34%. (2) libyul拥有499个测试合约，SOLL可通过393个，通过率79%. |
|  3   | 未实现的solidity语言功能                                     | (1) Library declaration 库的声明 (2) Modifier declaration 修饰符声明 (3) Multiple classes 多签 (4) Contract inheritance 合约的继承 (5) Constructor with parameter 带参数的构造函数 (6) Return multiple value 多值返回 (7) Inline assembly 在线组装 (8) Some special variable (like address.static_call) 一些特殊的变量（例如address.static_call) |
|  4   | SOLL编译工具开发阶段                                         | 开发早期阶段                                                 |
|  5   | [SSVM](https://github.com/second-state/SSVM/tree/0.5.1)是否需要部署在主网上 | [NewEVM](https://github.com/ethereum/go-ethereum/blob/release/1.9/core/vm/evm.go)，该[PR](https://github.com/ethereum/go-ethereum/pull/16957)当前处于暂停状态，因为目前尚无明显迹象表明强烈希望在eth1主网上部署ewasm. PR需要进行一些优化，但可以在专用网络上使用; 目前，以太坊并未合并它，因为这将迫使其在无法保证将要使用它的情况下继续维护它. 因此，它被搁置了 |

- 1.2 对于evm，ethermint

| 序号 | 描述             | 说明                               |
| :--: | :--------------- | :--------------------------------- |
|  1   | 实现Ethereum交易 | 在 cosmosSDK 里实现了Ethereum交易  |
|  2   | web3兼容         | 实现了web3兼容的API                |
|  3   | EVM实现          | EVM被分离为一个单独的cosmosSDK模块 |

### 2. 开源代码的完成度

- 2.1 ewasm

  ```
    项目：https://github.com/go-interpreter/wagon
  
    未经测试和优化的代码开发完成，但未合入主分支
  ```

- 2.2 evm

  项目：

  ```
    https://github.com/ChainSafe/ethermint
  
    https://github.com/ChainSafe/ethermint-deploy
  ```

```
    ethermint 处于积极开发阶段，API变动频繁，没有发布稳定版本。
    EVM 功能可用
    web3 API可用
```

### 3. 是否适应我们的场景

- 3.1 ewasm

  wagon 项目是对接 ethereum 的。对于目前我们基于 cosmosSDK 的项目来说，如果需要移植，开发时间和难度都会很高。并且 ethereum 项目自身目前也未使用 ewasm，所以与原生 evm 的对比也存在诸多不便

- 3.2 evm

  ethermint 是 cosmos 官方开源的基于 cosmosSDK 的 ethereum 克隆项目。已经实现了 EVM 和 ethereum 的交易。经测试，该项目可创建合约，部署合约。综合来说，移植成本较低。

### 4. 安全是否有保障

- 4.1 ewasm

  官方完全还未跑测试网，也未进行任何优化和安全测试。没有相应配套的安全审查工具。

- 4.2 evm

  处于积极开发阶段，安全性没有保障。官方表示现在未 alphine 版本，不建议上生产。

### 5. 开发周期预估

- 5.1 ewasm

  官方尚未完成，没有可用的用例。无法预估。

- 5.2 evm

  有可用的基于cosmosSDK的模块。且经测试，创建部署简单的合约是正常的。

  预估需要时间2周，需要人员2人。

  - 模块移植
  - 交易移植
  - web3 移植
  - 测试工作