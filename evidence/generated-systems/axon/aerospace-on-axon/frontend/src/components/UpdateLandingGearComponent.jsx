import React, { Component } from 'react'
import LandingGearService from '../services/LandingGearService';

class UpdateLandingGearComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                supplierPartNumber: '',
                gearType: ''
        }
        this.updateLandingGear = this.updateLandingGear.bind(this);

        this.changesupplierPartNumberHandler = this.changesupplierPartNumberHandler.bind(this);
        this.changeGearTypeHandler = this.changeGearTypeHandler.bind(this);
    }

    componentDidMount(){
        LandingGearService.getLandingGearById(this.state.id).then( (res) =>{
            let landingGear = res.data;
            this.setState({
                supplierPartNumber: landingGear.supplierPartNumber,
                gearType: landingGear.gearType
            });
        });
    }

    updateLandingGear = (e) => {
        e.preventDefault();
        let landingGear = {
            landingGearId: this.state.id,
            supplierPartNumber: this.state.supplierPartNumber,
            gearType: this.state.gearType
        };
        console.log('landingGear => ' + JSON.stringify(landingGear));
        console.log('id => ' + JSON.stringify(this.state.id));
        LandingGearService.updateLandingGear(landingGear).then( res => {
            this.props.history.push('/landingGears');
        });
    }

    changesupplierPartNumberHandler= (event) => {
        this.setState({supplierPartNumber: event.target.value});
    }
    changeGearTypeHandler= (event) => {
        this.setState({gearType: event.target.value});
    }

    cancel(){
        this.props.history.push('/landingGears');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update LandingGear</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> supplierPartNumber: </label>
                                                <input placeholder="supplierPartNumber" name="supplierPartNumber" className="form-control" value={this.state.supplierPartNumber} onChange={this.changesupplierPartNumberHandler}/>

                                            <label> GearType: </label>
                                                <select value={this.state.gearType} onChange={this.changeGearTypeHandler}>
                      <option name="GearType" className="form-control" >
                          Tricycle
                      </option>
                      <option name="GearType" className="form-control" >
                          Tandem
                      </option>
                      <option name="GearType" className="form-control" >
                          Taildragger
                      </option>
                      <option name="GearType" className="form-control" >
                          Skid
                      </option>
                      <option name="GearType" className="form-control" >
                          Floats
                      </option>
                      <option name="GearType" className="form-control" >
                          Retractable
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateLandingGear}>Save</button>
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

export default UpdateLandingGearComponent
