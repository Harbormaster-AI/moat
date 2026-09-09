import React, { Component } from 'react'
import VerifiedAddressService from '../services/VerifiedAddressService';

class CreateVerifiedAddressComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                address: '',
                verifiedAt: '',
                verificationStatus: ''
        }
        this.changeaddressHandler = this.changeaddressHandler.bind(this);
        this.changeverifiedAtHandler = this.changeverifiedAtHandler.bind(this);
        this.changeVerificationStatusHandler = this.changeVerificationStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            VerifiedAddressService.getVerifiedAddressById(this.state.id).then( (res) =>{
                let verifiedAddress = res.data;
                this.setState({
                    address: verifiedAddress.address,
                    verifiedAt: verifiedAddress.verifiedAt,
                    verificationStatus: verifiedAddress.verificationStatus
                });
            });
        }        
    }
    saveOrUpdateVerifiedAddress = (e) => {
        e.preventDefault();
        let verifiedAddress = {
                verifiedAddressId: this.state.id,
                address: this.state.address,
                verifiedAt: this.state.verifiedAt,
                verificationStatus: this.state.verificationStatus
            };
        console.log('verifiedAddress => ' + JSON.stringify(verifiedAddress));

        // step 5
        if(this.state.id === '_add'){
            verifiedAddress.verifiedAddressId=''
            VerifiedAddressService.createVerifiedAddress(verifiedAddress).then(res =>{
                this.props.history.push('/verifiedAddresss');
            });
        }else{
            VerifiedAddressService.updateVerifiedAddress(verifiedAddress).then( res => {
                this.props.history.push('/verifiedAddresss');
            });
        }
    }
    
    changeaddressHandler= (event) => {
        this.setState({address: event.target.value});
    }
    changeverifiedAtHandler= (event) => {
        this.setState({verifiedAt: event.target.value});
    }
    changeVerificationStatusHandler= (event) => {
        this.setState({verificationStatus: event.target.value});
    }

    cancel(){
        this.props.history.push('/verifiedAddresss');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add VerifiedAddress</h3>
        }else{
            return <h3 className="text-center">Update VerifiedAddress</h3>
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
                                            <label> address:&emsp; </label>
                                                <input placeholder="address" name="address" className="form-control" value={this.state.address} onChange={this.changeaddressHandler}/>

                                            <label> verifiedAt:&emsp; </label>
                                                <input type="time" placeholder="verifiedAt" name="verifiedAt" className="form-control" value={this.state.verifiedAt} onChange={this.changeverifiedAtHandler}/>

                                            <label> VerificationStatus:&emsp; </label>
                                                <select value={this.state.verificationStatus} onChange={this.changeVerificationStatusHandler}>
                      <option name="VerificationStatus" className="form-control" >
                          Unverified
                      </option>
                      <option name="VerificationStatus" className="form-control" >
                          Verified
                      </option>
                      <option name="VerificationStatus" className="form-control" >
                          Failed
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateVerifiedAddress}>Save</button>
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

export default CreateVerifiedAddressComponent
