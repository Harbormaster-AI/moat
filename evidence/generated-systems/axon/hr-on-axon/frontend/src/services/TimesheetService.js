import axios from 'axios';

const TIMESHEET_API_BASE_URL = "/Timesheet";

class TimesheetService {

    getTimesheets(){
        return axios.get(TIMESHEET_API_BASE_URL + '/' );
    }

    createTimesheet(timesheet){
        return axios.post(TIMESHEET_API_BASE_URL  + '/create', timesheet);
    }

    getTimesheetById(timesheetId){
        return axios.get(TIMESHEET_API_BASE_URL + '/load?timesheetId=' + timesheetId);
    }

    updateTimesheet(timesheet){
        return axios.put(TIMESHEET_API_BASE_URL + '/update', timesheet);
    }

    deleteTimesheet(timesheetId){
        return axios.delete(TIMESHEET_API_BASE_URL + '/delete?timesheetId=' + timesheetId);
    }
}

export default new TimesheetService()