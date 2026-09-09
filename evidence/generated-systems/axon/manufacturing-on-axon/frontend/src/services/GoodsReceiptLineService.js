import axios from 'axios';

const GOODSRECEIPTLINE_API_BASE_URL = "/GoodsReceiptLine";

class GoodsReceiptLineService {

    getGoodsReceiptLines(){
        return axios.get(GOODSRECEIPTLINE_API_BASE_URL + '/' );
    }

    createGoodsReceiptLine(goodsReceiptLine){
        return axios.post(GOODSRECEIPTLINE_API_BASE_URL  + '/create', goodsReceiptLine);
    }

    getGoodsReceiptLineById(goodsReceiptLineId){
        return axios.get(GOODSRECEIPTLINE_API_BASE_URL + '/load?goodsReceiptLineId=' + goodsReceiptLineId);
    }

    updateGoodsReceiptLine(goodsReceiptLine){
        return axios.put(GOODSRECEIPTLINE_API_BASE_URL + '/update', goodsReceiptLine);
    }

    deleteGoodsReceiptLine(goodsReceiptLineId){
        return axios.delete(GOODSRECEIPTLINE_API_BASE_URL + '/delete?goodsReceiptLineId=' + goodsReceiptLineId);
    }
}

export default new GoodsReceiptLineService()