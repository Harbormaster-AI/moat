import axios from 'axios';

const DASHBOARD_API_BASE_URL = "/Dashboard";

class DashboardService {

    getDashboards(){
        return axios.get(DASHBOARD_API_BASE_URL + '/' );
    }

    createDashboard(dashboard){
        return axios.post(DASHBOARD_API_BASE_URL  + '/create', dashboard);
    }

    getDashboardById(dashboardId){
        return axios.get(DASHBOARD_API_BASE_URL + '/load?dashboardId=' + dashboardId);
    }

    updateDashboard(dashboard){
        return axios.put(DASHBOARD_API_BASE_URL + '/update', dashboard);
    }

    deleteDashboard(dashboardId){
        return axios.delete(DASHBOARD_API_BASE_URL + '/delete?dashboardId=' + dashboardId);
    }
}

export default new DashboardService()