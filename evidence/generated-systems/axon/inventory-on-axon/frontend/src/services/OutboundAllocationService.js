import axios from 'axios';

const OUTBOUNDALLOCATION_API_BASE_URL = "/OutboundAllocation";

class OutboundAllocationService {

    getOutboundAllocations(){
        return axios.get(OUTBOUNDALLOCATION_API_BASE_URL + '/' );
    }

    createOutboundAllocation(outboundAllocation){
        return axios.post(OUTBOUNDALLOCATION_API_BASE_URL  + '/create', outboundAllocation);
    }

    getOutboundAllocationById(outboundAllocationId){
        return axios.get(OUTBOUNDALLOCATION_API_BASE_URL + '/load?outboundAllocationId=' + outboundAllocationId);
    }

    updateOutboundAllocation(outboundAllocation){
        return axios.put(OUTBOUNDALLOCATION_API_BASE_URL + '/update', outboundAllocation);
    }

    deleteOutboundAllocation(outboundAllocationId){
        return axios.delete(OUTBOUNDALLOCATION_API_BASE_URL + '/delete?outboundAllocationId=' + outboundAllocationId);
    }
}

export default new OutboundAllocationService()