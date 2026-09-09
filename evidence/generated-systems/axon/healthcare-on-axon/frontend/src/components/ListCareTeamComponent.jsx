import React, { Component } from 'react'
import CareTeamService from '../services/CareTeamService'

class ListCareTeamComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                careTeams: []
        }
        this.addCareTeam = this.addCareTeam.bind(this);
        this.editCareTeam = this.editCareTeam.bind(this);
        this.deleteCareTeam = this.deleteCareTeam.bind(this);
    }

    deleteCareTeam(id){
        CareTeamService.deleteCareTeam(id).then( res => {
            this.setState({careTeams: this.state.careTeams.filter(careTeam => careTeam.careTeamId !== id)});
        });
    }
    viewCareTeam(id){
        this.props.history.push(`/view-careTeam/${id}`);
    }
    editCareTeam(id){
        this.props.history.push(`/add-careTeam/${id}`);
    }

    componentDidMount(){
        CareTeamService.getCareTeams().then((res) => {
            this.setState({ careTeams: res.data});
        });
    }

    addCareTeam(){
        this.props.history.push('/add-careTeam/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">CareTeam List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCareTeam}> Add CareTeam</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> CareSetting </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.careTeams.map(
                                        careTeam => 
                                        <tr key = {careTeam.careTeamId}>
                                             <td> { careTeam.name } </td>
                                             <td> { careTeam.careSetting } </td>
                                             <td>
                                                 <button onClick={ () => this.editCareTeam(careTeam.careTeamId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCareTeam(careTeam.careTeamId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCareTeam(careTeam.careTeamId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListCareTeamComponent
