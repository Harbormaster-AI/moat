import React, { Component } from 'react'
import KYCProfileService from '../services/KYCProfileService';

class UpdateKYCProfileComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                profileId: '',
                createdAt: '',
                status: '',
                verificationLevel: ''
        }
        this.updateKYCProfile = this.updateKYCProfile.bind(this);

        this.changeprofileIdHandler = this.changeprofileIdHandler.bind(this);
        this.changecreatedAtHandler = this.changecreatedAtHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
        this.changeVerificationLevelHandler = this.changeVerificationLevelHandler.bind(this);
    }

    componentDidMount(){
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

    updateKYCProfile = (e) => {
        e.preventDefault();
        let kYCProfile = {
            kYCProfileId: this.state.id,
            profileId: this.state.profileId,
            createdAt: this.state.createdAt,
            status: this.state.status,
            verificationLevel: this.state.verificationLevel
        };
        console.log('kYCProfile => ' + JSON.stringify(kYCProfile));
        console.log('id => ' + JSON.stringify(this.state.id));
        KYCProfileService.updateKYCProfile(kYCProfile).then( res => {
            this.props.history.push('/kYCProfiles');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update KYCProfile</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> profileId: </label>
                                                <input placeholder="profileId" name="profileId" className="form-control" value={this.state.profileId} onChange={this.changeprofileIdHandler}/>

                                            <label> createdAt: </label>
                                                <input type="time" placeholder="createdAt" name="createdAt" className="form-control" value={this.state.createdAt} onChange={this.changecreatedAtHandler}/>

                                            <label> Status: </label>
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

                                            <label> VerificationLevel: </label>
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
                                        <button className="btn btn-success" onClick={this.updateKYCProfile}>Save</button>
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

export default UpdateKYCProfileComponent
