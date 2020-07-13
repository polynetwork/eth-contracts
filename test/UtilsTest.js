const UtilsMock = artifacts.require('./contractmocks/libs/utils/UtilsMock.sol');
const { expectRevert, expectEvent, constants, BN } = require('@openzeppelin/test-helpers');
const { expect } = require('chai');
const { SHA256, SHA3} = require('crypto-js')
contract('Utils', () => {

    before(async function () {
        this.utils = await UtilsMock.new();
    });

    describe('bytesToBytes32', function () {
        it('bytesToBytes32 correctly', async function () {
            const expected = web3.utils.randomHex(32);
            const result = await this.utils.bytesToBytes32(expected);
            assert.equal(result, expected);
        });
        it('throws an error on bytes length == 31 < 32', async function () {
            const expected = web3.utils.randomHex(31);
            await expectRevert(this.utils.bytesToBytes32.call(expected), "bytes length is not 32.");
        });
        it('throws an error on bytes length == 33 > 32', async function () {
            const expected = web3.utils.randomHex(33);
            await expectRevert(this.utils.bytesToBytes32.call(expected), "bytes length is not 32.");
        });
    });
    describe('bytesToUint256', function () {
        it('bytesToUint256 correctly', async function () {
            const randomHex32 = web3.utils.randomHex(32).slice(2);
            const expectedHex = '0x' + (randomHex32.charAt(0) & 7) + randomHex32.slice(1);
            const expected = web3.utils.toBN(expectedHex);
            const result = await this.utils.bytesToUint256(expectedHex);
            expect(result).to.be.bignumber.equal(expected);
        });
        it('throws an error on bytes length == 31 < 32', async function () {
            const randomHex32 = web3.utils.randomHex(31).slice(2);
            const expectedHex = '0x' + (randomHex32.charAt(0) & 7) + randomHex32.slice(1);
            await expectRevert(this.utils.bytesToUint256.call(expectedHex), "bytes length is not 32.");
        });
        it('throws an error on bytes length == 33 > 32', async function () {
            const randomHex32 = web3.utils.randomHex(33).slice(2);
            const expectedHex = '0x' + (randomHex32.charAt(0) & 7) + randomHex32.slice(1);
            await expectRevert(this.utils.bytesToUint256.call(expectedHex), "bytes length is not 32.");
        });
    });
    describe('uint256ToBytes', function () {
        it('uint256ToBytes correctly', async function () {
            const randomHex32 = web3.utils.randomHex(32).slice(2);
            const expected = '0x' + (randomHex32.charAt(0) & 7) + randomHex32.slice(1);
            const expectedNum = web3.utils.toBN(expected);
            const result = await this.utils.uint256ToBytes.call(expectedNum);
            expect(result).to.be.equal(expected);
        });
        it('throws an error on value exceeds the range', async function () {
            const randomHex32 = web3.utils.randomHex(32).slice(2);
            const expected = '0x' + '80' + randomHex32.slice(2);
            const expectedNum = web3.utils.toBN(expected);
            await expectRevert(this.utils.uint256ToBytes.call(expectedNum), "Value exceeds the range");
        });
    });
    describe('hashLeaf', function () {
        // TODO: check
        it('hashLeaf correctly', async function () {
            const leafData =web3.utils.randomHex(100);
            const expected = web3.utils.soliditySha3('0x00'+leafData.slice(2));
            const expected1 = web3.utils.sha3('0x00'+leafData.slice(2));
            const expected2 = SHA256('0x00'+leafData.slice(2)).toString();
            const expected3 = SHA3('0x00'+leafData.slice(2)).toString();
            const result = await this.utils.hashLeaf.call(leafData);
            console.log('expected = ', expected);
            console.log('expected1 = ', expected1);
            console.log('expected2 = ', expected2);
            console.log('expected3 = ', expected3);
            console.log('result = ', result);
            // expect(result).to.be.equal(expected);
        });
        
    });
    describe('slice', function () {
        bs = '0xab036729af8b8f9b610af4e11b14fa30c348f40c2c230cce92ef6ef37726fee7';
        bs2 = '0xab036729af';

        it('slice correctly', async function () {

            const result = await this.utils.slice(bs, 0, 5);
            assert.equal(result, bs2);
        });

        // it('throws an error on slice length greater than bytes length', async function () {

        //     await assertRevert(this.utils.slice(bs, 0, 100));
        // });

    });

    describe('addressToBytes and bytesToAddress correctly', function () {
        addr = web3.utils.randomHex(20);

        it('addressToBytes correctly', async function () {
            const b = await this.utils.addressToBytes(addr);
            assert.equal(b, addr);
        });

        it('bytesToAddress correctly', async function () {
            const b = await this.utils.addressToBytes(addr);
            const a = await this.utils.bytesToAddress(b);
            assert.equal(a.toLowerCase(), addr);
        });

    });
});