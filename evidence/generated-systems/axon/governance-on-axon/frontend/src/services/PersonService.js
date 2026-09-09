import axios from 'axios';

const PERSON_API_BASE_URL = "/Person";

class PersonService {

    getPersons(){
        return axios.get(PERSON_API_BASE_URL + '/' );
    }

    createPerson(person){
        return axios.post(PERSON_API_BASE_URL  + '/create', person);
    }

    getPersonById(personId){
        return axios.get(PERSON_API_BASE_URL + '/load?personId=' + personId);
    }

    updatePerson(person){
        return axios.put(PERSON_API_BASE_URL + '/update', person);
    }

    deletePerson(personId){
        return axios.delete(PERSON_API_BASE_URL + '/delete?personId=' + personId);
    }
}

export default new PersonService()