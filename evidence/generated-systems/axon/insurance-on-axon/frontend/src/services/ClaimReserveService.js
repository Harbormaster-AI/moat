import axios from 'axios';

const CLAIMRESERVE_API_BASE_URL = "/ClaimReserve";

class ClaimReserveService {

    getClaimReserves(){
        return axios.get(CLAIMRESERVE_API_BASE_URL + '/' );
    }

    createClaimReserve(claimReserve){
        return axios.post(CLAIMRESERVE_API_BASE_URL  + '/create', claimReserve);
    }

    getClaimReserveById(claimReserveId){
        return axios.get(CLAIMRESERVE_API_BASE_URL + '/load?claimReserveId=' + claimReserveId);
    }

    updateClaimReserve(claimReserve){
        return axios.put(CLAIMRESERVE_API_BASE_URL + '/update', claimReserve);
    }

    deleteClaimReserve(claimReserveId){
        return axios.delete(CLAIMRESERVE_API_BASE_URL + '/delete?claimReserveId=' + claimReserveId);
    }
}

export default new ClaimReserveService()