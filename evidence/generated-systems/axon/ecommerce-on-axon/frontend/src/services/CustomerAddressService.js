import axios from 'axios';

const CUSTOMERADDRESS_API_BASE_URL = "/CustomerAddress";

class CustomerAddressService {

    getCustomerAddresss(){
        return axios.get(CUSTOMERADDRESS_API_BASE_URL + '/' );
    }

    createCustomerAddress(customerAddress){
        return axios.post(CUSTOMERADDRESS_API_BASE_URL  + '/create', customerAddress);
    }

    getCustomerAddressById(customerAddressId){
        return axios.get(CUSTOMERADDRESS_API_BASE_URL + '/load?customerAddressId=' + customerAddressId);
    }

    updateCustomerAddress(customerAddress){
        return axios.put(CUSTOMERADDRESS_API_BASE_URL + '/update', customerAddress);
    }

    deleteCustomerAddress(customerAddressId){
        return axios.delete(CUSTOMERADDRESS_API_BASE_URL + '/delete?customerAddressId=' + customerAddressId);
    }
}

export default new CustomerAddressService()