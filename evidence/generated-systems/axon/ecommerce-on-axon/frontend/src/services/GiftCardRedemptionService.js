import axios from 'axios';

const GIFTCARDREDEMPTION_API_BASE_URL = "/GiftCardRedemption";

class GiftCardRedemptionService {

    getGiftCardRedemptions(){
        return axios.get(GIFTCARDREDEMPTION_API_BASE_URL + '/' );
    }

    createGiftCardRedemption(giftCardRedemption){
        return axios.post(GIFTCARDREDEMPTION_API_BASE_URL  + '/create', giftCardRedemption);
    }

    getGiftCardRedemptionById(giftCardRedemptionId){
        return axios.get(GIFTCARDREDEMPTION_API_BASE_URL + '/load?giftCardRedemptionId=' + giftCardRedemptionId);
    }

    updateGiftCardRedemption(giftCardRedemption){
        return axios.put(GIFTCARDREDEMPTION_API_BASE_URL + '/update', giftCardRedemption);
    }

    deleteGiftCardRedemption(giftCardRedemptionId){
        return axios.delete(GIFTCARDREDEMPTION_API_BASE_URL + '/delete?giftCardRedemptionId=' + giftCardRedemptionId);
    }
}

export default new GiftCardRedemptionService()