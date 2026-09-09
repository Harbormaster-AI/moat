import axios from 'axios';

const PLANT_API_BASE_URL = "/Plant";

class PlantService {

    getPlants(){
        return axios.get(PLANT_API_BASE_URL + '/' );
    }

    createPlant(plant){
        return axios.post(PLANT_API_BASE_URL  + '/create', plant);
    }

    getPlantById(plantId){
        return axios.get(PLANT_API_BASE_URL + '/load?plantId=' + plantId);
    }

    updatePlant(plant){
        return axios.put(PLANT_API_BASE_URL + '/update', plant);
    }

    deletePlant(plantId){
        return axios.delete(PLANT_API_BASE_URL + '/delete?plantId=' + plantId);
    }
}

export default new PlantService()