import axios from 'axios';

const LINEAGENODE_API_BASE_URL = "/LineageNode";

class LineageNodeService {

    getLineageNodes(){
        return axios.get(LINEAGENODE_API_BASE_URL + '/' );
    }

    createLineageNode(lineageNode){
        return axios.post(LINEAGENODE_API_BASE_URL  + '/create', lineageNode);
    }

    getLineageNodeById(lineageNodeId){
        return axios.get(LINEAGENODE_API_BASE_URL + '/load?lineageNodeId=' + lineageNodeId);
    }

    updateLineageNode(lineageNode){
        return axios.put(LINEAGENODE_API_BASE_URL + '/update', lineageNode);
    }

    deleteLineageNode(lineageNodeId){
        return axios.delete(LINEAGENODE_API_BASE_URL + '/delete?lineageNodeId=' + lineageNodeId);
    }
}

export default new LineageNodeService()