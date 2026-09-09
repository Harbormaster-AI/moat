import axios from 'axios';

const MAINTENANCEAPPOINTMENT_API_BASE_URL = "/MaintenanceAppointment";

class MaintenanceAppointmentService {

    getMaintenanceAppointments(){
        return axios.get(MAINTENANCEAPPOINTMENT_API_BASE_URL + '/' );
    }

    createMaintenanceAppointment(maintenanceAppointment){
        return axios.post(MAINTENANCEAPPOINTMENT_API_BASE_URL  + '/create', maintenanceAppointment);
    }

    getMaintenanceAppointmentById(maintenanceAppointmentId){
        return axios.get(MAINTENANCEAPPOINTMENT_API_BASE_URL + '/load?maintenanceAppointmentId=' + maintenanceAppointmentId);
    }

    updateMaintenanceAppointment(maintenanceAppointment){
        return axios.put(MAINTENANCEAPPOINTMENT_API_BASE_URL + '/update', maintenanceAppointment);
    }

    deleteMaintenanceAppointment(maintenanceAppointmentId){
        return axios.delete(MAINTENANCEAPPOINTMENT_API_BASE_URL + '/delete?maintenanceAppointmentId=' + maintenanceAppointmentId);
    }
}

export default new MaintenanceAppointmentService()