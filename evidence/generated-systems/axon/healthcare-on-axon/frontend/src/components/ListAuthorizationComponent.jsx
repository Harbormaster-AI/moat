import React, { Component } from 'react'
import AuthorizationService from '../services/AuthorizationService'

class ListAuthorizationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                authorizations: []
        }
        this.addAuthorization = this.addAuthorization.bind(this);
        this.editAuthorization = this.editAuthorization.bind(this);
        this.deleteAuthorization = this.deleteAuthorization.bind(this);
    }

    deleteAuthorization(id){
        AuthorizationService.deleteAuthorization(id).then( res => {
            this.setState({authorizations: this.state.authorizations.filter(authorization => authorization.authorizationId !== id)});
        });
    }
    viewAuthorization(id){
        this.props.history.push(`/view-authorization/${id}`);
    }
    editAuthorization(id){
        this.props.history.push(`/add-authorization/${id}`);
    }

    componentDidMount(){
        AuthorizationService.getAuthorizations().then((res) => {
            this.setState({ authorizations: res.data});
        });
    }

    addAuthorization(){
        this.props.history.push('/add-authorization/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Authorization List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAuthorization}> Add Authorization</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> AuthNumber </th>
                                    <th> RequestedService </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.authorizations.map(
                                        authorization => 
                                        <tr key = {authorization.authorizationId}>
                                             <td> { authorization.authNumber } </td>
                                             <td> { authorization.requestedService } </td>
                                             <td> { authorization.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editAuthorization(authorization.authorizationId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAuthorization(authorization.authorizationId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAuthorization(authorization.authorizationId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAuthorizationComponent
