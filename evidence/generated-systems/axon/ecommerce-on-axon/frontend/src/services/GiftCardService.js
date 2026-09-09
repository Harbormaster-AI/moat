import axios from 'axios';

const GIFTCARD_API_BASE_URL = "/GiftCard";

class GiftCardService {

    getGiftCards(){
        return axios.get(GIFTCARD_API_BASE_URL + '/' );
    }

    createGiftCard(giftCard){
        return axios.post(GIFTCARD_API_BASE_URL  + '/create', giftCard);
    }

    getGiftCardById(giftCardId){
        return axios.get(GIFTCARD_API_BASE_URL + '/load?giftCardId=' + giftCardId);
    }

    updateGiftCard(giftCard){
        return axios.put(GIFTCARD_API_BASE_URL + '/update', giftCard);
    }

    deleteGiftCard(giftCardId){
        return axios.delete(GIFTCARD_API_BASE_URL + '/delete?giftCardId=' + giftCardId);
    }
}

export default new GiftCardService()