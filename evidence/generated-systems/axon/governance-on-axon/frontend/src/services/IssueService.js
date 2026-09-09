import axios from 'axios';

const ISSUE_API_BASE_URL = "/Issue";

class IssueService {

    getIssues(){
        return axios.get(ISSUE_API_BASE_URL + '/' );
    }

    createIssue(issue){
        return axios.post(ISSUE_API_BASE_URL  + '/create', issue);
    }

    getIssueById(issueId){
        return axios.get(ISSUE_API_BASE_URL + '/load?issueId=' + issueId);
    }

    updateIssue(issue){
        return axios.put(ISSUE_API_BASE_URL + '/update', issue);
    }

    deleteIssue(issueId){
        return axios.delete(ISSUE_API_BASE_URL + '/delete?issueId=' + issueId);
    }
}

export default new IssueService()