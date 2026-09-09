import React, { Component } from 'react'
import DataSubjectRequestService from '../services/DataSubjectRequestService'

class ListDataSubjectRequestComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                dataSubjectRequests: []
        }
        this.addDataSubjectRequest = this.addDataSubjectRequest.bind(this);
        this.editDataSubjectRequest = this.editDataSubjectRequest.bind(this);
        this.deleteDataSubjectRequest = this.deleteDataSubjectRequest.bind(this);
    }

    deleteDataSubjectRequest(id){
        DataSubjectRequestService.deleteDataSubjectRequest(id).then( res => {
            this.setState({dataSubjectRequests: this.state.dataSubjectRequests.filter(dataSubjectRequest => dataSubjectRequest.dataSubjectRequestId !== id)});
        });
    }
    viewDataSubjectRequest(id){
        this.props.history.push(`/view-dataSubjectRequest/${id}`);
    }
    editDataSubjectRequest(id){
        this.props.history.push(`/add-dataSubjectRequest/${id}`);
    }

    componentDidMount(){
        DataSubjectRequestService.getDataSubjectRequests().then((res) => {
            this.setState({ dataSubjectRequests: res.data});
        });
    }

    addDataSubjectRequest(){
        this.props.history.push('/add-dataSubjectRequest/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">DataSubjectRequest List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addDataSubjectRequest}> Add DataSubjectRequest</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ReceivedDate </th>
                                    <th> DueDate </th>
                                    <th> RequesterCountry </th>
                                    <th> RequestType </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.dataSubjectRequests.map(
                                        dataSubjectRequest => 
                                        <tr key = {dataSubjectRequest.dataSubjectRequestId}>
                                             <td> { dataSubjectRequest.receivedDate } </td>
                                             <td> { dataSubjectRequest.dueDate } </td>
                                             <td> { dataSubjectRequest.requesterCountry } </td>
                                             <td> { dataSubjectRequest.requestType } </td>
                                             <td> { dataSubjectRequest.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editDataSubjectRequest(dataSubjectRequest.dataSubjectRequestId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteDataSubjectRequest(dataSubjectRequest.dataSubjectRequestId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewDataSubjectRequest(dataSubjectRequest.dataSubjectRequestId)} className="btn btn-outline-info btn-sm">View </button>
                                             </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>

                 </div>

            </div>
        )
    }
}

export default ListDataSubjectRequestComponent
