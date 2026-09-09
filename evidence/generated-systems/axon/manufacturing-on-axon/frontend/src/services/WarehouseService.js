import axios from 'axios';

const WAREHOUSE_API_BASE_URL = "/Warehouse";

class WarehouseService {

    getWarehouses(){
        return axios.get(WAREHOUSE_API_BASE_URL + '/' );
    }

    createWarehouse(warehouse){
        return axios.post(WAREHOUSE_API_BASE_URL  + '/create', warehouse);
    }

    getWarehouseById(warehouseId){
        return axios.get(WAREHOUSE_API_BASE_URL + '/load?warehouseId=' + warehouseId);
    }

    updateWarehouse(warehouse){
        return axios.put(WAREHOUSE_API_BASE_URL + '/update', warehouse);
    }

    deleteWarehouse(warehouseId){
        return axios.delete(WAREHOUSE_API_BASE_URL + '/delete?warehouseId=' + warehouseId);
    }
}

export default new WarehouseService()