import axios from 'axios';

const COVERAGE_API_BASE_URL = "/Coverage";

class CoverageService {

    getCoverages(){
        return axios.get(COVERAGE_API_BASE_URL + '/' );
    }

    createCoverage(coverage){
        return axios.post(COVERAGE_API_BASE_URL  + '/create', coverage);
    }

    getCoverageById(coverageId){
        return axios.get(COVERAGE_API_BASE_URL + '/load?coverageId=' + coverageId);
    }

    updateCoverage(coverage){
        return axios.put(COVERAGE_API_BASE_URL + '/update', coverage);
    }

    deleteCoverage(coverageId){
        return axios.delete(COVERAGE_API_BASE_URL + '/delete?coverageId=' + coverageId);
    }
}

export default new CoverageService()