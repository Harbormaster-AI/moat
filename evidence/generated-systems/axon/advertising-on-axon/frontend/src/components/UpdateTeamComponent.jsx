import React, { Component } from 'react'
import TeamService from '../services/TeamService';

class UpdateTeamComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: ''
        }
        this.updateTeam = this.updateTeam.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
    }

    componentDidMount(){
        TeamService.getTeamById(this.state.id).then( (res) =>{
            let team = res.data;
            this.setState({
                name: team.name
            });
        });
    }

    updateTeam = (e) => {
        e.preventDefault();
        let team = {
            teamId: this.state.id,
            name: this.state.name
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
