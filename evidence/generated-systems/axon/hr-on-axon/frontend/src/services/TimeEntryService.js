import axios from 'axios';

const TIMEENTRY_API_BASE_URL = "/TimeEntry";

class TimeEntryService {

    getTimeEntrys(){
        return axios.get(TIMEENTRY_API_BASE_URL + '/' );
    }

    createTimeEntry(timeEntry){
        return axios.post(TIMEENTRY_API_BASE_URL  + '/create', timeEntry);
    }

    getTimeEntryById(timeEntryId){
        return axios.get(TIMEENTRY_API_BASE_URL + '/load?timeEntryId=' + timeEntryId);
    }

    updateTimeEntry(timeEntry){
        return axios.put(TIMEENTRY_API_BASE_URL + '/update', timeEntry);
    }

    deleteTimeEntry(timeEntryId){
        return axios.delete(TIMEENTRY_API_BASE_URL + '/delete?timeEntryId=' + timeEntryId);
    }
}

export default new TimeEntryService()