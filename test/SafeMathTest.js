const { expectRevert, expectEvent, constants, BN} = require('@openzeppelin/test-helpers');
const SafeMathMock = artifacts.require('SafeMathMock');
const { MAX_UINT256 } = constants;
const { expect } = require('chai');
const BigNumber = web3.utils.BN;

contract('SafeMath', () => {

  before(async function () {
    this.safeMath = await SafeMathMock.new();
  });

  describe('add', function () {
    it('adds correctly', async function () {
      const a = new BigNumber(5678);
      const b = new BigNumber(1234);
      const c = a.add(b);

      const result = await this.safeMath.add(a, b);
      assert(result.eq(a.add(b)));
    });

    it('reverts on addition overflow', async function () {
      const a = MAX_UINT256;
      const b = new BigNumber(1);
        await expectRevert(this.safeMath.add.call(a, b), "SafeMath: addition overflow");
    });
  });

  describe('sub', function () {
    it('subtracts correctly', async function () {
      const a = new BigNumber(5678);
      const b = new BigNumber(1234);
      const result = await this.safeMath.sub(a, b);

      assert(result.eq(a.sub(b)));
    });

    it('throws an error if subtraction result would be negative', async function () {
      const a = new BigNumber(1234);
      const b = new BigNumber(5678);
      await expectRevert(this.safeMath.sub.call(a, b), "SafeMath: subtraction overflow");
    });
  });

  describe('mul', function () {
    it('multiplies correctly', async function () {
      const a = new BN('1234');
      const b = new BN('5678');
      let result = await this.safeMath.mul.call(a, b);
      let expected = a.mul(b);
    //   assert.equal(result.toString(), expected.toString());
      expect(result).to.be.bignumber.equal(expected);
    });

    it('multiplies by zero correctly', async function () {
      const a = new BN('0');
      const b = new BN('5678');
      let result = await this.safeMath.mul.call(a, b);
      let expected = a.mul(b);
      expect(result).to.be.bignumber.equal(expected);
    });

    it('reverts on multiplication overflow', async function () {
      const a = MAX_UINT256;
      const b = new BN('2');
      await expectRevert(this.safeMath.mul.call(a, b), "SafeMath: multiplication overflow");
    });
  });

  describe('div', function () {
    it('divides correctly', async function () {
      const a = new BN('5678');
      const b = new BN('5678');

      expect(await this.safeMath.div(a, b)).to.be.bignumber.equal(a.div(b));
    });

    it('divides zero correctly', async function () {
      const a = new BN('0');
      const b = new BN('5678');

      expect(await this.safeMath.div(a, b)).to.be.bignumber.equal('0');
    });

    it('returns complete number result on non-even division', async function () {
      const a = new BN('7000');
      const b = new BN('5678');

      expect(await this.safeMath.div(a, b)).to.be.bignumber.equal('1');
    });

    it('reverts on division by zero', async function () {
      const a = new BN('5678');
      const b = new BN('0');

      await expectRevert(this.safeMath.div(a, b), 'SafeMath: division by zero');
    });
  });

  describe('mod', function () {
    describe('modulos correctly', async function () {
      it('when the dividend is smaller than the divisor', async function () {
        const a = new BN('284');
        const b = new BN('5678');

        expect(await this.safeMath.mod(a, b)).to.be.bignumber.equal(a.mod(b));
      });

      it('when the dividend is equal to the divisor', async function () {
        const a = new BN('5678');
        const b = new BN('5678');

        expect(await this.safeMath.mod(a, b)).to.be.bignumber.equal(a.mod(b));
      });

      it('when the dividend is larger than the divisor', async function () {
        const a = new BN('7000');
        const b = new BN('5678');

        expect(await this.safeMath.mod(a, b)).to.be.bignumber.equal(a.mod(b));
      });

      it('when the dividend is a multiple of the divisor', async function () {
        const a = new BN('17034'); // 17034 == 5678 * 3
        const b = new BN('5678');

        expect(await this.safeMath.mod(a, b)).to.be.bignumber.equal(a.mod(b));
      });
    });

    it('reverts with a 0 divisor', async function () {
      const a = new BN('5678');
      const b = new BN('0');

      await expectRevert(this.safeMath.mod(a, b), 'SafeMath: modulo by zero');
    });
  });

});
