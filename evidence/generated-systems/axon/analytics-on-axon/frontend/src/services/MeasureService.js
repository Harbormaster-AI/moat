import axios from 'axios';

const MEASURE_API_BASE_URL = "/Measure";

class MeasureService {

    getMeasures(){
        return axios.get(MEASURE_API_BASE_URL + '/' );
    }

    createMeasure(measure){
        return axios.post(MEASURE_API_BASE_URL  + '/create', measure);
    }

    getMeasureById(measureId){
        return axios.get(MEASURE_API_BASE_URL + '/load?measureId=' + measureId);
    }

    updateMeasure(measure){
        return axios.put(MEASURE_API_BASE_URL + '/update', measure);
    }

    deleteMeasure(measureId){
        return axios.delete(MEASURE_API_BASE_URL + '/delete?measureId=' + measureId);
    }
}

export default new MeasureService()