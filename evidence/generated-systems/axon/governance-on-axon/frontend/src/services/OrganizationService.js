import axios from 'axios';

const ORGANIZATION_API_BASE_URL = "/Organization";

class OrganizationService {

    getOrganizations(){
        return axios.get(ORGANIZATION_API_BASE_URL + '/' );
    }

    createOrganization(organization){
        return axios.post(ORGANIZATION_API_BASE_URL  + '/create', organization);
    }

    getOrganizationById(organizationId){
        return axios.get(ORGANIZATION_API_BASE_URL + '/load?organizationId=' + organizationId);
    }

    updateOrganization(organization){
        return axios.put(ORGANIZATION_API_BASE_URL + '/update', organization);
    }

    deleteOrganization(organizationId){
        return axios.delete(ORGANIZATION_API_BASE_URL + '/delete?organizationId=' + organizationId);
    }
}

export default new OrganizationService()