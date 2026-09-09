import React, { Component } from 'react'
import KYCProfileService from '../services/KYCProfileService';

class CreateKYCProfileComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                profileId: '',
                createdAt: '',
                status: '',
                verificationLevel: ''
        }
        this.changeprofileIdHandler = this.changeprofileIdHandler.bind(this);
        this.changecreatedAtHandler = this.changecreatedAtHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
        this.changeVerificationLevelHandler = this.changeVerificationLevelHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            KYCProfileService.getKYCProfileById(this.state.id).then( (res) =>{
                let kYCProfile = res.data;
                this.setState({
                    profileId: kYCProfile.profileId,
                    createdAt: kYCProfile.createdAt,
                    status: kYCProfile.status,
                    verificationLevel: kYCProfile.verificationLevel
                });
            });
        }        
    }
    saveOrUpdateKYCProfile = (e) => {
        e.preventDefault();
        let kYCProfile = {
                kYCProfileId: this.state.id,
                profileId: this.state.profileId,
                createdAt: this.state.createdAt,
                status: this.state.status,
                verificationLevel: this.state.verificationLevel
            };
        console.log('kYCProfile => ' + JSON.stringify(kYCProfile));

        // step 5
        if(this.state.id === '_add'){
            kYCProfile.kYCProfileId=''
            KYCProfileService.createKYCProfile(kYCProfile).then(res =>{
                this.props.history.push('/kYCProfiles');
            });
        }else{
            KYCProfileService.updateKYCProfile(kYCProfile).then( res => {
                this.props.history.push('/kYCProfiles');
            });
        }
    }
    
    changeprofileIdHandler= (event) => {
        this.setState({profileId: event.target.value});
    }
    changecreatedAtHandler= (event) => {
        this.setState({createdAt: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }
    changeVerificationLevelHandler= (event) => {
        this.setState({verificationLevel: event.target.value});
    }

    cancel(){
        this.props.history.push('/kYCProfiles');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add KYCProfile</h3>
        }else{
            return <h3 className="text-center">Update KYCProfile</h3>
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
                                            <label> profileId:&emsp; </label>
                                                <input placeholder="profileId" name="profileId" className="form-control" value={this.state.profileId} onChange={this.changeprofileIdHandler}/>

                                            <label> createdAt:&emsp; </label>
                                                <input type="time" placeholder="createdAt" name="createdAt" className="form-control" value={this.state.createdAt} onChange={this.changecreatedAtHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Pending
                      </option>
                      <option name="Status" className="form-control" >
                          Verified
                      </option>
                      <option name="Status" className="form-control" >
                          Rejected
                      </option>
                      <option name="Status" className="form-control" >
                          Expired
                      </option>
                    </select>

                                            <label> VerificationLevel:&emsp; </label>
                                                <select value={this.state.verificationLevel} onChange={this.changeVerificationLevelHandler}>
                      <option name="VerificationLevel" className="form-control" >
                          Basic
                      </option>
                      <option name="VerificationLevel" className="form-control" >
                          Standard
                      </option>
                      <option name="VerificationLevel" className="form-control" >
                          Enhanced
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateKYCProfile}>Save</button>
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

export default CreateKYCProfileComponent
