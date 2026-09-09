import React, { Component } from 'react'
import WorkAuthorizationService from '../services/WorkAuthorizationService'

class ListWorkAuthorizationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                workAuthorizations: []
        }
        this.addWorkAuthorization = this.addWorkAuthorization.bind(this);
        this.editWorkAuthorization = this.editWorkAuthorization.bind(this);
        this.deleteWorkAuthorization = this.deleteWorkAuthorization.bind(this);
    }

    deleteWorkAuthorization(id){
        WorkAuthorizationService.deleteWorkAuthorization(id).then( res => {
            this.setState({workAuthorizations: this.state.workAuthorizations.filter(workAuthorization => workAuthorization.workAuthorizationId !== id)});
        });
    }
    viewWorkAuthorization(id){
        this.props.history.push(`/view-workAuthorization/${id}`);
    }
    editWorkAuthorization(id){
        this.props.history.push(`/add-workAuthorization/${id}`);
    }

    componentDidMount(){
        WorkAuthorizationService.getWorkAuthorizations().then((res) => {
            this.setState({ workAuthorizations: res.data});
        });
    }

    addWorkAuthorization(){
        this.props.history.push('/add-workAuthorization/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">WorkAuthorization List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addWorkAuthorization}> Add WorkAuthorization</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Country </th>
                                    <th> ExpirationDate </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.workAuthorizations.map(
                                        workAuthorization => 
                                        <tr key = {workAuthorization.workAuthorizationId}>
                                             <td> { workAuthorization.country } </td>
                                             <td> { workAuthorization.expirationDate } </td>
                                             <td> { workAuthorization.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editWorkAuthorization(workAuthorization.workAuthorizationId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteWorkAuthorization(workAuthorization.workAuthorizationId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewWorkAuthorization(workAuthorization.workAuthorizationId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListWorkAuthorizationComponent
