pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./../../../libs/math/SafeMath.sol";
import "./../../../libs/common/ZeroCopySource.sol";
import "./../../../libs/common/ZeroCopySink.sol";
import "./../../../libs/utils/Utils.sol";
import "./../upgrade/UpgradableECCM.sol";
import "./../libs/EthCrossChainUtils.sol";
import "./../interface/IEthCrossChainManager.sol";
import "./../interface/IEthCrossChainData.sol";
contract EthCrossChainManagerForRestartShort is UpgradableECCM {
    using SafeMath for uint256;

    event ChangeBookKeeperEvent(uint256 height, bytes rawHeader);
    constructor(
        address _eccd, 
        uint64 _chainId
    ) UpgradableECCM(_eccd,_chainId) public {}

    bool changed = false;
    uint32 constant newEpochStartHeight = 30321214; // mainnet
    bytes constant newBookerKeeperPubkeyList = hex"12050409d013d6adf0e56e439bde646ea9eff68467fa0a03a16d227d79f98c5a6725f64011909c67b64362ceaaf7032d177b453c0bf57a8ed8b07c17b00dc6525ec751120504288bdebfab545852b31638298c7100a9c26ad325f246b0939c661b9b112722c188f1611de3c1c4ed0c68bedaa0c2e6f771e306ad397c8fa571d44b2856fbaece120504cf256247dda995ab1ec68ef90865b5d78f6d805339b40e736e913493bb25446d159039a2a07c9b4ac01a9d1201defbd444226afe45a122265b14bd0c91b83c0f120504ecb8c0737073522c6f80dd8c7f1d8f0ec19c320010a27f9568bcdebefe8196d4898a17e9aeba611bfce20dc542c1152cea665fd426d3cbb5048bd9bbca23b1b3"; 
    function changeBookKeeper() public returns(bool) {
        require(!changed, "already changed bookkeeper");
        changed = true;
        IEthCrossChainData eccd = IEthCrossChainData(EthCrossChainDataAddress);
        
        (, address[] memory keepers) = ECCUtils.verifyPubkey(newBookerKeeperPubkeyList);
        
        require(eccd.putCurEpochStartHeight(newEpochStartHeight), "Save MC LatestHeight to Data contract failed!");
        require(eccd.putCurEpochConPubKeyBytes(ECCUtils.serializeKeepers(keepers)), "Save Poly chain book keepers bytes to Data contract failed!");
    
        return true;
    }
}