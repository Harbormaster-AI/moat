import React, { Component } from 'react'
import APIClientService from '../services/APIClientService'

class ListAPIClientComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                aPIClients: []
        }
        this.addAPIClient = this.addAPIClient.bind(this);
        this.editAPIClient = this.editAPIClient.bind(this);
        this.deleteAPIClient = this.deleteAPIClient.bind(this);
    }

    deleteAPIClient(id){
        APIClientService.deleteAPIClient(id).then( res => {
            this.setState({aPIClients: this.state.aPIClients.filter(aPIClient => aPIClient.aPIClientId !== id)});
        });
    }
    viewAPIClient(id){
        this.props.history.push(`/view-aPIClient/${id}`);
    }
    editAPIClient(id){
        this.props.history.push(`/add-aPIClient/${id}`);
    }

    componentDidMount(){
        APIClientService.getAPIClients().then((res) => {
            this.setState({ aPIClients: res.data});
        });
    }

    addAPIClient(){
        this.props.history.push('/add-aPIClient/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">APIClient List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAPIClient}> Add APIClient</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> ClientId </th>
                                    <th> RedirectUri </th>
                                    <th> ClientType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.aPIClients.map(
                                        aPIClient => 
                                        <tr key = {aPIClient.aPIClientId}>
                                             <td> { aPIClient.name } </td>
                                             <td> { aPIClient.clientId } </td>
                                             <td> { aPIClient.redirectUri } </td>
                                             <td> { aPIClient.clientType } </td>
                                             <td>
                                                 <button onClick={ () => this.editAPIClient(aPIClient.aPIClientId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAPIClient(aPIClient.aPIClientId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAPIClient(aPIClient.aPIClientId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAPIClientComponent
