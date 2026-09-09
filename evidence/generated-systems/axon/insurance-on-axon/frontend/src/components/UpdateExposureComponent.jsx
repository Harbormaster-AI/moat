import React, { Component } from 'react'
import ExposureService from '../services/ExposureService';

class UpdateExposureComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                exposureType: '',
                status: ''
        }
        this.updateExposure = this.updateExposure.bind(this);

        this.changeExposureTypeHandler = this.changeExposureTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        ExposureService.getExposureById(this.state.id).then( (res) =>{
            let exposure = res.data;
            this.setState({
                exposureType: exposure.exposureType,
                status: exposure.status
            });
        });
    }

    updateExposure = (e) => {
        e.preventDefault();
        let exposure = {
            exposureId: this.state.id,
            exposureType: this.state.exposureType,
            status: this.state.status
        };
        console.log('exposure => ' + JSON.stringify(exposure));
        console.log('id => ' + JSON.stringify(this.state.id));
        ExposureService.updateExposure(exposure).then( res => {
            this.props.history.push('/exposures');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Exposure</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> ExposureType: </label>
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

                                            <label> Status: </label>
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
                                        <button className="btn btn-success" onClick={this.updateExposure}>Save</button>
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

export default UpdateExposureComponent
