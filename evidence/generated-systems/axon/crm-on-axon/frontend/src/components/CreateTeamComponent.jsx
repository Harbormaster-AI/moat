import React, { Component } from 'react'
import TeamService from '../services/TeamService';

class CreateTeamComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                teamType: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeTeamTypeHandler = this.changeTeamTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            TeamService.getTeamById(this.state.id).then( (res) =>{
                let team = res.data;
                this.setState({
                    name: team.name,
                    teamType: team.teamType
                });
            });
        }        
    }
    saveOrUpdateTeam = (e) => {
        e.preventDefault();
        let team = {
                teamId: this.state.id,
                name: this.state.name,
                teamType: this.state.teamType
            };
        console.log('team => ' + JSON.stringify(team));

        // step 5
        if(this.state.id === '_add'){
            team.teamId=''
            TeamService.createTeam(team).then(res =>{
                this.props.history.push('/teams');
            });
        }else{
            TeamService.updateTeam(team).then( res => {
                this.props.history.push('/teams');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Team</h3>
        }else{
            return <h3 className="text-center">Update Team</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> TeamType:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateTeam}>Save</button>
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

export default CreateTeamComponent
