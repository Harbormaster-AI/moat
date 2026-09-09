import axios from 'axios';

const CANDIDATE_API_BASE_URL = "/Candidate";

class CandidateService {

    getCandidates(){
        return axios.get(CANDIDATE_API_BASE_URL + '/' );
    }

    createCandidate(candidate){
        return axios.post(CANDIDATE_API_BASE_URL  + '/create', candidate);
    }

    getCandidateById(candidateId){
        return axios.get(CANDIDATE_API_BASE_URL + '/load?candidateId=' + candidateId);
    }

    updateCandidate(candidate){
        return axios.put(CANDIDATE_API_BASE_URL + '/update', candidate);
    }

    deleteCandidate(candidateId){
        return axios.delete(CANDIDATE_API_BASE_URL + '/delete?candidateId=' + candidateId);
    }
}

export default new CandidateService()