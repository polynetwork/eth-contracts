const { ethers } = require("hardhat");
const hre = require("hardhat");
const fs = require("fs");
const Web3 = require("web3");
const colors = require('colors');
hre.web3 = new Web3(hre.network.provider);

async function main() {
    await printWhiteListInfo("eth-main");
    await printWhiteListInfo("bsc-main");
    await printWhiteListInfo("heco-main");
    await printWhiteListInfo("ok-main");
    await printWhiteListInfo("polygon-main");
    await printWhiteListInfo("arbitrum-main");
}

async function printWhiteListInfo(networkName) {

  //let networkName = 'eth-main';

  let fromContractWhiteList;
  let contractMethodWhiteList;

  try {
    [fromContractWhiteList,contractMethodWhiteList] = getWhiteListInfo(networkName);
  } catch(error) {
    throw error;
  }
  
  const Encoder = await hre.ethers.getContractFactory("AbiEncoder");
  let encoder = await Encoder.deploy();
  for (let i=0;i<contractMethodWhiteList.length;i++) {
    contractMethodWhiteList[i] = await encoder.encodeWhiteList(contractMethodWhiteList[i][0], contractMethodWhiteList[i][1]);
  }
  console.log(networkName," whiteListInfo: ".magenta);
  console.log("fromContractWhiteList: ".green,fromContractWhiteList);
  console.log("contractMethodWhiteList: ".green,contractMethodWhiteList);

  console.log("\nDone.\n".magenta);

}

