import axios from 'axios';

const KPI_API_BASE_URL = "/KPI";

class KPIService {

    getKPIs(){
        return axios.get(KPI_API_BASE_URL + '/' );
    }

    createKPI(kPI){
        return axios.post(KPI_API_BASE_URL  + '/create', kPI);
    }

    getKPIById(kPIId){
        return axios.get(KPI_API_BASE_URL + '/load?kPIId=' + kPIId);
    }

    updateKPI(kPI){
        return axios.put(KPI_API_BASE_URL + '/update', kPI);
    }

    deleteKPI(kPIId){
        return axios.delete(KPI_API_BASE_URL + '/delete?kPIId=' + kPIId);
    }
}

export default new KPIService()