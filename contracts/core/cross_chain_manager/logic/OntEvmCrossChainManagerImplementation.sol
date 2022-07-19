pragma solidity ^0.5.0;

import "./Const.sol";
import "../interface/IEthCrossChainData.sol";
import "../interface/IEventWitness.sol";
import "../data/EthCrossChainData.sol";
import "../caller/CallerFactory.sol";
import "../caller/EthCrossChainCaller.sol"; 
import "../caller/EthCrossChainCaller.sol"; 
import "../libs/ECCUtils/EthCrossChainUtils.sol";
import "./../../../libs/math/SafeMath.sol";

contract OntEvmCrossChainManagerImplementation is Const, OntConst {
    using SafeMath for uint256;

    event InitGenesisBlockEvent(uint256 height, bytes rawHeader);
    event ChangeEpochEvent(uint256 height, bytes rawHeader, address[] oldValidators, address[] newValidators);
    event CrossChainEvent(address indexed sender, bytes txId, address proxyOrAssetContract, uint64 toChainId, bytes toContract, bytes rawdata);
    event VerifyHeaderAndExecuteTxEvent(uint64 fromChainID, bytes toContract, bytes crossChainTxHash, bytes fromChainTxHash);
    
    // see in Const.sol
    // address constant EthCrossChainDataAddress = 0x0000000000000000000000000000000000000000;
    // address constant EthCrossChainCallerFactoryAddress = 0x0000000000000000000000000000000000000000; 
    // bytes constant ZionCrossChainManagerAddress = "0x000000000000000000000000000000000000000000";
    // uint constant chainId = 0;

    // see in OntConst.sol
    // address constant EVENT_WITNESS = 0xA8a1c2E2739725a14072B0bB1C6FAb0B36C15952; // testnet
    // address constant EVENT_WITNESS = 0x2b1143484bf5097A29678FD9592f75FE4639CA08; // mainnet
    
    function getZionChainId() public pure returns(uint) {
        return chainId;
    }    
    function getEthCrossChainDataAddress() public pure returns(address) {
        return EthCrossChainDataAddress;
    }    
    function getEthCrossChainManager() public view returns(address) {
        return address(this);
    }
    function getEthCrossChainCallerFactoryAddress() public pure returns(address) {
        return EthCrossChainCallerFactoryAddress;
    }
    
    function initGenesisBlock(
        bytes memory rawHeader
    ) public returns(bool) {
        ECCUtils.Header memory header = ECCUtils.decodeHeader(rawHeader);
        IEthCrossChainData eccd = IEthCrossChainData(EthCrossChainDataAddress);
        
        // verify isInit
        require(eccd.getCurEpochValidatorPkBytes().length == 0, "EthCrossChainData contract has already been initialized!");
        
        // get validators
        address[] memory validators = ECCUtils.getHeaderValidators(rawHeader);
        require(validators.length != 0, "Given block header does not contain any validator");

        // put epoch information
        require(eccd.putCurEpochStartHeight(uint64(header.number + 1)), "Save Zion current epoch start height to Data contract failed!");
        require(eccd.putCurEpochValidatorPkBytes(ECCUtils.encodeValidators(validators)), "Save Zion current epoch validators to Data contract failed!");
        
        emit InitGenesisBlockEvent(header.number, rawHeader);
        return true;
    }
    
    function changeEpoch(
        bytes memory rawHeader, 
        bytes memory rawSeals
    ) public returns(bool) {
        ECCUtils.Header memory header = ECCUtils.decodeHeader(rawHeader);
        IEthCrossChainData eccd = IEthCrossChainData(EthCrossChainDataAddress);
        
        // verify block.height
        require(header.number>=eccd.getCurEpochStartHeight(), "Given block height is lower than current epoch start height");
        
        // verify header
        bytes memory curPkBytes = eccd.getCurEpochValidatorPkBytes();
        address[] memory validators = ECCUtils.decodeValidators(curPkBytes);
        address[] memory newValidators = ECCUtils.getHeaderValidators(rawHeader);
        require(newValidators.length != 0, "Given block header does not contain any validator");
        require(ECCUtils.verifyHeader(keccak256(rawHeader), rawSeals, validators), "Verify header failed");
        
        // put new epoch info
        require(eccd.putCurEpochStartHeight(uint64(header.number + 1)), "Save Zion next epoch height to Data contract failed!");
        require(eccd.putCurEpochValidatorPkBytes(ECCUtils.encodeValidators(newValidators)), "Save Zion next epoch validators to Data contract failed!");
        
        emit ChangeEpochEvent(header.number, rawHeader, validators, newValidators);
        return true;
    }
    
    function crossChain(
        uint64 toChainId, 
        bytes calldata toContract, 
        bytes calldata method, 
        bytes calldata txData
    ) external returns (bool) {
        require(CallerFactory(EthCrossChainCallerFactoryAddress).isChild(msg.sender), "The caller is child of the caller factory!");
        uint256 txHashIndex = IEthCrossChainData(EthCrossChainDataAddress).getEthTxHashIndex();
        bytes memory paramTxHash = ECCUtils.uint256ToBytes(txHashIndex);
        bytes memory crossChainId = abi.encodePacked(sha256(abi.encodePacked(address(this), paramTxHash)));
        bytes memory rawParam = 
        ECCUtils.encodeTxParam(
            paramTxHash,
            crossChainId,
            ECCUtils.addressToBytes(msg.sender),
            toChainId,
            toContract,
            method,
            txData
        );
        
        require(IEthCrossChainData(EthCrossChainDataAddress).putEthTxHash(keccak256(rawParam)), "Save ethTxHash by index to Data contract failed!");
        
        emit CrossChainEvent(tx.origin, paramTxHash, msg.sender, toChainId, toContract, rawParam);

        // call EventWitness for ont-evm
        IEventWitness(EVENT_WITNESS).witness(rawParam);

        return true;
    }

    function verifyHeaderAndExecuteTx(
        bytes memory rawHeader,
        bytes memory rawSeals,
        bytes memory accountProof, 
        bytes memory storageProof,
        bytes memory rawCrossTx
    ) public returns (bool)
    {
        ECCUtils.Header memory header = ECCUtils.decodeHeader(rawHeader);
        ECCUtils.CrossTx memory crossTx = ECCUtils.decodeCrossTx(rawCrossTx);
        IEthCrossChainData eccd = IEthCrossChainData(EthCrossChainDataAddress);
        
        address[] memory validators = ECCUtils.decodeValidators(eccd.getCurEpochValidatorPkBytes());
        
        // verify block.height
        require(header.number>=eccd.getCurEpochStartHeight(), "Invalid block height");
        
        // verify header
        require(ECCUtils.verifyHeader(keccak256(rawHeader), rawSeals, validators), "Verify header failed");
        
        // verify proof
        bytes memory storageIndex = ECCUtils.getCrossTxStorageSlot(crossTx);
        bytes memory storageValue = ECCUtils.verifyAccountProof(accountProof, header.root, ZionCrossChainManagerAddress, storageProof, storageIndex);
        require(ECCUtils.bytesToBytes32(storageValue) == keccak256(rawCrossTx), "Verify proof failed");
        
        // check & put tx exection information
        require(!eccd.checkIfFromChainTxExist(crossTx.fromChainID, ECCUtils.bytesToBytes32(crossTx.txHash)), "the transaction has been executed!");
        require(eccd.markFromChainTxExist(crossTx.fromChainID, ECCUtils.bytesToBytes32(crossTx.txHash)), "Save crosschain tx exist failed!");
        require(crossTx.crossTxParam.toChainId == chainId, "This Tx is not aiming at this network!");

        address toContract = ECCUtils.bytesToAddress(crossTx.crossTxParam.toContract);
        
        require(_executeCrossChainTx(toContract, crossTx.crossTxParam.method, crossTx.crossTxParam.args, crossTx.crossTxParam.fromContract, crossTx.fromChainID), "Execute CrossChain Tx failed!");

        emit VerifyHeaderAndExecuteTxEvent(crossTx.fromChainID, crossTx.crossTxParam.toContract, crossTx.txHash, crossTx.crossTxParam.txHash);

        return true;
    }

    function _executeCrossChainTx(
        address _toContract, bytes memory _method, bytes memory _args, bytes memory _fromContractAddr, uint64 _fromChainId
    ) internal returns (bool)
    {   
        // verify to contract valid
        require(CallerFactory(EthCrossChainCallerFactoryAddress).isChild(_toContract), "The passed in address is not from the factory!");
        require(_toContract!=EthCrossChainDataAddress, "Don't try to call eccd!");
        
        (bool success, bytes memory returnData) = _toContract.call(abi.encodePacked(bytes4(keccak256(abi.encodePacked(_method, "(bytes,bytes,uint64)"))), abi.encode(_args, _fromContractAddr, _fromChainId)));
        
        require(success == true, "EthCrossChain call business contract failed");
        
        require(returnData.length != 0, "No return value from business contract!");
        bool res = abi.decode(returnData, (bool));
        require(res == true, "EthCrossChain call business contract return is not true");
        
        return true;
    }
    
    function fallback() public payable {
        revert("Unsupported function");
    }
}