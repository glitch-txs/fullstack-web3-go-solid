import { ethers } from 'ethers'
import create from 'zustand'
import { connectToMetamask } from '../utils/metamask/connectMetamask'
import { removeEventsMetamask } from '../utils/metamask/helpers/eventListeners'
import { metamaskInit } from '../utils/metamask/metamaskInit'
import { checkChainAndAccount } from '../utils/walletconnect/helper/checkChainAndAccount'
import { openWCModal } from '../utils/walletconnect/WCConnect'
import { WCInit } from '../utils/walletconnect/WCInit'

//WC stands for Walletconnect

export type ContractInfo = {
    address: string
    abi: any
}

export type CallInfo = {
    name: string
    params: any[]
    action: 'read' | 'write'  
}

interface Web3Store {
    //We need a time for the WC init to load
    isLoading: boolean
    // Modal will trigger the modal and show specific warning depending on the web3 states status.
    modal: '' | 'provider' | 'chain' | 'connect'
    isProvider: boolean
    //if WC connect init fails to connect this will be false
    WCInitFailed: boolean
    userAccount: string
    chainId: boolean
    Provider: any
    childProvider: any

    clearModal: ()=>void

    web3Init: ()=> void
    connectMetamask: ()=> void
    connectWC: ()=> void
    disconnectWC: ()=> void
    callContract: (contractInfo: ContractInfo, callInfo: CallInfo, setStatus?: (status: string)=> void)=> any
    restartWeb3: ()=> void
}


export const useWeb3Store = create<Web3Store>()((set, get) => ({
    isLoading: true,
    modal: '',
    isProvider: true,
    WCInitFailed: false,
    userAccount: '',
    chainId: false,
    Provider: null,
    childProvider: null,

    
}))