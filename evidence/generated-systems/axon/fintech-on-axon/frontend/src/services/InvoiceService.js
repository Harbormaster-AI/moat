import axios from 'axios';

const INVOICE_API_BASE_URL = "/Invoice";

class InvoiceService {

    getInvoices(){
        return axios.get(INVOICE_API_BASE_URL + '/' );
    }

    createInvoice(invoice){
        return axios.post(INVOICE_API_BASE_URL  + '/create', invoice);
    }

    getInvoiceById(invoiceId){
        return axios.get(INVOICE_API_BASE_URL + '/load?invoiceId=' + invoiceId);
    }

    updateInvoice(invoice){
        return axios.put(INVOICE_API_BASE_URL + '/update', invoice);
    }

    deleteInvoice(invoiceId){
        return axios.delete(INVOICE_API_BASE_URL + '/delete?invoiceId=' + invoiceId);
    }
}

export default new InvoiceService()