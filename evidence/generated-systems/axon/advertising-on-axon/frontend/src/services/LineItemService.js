import axios from 'axios';

const LINEITEM_API_BASE_URL = "/LineItem";

class LineItemService {

    getLineItems(){
        return axios.get(LINEITEM_API_BASE_URL + '/' );
    }

    createLineItem(lineItem){
        return axios.post(LINEITEM_API_BASE_URL  + '/create', lineItem);
    }

    getLineItemById(lineItemId){
        return axios.get(LINEITEM_API_BASE_URL + '/load?lineItemId=' + lineItemId);
    }

    updateLineItem(lineItem){
        return axios.put(LINEITEM_API_BASE_URL + '/update', lineItem);
    }

    deleteLineItem(lineItemId){
        return axios.delete(LINEITEM_API_BASE_URL + '/delete?lineItemId=' + lineItemId);
    }
}

export default new LineItemService()