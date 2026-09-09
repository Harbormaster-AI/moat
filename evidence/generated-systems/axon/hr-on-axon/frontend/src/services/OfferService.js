import axios from 'axios';

const OFFER_API_BASE_URL = "/Offer";

class OfferService {

    getOffers(){
        return axios.get(OFFER_API_BASE_URL + '/' );
    }

    createOffer(offer){
        return axios.post(OFFER_API_BASE_URL  + '/create', offer);
    }

    getOfferById(offerId){
        return axios.get(OFFER_API_BASE_URL + '/load?offerId=' + offerId);
    }

    updateOffer(offer){
        return axios.put(OFFER_API_BASE_URL + '/update', offer);
    }

    deleteOffer(offerId){
        return axios.delete(OFFER_API_BASE_URL + '/delete?offerId=' + offerId);
    }
}

export default new OfferService()