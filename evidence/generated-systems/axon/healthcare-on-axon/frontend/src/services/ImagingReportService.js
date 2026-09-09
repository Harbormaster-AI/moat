import axios from 'axios';

const IMAGINGREPORT_API_BASE_URL = "/ImagingReport";

class ImagingReportService {

    getImagingReports(){
        return axios.get(IMAGINGREPORT_API_BASE_URL + '/' );
    }

    createImagingReport(imagingReport){
        return axios.post(IMAGINGREPORT_API_BASE_URL  + '/create', imagingReport);
    }

    getImagingReportById(imagingReportId){
        return axios.get(IMAGINGREPORT_API_BASE_URL + '/load?imagingReportId=' + imagingReportId);
    }

    updateImagingReport(imagingReport){
        return axios.put(IMAGINGREPORT_API_BASE_URL + '/update', imagingReport);
    }

    deleteImagingReport(imagingReportId){
        return axios.delete(IMAGINGREPORT_API_BASE_URL + '/delete?imagingReportId=' + imagingReportId);
    }
}

export default new ImagingReportService()