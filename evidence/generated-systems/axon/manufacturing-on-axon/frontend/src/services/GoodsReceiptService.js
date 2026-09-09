import axios from 'axios';

const GOODSRECEIPT_API_BASE_URL = "/GoodsReceipt";

class GoodsReceiptService {

    getGoodsReceipts(){
        return axios.get(GOODSRECEIPT_API_BASE_URL + '/' );
    }

    createGoodsReceipt(goodsReceipt){
        return axios.post(GOODSRECEIPT_API_BASE_URL  + '/create', goodsReceipt);
    }

    getGoodsReceiptById(goodsReceiptId){
        return axios.get(GOODSRECEIPT_API_BASE_URL + '/load?goodsReceiptId=' + goodsReceiptId);
    }

    updateGoodsReceipt(goodsReceipt){
        return axios.put(GOODSRECEIPT_API_BASE_URL + '/update', goodsReceipt);
    }

    deleteGoodsReceipt(goodsReceiptId){
        return axios.delete(GOODSRECEIPT_API_BASE_URL + '/delete?goodsReceiptId=' + goodsReceiptId);
    }
}

export default new GoodsReceiptService()