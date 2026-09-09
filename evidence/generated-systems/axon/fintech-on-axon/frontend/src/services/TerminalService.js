import axios from 'axios';

const TERMINAL_API_BASE_URL = "/Terminal";

class TerminalService {

    getTerminals(){
        return axios.get(TERMINAL_API_BASE_URL + '/' );
    }

    createTerminal(terminal){
        return axios.post(TERMINAL_API_BASE_URL  + '/create', terminal);
    }

    getTerminalById(terminalId){
        return axios.get(TERMINAL_API_BASE_URL + '/load?terminalId=' + terminalId);
    }

    updateTerminal(terminal){
        return axios.put(TERMINAL_API_BASE_URL + '/update', terminal);
    }

    deleteTerminal(terminalId){
        return axios.delete(TERMINAL_API_BASE_URL + '/delete?terminalId=' + terminalId);
    }
}

export default new TerminalService()