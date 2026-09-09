import React, { Component } from 'react'
import AnalyticsWorkspaceService from '../services/AnalyticsWorkspaceService';

class CreateAnalyticsWorkspaceComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                businessDomain: '',
                ownerTeam: '',
                governanceTier: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changebusinessDomainHandler = this.changebusinessDomainHandler.bind(this);
        this.changeownerTeamHandler = this.changeownerTeamHandler.bind(this);
        this.changeGovernanceTierHandler = this.changeGovernanceTierHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            AnalyticsWorkspaceService.getAnalyticsWorkspaceById(this.state.id).then( (res) =>{
                let analyticsWorkspace = res.data;
                this.setState({
                    name: analyticsWorkspace.name,
                    businessDomain: analyticsWorkspace.businessDomain,
                    ownerTeam: analyticsWorkspace.ownerTeam,
                    governanceTier: analyticsWorkspace.governanceTier
                });
            });
        }        
    }
    saveOrUpdateAnalyticsWorkspace = (e) => {
        e.preventDefault();
        let analyticsWorkspace = {
                analyticsWorkspaceId: this.state.id,
                name: this.state.name,
                businessDomain: this.state.businessDomain,
                ownerTeam: this.state.ownerTeam,
                governanceTier: this.state.governanceTier
            };
        console.log('analyticsWorkspace => ' + JSON.stringify(analyticsWorkspace));

        // step 5
        if(this.state.id === '_add'){
            analyticsWorkspace.analyticsWorkspaceId=''
            AnalyticsWorkspaceService.createAnalyticsWorkspace(analyticsWorkspace).then(res =>{
                this.props.history.push('/analyticsWorkspaces');
            });
        }else{
            AnalyticsWorkspaceService.updateAnalyticsWorkspace(analyticsWorkspace).then( res => {
                this.props.history.push('/analyticsWorkspaces');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changebusinessDomainHandler= (event) => {
        this.setState({businessDomain: event.target.value});
    }
    changeownerTeamHandler= (event) => {
        this.setState({ownerTeam: event.target.value});
    }
    changeGovernanceTierHandler= (event) => {
        this.setState({governanceTier: event.target.value});
    }

    cancel(){
        this.props.history.push('/analyticsWorkspaces');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add AnalyticsWorkspace</h3>
        }else{
            return <h3 className="text-center">Update AnalyticsWorkspace</h3>
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

                                            <label> businessDomain:&emsp; </label>
                                                <input placeholder="businessDomain" name="businessDomain" className="form-control" value={this.state.businessDomain} onChange={this.changebusinessDomainHandler}/>

                                            <label> ownerTeam:&emsp; </label>
                                                <input placeholder="ownerTeam" name="ownerTeam" className="form-control" value={this.state.ownerTeam} onChange={this.changeownerTeamHandler}/>

                                            <label> GovernanceTier:&emsp; </label>
                                                <select value={this.state.governanceTier} onChange={this.changeGovernanceTierHandler}>
                      <option name="GovernanceTier" className="form-control" >
                          Open
                      </option>
                      <option name="GovernanceTier" className="form-control" >
                          Internal
                      </option>
                      <option name="GovernanceTier" className="form-control" >
                          Restricted
                      </option>
                      <option name="GovernanceTier" className="form-control" >
                          Confidential
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateAnalyticsWorkspace}>Save</button>
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

export default CreateAnalyticsWorkspaceComponent
