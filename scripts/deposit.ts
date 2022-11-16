import { ethers } from 'hardhat';

async function main() {
  const tokenAddressBG1 = '0xEfA08bCF0EF4BB63619f79aA92a753E245DfDDce'; // set address
  const erc20BridgeAddressBG1 = '0xc2504cD8ca96723c248df5d696B921a3332f1d38'; // set address

    const [receiver] = await ethers.getSigners();

    const chainIdBG2 = '5002'

    const erc20bridge = await ethers.getContractAt('ERC20Bridge', erc20BridgeAddressBG1);

    console.log("Deposit");
    await(await erc20bridge.deposit(tokenAddressBG1, chainIdBG2, receiver.address, '1000000000000000000')).wait();
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});