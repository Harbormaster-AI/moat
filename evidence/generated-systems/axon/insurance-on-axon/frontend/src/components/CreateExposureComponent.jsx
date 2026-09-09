import React, { Component } from 'react'
import ExposureService from '../services/ExposureService';

class CreateExposureComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                exposureType: '',
                status: ''
        }
        this.changeExposureTypeHandler = this.changeExposureTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ExposureService.getExposureById(this.state.id).then( (res) =>{
                let exposure = res.data;
                this.setState({
                    exposureType: exposure.exposureType,
                    status: exposure.status
                });
            });
        }        
    }
    saveOrUpdateExposure = (e) => {
        e.preventDefault();
        let exposure = {
                exposureId: this.state.id,
                exposureType: this.state.exposureType,
                status: this.state.status
            };
        console.log('exposure => ' + JSON.stringify(exposure));

        // step 5
        if(this.state.id === '_add'){
            exposure.exposureId=''
            ExposureService.createExposure(exposure).then(res =>{
                this.props.history.push('/exposures');
            });
        }else{
            ExposureService.updateExposure(exposure).then( res => {
                this.props.history.push('/exposures');
            });
        }
    }
    
    changeExposureTypeHandler= (event) => {
        this.setState({exposureType: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/exposures');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Exposure</h3>
        }else{
            return <h3 className="text-center">Update Exposure</h3>
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
                                            <label> ExposureType:&emsp; </label>
                                                <select value={this.state.exposureType} onChange={this.changeExposureTypeHandler}>
                      <option name="ExposureType" className="form-control" >
                          BodilyInjury
                      </option>
                      <option name="ExposureType" className="form-control" >
                          PropertyDamage
                      </option>
                      <option name="ExposureType" className="form-control" >
                          Medical
                      </option>
                      <option name="ExposureType" className="form-control" >
                          UninsuredMotorist
                      </option>
                      <option name="ExposureType" className="form-control" >
                          PersonalInjuryProtection
                      </option>
                      <option name="ExposureType" className="form-control" >
                          DwellingDamage
                      </option>
                      <option name="ExposureType" className="form-control" >
                          ContentsDamage
                      </option>
                      <option name="ExposureType" className="form-control" >
                          BusinessIncome
                      </option>
                    </select>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Open
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
                      </option>
                      <option name="Status" className="form-control" >
                          Pending
                      </option>
                      <option name="Status" className="form-control" >
                          Reserved
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateExposure}>Save</button>
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

export default CreateExposureComponent
