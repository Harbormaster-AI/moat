import React, { Component } from 'react'
import TeamService from '../services/TeamService';

class UpdateTeamComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                teamType: ''
        }
        this.updateTeam = this.updateTeam.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeTeamTypeHandler = this.changeTeamTypeHandler.bind(this);
    }

    componentDidMount(){
        TeamService.getTeamById(this.state.id).then( (res) =>{
            let team = res.data;
            this.setState({
                name: team.name,
                teamType: team.teamType
            });
        });
    }

    updateTeam = (e) => {
        e.preventDefault();
        let team = {
            teamId: this.state.id,
            name: this.state.name,
            teamType: this.state.teamType
        };
        console.log('team => ' + JSON.stringify(team));
        console.log('id => ' + JSON.stringify(this.state.id));
        TeamService.updateTeam(team).then( res => {
            this.props.history.push('/teams');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeTeamTypeHandler= (event) => {
        this.setState({teamType: event.target.value});
    }

    cancel(){
        this.props.history.push('/teams');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Team</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> TeamType: </label>
                                                <select value={this.state.teamType} onChange={this.changeTeamTypeHandler}>
                      <option name="TeamType" className="form-control" >
                          Sales
                      </option>
                      <option name="TeamType" className="form-control" >
                          Service
                      </option>
                      <option name="TeamType" className="form-control" >
                          Marketing
                      </option>
                      <option name="TeamType" className="form-control" >
                          AccountTeam
                      </option>
                      <option name="TeamType" className="form-control" >
                          DealDesk
                      </option>
                      <option name="TeamType" className="form-control" >
                          CrossFunctional
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateTeam}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>

                   </div>
            </div>
        )
    }
}

export default UpdateTeamComponent
