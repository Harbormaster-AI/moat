import axios from 'axios';

const TERRITORY_API_BASE_URL = "/Territory";

class TerritoryService {

    getTerritorys(){
        return axios.get(TERRITORY_API_BASE_URL + '/' );
    }

    createTerritory(territory){
        return axios.post(TERRITORY_API_BASE_URL  + '/create', territory);
    }

    getTerritoryById(territoryId){
        return axios.get(TERRITORY_API_BASE_URL + '/load?territoryId=' + territoryId);
    }

    updateTerritory(territory){
        return axios.put(TERRITORY_API_BASE_URL + '/update', territory);
    }

    deleteTerritory(territoryId){
        return axios.delete(TERRITORY_API_BASE_URL + '/delete?territoryId=' + territoryId);
    }
}

export default new TerritoryService()