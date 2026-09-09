import axios from 'axios';

const SALESREGION_API_BASE_URL = "/SalesRegion";

class SalesRegionService {

    getSalesRegions(){
        return axios.get(SALESREGION_API_BASE_URL + '/' );
    }

    createSalesRegion(salesRegion){
        return axios.post(SALESREGION_API_BASE_URL  + '/create', salesRegion);
    }

    getSalesRegionById(salesRegionId){
        return axios.get(SALESREGION_API_BASE_URL + '/load?salesRegionId=' + salesRegionId);
    }

    updateSalesRegion(salesRegion){
        return axios.put(SALESREGION_API_BASE_URL + '/update', salesRegion);
    }

    deleteSalesRegion(salesRegionId){
        return axios.delete(SALESREGION_API_BASE_URL + '/delete?salesRegionId=' + salesRegionId);
    }
}

export default new SalesRegionService()