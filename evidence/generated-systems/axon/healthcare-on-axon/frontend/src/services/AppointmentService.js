import axios from 'axios';

const APPOINTMENT_API_BASE_URL = "/Appointment";

class AppointmentService {

    getAppointments(){
        return axios.get(APPOINTMENT_API_BASE_URL + '/' );
    }

    createAppointment(appointment){
        return axios.post(APPOINTMENT_API_BASE_URL  + '/create', appointment);
    }

    getAppointmentById(appointmentId){
        return axios.get(APPOINTMENT_API_BASE_URL + '/load?appointmentId=' + appointmentId);
    }

    updateAppointment(appointment){
        return axios.put(APPOINTMENT_API_BASE_URL + '/update', appointment);
    }

    deleteAppointment(appointmentId){
        return axios.delete(APPOINTMENT_API_BASE_URL + '/delete?appointmentId=' + appointmentId);
    }
}

export default new AppointmentService()