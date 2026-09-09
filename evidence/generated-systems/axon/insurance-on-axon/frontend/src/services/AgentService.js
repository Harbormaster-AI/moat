import axios from 'axios';

const AGENT_API_BASE_URL = "/Agent";

class AgentService {

    getAgents(){
        return axios.get(AGENT_API_BASE_URL + '/' );
    }

    createAgent(agent){
        return axios.post(AGENT_API_BASE_URL  + '/create', agent);
    }

    getAgentById(agentId){
        return axios.get(AGENT_API_BASE_URL + '/load?agentId=' + agentId);
    }

    updateAgent(agent){
        return axios.put(AGENT_API_BASE_URL + '/update', agent);
    }

    deleteAgent(agentId){
        return axios.delete(AGENT_API_BASE_URL + '/delete?agentId=' + agentId);
    }
}

export default new AgentService()