import { ethers } from "ethers"
import { setWeb3Store } from "../../../store/useWeb3"

export const checkChainAndAccount = async (provider: any)=>{

    const web3Provider = new ethers.providers.Web3Provider(provider)

    const signer = web3Provider.getSigner()

    const address = await signer.getAddress()

    console.log('WC: user account ',address)

    setWeb3Store({ userAccount: address })

    const chainId = await signer.getChainId()

    if(chainId != 56){
        setWeb3Store({ chainId: false })
        console.log('WC: invalid chain id')
    }else if(chainId == 56){
        setWeb3Store({ chainId: true })
        console.log('WC: valid chain id')
    }
}