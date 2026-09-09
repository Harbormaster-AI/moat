import axios from 'axios';

const POSITION_API_BASE_URL = "/Position";

class PositionService {

    getPositions(){
        return axios.get(POSITION_API_BASE_URL + '/' );
    }

    createPosition(position){
        return axios.post(POSITION_API_BASE_URL  + '/create', position);
    }

    getPositionById(positionId){
        return axios.get(POSITION_API_BASE_URL + '/load?positionId=' + positionId);
    }

    updatePosition(position){
        return axios.put(POSITION_API_BASE_URL + '/update', position);
    }

    deletePosition(positionId){
        return axios.delete(POSITION_API_BASE_URL + '/delete?positionId=' + positionId);
    }
}

export default new PositionService()