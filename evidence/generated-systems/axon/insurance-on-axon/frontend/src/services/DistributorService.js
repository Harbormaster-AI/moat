import axios from 'axios';

const DISTRIBUTOR_API_BASE_URL = "/Distributor";

class DistributorService {

    getDistributors(){
        return axios.get(DISTRIBUTOR_API_BASE_URL + '/' );
    }

    createDistributor(distributor){
        return axios.post(DISTRIBUTOR_API_BASE_URL  + '/create', distributor);
    }

    getDistributorById(distributorId){
        return axios.get(DISTRIBUTOR_API_BASE_URL + '/load?distributorId=' + distributorId);
    }

    updateDistributor(distributor){
        return axios.put(DISTRIBUTOR_API_BASE_URL + '/update', distributor);
    }

    deleteDistributor(distributorId){
        return axios.delete(DISTRIBUTOR_API_BASE_URL + '/delete?distributorId=' + distributorId);
    }
}

export default new DistributorService()