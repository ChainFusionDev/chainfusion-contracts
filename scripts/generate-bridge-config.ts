import { readContractsConfig, createBridgeConfig, writeBridgeConfig } from './deploy/bridge-config';
import { ContractsConfig } from './deploy/config';
import glob from 'glob';

async function main() {
   glob('contracts-*.json', async function (err, files) {
    var configs: ContractsConfig[] = [];

    if (err) {
      console.log(err);
      return err
    } else {
      for (const path of files) {
        const config = await readContractsConfig(path);

        configs.push(config);
      }

    const bConfig = await createBridgeConfig(configs)

    await writeBridgeConfig(bConfig)
    }
  });
}
  
  main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
  });