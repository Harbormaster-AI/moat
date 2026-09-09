import axios from 'axios';

const KYCPROFILE_API_BASE_URL = "/KYCProfile";

class KYCProfileService {

    getKYCProfiles(){
        return axios.get(KYCPROFILE_API_BASE_URL + '/' );
    }

    createKYCProfile(kYCProfile){
        return axios.post(KYCPROFILE_API_BASE_URL  + '/create', kYCProfile);
    }

    getKYCProfileById(kYCProfileId){
        return axios.get(KYCPROFILE_API_BASE_URL + '/load?kYCProfileId=' + kYCProfileId);
    }

    updateKYCProfile(kYCProfile){
        return axios.put(KYCPROFILE_API_BASE_URL + '/update', kYCProfile);
    }

    deleteKYCProfile(kYCProfileId){
        return axios.delete(KYCPROFILE_API_BASE_URL + '/delete?kYCProfileId=' + kYCProfileId);
    }
}

export default new KYCProfileService()