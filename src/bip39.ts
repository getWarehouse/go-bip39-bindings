import { dirname, Go, join } from "../deps.ts";

let go = new Go();
let inst = await WebAssembly.instantiate(
  Deno.readFileSync(
    join(dirname(import.meta.url), "bip39.wasm").replace("file:", ""),
  ),
  go.importObject,
);
let promise = go.run(inst.instance);

type Key = {
  Key: string;
  Version: string;
  ChildNumber: string;
  FingerPrint: string;
  ChainCode: string;
  Depth: number;
  IsPrivate: boolean;
};

class Bip39 {
  constructor() {
  }

  static NewCrypto(size: 256 | 128, passPhrase: string): {
    "entropy": string;
    "mnemonic": string;
    "seed": string;
    "masterKey": Key;
    "publicKey": Key;
  } {
    return JSON.parse(go.exports.NewCrypto(size, passPhrase));
  }
}

export { Bip39 };