function getWhiteListInfo(networkName) {
    switch (networkName) {
        case 'eth-main' :
            fromContractWhiteList = [
                // lockproxy:
                '0x250e76987d838a75310c34bf422ea9f1AC4Cc906',
                // swapper:
                '0xaf83ce8d461e8834de03a3803c968615013c6b3d',
                // nftptoxy:
                '0x2cdfc90250EF967036838DA601099656e74bCfc5',
                // switcheo:
                '0x9a016ce184a22dbf6c17daa59eb7d3140dbd1c54'
                // flux: nil
            ];
            contractMethodWhiteList = [
                // lockproxy: (method: unlock)
                ['0x250e76987d838a75310c34bf422ea9f1AC4Cc906', ['0x756e6c6f636b']],
                // swapper: nil
                // nftptoxy: (method: unlock)
                ['0x2cdfc90250EF967036838DA601099656e74bCfc5', ['0x756e6c6f636b']],
                // switcheo: (method: addExtension , removeExtension , registerAsset , unlock)
                ['0x9a016ce184a22dbf6c17daa59eb7d3140dbd1c54', ['0x616464457874656e73696f6e','0x72656d6f7665457874656e73696f6e','0x72656769737465724173736574','0x756e6c6f636b']]
                // flux: nil
            ];
            return [fromContractWhiteList,contractMethodWhiteList];

        case 'bsc-main' :
            fromContractWhiteList = [
                // lockproxy:
                '0x2f7ac9436ba4B548f9582af91CA1Ef02cd2F1f03',
                // swapper:
                '0x00b93851e3135663AAeC351555EddEE5B01325e6',
                // nftptoxy:
                '0x2cdfc90250EF967036838DA601099656e74bCfc5',
                // switcheo:
                '0xb5d4f343412dc8efb6ff599d790074d0f1e8d430',
                // flux:
                "0x6B7C026621d53d092FB446C2162D1DFc048bC9D7",
                "0xf14a7d43bedaC1C17Da9C22cf3D8c900AA8DEcA9",
                "0x9800c0EB44cCb5423bb95205921fCb5FFa066a62",
                "0xa61E289C29c47A2EEA7092355602De8A0dCB16e2",
                "0x3803e2A56F657922b76211ab22eAFE6C0Aa4F95f",
                "0x303E57962D735A8dfC7CA5d30676d48A63AF7338",
                "0x5601D189D5751e4f9dde2E598999228B4a5DB9e2",
                "0x58EEac66d46d90Cc1c563CABc2c77D610B91284c",
                "0xd397a9F470212d419AcEC535783aeD8145E58853",
                "0xaeD9013944dFD09a16e5c4115ce45c48520383CE",
                "0xB8c221dc3EE9132015AE6816cEF428089498fB82",
                "0x6cE07781968c5f4747AcDd46D0975A797a8B1143",
                "0x54CF38E6B75859cD7958d3B7bc020a1b89482d7E",
                "0xc7d0d0E5bcfe73799EEcc1aD7204e9c52ea014eC",
                "0x19630C41EeD3d437348FF8A6ddcB9FAcF550AC41"
            ];
            contractMethodWhiteList = [
                // lockproxy: (method: unlock)
                ['0x2f7ac9436ba4B548f9582af91CA1Ef02cd2F1f03', ['0x756e6c6f636b']],
                // swapper: nil
                // nftptoxy: (method: unlock)
                ['0x2cdfc90250EF967036838DA601099656e74bCfc5', ['0x756e6c6f636b']],
                // switcheo:
                ['0xb5d4f343412dc8efb6ff599d790074d0f1e8d430', ['0x616464457874656e73696f6e','0x72656d6f7665457874656e73696f6e','0x72656769737465724173736574','0x756e6c6f636b']],
                // flux: (method : onCrossTransfer)
                ['0x6B7C026621d53d092FB446C2162D1DFc048bC9D7', ['0x6f6e43726f73735472616e73666572']],
                ['0xf14a7d43bedaC1C17Da9C22cf3D8c900AA8DEcA9', ['0x6f6e43726f73735472616e73666572']],
                ['0x9800c0EB44cCb5423bb95205921fCb5FFa066a62', ['0x6f6e43726f73735472616e73666572']],
                ['0xa61E289C29c47A2EEA7092355602De8A0dCB16e2', ['0x6f6e43726f73735472616e73666572']],
                ['0x3803e2A56F657922b76211ab22eAFE6C0Aa4F95f', ['0x6f6e43726f73735472616e73666572']],
                ['0x303E57962D735A8dfC7CA5d30676d48A63AF7338', ['0x6f6e43726f73735472616e73666572']],
                ['0x5601D189D5751e4f9dde2E598999228B4a5DB9e2', ['0x6f6e43726f73735472616e73666572']],
                ['0x58EEac66d46d90Cc1c563CABc2c77D610B91284c', ['0x6f6e43726f73735472616e73666572']],
                ['0xd397a9F470212d419AcEC535783aeD8145E58853', ['0x6f6e43726f73735472616e73666572']],
                ['0xaeD9013944dFD09a16e5c4115ce45c48520383CE', ['0x6f6e43726f73735472616e73666572']],
                ['0xB8c221dc3EE9132015AE6816cEF428089498fB82', ['0x6f6e43726f73735472616e73666572']],
                ['0x6cE07781968c5f4747AcDd46D0975A797a8B1143', ['0x6f6e43726f73735472616e73666572']],
                ['0x54CF38E6B75859cD7958d3B7bc020a1b89482d7E', ['0x6f6e43726f73735472616e73666572']],
                ['0xc7d0d0E5bcfe73799EEcc1aD7204e9c52ea014eC', ['0x6f6e43726f73735472616e73666572']],
                ['0x19630C41EeD3d437348FF8A6ddcB9FAcF550AC41', ['0x6f6e43726f73735472616e73666572']]
            ];
            return [fromContractWhiteList,contractMethodWhiteList];

        case 'heco-main' :
            fromContractWhiteList = [
                // lockproxy:
                '0x020c15e7d08A8Ec7D35bCf3AC3CCbF0BBf2704e6',
                // swapper:
                '0xD98Ee7Ca1B33e60C75E3cd9493c566fc857592c8',
                // nftptoxy:
                '0x2cdfc90250EF967036838DA601099656e74bCfc5',
                // switcheo: nil
                // flux:
                "0xA38a2038136a57B6B1019a50A91Ab66738905a8A",
                "0xfa8440025d7EE21925be5A4adfd4aC35A156e0ed",
                "0x8b963E30f67C9a6C3F88306e799Fe2bBC3626f2E",
                "0xdf42B52157c10aaD666886dB0AB8Bb99b6485ae6",
                "0xAfFb645d73A784Ab53D074488349fF4dC22D3435",
                "0x93D3B6713c08F50C576490c7C975E1D315927EF3",
                "0x6F5E4F841DAbEac07A9270cdD26421aA869A59Bb",
                "0x64840C731022E51b20Dc5aeA6180F1b2299b5De1",
                "0x5C3CD4846d07548020eD9F40a08Fd66D16f71fE7",
                "0xB6626Eb5BB5918D01593Ad8504E4513f4385b62b",
                "0x5902BDbB4FDD0d44e2871577eaA46Fd0b3424B3e",
                "0xb4F8FeBb7372F47bEa2B06B6F9D878093a3B014E",
                "0x22A3f421115FbF2eda8AbCb56CE63219481C2A6C",
                "0x37De929B55B811FEe16a9244a0De6aac008c3F3C",
                "0xC5293535d05A98f73f5e2Af9A804363FA39BE5f3"
            ];
            contractMethodWhiteList = [
                // lockproxy: (method: unlock)
                ['0x020c15e7d08A8Ec7D35bCf3AC3CCbF0BBf2704e6', ['0x756e6c6f636b']],
                // swapper: nil
                // nftptoxy: (method: unlock)
                ['0x2cdfc90250EF967036838DA601099656e74bCfc5', ['0x756e6c6f636b']],
                // switcheo: nil
                // flux: (method : onCrossTransfer)
                ['0xA38a2038136a57B6B1019a50A91Ab66738905a8A', ['0x6f6e43726f73735472616e73666572']],
                ['0xfa8440025d7EE21925be5A4adfd4aC35A156e0ed', ['0x6f6e43726f73735472616e73666572']],
                ['0x8b963E30f67C9a6C3F88306e799Fe2bBC3626f2E', ['0x6f6e43726f73735472616e73666572']],
                ['0xdf42B52157c10aaD666886dB0AB8Bb99b6485ae6', ['0x6f6e43726f73735472616e73666572']],
                ['0xAfFb645d73A784Ab53D074488349fF4dC22D3435', ['0x6f6e43726f73735472616e73666572']],
                ['0x93D3B6713c08F50C576490c7C975E1D315927EF3', ['0x6f6e43726f73735472616e73666572']],
                ['0x6F5E4F841DAbEac07A9270cdD26421aA869A59Bb', ['0x6f6e43726f73735472616e73666572']],
                ['0x64840C731022E51b20Dc5aeA6180F1b2299b5De1', ['0x6f6e43726f73735472616e73666572']],
                ['0x5C3CD4846d07548020eD9F40a08Fd66D16f71fE7', ['0x6f6e43726f73735472616e73666572']],
                ['0xB6626Eb5BB5918D01593Ad8504E4513f4385b62b', ['0x6f6e43726f73735472616e73666572']],
                ['0x5902BDbB4FDD0d44e2871577eaA46Fd0b3424B3e', ['0x6f6e43726f73735472616e73666572']],
                ['0xb4F8FeBb7372F47bEa2B06B6F9D878093a3B014E', ['0x6f6e43726f73735472616e73666572']],
                ['0x22A3f421115FbF2eda8AbCb56CE63219481C2A6C', ['0x6f6e43726f73735472616e73666572']],
                ['0x37De929B55B811FEe16a9244a0De6aac008c3F3C', ['0x6f6e43726f73735472616e73666572']],
                ['0xC5293535d05A98f73f5e2Af9A804363FA39BE5f3', ['0x6f6e43726f73735472616e73666572']]
            ];
            return [fromContractWhiteList,contractMethodWhiteList];

        case 'polygon-main' :
            fromContractWhiteList = [
                // lockproxy:
                '0x28ff66a1b95d7cacf8eded2e658f768f44841212',
                // swapper:
                '0xaC57280B3A657A2e8D1180493C519a476D208F61',
                // nftptoxy: nil
                // switcheo: nil
                // flux:
                "0xf14a7d43bedaC1C17Da9C22cf3D8c900AA8DEcA9",
                "0x9800c0EB44cCb5423bb95205921fCb5FFa066a62",
                "0xa61E289C29c47A2EEA7092355602De8A0dCB16e2",
                "0x3803e2A56F657922b76211ab22eAFE6C0Aa4F95f",
                "0x303E57962D735A8dfC7CA5d30676d48A63AF7338",
                "0x5601D189D5751e4f9dde2E598999228B4a5DB9e2",
                "0xaca1227d3526D78a7Df361dF3495Aa2A64042808",
                "0xdE8B716F43dcD94e4D4afE73D040825b2E990E6D",
                "0x35b161F6F1Aa5787045505bd2E8307195485725D",
                "0x2d7Fe908524180D17C058e6C5388738C1b30d198",
                "0x487Fd5C13f5764b798D4AecaEA79B2e0DF5d3952",
                "0x1B49485342C38C83f5376efF64800b52ee35787F",
                "0x6e4475B49f5b2d0f06D6B7218C80C2A588252889",
                "0x039DB2Fc6377b20A30ac9F05eb6477ce7461cB0b",
                "0x32cAf282bdb48147C157285E0eb715fbFa07826a"
            ];
            contractMethodWhiteList = [
                // lockproxy: (method: unlock)
                ['0x28ff66a1b95d7cacf8eded2e658f768f44841212', ['0x756e6c6f636b']],
                // swapper: nil
                // nftptoxy: nil
                // switcheo: nil
                // flux: (method : onCrossTransfer)
                ['0xf14a7d43bedaC1C17Da9C22cf3D8c900AA8DEcA9', ['0x6f6e43726f73735472616e73666572']],
                ['0x9800c0EB44cCb5423bb95205921fCb5FFa066a62', ['0x6f6e43726f73735472616e73666572']],
                ['0xa61E289C29c47A2EEA7092355602De8A0dCB16e2', ['0x6f6e43726f73735472616e73666572']],
                ['0x3803e2A56F657922b76211ab22eAFE6C0Aa4F95f', ['0x6f6e43726f73735472616e73666572']],
                ['0x303E57962D735A8dfC7CA5d30676d48A63AF7338', ['0x6f6e43726f73735472616e73666572']],
                ['0x5601D189D5751e4f9dde2E598999228B4a5DB9e2', ['0x6f6e43726f73735472616e73666572']],
                ['0xaca1227d3526D78a7Df361dF3495Aa2A64042808', ['0x6f6e43726f73735472616e73666572']],
                ['0xdE8B716F43dcD94e4D4afE73D040825b2E990E6D', ['0x6f6e43726f73735472616e73666572']],
                ['0x35b161F6F1Aa5787045505bd2E8307195485725D', ['0x6f6e43726f73735472616e73666572']],
                ['0x2d7Fe908524180D17C058e6C5388738C1b30d198', ['0x6f6e43726f73735472616e73666572']],
                ['0x487Fd5C13f5764b798D4AecaEA79B2e0DF5d3952', ['0x6f6e43726f73735472616e73666572']],
                ['0x1B49485342C38C83f5376efF64800b52ee35787F', ['0x6f6e43726f73735472616e73666572']],
                ['0x6e4475B49f5b2d0f06D6B7218C80C2A588252889', ['0x6f6e43726f73735472616e73666572']],
                ['0x039DB2Fc6377b20A30ac9F05eb6477ce7461cB0b', ['0x6f6e43726f73735472616e73666572']],
                ['0x32cAf282bdb48147C157285E0eb715fbFa07826a', ['0x6f6e43726f73735472616e73666572']]
            ];
            return [fromContractWhiteList,contractMethodWhiteList];

        case 'ok-main' :
            fromContractWhiteList = [
                // lockproxy:
                '0x9a3658864Aa2Ccc63FA61eAAD5e4f65fA490cA7D',
                // swapper: nil
                // nftptoxy: nil
                // switcheo: nil
                // flux:
                "0x6d981C8535af305915c7eC2FceFA126d6286c0F5",
                "0xA7FeD91A00fE5847cc4970a1122fA5bCf2700291",
                "0x8b0963Ac00776111160f9bAB0E1835719ce5ff6C",
                "0x1ee6dCa80F09E7d387DbFA8EB2Dea08e83FaD182",
                "0xBe7F5e9eC0b1911251D61Ac64860a315a5185A18",
                "0xA08584E87703Da3E454133080362FfBE545dEf11",
                "0xf2575b39ADF3fC2d9370862e17efDCC79D935c00",
                "0xaca1227d3526D78a7Df361dF3495Aa2A64042808",
                "0xdE8B716F43dcD94e4D4afE73D040825b2E990E6D",
                "0x35b161F6F1Aa5787045505bd2E8307195485725D",
                "0xf195EeC267130FFF260b107a26b94976f59e09e2",
                "0x353881c2afda8D07Db6c01f59D91Bef75570d4D2",
                "0xA86E17e1B40F251d50Ce64Da1f2F824Ff9520E44",
                "0xE5bbC1FaF1dD83251B2949c08dE53Ba86BFf4Fb8",
                "0x5CB3Ff30Ae417fddC9301b64C35A211f40f801C2"
            ];
            contractMethodWhiteList = [
                // lockproxy: (method: unlock)
                ['0x9a3658864Aa2Ccc63FA61eAAD5e4f65fA490cA7D', ['0x756e6c6f636b']],
                // swapper: nil
                // nftptoxy: nil
                // switcheo: nil
                // flux: (method : onCrossTransfer)
                ['0x6d981C8535af305915c7eC2FceFA126d6286c0F5', ['0x6f6e43726f73735472616e73666572']],
                ['0xA7FeD91A00fE5847cc4970a1122fA5bCf2700291', ['0x6f6e43726f73735472616e73666572']],
                ['0x8b0963Ac00776111160f9bAB0E1835719ce5ff6C', ['0x6f6e43726f73735472616e73666572']],
                ['0x1ee6dCa80F09E7d387DbFA8EB2Dea08e83FaD182', ['0x6f6e43726f73735472616e73666572']],
                ['0xBe7F5e9eC0b1911251D61Ac64860a315a5185A18', ['0x6f6e43726f73735472616e73666572']],
                ['0xA08584E87703Da3E454133080362FfBE545dEf11', ['0x6f6e43726f73735472616e73666572']],
                ['0xf2575b39ADF3fC2d9370862e17efDCC79D935c00', ['0x6f6e43726f73735472616e73666572']],
                ['0xaca1227d3526D78a7Df361dF3495Aa2A64042808', ['0x6f6e43726f73735472616e73666572']],
                ['0xdE8B716F43dcD94e4D4afE73D040825b2E990E6D', ['0x6f6e43726f73735472616e73666572']],
                ['0x35b161F6F1Aa5787045505bd2E8307195485725D', ['0x6f6e43726f73735472616e73666572']],
                ['0xf195EeC267130FFF260b107a26b94976f59e09e2', ['0x6f6e43726f73735472616e73666572']],
                ['0x353881c2afda8D07Db6c01f59D91Bef75570d4D2', ['0x6f6e43726f73735472616e73666572']],
                ['0xA86E17e1B40F251d50Ce64Da1f2F824Ff9520E44', ['0x6f6e43726f73735472616e73666572']],
                ['0xE5bbC1FaF1dD83251B2949c08dE53Ba86BFf4Fb8', ['0x6f6e43726f73735472616e73666572']],
                ['0x5CB3Ff30Ae417fddC9301b64C35A211f40f801C2', ['0x6f6e43726f73735472616e73666572']]
            ];
            return [fromContractWhiteList,contractMethodWhiteList];
        case 'arbitrum-main' :
            fromContractWhiteList = [
                // lockproxy:
                '0x2f7ac9436ba4B548f9582af91CA1Ef02cd2F1f03',
                // swapper:
                '0x7E418a9926c8D1cbd09CC93E8051cC3BbdfE3854'
                // nftptoxy: nil
                // switcheo: nil
                // flux: nil
            ];
            contractMethodWhiteList = [
                // lockproxy: (method: unlock)
                ['0x2f7ac9436ba4B548f9582af91CA1Ef02cd2F1f03', ['0x756e6c6f636b']]
                // swapper: nil
                // nftptoxy: nil
                // switcheo: nil
                // flux: nil
            ];
            return [fromContractWhiteList,contractMethodWhiteList];
        default :
            throw new Error("fail to get WhiteListInfo, unknown Network: "+networkName);
    }
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });