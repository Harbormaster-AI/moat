import axios from 'axios';

const PROCEDURE_API_BASE_URL = "/Procedure";

class ProcedureService {

    getProcedures(){
        return axios.get(PROCEDURE_API_BASE_URL + '/' );
    }

    createProcedure(procedure){
        return axios.post(PROCEDURE_API_BASE_URL  + '/create', procedure);
    }

    getProcedureById(procedureId){
        return axios.get(PROCEDURE_API_BASE_URL + '/load?procedureId=' + procedureId);
    }

    updateProcedure(procedure){
        return axios.put(PROCEDURE_API_BASE_URL + '/update', procedure);
    }

    deleteProcedure(procedureId){
        return axios.delete(PROCEDURE_API_BASE_URL + '/delete?procedureId=' + procedureId);
    }
}

export default new ProcedureService()