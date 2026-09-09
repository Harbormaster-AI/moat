import axios from 'axios';

const VERIFIEDADDRESS_API_BASE_URL = "/VerifiedAddress";

class VerifiedAddressService {

    getVerifiedAddresss(){
        return axios.get(VERIFIEDADDRESS_API_BASE_URL + '/' );
    }

    createVerifiedAddress(verifiedAddress){
        return axios.post(VERIFIEDADDRESS_API_BASE_URL  + '/create', verifiedAddress);
    }

    getVerifiedAddressById(verifiedAddressId){
        return axios.get(VERIFIEDADDRESS_API_BASE_URL + '/load?verifiedAddressId=' + verifiedAddressId);
    }

    updateVerifiedAddress(verifiedAddress){
        return axios.put(VERIFIEDADDRESS_API_BASE_URL + '/update', verifiedAddress);
    }

    deleteVerifiedAddress(verifiedAddressId){
        return axios.delete(VERIFIEDADDRESS_API_BASE_URL + '/delete?verifiedAddressId=' + verifiedAddressId);
    }
}

export default new VerifiedAddressService()