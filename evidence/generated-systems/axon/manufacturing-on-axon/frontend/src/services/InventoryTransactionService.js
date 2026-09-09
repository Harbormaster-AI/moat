import axios from 'axios';

const INVENTORYTRANSACTION_API_BASE_URL = "/InventoryTransaction";

class InventoryTransactionService {

    getInventoryTransactions(){
        return axios.get(INVENTORYTRANSACTION_API_BASE_URL + '/' );
    }

    createInventoryTransaction(inventoryTransaction){
        return axios.post(INVENTORYTRANSACTION_API_BASE_URL  + '/create', inventoryTransaction);
    }

    getInventoryTransactionById(inventoryTransactionId){
        return axios.get(INVENTORYTRANSACTION_API_BASE_URL + '/load?inventoryTransactionId=' + inventoryTransactionId);
    }

    updateInventoryTransaction(inventoryTransaction){
        return axios.put(INVENTORYTRANSACTION_API_BASE_URL + '/update', inventoryTransaction);
    }

    deleteInventoryTransaction(inventoryTransactionId){
        return axios.delete(INVENTORYTRANSACTION_API_BASE_URL + '/delete?inventoryTransactionId=' + inventoryTransactionId);
    }
}

export default new InventoryTransactionService()