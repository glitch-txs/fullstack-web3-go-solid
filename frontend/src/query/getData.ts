import api from './api'

type Props = {
    contractAddress: string,
    userAddress: string
}

export const getBalance = async ({ contractAddress, userAddress }: Props)=>{

    try {
        const res = await api.get(`/balance?contract=${contractAddress}&user=${userAddress}`)

        return res.data

    } catch (er){
        console.error(er)
        return '-'
    }

}
