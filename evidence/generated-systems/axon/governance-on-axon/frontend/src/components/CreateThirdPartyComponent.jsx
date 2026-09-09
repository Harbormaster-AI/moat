import React, { Component } from 'react'
import ThirdPartyService from '../services/ThirdPartyService';

class CreateThirdPartyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                country: '',
                contactEmail: '',
                thirdPartyType: '',
                criticality: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changecountryHandler = this.changecountryHandler.bind(this);
        this.changecontactEmailHandler = this.changecontactEmailHandler.bind(this);
        this.changeThirdPartyTypeHandler = this.changeThirdPartyTypeHandler.bind(this);
        this.changeCriticalityHandler = this.changeCriticalityHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ThirdPartyService.getThirdPartyById(this.state.id).then( (res) =>{
                let thirdParty = res.data;
                this.setState({
                    name: thirdParty.name,
                    country: thirdParty.country,
                    contactEmail: thirdParty.contactEmail,
                    thirdPartyType: thirdParty.thirdPartyType,
                    criticality: thirdParty.criticality
                });
            });
        }        
    }
    saveOrUpdateThirdParty = (e) => {
        e.preventDefault();
        let thirdParty = {
                thirdPartyId: this.state.id,
                name: this.state.name,
                country: this.state.country,
                contactEmail: this.state.contactEmail,
                thirdPartyType: this.state.thirdPartyType,
                criticality: this.state.criticality
            };
        console.log('thirdParty => ' + JSON.stringify(thirdParty));

        // step 5
        if(this.state.id === '_add'){
            thirdParty.thirdPartyId=''
            ThirdPartyService.createThirdParty(thirdParty).then(res =>{
                this.props.history.push('/thirdPartys');
            });
        }else{
            ThirdPartyService.updateThirdParty(thirdParty).then( res => {
                this.props.history.push('/thirdPartys');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changecountryHandler= (event) => {
        this.setState({country: event.target.value});
    }
    changecontactEmailHandler= (event) => {
        this.setState({contactEmail: event.target.value});
    }
    changeThirdPartyTypeHandler= (event) => {
        this.setState({thirdPartyType: event.target.value});
    }
    changeCriticalityHandler= (event) => {
        this.setState({criticality: event.target.value});
    }

    cancel(){
        this.props.history.push('/thirdPartys');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add ThirdParty</h3>
        }else{
            return <h3 className="text-center">Update ThirdParty</h3>
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

                                            <label> country:&emsp; </label>
                                                <input placeholder="country" name="country" className="form-control" value={this.state.country} onChange={this.changecountryHandler}/>

                                            <label> contactEmail:&emsp; </label>
                                                <input placeholder="contactEmail" name="contactEmail" className="form-control" value={this.state.contactEmail} onChange={this.changecontactEmailHandler}/>

                                            <label> ThirdPartyType:&emsp; </label>
                                                <select value={this.state.thirdPartyType} onChange={this.changeThirdPartyTypeHandler}>
                      <option name="ThirdPartyType" className="form-control" >
                          Vendor
                      </option>
                      <option name="ThirdPartyType" className="form-control" >
                          Processor
                      </option>
                      <option name="ThirdPartyType" className="form-control" >
                          JointController
                      </option>
                      <option name="ThirdPartyType" className="form-control" >
                          Subprocessor
                      </option>
                      <option name="ThirdPartyType" className="form-control" >
                          Partner
                      </option>
                      <option name="ThirdPartyType" className="form-control" >
                          Consultant
                      </option>
                    </select>

                                            <label> Criticality:&emsp; </label>
                                                <select value={this.state.criticality} onChange={this.changeCriticalityHandler}>
                      <option name="Criticality" className="form-control" >
                          Low
                      </option>
                      <option name="Criticality" className="form-control" >
                          Medium
                      </option>
                      <option name="Criticality" className="form-control" >
                          High
                      </option>
                      <option name="Criticality" className="form-control" >
                          Critical
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateThirdParty}>Save</button>
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

export default CreateThirdPartyComponent
