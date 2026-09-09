import axios from 'axios';

const MEDICALDEVICE_API_BASE_URL = "/MedicalDevice";

class MedicalDeviceService {

    getMedicalDevices(){
        return axios.get(MEDICALDEVICE_API_BASE_URL + '/' );
    }

    createMedicalDevice(medicalDevice){
        return axios.post(MEDICALDEVICE_API_BASE_URL  + '/create', medicalDevice);
    }

    getMedicalDeviceById(medicalDeviceId){
        return axios.get(MEDICALDEVICE_API_BASE_URL + '/load?medicalDeviceId=' + medicalDeviceId);
    }

    updateMedicalDevice(medicalDevice){
        return axios.put(MEDICALDEVICE_API_BASE_URL + '/update', medicalDevice);
    }

    deleteMedicalDevice(medicalDeviceId){
        return axios.delete(MEDICALDEVICE_API_BASE_URL + '/delete?medicalDeviceId=' + medicalDeviceId);
    }
}

export default new MedicalDeviceService()