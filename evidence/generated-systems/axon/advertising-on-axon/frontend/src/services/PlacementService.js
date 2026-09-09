import axios from 'axios';

const PLACEMENT_API_BASE_URL = "/Placement";

class PlacementService {

    getPlacements(){
        return axios.get(PLACEMENT_API_BASE_URL + '/' );
    }

    createPlacement(placement){
        return axios.post(PLACEMENT_API_BASE_URL  + '/create', placement);
    }

    getPlacementById(placementId){
        return axios.get(PLACEMENT_API_BASE_URL + '/load?placementId=' + placementId);
    }

    updatePlacement(placement){
        return axios.put(PLACEMENT_API_BASE_URL + '/update', placement);
    }

    deletePlacement(placementId){
        return axios.delete(PLACEMENT_API_BASE_URL + '/delete?placementId=' + placementId);
    }
}

export default new PlacementService()