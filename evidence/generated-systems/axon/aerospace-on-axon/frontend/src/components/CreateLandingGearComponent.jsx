import React, { Component } from 'react'
import LandingGearService from '../services/LandingGearService';

class CreateLandingGearComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                supplierPartNumber: '',
                gearType: ''
        }
        this.changesupplierPartNumberHandler = this.changesupplierPartNumberHandler.bind(this);
        this.changeGearTypeHandler = this.changeGearTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            LandingGearService.getLandingGearById(this.state.id).then( (res) =>{
                let landingGear = res.data;
                this.setState({
                    supplierPartNumber: landingGear.supplierPartNumber,
                    gearType: landingGear.gearType
                });
            });
        }        
    }
    saveOrUpdateLandingGear = (e) => {
        e.preventDefault();
        let landingGear = {
                landingGearId: this.state.id,
                supplierPartNumber: this.state.supplierPartNumber,
                gearType: this.state.gearType
            };
        console.log('landingGear => ' + JSON.stringify(landingGear));

        // step 5
        if(this.state.id === '_add'){
            landingGear.landingGearId=''
            LandingGearService.createLandingGear(landingGear).then(res =>{
                this.props.history.push('/landingGears');
            });
        }else{
            LandingGearService.updateLandingGear(landingGear).then( res => {
                this.props.history.push('/landingGears');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add LandingGear</h3>
        }else{
            return <h3 className="text-center">Update LandingGear</h3>
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
                                            <label> supplierPartNumber:&emsp; </label>
                                                <input placeholder="supplierPartNumber" name="supplierPartNumber" className="form-control" value={this.state.supplierPartNumber} onChange={this.changesupplierPartNumberHandler}/>

                                            <label> GearType:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateLandingGear}>Save</button>
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

export default CreateLandingGearComponent
