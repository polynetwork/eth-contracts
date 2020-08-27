const ZeroCopySinkMock = artifacts.require('./contractmocks/libs/common/ZeroCopySinkMock.sol');
const { expectRevert, expectEvent, constants, BN } = require('@openzeppelin/test-helpers');
const ontsdk = require("ontology-ts-sdk");
const ontUtils = ontsdk.utils;
const { expect } = require('chai');
const { bigNumberify } = require('ethers/utils');
// const BN = require('bn.js');
contract('ZeroCopySink', () => {

    before(async function () {
        this.zeroCopy = await ZeroCopySinkMock.new();
    });

    describe('WriteBool', function () {

        it('WriteBool correctly', async function () {
            t = '0x01';
            f = '0x00';

            const b1 = await this.zeroCopy.WriteBool.call(true);
            assert.equal(b1, t);
            {
                const {receipt} = await this.zeroCopy.WriteBool(true);
                console.log("WriteBool(true) gas:", receipt.gasUsed);
            }
            
            
            const b2 = await this.zeroCopy.WriteBool.call(false);
            assert.equal(b2, f);
            {
                const {receipt} = await this.zeroCopy.WriteBool(false);
                console.log("WriteBool(false) gas:", receipt.gasUsed);
            }
        });

    });

    describe('WriteByte', function () {

        it('WriteByte correctly', async function () {
            b = web3.utils.numberToHex(128);
            let bs = await this.zeroCopy.WriteByte.call(b);
            assert.equal(b, bs);
            {
                const {receipt} = await this.zeroCopy.WriteByte(b);
                console.log("WriteByte(", b, ") gas:", receipt.gasUsed);
            }

        });

        it('WriteByte throws error on byte number greater than 2 ** 8', async function () {
            b = web3.utils.numberToHex(256);
            const bs = await this.zeroCopy.WriteByte.call(b);
            assert.notEqual(b, bs);
        });

    });

    describe('WriteUint8', function () {
        const MAX_UINT_8 = new web3.utils.toBN('0xff');
        assert.equal(MAX_UINT_8.toString(), "255");
        expect(MAX_UINT_8).to.be.bignumber.equal(new BN('2').pow(new BN('8')).sub(new BN('1')));
        const hexCharacterAmount = 2;
        it('WriteUint8(MAX_UINT_8) correctly ', async function () {
            let num = MAX_UINT_8;
            let uint8_v_hex = web3.utils.toHex(num);
            let expected = '0x' + ontUtils.reverseHex(web3.utils.padLeft(uint8_v_hex, hexCharacterAmount).slice(2));
            const result = await this.zeroCopy.WriteUint8.call(num);
            assert.equal(result, expected);
            {
                const {receipt} = await this.zeroCopy.WriteUint8(num);
                console.log("WriteUint8(", num.toString(), ") gas:", receipt.gasUsed);
            }
        });
        it('WriteUint8(MAX_UINT_8 - 1) correctly ', async function () {
            let num = MAX_UINT_8.sub(new web3.utils.BN('1'));
            let uint8_v_hex = web3.utils.toHex(num);
            let expected = '0x' + ontUtils.reverseHex(web3.utils.padLeft(uint8_v_hex, hexCharacterAmount).slice(2));
            const result = await this.zeroCopy.WriteUint8.call(num);
            assert.equal(result, expected);
        });
        it('WriteUint8(100) correctly ', async function () {
            let num = new web3.utils.BN('100');
            let uint8_v_hex = web3.utils.toHex(num);
            let expected = '0x' + ontUtils.reverseHex(web3.utils.padLeft(uint8_v_hex, hexCharacterAmount).slice(2));
            const result = await this.zeroCopy.WriteUint8.call(num);
            assert.equal(result, expected);
        });
        it('WriteUint8(MAX_UINT_8 + increment) equal to WriteUint8(increment - 1)', async function () {
            let increment = new web3.utils.BN('10');
            let num = MAX_UINT_8.add(increment);
            const result = await this.zeroCopy.WriteUint8.call(num);

            let num1 = increment.sub(new web3.utils.BN('1'));
            let uint8_v_hex = web3.utils.toHex(num1);
            let expected1 = '0x' + ontUtils.reverseHex(web3.utils.padLeft(uint8_v_hex, hexCharacterAmount).slice(2));
            const result1 = await this.zeroCopy.WriteUint8.call(num1);

            assert.equal(result, expected1);
            assert.equal(result1, expected1);
        });
    });

    describe('WriteUint16', function () {
        // const MAX_UINT_16 = new web3.utils.BN('65535');
        const MAX_UINT_16 = new web3.utils.toBN('0xffff');
        assert.equal(MAX_UINT_16.toString(), "65535");
        expect(MAX_UINT_16).to.be.bignumber.equal(new BN('2').pow(new BN('16')).sub(new BN('1')));
        const hexCharacterAmount = 4;
        it('WriteUint16(MAX_UINT_16) correctly ', async function () {
            let num = MAX_UINT_16;
            let uint16_v_hex = web3.utils.toHex(num);
            let expected = '0x' + ontUtils.reverseHex(web3.utils.padLeft(uint16_v_hex, hexCharacterAmount).slice(2));
            const result = await this.zeroCopy.WriteUint16.call(num);
            assert.equal(result, expected);
            {
                const {receipt} = await this.zeroCopy.WriteUint16(num);
                console.log("WriteUint16(", num.toString(), ") gas:", receipt.gasUsed);
            }
        });
        it('WriteUint16(MAX_UINT_16 - 1) correctly ', async function () {
            let num = MAX_UINT_16.sub(new web3.utils.BN('1'));
            let uint16_v_hex = web3.utils.toHex(num);
            let expected = '0x' + ontUtils.reverseHex(web3.utils.padLeft(uint16_v_hex, hexCharacterAmount).slice(2));
            const result = await this.zeroCopy.WriteUint16.call(num);
            assert.equal(result, expected);
        });
        it('WriteUint16(1000) correctly ', async function () {
            let num = new web3.utils.BN('1000');
            let uint16_v_hex = web3.utils.toHex(num);
            let expected = '0x' + ontUtils.reverseHex(web3.utils.padLeft(uint16_v_hex, hexCharacterAmount).slice(2));
            const result = await this.zeroCopy.WriteUint16.call(num);
            assert.equal(result, expected);
        });
        it('WriteUint16(MAX_UINT_16 + increment) equal to WriteUint16(increment - 1)', async function () {
            let increment = new web3.utils.BN('100');
            let num = MAX_UINT_16.add(increment);
            const result = await this.zeroCopy.WriteUint16.call(num);

            let num1 = increment.sub(new web3.utils.BN('1'));
            let uint16_v_hex = web3.utils.toHex(num1);
            let expected1 = '0x' + ontUtils.reverseHex(web3.utils.padLeft(uint16_v_hex, hexCharacterAmount).slice(2));
            const result1 = await this.zeroCopy.WriteUint16.call(num1);

            assert.equal(result, expected1);
            assert.equal(result1, expected1);
        });
    });

    describe('WriteUint32', function () {
        // const MAX_UINT_32 = new web3.utils.BN('4294967295‬');
        const MAX_UINT_32 = new web3.utils.toBN('0xffffffff');
        // assert.equal(MAX_UINT_32.toString(), (new web3.utils.BN('4294967295‬')).toString());
        // expect(MAX_UINT_32).to.be.bignumber.equal(bigNumberify('4294967295‬'));
        // console.log("bigNumberify(4294967295‬) = ", bigNumberify("4294967295‬").toString())
        expect(MAX_UINT_32).to.be.bignumber.equal(new BN('2').pow(new BN('32')).sub(new BN('1')));
        const hexCharacterAmount = 8;
        it('WriteUint32(MAX_UINT_32) correctly ', async function () {
            let num = MAX_UINT_32;
            let uint32_v_hex = web3.utils.toHex(num);
            let expected = '0x' + ontUtils.reverseHex(web3.utils.padLeft(uint32_v_hex, hexCharacterAmount).slice(2));
            const result = await this.zeroCopy.WriteUint32.call(num);
            assert.equal(result, expected);
            {
                const {receipt} = await this.zeroCopy.WriteUint32(num);
                console.log("WriteUint32(", num.toString(), ") gas:", receipt.gasUsed);
            }
        });
        it('WriteUint32(MAX_UINT_32 - 1) correctly ', async function () {
            let num = MAX_UINT_32.sub(new web3.utils.BN('1'));
            let uint32_v_hex = web3.utils.toHex(num);
            let expected = '0x' + ontUtils.reverseHex(web3.utils.padLeft(uint32_v_hex, hexCharacterAmount).slice(2));
            const result = await this.zeroCopy.WriteUint32.call(num);
            assert.equal(result, expected);
        });
        it('WriteUint32(1000) correctly ', async function () {
            let num = new web3.utils.BN('1000');
            let uint32_v_hex = web3.utils.toHex(num);
            let expected = '0x' + ontUtils.reverseHex(web3.utils.padLeft(uint32_v_hex, hexCharacterAmount).slice(2));
            const result = await this.zeroCopy.WriteUint32.call(num);
            assert.equal(result, expected);
        });
        it('WriteUint32(MAX_UINT_32 + increment) equal to WriteUint32(increment - 1)', async function () {
            let increment = new web3.utils.BN('10');
            let num = MAX_UINT_32.add(increment);
            const result = await this.zeroCopy.WriteUint32.call(num);

            let num1 = increment.sub(new web3.utils.BN('1'));
            let uint32_v_hex = web3.utils.toHex(num1);
            let expected1 = '0x' + ontUtils.reverseHex(web3.utils.padLeft(uint32_v_hex, hexCharacterAmount).slice(2));
            const result1 = await this.zeroCopy.WriteUint32.call(num1);

            assert.equal(result, expected1);
            assert.equal(result1, expected1);
        });
    });

    describe('WriteUint64', function () {
        // const MAX_UINT_64 = new web3.utils.BN('184467440737095516165');
        const MAX_UINT_64 = new web3.utils.toBN('0xffffffffffffffff');
        // assert.equal(MAX_UINT_64.toString(), "184467440737095516165");
        expect(MAX_UINT_64).to.be.bignumber.equal(new BN('2').pow(new BN('64')).sub(new BN('1')));
        const hexCharacterAmount = 16;
        it('WriteUint64(MAX_UINT_64) correctly ', async function () {
            let num = MAX_UINT_64;
            let uint64_v_hex = web3.utils.toHex(num);
            let expected = '0x' + ontUtils.reverseHex(web3.utils.padLeft(uint64_v_hex, hexCharacterAmount).slice(2));
            const result = await this.zeroCopy.WriteUint64.call(num);
            assert.equal(result, expected);
            {
                const {receipt} = await this.zeroCopy.WriteUint64(num);
                console.log("WriteUint64(", num.toString(), ") gas:", receipt.gasUsed);
            }
        });
        it('WriteUint64(MAX_UINT_64 - 1) correctly ', async function () {
            let num = MAX_UINT_64.sub(new web3.utils.BN('1'));
            let uint64_v_hex = web3.utils.toHex(num);
            let expected = '0x' + ontUtils.reverseHex(web3.utils.padLeft(uint64_v_hex, hexCharacterAmount).slice(2));
            const result = await this.zeroCopy.WriteUint64.call(num);
            assert.equal(result, expected);
        });
        it('WriteUint64(1000) correctly ', async function () {
            let num = new web3.utils.BN('1000');
            let uint64_v_hex = web3.utils.toHex(num);
            let expected = '0x' + ontUtils.reverseHex(web3.utils.padLeft(uint64_v_hex, hexCharacterAmount).slice(2));
            const result = await this.zeroCopy.WriteUint64.call(num);
            assert.equal(result, expected);
        });
        it('WriteUint64(MAX_UINT_64 + increment) equal to WriteUint64(increment - 1)', async function () {
            let increment = new web3.utils.BN('10');
            let num = MAX_UINT_64.add(increment);
            const result = await this.zeroCopy.WriteUint64.call(num);

            let num1 = increment.sub(new web3.utils.BN('1'));
            let uint64_v_hex = web3.utils.toHex(num1);
            let expected1 = '0x' + ontUtils.reverseHex(web3.utils.padLeft(uint64_v_hex, hexCharacterAmount).slice(2));
            const result1 = await this.zeroCopy.WriteUint64.call(num1);

            assert.equal(result, expected1);
            assert.equal(result1, expected1);
        });
    });

    describe('WriteUint255', function () {
        // //  in hex format should be 0x7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff
        const MAX_UINT_255 = new web3.utils.toBN('0x7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff');
        // assert.equal(MAX_UINT_255.toString(), "57896044618658097711785492504343953926634992332820282019728792003956564819967");
        expect(MAX_UINT_255).to.be.bignumber.equal(new BN('2').pow(new BN('255')).sub(new BN('1')));
        const hexCharacterAmount = 64;
        it('WriteUint255(MAX_UINT_255) correctly', async function () {
            let uint255_v = MAX_UINT_255;
            let uint255_v_hex = web3.utils.toHex(uint255_v);
            let expected = '0x' + ontUtils.reverseHex(web3.utils.padLeft(uint255_v_hex, hexCharacterAmount).slice(2));
            let result = await this.zeroCopy.WriteUint255.call(uint255_v);
            assert.equal(result, expected);
            {
                const {receipt} = await this.zeroCopy.WriteUint255(uint255_v);
                console.log("WriteUint255(", uint255_v.toString(), ") gas:", receipt.gasUsed);
            }
        });
        it('WriteUint255(MAX_UINT_255 - 1) correctly', async function () {
            let uint256_v = MAX_UINT_255.sub(new web3.utils.BN('1'));
            let uint256_v_hex = web3.utils.toHex(uint256_v);
            let expected = '0x' + ontUtils.reverseHex(web3.utils.padLeft(uint256_v_hex, hexCharacterAmount).slice(2));
            let result = await this.zeroCopy.WriteUint255.call(uint256_v);
            assert.equal(result, expected);
        });
        it('WriteUint255(0) correctly', async function () {
            let uint256_v = new web3.utils.BN('0');
            let uint256_v_hex = web3.utils.toHex(uint256_v);
            let expected = '0x' + ontUtils.reverseHex(web3.utils.padLeft(uint256_v_hex, hexCharacterAmount).slice(2));
            let result = await this.zeroCopy.WriteUint255.call(uint256_v);
            assert.equal(result, expected);
        });

        it('WriteUint255(-1) incorrectly', async function () {
            uint256_v = new web3.utils.BN('-1');
            await expectRevert(this.zeroCopy.WriteUint255.call(uint256_v), "Value exceeds uint255 range");
        });
        it('WriteUint255(MAX_UINT_255 + 1) incorrectly', async function () {
            let uint256_v = MAX_UINT_255.add(new web3.utils.BN('1'));
            await expectRevert(this.zeroCopy.WriteUint255.call(uint256_v), "Value exceeds uint255 range");
        });
    });
    
    describe('WriteVarUint', function() {
        it('WriteVarUint value < 0xFD === 253', async function() {
            let num = 100;
            let num_hex = web3.utils.toHex(num);
            let expected = '0x' + ontUtils.reverseHex(web3.utils.padLeft(num_hex, 2).slice(2));
            const result = await this.zeroCopy.WriteVarUint.call(num);
            assert.equal(result, expected);
            {
                const {receipt} = await this.zeroCopy.WriteVarUint(num);
                console.log("WriteVarUint(", num.toString(), ") gas:", receipt.gasUsed);
            }
        });
        it('WriteVarUint value == 0xFD === 253', async function() {
            let num = 0xfd;
            let num_hex = web3.utils.toHex(num);
            let expected = '0xfd' + ontUtils.reverseHex(web3.utils.padLeft(num_hex, 4).slice(2));
            const result = await this.zeroCopy.WriteVarUint.call(num);
            assert.equal(result, expected);
            {
                const {receipt} = await this.zeroCopy.WriteVarUint(num);
                console.log("WriteVarUint(", num.toString(), ") gas:", receipt.gasUsed);
            }
        });
        it('WriteVarUint 0xFD <= value == 1000 <= 0xFFFF', async function() {
            let num = 1000;
            let num_hex = web3.utils.toHex(num);
            let expected = '0xfd' + ontUtils.reverseHex(web3.utils.padLeft(num_hex, 4).slice(2));
            const result = await this.zeroCopy.WriteVarUint.call(num);
            assert.equal(result, expected);
            {
                const {receipt} = await this.zeroCopy.WriteVarUint(num);
                console.log("WriteVarUint(", num.toString(), ") gas:", receipt.gasUsed);
            }
        });
        it('WriteVarUint 0xFD <= value == 0xFFFF <= 0xFFFF', async function() {
            let num = 0xFFFF;
            let num_hex = web3.utils.toHex(num);
            let expected = '0xfd' + ontUtils.reverseHex(web3.utils.padLeft(num_hex, 4).slice(2));
            const result = await this.zeroCopy.WriteVarUint.call(num);
            assert.equal(result, expected);
            {
                const {receipt} = await this.zeroCopy.WriteVarUint(num);
                console.log("WriteVarUint(", num.toString(), ") gas:", receipt.gasUsed);
            }
        });
        it('WriteVarUint 0xFFFF < value == 655360 <= 0xFFFFFFFF', async function() {
            let num = 655360;
            let num_hex = web3.utils.toHex(num);
            let expected = '0xfe' + ontUtils.reverseHex(web3.utils.padLeft(num_hex, 8).slice(2));
            const result = await this.zeroCopy.WriteVarUint.call(num);
            assert.equal(result, expected);
            {
                const {receipt} = await this.zeroCopy.WriteVarUint(num);
                console.log("WriteVarUint(", num.toString(), ") gas:", receipt.gasUsed);
            }
        });
        it('WriteVarUint 0xFFFF < value == 0xFFFFFFFF <= 0xFFFFFFFF', async function() {
            let num = 0xFFFFFFFF;
            let num_hex = web3.utils.toHex(num);
            let expected = '0xfe' + ontUtils.reverseHex(web3.utils.padLeft(num_hex, 8).slice(2));
            const result = await this.zeroCopy.WriteVarUint.call(num);
            assert.equal(result, expected);
            {
                const {receipt} = await this.zeroCopy.WriteVarUint(num);
                console.log("WriteVarUint(", num.toString(), ") gas:", receipt.gasUsed);
            }
        });
        it('WriteVarUint 0xFFFF < value == 0xFFFFFFFF <= 0xFFFFFFFF', async function() {
            let num = 42949672950;
            let num_hex = web3.utils.toHex(num);
            let expected = '0xff' + ontUtils.reverseHex(web3.utils.padLeft(num_hex, 16).slice(2));
            const result = await this.zeroCopy.WriteVarUint.call(num);
            assert.equal(result, expected);
            {
                const {receipt} = await this.zeroCopy.WriteVarUint(num);
                console.log("WriteVarUint(", num.toString(), ") gas:", receipt.gasUsed);
            }
        });
    });
    describe('WriteVarBytes', function () {
        // len < (), which is the ontology encoded hex number
        const str_11 = web3.utils.randomHex(11); // len < 0xFD
        const str_253 = web3.utils.randomHex(0xfd); // len == 0xFD
        const str_325 = web3.utils.randomHex(325); // 0xFD < len < 0XFFFF
        const str_10000 = web3.utils.randomHex(10000); // 0xFD < len < 0XFFFF
        // const str_16776960 = web3.utils.randomHex(0xffff); // len == 0xFFFF
        // const str_4294905600 = web3.utils.randomHex(0x0fffff); // 0xFFFF < len < 0xFFFFFFFF
        // const str_1099511627520 = web3.utils.randomHex(0xffffffff); // len == 0xFFFFFFFF
        // const str_1099511627521 = web3.utils.randomHex(0x01ffffffff); // len > 0xFFFFFFFF
        describe('length < 0xFD', async function () {
            const padByteLen = 1;
            it('WriteVarBytes length < 0xFD correctly', async function () {
                const str = str_11;
                const input = web3.utils.hexToBytes(str);
                let result = await this.zeroCopy.WriteVarBytes.call(input);
                let lenHex = web3.utils.padLeft(web3.utils.numberToHex((str.length - 2) / 2), padByteLen * 2);
                let reversedLenHex = ontUtils.reverseHex(lenHex.slice(2));
                let expected = '0x' + reversedLenHex + str.slice(2);
                assert.equal(result, expected);
                {
                    const {receipt} = await this.zeroCopy.WriteVarBytes(input);
                    console.log("WriteVarBytes(", str, ") gas:", receipt.gasUsed);
                }
            });
        });
        describe('WriteVarBytes 0xFD <= length < 0xFFFF', async function () {
            const padByteLen = 2;
            it('WriteVarBytes length == 0xFD correctly', async function () {
                const str = str_253;
                const input = web3.utils.hexToBytes(str);
                let result = await this.zeroCopy.WriteVarBytes.call(input);
                let lenHex = web3.utils.padLeft(web3.utils.numberToHex((str.length - 2) / 2), padByteLen * 2);
                let reversedLenHex = ontUtils.reverseHex(lenHex.slice(2));
                let expected = '0xfd' + reversedLenHex + str.slice(2);
                assert.equal(result, expected);
                {
                    const {receipt} = await this.zeroCopy.WriteVarBytes(input);
                    console.log("WriteVarBytes(", str, ") gas:", receipt.gasUsed);
                }
            });
            it('WriteVarBytes 0xFD < length < 0xFFFF correctly', async function () {
                const str = str_325;
                const input = web3.utils.hexToBytes(str);
                let result = await this.zeroCopy.WriteVarBytes.call(input);
                let lenHex = web3.utils.padLeft(web3.utils.numberToHex((str.length - 2) / 2), padByteLen * 2);
                let reversedLenHex = ontUtils.reverseHex(lenHex.slice(2));
                let expected = '0xfd' + reversedLenHex + str.slice(2);
                assert.equal(result, expected);
                {
                    const {receipt} = await this.zeroCopy.WriteVarBytes(input);
                    console.log("WriteVarBytes(", str, ") gas:", receipt.gasUsed);
                }
            });
            it('WriteVarBytes 0xFD < length = 10000 < 0xFFFF correctly', async function () {
                const str = str_10000;
                const input = web3.utils.hexToBytes(str);
                let result = await this.zeroCopy.WriteVarBytes.call(input);
                let lenHex = web3.utils.padLeft(web3.utils.numberToHex((str.length - 2) / 2), padByteLen * 2);
                let reversedLenHex = ontUtils.reverseHex(lenHex.slice(2));
                let expected = '0xfd' + reversedLenHex + str.slice(2);
                assert.equal(result, expected);
                {
                    const {receipt} = await this.zeroCopy.WriteVarBytes(input);
                    console.log("WriteVarBytes(", str, ") gas:", receipt.gasUsed);
                }
            });
        });
        // describe('WriteVarBytes 0xFFFF <= length < 0xFFFFFFFF', async function () {
        //     const padByteLen = 4;
        //     it('WriteVarBytes length == 0xFFFF correctly', async function () {
        //         const str = str_16776960;
        //         let result = await this.zeroCopy.WriteVarBytes(web3.utils.hexToBytes(str));
        //         let lenHex = web3.utils.padLeft(web3.utils.numberToHex((str.length - 2) / 2), padByteLen * 2);
        //         let reversedLenHex = ontUtils.reverseHex(lenHex.slice(2));
        //         let expected = '0xfe' + reversedLenHex + str.slice(2);
        //         assert.equal(result, expected);
        //     });
        //     //  the following would have error:  Requested too many random bytes
        //     it('WriteVarBytes 0xFFFF < length < 0xFFFFFFFF correctly', async function () {
        //         const str = str_4294905600;
        //         let result = await this.zeroCopy.WriteVarBytes(web3.utils.hexToBytes(str));
        //         let lenHex = web3.utils.padLeft(web3.utils.numberToHex((str.length - 2) / 2), padByteLen * 2);
        //         let reversedLenHex = ontUtils.reverseHex(lenHex.slice(2));
        //         let expected = '0xfe' + reversedLenHex + str.slice(2);
        //         assert.equal(result, expected);
        //     });
        //     it('WriteVarBytes length == 0xFFFFFFFF correctly', async function () {
        //         const str = str_1099511627520;
        //         let result = await this.zeroCopy.WriteVarBytes(web3.utils.hexToBytes(str));
        //         let lenHex = web3.utils.padLeft(web3.utils.numberToHex((str.length - 2) / 2), padByteLen * 2);
        //         let reversedLenHex = ontUtils.reverseHex(lenHex.slice(2));
        //         let expected = '0xfe' + reversedLenHex + str.slice(2);
        //         assert.equal(result, expected);
        //     });
        // });
        // describe('WriteVarBytes length > 0xFFFFFFFF', async function () {
        //     it('WriteVarBytes length > 0xFFFFFFFF correctly', async function () {
        //         const str = str_1099511627521;
        //         const padByteLen = 8;
        //         let result = await this.zeroCopy.WriteVarBytes(web3.utils.hexToBytes(str));
        //         let lenHex = web3.utils.padLeft(web3.utils.numberToHex((str.length - 2) / 2), padByteLen * 2);
        //         let reversedLenHex = ontUtils.reverseHex(lenHex.slice(2));
        //         let expected = '0xff' + reversedLenHex + str.slice(2);
        //         assert.equal(result, expected);
        //     });
        // });
    });

});

