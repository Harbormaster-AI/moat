import axios from 'axios';

const LOT_API_BASE_URL = "/Lot";

class LotService {

    getLots(){
        return axios.get(LOT_API_BASE_URL + '/' );
    }

    createLot(lot){
        return axios.post(LOT_API_BASE_URL  + '/create', lot);
    }

    getLotById(lotId){
        return axios.get(LOT_API_BASE_URL + '/load?lotId=' + lotId);
    }

    updateLot(lot){
        return axios.put(LOT_API_BASE_URL + '/update', lot);
    }

    deleteLot(lotId){
        return axios.delete(LOT_API_BASE_URL + '/delete?lotId=' + lotId);
    }
}

export default new LotService()