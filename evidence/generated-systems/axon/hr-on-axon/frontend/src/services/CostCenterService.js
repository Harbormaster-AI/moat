import axios from 'axios';

const COSTCENTER_API_BASE_URL = "/CostCenter";

class CostCenterService {

    getCostCenters(){
        return axios.get(COSTCENTER_API_BASE_URL + '/' );
    }

    createCostCenter(costCenter){
        return axios.post(COSTCENTER_API_BASE_URL  + '/create', costCenter);
    }

    getCostCenterById(costCenterId){
        return axios.get(COSTCENTER_API_BASE_URL + '/load?costCenterId=' + costCenterId);
    }

    updateCostCenter(costCenter){
        return axios.put(COSTCENTER_API_BASE_URL + '/update', costCenter);
    }

    deleteCostCenter(costCenterId){
        return axios.delete(COSTCENTER_API_BASE_URL + '/delete?costCenterId=' + costCenterId);
    }
}

export default new CostCenterService()