import axios from 'axios';

const WALLET_API_BASE_URL = "/Wallet";

class WalletService {

    getWallets(){
        return axios.get(WALLET_API_BASE_URL + '/' );
    }

    createWallet(wallet){
        return axios.post(WALLET_API_BASE_URL  + '/create', wallet);
    }

    getWalletById(walletId){
        return axios.get(WALLET_API_BASE_URL + '/load?walletId=' + walletId);
    }

    updateWallet(wallet){
        return axios.put(WALLET_API_BASE_URL + '/update', wallet);
    }

    deleteWallet(walletId){
        return axios.delete(WALLET_API_BASE_URL + '/delete?walletId=' + walletId);
    }
}

export default new WalletService()