import React, { Component } from 'react'
import CareTeamService from '../services/CareTeamService';

class CreateCareTeamComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                careSetting: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeCareSettingHandler = this.changeCareSettingHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            CareTeamService.getCareTeamById(this.state.id).then( (res) =>{
                let careTeam = res.data;
                this.setState({
                    name: careTeam.name,
                    careSetting: careTeam.careSetting
                });
            });
        }        
    }
    saveOrUpdateCareTeam = (e) => {
        e.preventDefault();
        let careTeam = {
                careTeamId: this.state.id,
                name: this.state.name,
                careSetting: this.state.careSetting
            };
        console.log('careTeam => ' + JSON.stringify(careTeam));

        // step 5
        if(this.state.id === '_add'){
            careTeam.careTeamId=''
            CareTeamService.createCareTeam(careTeam).then(res =>{
                this.props.history.push('/careTeams');
            });
        }else{
            CareTeamService.updateCareTeam(careTeam).then( res => {
                this.props.history.push('/careTeams');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add CareTeam</h3>
        }else{
            return <h3 className="text-center">Update CareTeam</h3>
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

                                            <label> CareSetting:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateCareTeam}>Save</button>
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

export default CreateCareTeamComponent
