const ZeroCopySourceMock = artifacts.require('./contractmocks/libs/common/ZeroCopySourceMock.sol');
const { expectRevert, expectEvent, constants, BN } = require('@openzeppelin/test-helpers');
const { expect } = require('chai');
const ontsdk = require("ontology-ts-sdk");
const ontUtils = ontsdk.utils;

contract('ZeroCopySource', () => {

    before(async function () {
        this.zeroCopySource = await ZeroCopySourceMock.new();
    });

    describe('ZeroCopySource Test', function () {

        const rawHeader = '0x010000000400000000000000386916e6b10e902a0bec78be62532262148c4a5e3952db4bfed00d27e5c779c30000000000000000000000000000000000000000000000000000000000000000e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855ae0798d0ecaed2b778eddebf18f071a561c53658c05e76cedecc27cafbdbc57743bab65d01000000a1e500aac88dbc1efd0c017b226c6561646572223a362c227672665f76616c7565223a22424543366e384b56653635556f6b474d79766e656a6f50763167586e46684a56776b35673141384d484b4b646768486e5478486e43497466484161706f686c5535666c314d79576a524d51717245586752556d38564f593d222c227672665f70726f6f66223a2246456a4d5266564e6c336a5937396650735130474d4e3261432b36554c487768742f4f697a6e4a424d7a394d634d64412b6b7232632f7852533354687979623466436366715a35766678684f4d6c474d5036417472413d3d222c226c6173745f636f6e6669675f626c6f636b5f6e756d223a302c226e65775f636861696e5f636f6e666967223a6e756c6c7d0000000000000000000000000000000000000000'

        const version = 1;
        const chainId = 4;
        const prevBlockHash = '0x386916e6b10e902a0bec78be62532262148c4a5e3952db4bfed00d27e5c779c3';
        const transactionsRoot = '0x0000000000000000000000000000000000000000000000000000000000000000';
        const crossStatesRoot = '0xe3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855';
        const blockRoot = '0xae0798d0ecaed2b778eddebf18f071a561c53658c05e76cedecc27cafbdbc577';
        const timestamp = 1572256323;
        const height = 1;
        const consensusData = 2214801009744602529;
        const consensusPayload = '0x7b226c6561646572223a362c227672665f76616c7565223a22424543366e384b56653635556f6b474d79766e656a6f50763167586e46684a56776b35673141384d484b4b646768486e5478486e43497466484161706f686c5535666c314d79576a524d51717245586752556d38564f593d222c227672665f70726f6f66223a2246456a4d5266564e6c336a5937396650735130474d4e3261432b36554c487768742f4f697a6e4a424d7a394d634d64412b6b7232632f7852533354687979623466436366715a35766678684f4d6c474d5036417472413d3d222c226c6173745f636f6e6669675f626c6f636b5f6e756d223a302c226e65775f636861696e5f636f6e666967223a6e756c6c7d';
        const nextBookkeeper = '0x0000000000000000000000000000000000000000';

        let off = 0;
        describe('NextBool', function () {

            it('NextBool read true correctly', async function () {
                const result = await this.zeroCopySource.NextBool.call('0x01', 0);
                assert.equal(result[0], true);
                assert.equal(result[1].toString(), '1');
                {
                    const {receipt} = await this.zeroCopySource.NextBool('0x01', 0);
                    console.log("NextBool(), gas:", receipt.gasUsed);
                }
            });
            it('NextBool read false correctly', async function () {
                const result = await this.zeroCopySource.NextBool.call('0x0101010001', 3);
                assert.equal(result[0], false);
                assert.equal(result[1].toString(), '4');
            });
            it('NextBool read boolean raise error since byte value = 02', async function () {
                await expectRevert(this.zeroCopySource.NextBool.call('0x0101010201', 3), "NextBool value error");
            });
            it('NextBool read boolean raise error since offset exceeds limit', async function () {
                await expectRevert(this.zeroCopySource.NextBool.call('0x0101010201', 5), "Offset exceeds limit");
            });

        });
        describe('NextByte', function () {
            it('NextByte read byte correctly', async function () {
                const result = await this.zeroCopySource.NextByte.call('0x23ff', 1);
                assert.equal(result[0], '0xff');
                assert.equal(result[1].toString(), '2');
            });
            it('NextByte read byte raise error since Offset exceeds maximum', async function () {
                await expectRevert(this.zeroCopySource.NextByte.call('0x0101010201', 5), "Offset exceeds maximum");
            });
        });
        describe('NextUint8', function () {
            it('NextUint8 read byte correctly', async function () {
                const result = await this.zeroCopySource.NextUint8.call('0xff11', 0);
                assert.equal(result[0], 0xff);
                assert.equal(result[1].toString(), '1');
            });
            it('NextUint8 read byte correctly', async function () {
                const result = await this.zeroCopySource.NextUint8.call('0x230011', 1);
                assert.equal(result[0], 0x00);
                assert.equal(result[1].toString(), '2');
            });
            it('NextUint8 read uint8 raise error since Offset exceeds maximum', async function () {
                await expectRevert(this.zeroCopySource.NextByte.call('0x0101010201', 5), "Offset exceeds maximum");
            });
        });


        describe('NextUint16 correctly', function () {
            // const MAX_UINT_16 = new web3.utils.BN('65535');
            const MAX_UINT_16 = new web3.utils.toBN('0xffff');
            assert.equal(MAX_UINT_16.toString(), "65535");
            expect(MAX_UINT_16).to.be.bignumber.equal(new BN('2').pow(new BN('16')).sub(new BN('1')));
            const hexCharacterAmount = 4;
            let bytes = '';
            let offset = 0;
            it('NextUint16(MAX_UINT_16) correctly ', async function () {
                let expected = MAX_UINT_16;
                let uint16_v_hex = web3.utils.toHex(expected);
                let param = ontUtils.reverseHex(web3.utils.padLeft(uint16_v_hex, hexCharacterAmount).slice(2));
                bytes = bytes + param;
                const result = await this.zeroCopySource.NextUint16.call('0x' + bytes, offset);
                offset = offset + hexCharacterAmount / 2;
                assert.equal(result[0].toString(), expected.toString());
                // expect(result[0]).to.be.bignumber.equal(expected);
                assert.equal(result[1].toString(), String(offset));
                
            });
            it('NextUint16(MAX_UINT_16 - 1) correctly ', async function () {
                let expected = MAX_UINT_16.sub(new web3.utils.BN('1'));
                let uint16_v_hex = web3.utils.toHex(expected);
                let param = ontUtils.reverseHex(web3.utils.padLeft(uint16_v_hex, hexCharacterAmount).slice(2));
                bytes = bytes + param;
                const result = await this.zeroCopySource.NextUint16.call('0x' + bytes, offset);
                offset = offset + hexCharacterAmount / 2;
                expect(result[0]).to.be.bignumber.equal(expected);
                assert.equal(result[1].toString(), String(offset));
            });
            it('NextUint16(1000) correctly ', async function () {
                let expected = new web3.utils.BN('1000');
                let uint16_v_hex = web3.utils.toHex(expected);
                let param = ontUtils.reverseHex(web3.utils.padLeft(uint16_v_hex, hexCharacterAmount).slice(2));
                bytes = bytes + param;
                const result = await this.zeroCopySource.NextUint16.call('0x' + bytes, offset);
                offset = offset + hexCharacterAmount / 2;
                expect(result[0]).to.be.bignumber.equal(expected);
                assert.equal(result[1].toString(), String(offset));
            });
            it('NextUint16() offset exceeds maximum ', async function () {
                await expectRevert(this.zeroCopySource.NextUint32.call('0x' + bytes, offset - 1), "offset exceeds maximum");
            });
        });

        describe('NextUint32 correctly', function () {
            const MAX_UINT_32 = new web3.utils.toBN('0xffffffff');
            expect(MAX_UINT_32).to.be.bignumber.equal(new BN('2').pow(new BN('32')).sub(new BN('1')));
            const hexCharacterAmount = 8;
            let bytes = '';
            let offset = 0;
            it('NextUint32(MAX_UINT_32) correctly ', async function () {
                let expected = MAX_UINT_32;
                let uint32_v_hex = web3.utils.toHex(expected);
                let param = ontUtils.reverseHex(web3.utils.padLeft(uint32_v_hex, hexCharacterAmount).slice(2));
                bytes = bytes + param;
                const result = await this.zeroCopySource.NextUint32.call('0x' + bytes, offset);
                {
                    const {receipt} = await this.zeroCopySource.NextUint32('0x' + bytes, offset);
                    console.log("NextUint32(), gas:", receipt.gasUsed);
                }
                offset = offset + hexCharacterAmount / 2;
                assert.equal(result[0].toString(), expected.toString());
                // expect(result[0]).to.be.bignumber.equal(expected);
                assert.equal(result[1].toString(), String(offset));
                
                
            });
            it('NextUint32(MAX_UINT_32 - 1) correctly ', async function () {
                let expected = MAX_UINT_32.sub(new web3.utils.BN('1'));
                let uint32_v_hex = web3.utils.toHex(expected);
                let param = ontUtils.reverseHex(web3.utils.padLeft(uint32_v_hex, hexCharacterAmount).slice(2));
                bytes = bytes + param;
                const result = await this.zeroCopySource.NextUint32.call('0x' + bytes, offset);
                offset = offset + hexCharacterAmount / 2;
                expect(result[0]).to.be.bignumber.equal(expected);
                assert.equal(result[1].toString(), String(offset));
            });
            it('NextUint32(1000) correctly ', async function () {
                let expected = new web3.utils.BN('1000');
                let uint32_v_hex = web3.utils.toHex(expected);
                let param = ontUtils.reverseHex(web3.utils.padLeft(uint32_v_hex, hexCharacterAmount).slice(2));
                bytes = bytes + param;
                const result = await this.zeroCopySource.NextUint32.call('0x' + bytes, offset);
                offset = offset + hexCharacterAmount / 2;
                expect(result[0]).to.be.bignumber.equal(expected);
                assert.equal(result[1].toString(), String(offset));
            });
            it('NextUint32() offset exceeds maximum ', async function () {
                await expectRevert(this.zeroCopySource.NextUint32.call('0x' + bytes, offset - 3), "offset exceeds maximum");
            });
        });

        describe('NextUint64 correctly', function () {
            const MAX_UINT_64 = new web3.utils.toBN('0xffffffffffffffff');
            expect(MAX_UINT_64).to.be.bignumber.equal(new BN('2').pow(new BN('64')).sub(new BN('1')));
            const hexCharacterAmount = 16;
            let bytes = '';
            let offset = 0;
            it('NextUint64(MAX_UINT_64) correctly ', async function () {
                let expected = MAX_UINT_64;
                let uint64_v_hex = web3.utils.toHex(expected);
                let param = ontUtils.reverseHex(web3.utils.padLeft(uint64_v_hex, hexCharacterAmount).slice(2));
                bytes = bytes + param;
                const result = await this.zeroCopySource.NextUint64.call('0x' + bytes, offset);
                offset = offset + hexCharacterAmount / 2;
                assert.equal(result[0].toString(), expected.toString());
                // expect(result[0]).to.be.bignumber.equal(expected);
                assert.equal(result[1].toString(), String(offset));
                
            });
            it('NextUint64(MAX_UINT_64 - 1) correctly ', async function () {
                let expected = MAX_UINT_64.sub(new web3.utils.BN('1'));
                let uint64_v_hex = web3.utils.toHex(expected);
                let param = ontUtils.reverseHex(web3.utils.padLeft(uint64_v_hex, hexCharacterAmount).slice(2));
                bytes = bytes + param;
                const result = await this.zeroCopySource.NextUint64.call('0x' + bytes, offset);
                offset = offset + hexCharacterAmount / 2;
                expect(result[0]).to.be.bignumber.equal(expected);
                assert.equal(result[1].toString(), String(offset));
            });
            it('NextUint64(1000) correctly ', async function () {
                let expected = new web3.utils.BN('1000');
                let uint64_v_hex = web3.utils.toHex(expected);
                let param = ontUtils.reverseHex(web3.utils.padLeft(uint64_v_hex, hexCharacterAmount).slice(2));
                bytes = bytes + param;
                const result = await this.zeroCopySource.NextUint64.call('0x' + bytes, offset);
                offset = offset + hexCharacterAmount / 2;
                expect(result[0]).to.be.bignumber.equal(expected);
                assert.equal(result[1].toString(), String(offset));
            });
            it('NextUint64() offset exceeds maximum ', async function () {
                await expectRevert(this.zeroCopySource.NextUint64.call('0x' + bytes, offset - 7), "offset exceeds maximum");
            });
        });
        describe('NextUint255 correctly', function () {
            const MAX_UINT_255 = new web3.utils.toBN('0x7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff');
            expect(MAX_UINT_255).to.be.bignumber.equal(new BN('2').pow(new BN('255')).sub(new BN('1')));
            const hexCharacterAmount = 64;
            let bytes = '';
            let offset = 0;
            it('NextUint255(MAX_UINT_255) correctly ', async function () {
                let expected = MAX_UINT_255;
                let uint64_v_hex = web3.utils.toHex(expected);
                let param = ontUtils.reverseHex(web3.utils.padLeft(uint64_v_hex, hexCharacterAmount).slice(2));
                bytes = bytes + param;
                const result = await this.zeroCopySource.NextUint255.call('0x' + bytes, offset);
                {
                    const {receipt} = await this.zeroCopySource.NextUint255('0x' + bytes, offset);
                    console.log("NextUint255(), gas:", receipt.gasUsed);
                }
                offset = offset + hexCharacterAmount / 2;
                assert.equal(result[0].toString(), expected.toString());
                // expect(result[0]).to.be.bignumber.equal(expected);
                assert.equal(result[1].toString(), String(offset));
                
            });
            it('NextUint255(MAX_UINT_255 - 1) correctly ', async function () {
                let expected = MAX_UINT_255.sub(new web3.utils.BN('1'));
                let uint64_v_hex = web3.utils.toHex(expected);
                let param = ontUtils.reverseHex(web3.utils.padLeft(uint64_v_hex, hexCharacterAmount).slice(2));
                bytes = bytes + param;
                const result = await this.zeroCopySource.NextUint255.call('0x' + bytes, offset);
                offset = offset + hexCharacterAmount / 2;
                expect(result[0]).to.be.bignumber.equal(expected);
                assert.equal(result[1].toString(), String(offset));
            });
            it('NextUint255(1000) correctly ', async function () {
                let expected = new web3.utils.BN('1000');
                let uint64_v_hex = web3.utils.toHex(expected);
                let param = ontUtils.reverseHex(web3.utils.padLeft(uint64_v_hex, hexCharacterAmount).slice(2));
                bytes = bytes + param;
                const result = await this.zeroCopySource.NextUint255.call('0x' + bytes, offset);
                offset = offset + hexCharacterAmount / 2;
                expect(result[0]).to.be.bignumber.equal(expected);
                assert.equal(result[1].toString(), String(offset));
            });
            it('NextUint255() offset exceeds maximum ', async function () {
                await expectRevert(this.zeroCopySource.NextUint255.call('0x' + bytes, offset - 31), "offset exceeds maximum");
            });
        });

        describe('NextVarUint', function() {
            let bytes = '';
            let offset = 0;
            it('NextVarUint value < 0xFD', async function() {
                let expected = new web3.utils.toBN('100');
                let num_hex = web3.utils.toHex(expected);
                let param = ontUtils.reverseHex(web3.utils.padLeft(num_hex, 2).slice(2));
                bytes = bytes + param;
                const result = await this.zeroCopySource.NextVarUint.call('0x' + bytes, offset);
                assert.equal(result[0].toString(), expected.toString());
                offset = offset + param.length/2;
                assert.equal(result[1].toString(), String(offset));
            });
            it('NextVarUint value == 0xFD', async function() {
                let expected = new web3.utils.toBN('0xfd');
                let num_hex = web3.utils.toHex(expected);
                let param = 'fd' + ontUtils.reverseHex(web3.utils.padLeft(num_hex, 4).slice(2));
                bytes = bytes + param;
                const result = await this.zeroCopySource.NextVarUint.call('0x' + bytes, offset);
                assert.equal(result[0].toString(), expected.toString());
                offset = offset + param.length/2;
                assert.equal(result[1].toString(), String(offset));
            });
            it('NextVarUint 0xFD < value < 0xFFFF', async function() {
                let expected = new web3.utils.toBN('1000');
                let num_hex = web3.utils.toHex(expected);
                let param = 'fd' + ontUtils.reverseHex(web3.utils.padLeft(num_hex, 4).slice(2));
                bytes = bytes + param;
                const result = await this.zeroCopySource.NextVarUint.call('0x' + bytes, offset);
                assert.equal(result[0].toString(), expected.toString());
                offset = offset + param.length/2;
                assert.equal(result[1].toString(), String(offset));
            });
            it('NextVarUint value == 0xFFFF', async function() {
                let expected = new web3.utils.toBN('0xFFFF');
                let num_hex = web3.utils.toHex(expected);
                let param = 'fd' + ontUtils.reverseHex(web3.utils.padLeft(num_hex, 4).slice(2));
                bytes = bytes + param;
                const result = await this.zeroCopySource.NextVarUint.call('0x' + bytes, offset);
                assert.equal(result[0].toString(), expected.toString());
                offset = offset + param.length/2;
                assert.equal(result[1].toString(), String(offset));
            });
            it('NextVarUint 0xFFFF < value < 0xFFFFFFFF', async function() {
                let expected = new web3.utils.toBN('0xFFFF').add(new web3.utils.toBN('100000'));
                let num_hex = web3.utils.toHex(expected);
                let param = 'fe' + ontUtils.reverseHex(web3.utils.padLeft(num_hex, 8).slice(2));
                bytes = bytes + param;
                const result = await this.zeroCopySource.NextVarUint.call('0x' + bytes, offset);
                assert.equal(result[0].toString(), expected.toString());
                offset = offset + param.length/2;
                assert.equal(result[1].toString(), String(offset));
            });
            it('NextVarUint value == 0xFFFFFFFF', async function() {
                let expected = new web3.utils.toBN('0xFFFFFFFF');
                let num_hex = web3.utils.toHex(expected);
                let param = 'fe' + ontUtils.reverseHex(web3.utils.padLeft(num_hex, 8).slice(2));
                bytes = bytes + param;
                const result = await this.zeroCopySource.NextVarUint.call('0x' + bytes, offset);
                assert.equal(result[0].toString(), expected.toString());
                offset = offset + param.length/2;
                assert.equal(result[1].toString(), String(offset));
            });
            it('NextVarUint value > 0xFFFFFFFF', async function() {
                let expected = new web3.utils.toBN('0xFFFFFFFF').add(new web3.utils.toBN('10000000000'));
                let num_hex = web3.utils.toHex(expected);
                let param = 'ff' + ontUtils.reverseHex(web3.utils.padLeft(num_hex, 16).slice(2));
                bytes = bytes + param;
                const result = await this.zeroCopySource.NextVarUint.call('0x' + bytes, offset);
                assert.equal(result[0].toString(), expected.toString());
                offset = offset + param.length/2;
                assert.equal(result[1].toString(), String(offset));
            });
        });

        describe('NextVarBytes', function () {
            // len < (), which is the ontology encoded hex number
            const str_11 = web3.utils.randomHex(11); // len < 0xFD
            const str_253 = web3.utils.randomHex(0xfd); // len == 0xFD
            const str_325 = web3.utils.randomHex(325); // 0xFD < len < 0XFFFF
            const str_10000 = web3.utils.randomHex(10000); // 0xFD < len < 0XFFFF
            describe('length < 0xFD', async function () {
                const padByteLen = 1;
                let bytes = '';
                let offset = 0;
                it('NextVarBytes length < 0xFD correctly', async function () {
                    const expected = str_11;
                    let lenHex = web3.utils.padLeft(web3.utils.numberToHex((expected.length - 2) / 2), padByteLen * 2);
                    let reversedLenHex = ontUtils.reverseHex(lenHex.slice(2));
                    let param = reversedLenHex + expected.slice(2);
                    bytes = bytes + param;
                    const result = await this.zeroCopySource.NextVarBytes.call('0x' + bytes, offset);
                    offset = offset + param.length / 2;
                    expect(result[0]).to.be.equal(expected);
                    assert.equal(result[1].toString(), String(offset));
                });
            });
            describe('NextVarBytes 0xFD <= length < 0xFFFF', async function () {
                const padByteLen = 2;
                let bytes = '';
                let offset = 0;
                it('NextVarBytes length == 0xFD correctly', async function () {
                    const expected = str_253;
                    let lenHex = web3.utils.padLeft(web3.utils.numberToHex((expected.length - 2) / 2), padByteLen * 2);
                    let reversedLenHex = ontUtils.reverseHex(lenHex.slice(2));
                    let param = 'fd' + reversedLenHex + expected.slice(2);
                    bytes = bytes + param;
                    const result = await this.zeroCopySource.NextVarBytes.call('0x' + bytes, offset);
                    offset = offset + param.length / 2;
                    expect(result[0]).to.be.equal(expected);
                    assert.equal(result[1].toString(), String(offset));
                });
                it('NextVarBytes 0xFD < length < 0xFFFF correctly', async function () {
                    const expected = str_325;
                    let lenHex = web3.utils.padLeft(web3.utils.numberToHex((expected.length - 2) / 2), padByteLen * 2);
                    let reversedLenHex = ontUtils.reverseHex(lenHex.slice(2));
                    let param = 'fd' + reversedLenHex + expected.slice(2);
                    bytes = bytes + param;
                    const result = await this.zeroCopySource.NextVarBytes.call('0x' + bytes, offset);
                    offset = offset + param.length / 2;
                    expect(result[0]).to.be.equal(expected);
                    assert.equal(result[1].toString(), String(offset));
                });
                it('NextVarBytes 0xFD < length = 10000 < 0xFFFF correctly', async function () {
                    const expected = str_10000;
                    let lenHex = web3.utils.padLeft(web3.utils.numberToHex((expected.length - 2) / 2), padByteLen * 2);
                    let reversedLenHex = ontUtils.reverseHex(lenHex.slice(2));
                    let param = 'fd' + reversedLenHex + expected.slice(2);
                    bytes = bytes + param;
                    const result = await this.zeroCopySource.NextVarBytes.call('0x' + bytes, offset);
                    offset = offset + param.length / 2;
                    expect(result[0]).to.be.equal(expected);
                    assert.equal(result[1].toString(), String(offset));
                });
                it('NextVarBytes, offset exceeds maximum', async function () {
                    const expected = str_10000;
                    let lenHex = web3.utils.padLeft(web3.utils.numberToHex((expected.length - 2) / 2), padByteLen * 2);
                    let reversedLenHex = ontUtils.reverseHex(lenHex.slice(2));
                    let param = 'fd' + reversedLenHex + expected.slice(2);
                    await expectRevert(this.zeroCopySource.NextVarBytes.call('0x' + bytes, offset - 1), "offset exceeds maximum");
                });
            });
        });
        describe('NextHash correctly', function () {
            it('NextHash correctly', async function () {
                const expected = web3.utils.randomHex(32);
                const result = await this.zeroCopySource.NextHash.call(expected, 0);
                expect(result[0]).to.be.equal(expected);
                assert.equal(result[1].toString(), String(32));
            });
            it('NextHash, offset exceeds maximum', async function () {
                await expectRevert(this.zeroCopySource.NextHash.call(web3.utils.randomHex(32), 1), "offset exceeds maximum");
            });
        });
        describe('NextBytes20 correctly', function () {
            it('NextBytes20 correctly', async function () {
                const expected = web3.utils.randomHex(20);
                const result = await this.zeroCopySource.NextBytes20.call(expected, 0);
                expect(result[0]).to.be.equal(expected);
                assert.equal(result[1].toString(), String(20));
            });
            it('NextBytes20, offset exceeds maximum', async function () {
                await expectRevert(this.zeroCopySource.NextBytes20.call(web3.utils.randomHex(20), 1), "offset exceeds maximum");
            });
        });
    });
});