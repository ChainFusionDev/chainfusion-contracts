import { ethers } from 'hardhat';

async function main() {
    const tokenAddressBG1 = '0xEfA08bCF0EF4BB63619f79aA92a753E245DfDDce'; // set address
    const erc20BridgeAddressBG1 = '0xc2504cD8ca96723c248df5d696B921a3332f1d38'; // set address

    const tokenAddressBG2 = '0xa2a0D69221829d6005E31Bb187A0A5DEBEaD8331'; // set address
    const erc20BridgeAddressBG2 = '0x6B99600daD0a1998337357696827381D122825F3'; // set address

    const bridgeAppAddress = '0x3B02fF1e626Ed7a8fd6eC5299e2C54e1421B626B';
    const mediatorAddress = '0xa513E6E4b8f2a923D98304ec87F64353C4D5C853';

    const chainIdBG1 = '5001'
    const chainIdBG2 = '5002'

    const bridgeApp = await ethers.getContractAt('BridgeApp', bridgeAppAddress);
    const mediator = await ethers.getContractAt('ERC20BridgeMediator', mediatorAddress);

    console.log("\nSet contract address");

    await(await bridgeApp.setContractAddress(chainIdBG1, erc20BridgeAddressBG1)).wait()
    await(await bridgeApp.setContractAddress(chainIdBG2, erc20BridgeAddressBG2)).wait()

    console.log("\nSet token");

    await(await mediator.addToken('CFN', chainIdBG1, tokenAddressBG1)).wait()
    await(await mediator.addToken('CFN', chainIdBG2, tokenAddressBG2)).wait()
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});