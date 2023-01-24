import type { Component } from 'solid-js';
import styles from './App.module.css';
import { UniversalProvider } from "@walletconnect/universal-provider";
import { Web3Modal } from "@web3modal/standalone";
import { getBalance } from './query/getData';

const web3Modal = new Web3Modal({ 
  // projectId: import.meta.env.VITE_PROJECT_ID
})

web3Modal.setTheme({
  themeMode: "light",
  themeColor: "blackWhite",
  themeBackground: "gradient",
});

const App: Component = () => {

  const callAPI = async ()=>{

    const options = {
      userAddress: "0x8245928bf6bef0b87d893965e8a5b008561020e4",
      contractAddress: "0x1fe84fE4e1ae96F9b202188f7a6835dB3D27a264"
    }
    const data = await getBalance(options)
    console.log(data)
  }

  return (
    <div class={styles.App}>
        <button onClick={callAPI} class={styles.button} >Connect Wallet</button>
    </div>
  );
};

export default App;
