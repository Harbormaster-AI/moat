import axios from 'axios';

const BUSINESSUNIT_API_BASE_URL = "/BusinessUnit";

class BusinessUnitService {

    getBusinessUnits(){
        return axios.get(BUSINESSUNIT_API_BASE_URL + '/' );
    }

    createBusinessUnit(businessUnit){
        return axios.post(BUSINESSUNIT_API_BASE_URL  + '/create', businessUnit);
    }

    getBusinessUnitById(businessUnitId){
        return axios.get(BUSINESSUNIT_API_BASE_URL + '/load?businessUnitId=' + businessUnitId);
    }

    updateBusinessUnit(businessUnit){
        return axios.put(BUSINESSUNIT_API_BASE_URL + '/update', businessUnit);
    }

    deleteBusinessUnit(businessUnitId){
        return axios.delete(BUSINESSUNIT_API_BASE_URL + '/delete?businessUnitId=' + businessUnitId);
    }
}

export default new BusinessUnitService()