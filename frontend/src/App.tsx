import type { Component } from 'solid-js';
import styles from './App.module.css';
import { UniversalProvider } from "@walletconnect/universal-provider";
import { Web3Modal } from "@web3modal/standalone";

const web3Modal = new Web3Modal({ 
  // projectId: import.meta.env.VITE_PROJECT_ID
})

web3Modal.setTheme({
  themeMode: "light",
  themeColor: "blackWhite",
  themeBackground: "gradient",
});

const App: Component = () => {


    
console.log(import.meta.env.VITE_PROJECT_ID)

  const start = async()=>{

    console.log("start")

    const provider = await UniversalProvider.init({
      projectId: import.meta.env.VITE_PROJECT_ID,
      metadata: {
        name: "Glitch Dapp",
        description: "Glitch Dapp",
        url: "mywebsite.com",
        icons: ["https://lh3.googleusercontent.com/ogw/AOh-ky0c2alK5GAwefGWkwQHVpcJR637KRzHSZx9dV31rg=s32-c-mo"],
      },
      //This catch is agressive, as it prevents the use of Metamask if WC failed to connect
    }).catch(e=>console.log(e))
    
    provider?.on("display_uri", async (uri: any) => {
      web3Modal?.openModal({ uri });
    });

    console.log("second", provider)

    await provider?.connect({
      namespaces: {
        eip155: {
          methods: [
            "eth_sendTransaction",
            "eth_signTransaction",
            "eth_sign",
            "personal_sign",
            "eth_signTypedData",
          ],
          chains: ["eip155:56"],
          events: ["chainChanged", "accountsChanged"],
          rpcMap: {
            56: 'https://bsc-dataseed1.binance.org/',
          },
        },
      },
    }).catch((e: any)=> console.log(e))
  
    web3Modal?.closeModal();
  }

  console.log("I rendered!")


  return (
    <div class={styles.App}>
        <button onClick={start} class={styles.button} >Connect Wallet</button>
    </div>
  );
};

export default App;
