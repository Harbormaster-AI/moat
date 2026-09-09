import axios from 'axios';

const DIMENSION_API_BASE_URL = "/Dimension";

class DimensionService {

    getDimensions(){
        return axios.get(DIMENSION_API_BASE_URL + '/' );
    }

    createDimension(dimension){
        return axios.post(DIMENSION_API_BASE_URL  + '/create', dimension);
    }

    getDimensionById(dimensionId){
        return axios.get(DIMENSION_API_BASE_URL + '/load?dimensionId=' + dimensionId);
    }

    updateDimension(dimension){
        return axios.put(DIMENSION_API_BASE_URL + '/update', dimension);
    }

    deleteDimension(dimensionId){
        return axios.delete(DIMENSION_API_BASE_URL + '/delete?dimensionId=' + dimensionId);
    }
}

export default new DimensionService()