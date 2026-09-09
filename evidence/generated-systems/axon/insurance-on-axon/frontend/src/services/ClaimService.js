import axios from 'axios';

const CLAIM_API_BASE_URL = "/Claim";

class ClaimService {

    getClaims(){
        return axios.get(CLAIM_API_BASE_URL + '/' );
    }

    createClaim(claim){
        return axios.post(CLAIM_API_BASE_URL  + '/create', claim);
    }

    getClaimById(claimId){
        return axios.get(CLAIM_API_BASE_URL + '/load?claimId=' + claimId);
    }

    updateClaim(claim){
        return axios.put(CLAIM_API_BASE_URL + '/update', claim);
    }

    deleteClaim(claimId){
        return axios.delete(CLAIM_API_BASE_URL + '/delete?claimId=' + claimId);
    }
}

export default new ClaimService()