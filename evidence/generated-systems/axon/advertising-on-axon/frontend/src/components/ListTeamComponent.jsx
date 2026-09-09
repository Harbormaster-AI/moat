import React, { Component } from 'react'
import TeamService from '../services/TeamService'

class ListTeamComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                teams: []
        }
        this.addTeam = this.addTeam.bind(this);
        this.editTeam = this.editTeam.bind(this);
        this.deleteTeam = this.deleteTeam.bind(this);
    }

    deleteTeam(id){
        TeamService.deleteTeam(id).then( res => {
            this.setState({teams: this.state.teams.filter(team => team.teamId !== id)});
        });
    }
    viewTeam(id){
        this.props.history.push(`/view-team/${id}`);
    }
    editTeam(id){
        this.props.history.push(`/add-team/${id}`);
    }

    componentDidMount(){
        TeamService.getTeams().then((res) => {
            this.setState({ teams: res.data});
        });
    }

    addTeam(){
        this.props.history.push('/add-team/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Team List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addTeam}> Add Team</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.teams.map(
                                        team => 
                                        <tr key = {team.teamId}>
                                             <td> { team.name } </td>
                                             <td>
                                                 <button onClick={ () => this.editTeam(team.teamId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteTeam(team.teamId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewTeam(team.teamId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListTeamComponent
