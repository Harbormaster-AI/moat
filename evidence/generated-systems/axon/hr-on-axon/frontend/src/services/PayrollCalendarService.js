import axios from 'axios';

const PAYROLLCALENDAR_API_BASE_URL = "/PayrollCalendar";

class PayrollCalendarService {

    getPayrollCalendars(){
        return axios.get(PAYROLLCALENDAR_API_BASE_URL + '/' );
    }

    createPayrollCalendar(payrollCalendar){
        return axios.post(PAYROLLCALENDAR_API_BASE_URL  + '/create', payrollCalendar);
    }

    getPayrollCalendarById(payrollCalendarId){
        return axios.get(PAYROLLCALENDAR_API_BASE_URL + '/load?payrollCalendarId=' + payrollCalendarId);
    }

    updatePayrollCalendar(payrollCalendar){
        return axios.put(PAYROLLCALENDAR_API_BASE_URL + '/update', payrollCalendar);
    }

    deletePayrollCalendar(payrollCalendarId){
        return axios.delete(PAYROLLCALENDAR_API_BASE_URL + '/delete?payrollCalendarId=' + payrollCalendarId);
    }
}

export default new PayrollCalendarService()