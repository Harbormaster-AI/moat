import axios from 'axios';

const ATTESTATION_API_BASE_URL = "/Attestation";

class AttestationService {

    getAttestations(){
        return axios.get(ATTESTATION_API_BASE_URL + '/' );
    }

    createAttestation(attestation){
        return axios.post(ATTESTATION_API_BASE_URL  + '/create', attestation);
    }

    getAttestationById(attestationId){
        return axios.get(ATTESTATION_API_BASE_URL + '/load?attestationId=' + attestationId);
    }

    updateAttestation(attestation){
        return axios.put(ATTESTATION_API_BASE_URL + '/update', attestation);
    }

    deleteAttestation(attestationId){
        return axios.delete(ATTESTATION_API_BASE_URL + '/delete?attestationId=' + attestationId);
    }
}

export default new AttestationService()