import axios from 'axios';

const ANALYTICSWORKSPACE_API_BASE_URL = "/AnalyticsWorkspace";

class AnalyticsWorkspaceService {

    getAnalyticsWorkspaces(){
        return axios.get(ANALYTICSWORKSPACE_API_BASE_URL + '/' );
    }

    createAnalyticsWorkspace(analyticsWorkspace){
        return axios.post(ANALYTICSWORKSPACE_API_BASE_URL  + '/create', analyticsWorkspace);
    }

    getAnalyticsWorkspaceById(analyticsWorkspaceId){
        return axios.get(ANALYTICSWORKSPACE_API_BASE_URL + '/load?analyticsWorkspaceId=' + analyticsWorkspaceId);
    }

    updateAnalyticsWorkspace(analyticsWorkspace){
        return axios.put(ANALYTICSWORKSPACE_API_BASE_URL + '/update', analyticsWorkspace);
    }

    deleteAnalyticsWorkspace(analyticsWorkspaceId){
        return axios.delete(ANALYTICSWORKSPACE_API_BASE_URL + '/delete?analyticsWorkspaceId=' + analyticsWorkspaceId);
    }
}

export default new AnalyticsWorkspaceService()