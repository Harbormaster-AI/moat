import axios from 'axios';

const VISUALIZATION_API_BASE_URL = "/Visualization";

class VisualizationService {

    getVisualizations(){
        return axios.get(VISUALIZATION_API_BASE_URL + '/' );
    }

    createVisualization(visualization){
        return axios.post(VISUALIZATION_API_BASE_URL  + '/create', visualization);
    }

    getVisualizationById(visualizationId){
        return axios.get(VISUALIZATION_API_BASE_URL + '/load?visualizationId=' + visualizationId);
    }

    updateVisualization(visualization){
        return axios.put(VISUALIZATION_API_BASE_URL + '/update', visualization);
    }

    deleteVisualization(visualizationId){
        return axios.delete(VISUALIZATION_API_BASE_URL + '/delete?visualizationId=' + visualizationId);
    }
}

export default new VisualizationService()