import React, { Component } from 'react'
import RegistrationService from '../services/RegistrationService'

class ListRegistrationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                registrations: []
        }
        this.addRegistration = this.addRegistration.bind(this);
        this.editRegistration = this.editRegistration.bind(this);
        this.deleteRegistration = this.deleteRegistration.bind(this);
    }

    deleteRegistration(id){
        RegistrationService.deleteRegistration(id).then( res => {
            this.setState({registrations: this.state.registrations.filter(registration => registration.registrationId !== id)});
        });
    }
    viewRegistration(id){
        this.props.history.push(`/view-registration/${id}`);
    }
    editRegistration(id){
        this.props.history.push(`/add-registration/${id}`);
    }

    componentDidMount(){
        RegistrationService.getRegistrations().then((res) => {
            this.setState({ registrations: res.data});
        });
    }

    addRegistration(){
        this.props.history.push('/add-registration/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Registration List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addRegistration}> Add Registration</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> TailNumber </th>
                                    <th> RegistryCountry </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.registrations.map(
                                        registration => 
                                        <tr key = {registration.registrationId}>
                                             <td> { registration.tailNumber } </td>
                                             <td> { registration.registryCountry } </td>
                                             <td>
                                                 <button onClick={ () => this.editRegistration(registration.registrationId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteRegistration(registration.registrationId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewRegistration(registration.registrationId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListRegistrationComponent
