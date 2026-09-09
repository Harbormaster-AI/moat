import axios from 'axios';

const REPORT_API_BASE_URL = "/Report";

class ReportService {

    getReports(){
        return axios.get(REPORT_API_BASE_URL + '/' );
    }

    createReport(report){
        return axios.post(REPORT_API_BASE_URL  + '/create', report);
    }

    getReportById(reportId){
        return axios.get(REPORT_API_BASE_URL + '/load?reportId=' + reportId);
    }

    updateReport(report){
        return axios.put(REPORT_API_BASE_URL + '/update', report);
    }

    deleteReport(reportId){
        return axios.delete(REPORT_API_BASE_URL + '/delete?reportId=' + reportId);
    }
}

export default new ReportService()