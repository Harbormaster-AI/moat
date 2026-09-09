import axios from 'axios';

const AUDIENCESEGMENT_API_BASE_URL = "/AudienceSegment";

class AudienceSegmentService {

    getAudienceSegments(){
        return axios.get(AUDIENCESEGMENT_API_BASE_URL + '/' );
    }

    createAudienceSegment(audienceSegment){
        return axios.post(AUDIENCESEGMENT_API_BASE_URL  + '/create', audienceSegment);
    }

    getAudienceSegmentById(audienceSegmentId){
        return axios.get(AUDIENCESEGMENT_API_BASE_URL + '/load?audienceSegmentId=' + audienceSegmentId);
    }

    updateAudienceSegment(audienceSegment){
        return axios.put(AUDIENCESEGMENT_API_BASE_URL + '/update', audienceSegment);
    }

    deleteAudienceSegment(audienceSegmentId){
        return axios.delete(AUDIENCESEGMENT_API_BASE_URL + '/delete?audienceSegmentId=' + audienceSegmentId);
    }
}

export default new AudienceSegmentService()