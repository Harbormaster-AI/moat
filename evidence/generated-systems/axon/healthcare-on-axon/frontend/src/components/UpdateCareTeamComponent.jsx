import React, { Component } from 'react'
import CareTeamService from '../services/CareTeamService';

class UpdateCareTeamComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                careSetting: ''
        }
        this.updateCareTeam = this.updateCareTeam.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeCareSettingHandler = this.changeCareSettingHandler.bind(this);
    }

    componentDidMount(){
        CareTeamService.getCareTeamById(this.state.id).then( (res) =>{
            let careTeam = res.data;
            this.setState({
                name: careTeam.name,
                careSetting: careTeam.careSetting
            });
        });
    }

    updateCareTeam = (e) => {
        e.preventDefault();
        let careTeam = {
            careTeamId: this.state.id,
            name: this.state.name,
            careSetting: this.state.careSetting
        };
        console.log('careTeam => ' + JSON.stringify(careTeam));
        console.log('id => ' + JSON.stringify(this.state.id));
        CareTeamService.updateCareTeam(careTeam).then( res => {
            this.props.history.push('/careTeams');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeCareSettingHandler= (event) => {
        this.setState({careSetting: event.target.value});
    }

    cancel(){
        this.props.history.push('/careTeams');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update CareTeam</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> CareSetting: </label>
                                                <select value={this.state.careSetting} onChange={this.changeCareSettingHandler}>
                      <option name="CareSetting" className="form-control" >
                          Inpatient
                      </option>
                      <option name="CareSetting" className="form-control" >
                          Outpatient
                      </option>
                      <option name="CareSetting" className="form-control" >
                          Emergency
                      </option>
                      <option name="CareSetting" className="form-control" >
                          HomeHealth
                      </option>
                      <option name="CareSetting" className="form-control" >
                          Telehealth
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateCareTeam}>Save</button>
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

export default UpdateCareTeamComponent
