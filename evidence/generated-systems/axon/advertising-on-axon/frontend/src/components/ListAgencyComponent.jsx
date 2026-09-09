import React, { Component } from 'react'
import AgencyService from '../services/AgencyService'

class ListAgencyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                agencys: []
        }
        this.addAgency = this.addAgency.bind(this);
        this.editAgency = this.editAgency.bind(this);
        this.deleteAgency = this.deleteAgency.bind(this);
    }

    deleteAgency(id){
        AgencyService.deleteAgency(id).then( res => {
            this.setState({agencys: this.state.agencys.filter(agency => agency.agencyId !== id)});
        });
    }
    viewAgency(id){
        this.props.history.push(`/view-agency/${id}`);
    }
    editAgency(id){
        this.props.history.push(`/add-agency/${id}`);
    }

    componentDidMount(){
        AgencyService.getAgencys().then((res) => {
            this.setState({ agencys: res.data});
        });
    }

    addAgency(){
        this.props.history.push('/add-agency/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Agency List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAgency}> Add Agency</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> LegalName </th>
                                    <th> HeadquartersCountry </th>
                                    <th> Website </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.agencys.map(
                                        agency => 
                                        <tr key = {agency.agencyId}>
                                             <td> { agency.name } </td>
                                             <td> { agency.legalName } </td>
                                             <td> { agency.headquartersCountry } </td>
                                             <td> { agency.website } </td>
                                             <td>
                                                 <button onClick={ () => this.editAgency(agency.agencyId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAgency(agency.agencyId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAgency(agency.agencyId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAgencyComponent
