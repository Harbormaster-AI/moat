import React, { Component } from 'react'
import FacilityService from '../services/FacilityService';

class UpdateFacilityComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                facilityCode: '',
                address: '',
                facilityType: ''
        }
        this.updateFacility = this.updateFacility.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changefacilityCodeHandler = this.changefacilityCodeHandler.bind(this);
        this.changeaddressHandler = this.changeaddressHandler.bind(this);
        this.changeFacilityTypeHandler = this.changeFacilityTypeHandler.bind(this);
    }

    componentDidMount(){
        FacilityService.getFacilityById(this.state.id).then( (res) =>{
            let facility = res.data;
            this.setState({
                name: facility.name,
                facilityCode: facility.facilityCode,
                address: facility.address,
                facilityType: facility.facilityType
            });
        });
    }

    updateFacility = (e) => {
        e.preventDefault();
        let facility = {
            facilityId: this.state.id,
            name: this.state.name,
            facilityCode: this.state.facilityCode,
            address: this.state.address,
            facilityType: this.state.facilityType
        };
        console.log('facility => ' + JSON.stringify(facility));
        console.log('id => ' + JSON.stringify(this.state.id));
        FacilityService.updateFacility(facility).then( res => {
            this.props.history.push('/facilitys');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changefacilityCodeHandler= (event) => {
        this.setState({facilityCode: event.target.value});
    }
    changeaddressHandler= (event) => {
        this.setState({address: event.target.value});
    }
    changeFacilityTypeHandler= (event) => {
        this.setState({facilityType: event.target.value});
    }

    cancel(){
        this.props.history.push('/facilitys');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Facility</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> facilityCode: </label>
                                                <input placeholder="facilityCode" name="facilityCode" className="form-control" value={this.state.facilityCode} onChange={this.changefacilityCodeHandler}/>

                                            <label> address: </label>
                                                <input placeholder="address" name="address" className="form-control" value={this.state.address} onChange={this.changeaddressHandler}/>

                                            <label> FacilityType: </label>
                                                <select value={this.state.facilityType} onChange={this.changeFacilityTypeHandler}>
                      <option name="FacilityType" className="form-control" >
                          Hospital
                      </option>
                      <option name="FacilityType" className="form-control" >
                          Clinic
                      </option>
                      <option name="FacilityType" className="form-control" >
                          AmbulatorySurgeryCenter
                      </option>
                      <option name="FacilityType" className="form-control" >
                          UrgentCare
                      </option>
                      <option name="FacilityType" className="form-control" >
                          Laboratory
                      </option>
                      <option name="FacilityType" className="form-control" >
                          ImagingCenter
                      </option>
                      <option name="FacilityType" className="form-control" >
                          Pharmacy
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateFacility}>Save</button>
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

export default UpdateFacilityComponent
