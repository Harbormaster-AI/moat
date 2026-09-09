import axios from 'axios';

const EVIDENCE_API_BASE_URL = "/Evidence";

class EvidenceService {

    getEvidences(){
        return axios.get(EVIDENCE_API_BASE_URL + '/' );
    }

    createEvidence(evidence){
        return axios.post(EVIDENCE_API_BASE_URL  + '/create', evidence);
    }

    getEvidenceById(evidenceId){
        return axios.get(EVIDENCE_API_BASE_URL + '/load?evidenceId=' + evidenceId);
    }

    updateEvidence(evidence){
        return axios.put(EVIDENCE_API_BASE_URL + '/update', evidence);
    }

    deleteEvidence(evidenceId){
        return axios.delete(EVIDENCE_API_BASE_URL + '/delete?evidenceId=' + evidenceId);
    }
}

export default new EvidenceService()