## Deploy EthCrossChainManager

1. Deploy ./data/EthCrossChainData.sol
2. Deploy ./caller/CallerFactory
3. Set up constants at ./logic/Const/sol
4. Deploy ./logic/EthCrossChainManagerImplementation.sol
5. Deploy ./upragde/EthCrossChainManager ; set contract address at step 4 as _logic

## Deploy crossChain project

1. Deploy your logic contract
2. Deploy CallerProxy via CallerFactory , contruct with your logic contract address

