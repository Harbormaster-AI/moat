import axios from 'axios';

const ADSLOT_API_BASE_URL = "/AdSlot";

class AdSlotService {

    getAdSlots(){
        return axios.get(ADSLOT_API_BASE_URL + '/' );
    }

    createAdSlot(adSlot){
        return axios.post(ADSLOT_API_BASE_URL  + '/create', adSlot);
    }

    getAdSlotById(adSlotId){
        return axios.get(ADSLOT_API_BASE_URL + '/load?adSlotId=' + adSlotId);
    }

    updateAdSlot(adSlot){
        return axios.put(ADSLOT_API_BASE_URL + '/update', adSlot);
    }

    deleteAdSlot(adSlotId){
        return axios.delete(ADSLOT_API_BASE_URL + '/delete?adSlotId=' + adSlotId);
    }
}

export default new AdSlotService()