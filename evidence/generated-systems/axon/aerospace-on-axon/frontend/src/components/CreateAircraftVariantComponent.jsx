import React, { Component } from 'react'
import AircraftVariantService from '../services/AircraftVariantService';

class CreateAircraftVariantComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                variantCode: '',
                rangeNm: '',
                maxTakeoffWeightKg: ''
        }
        this.changevariantCodeHandler = this.changevariantCodeHandler.bind(this);
        this.changerangeNmHandler = this.changerangeNmHandler.bind(this);
        this.changemaxTakeoffWeightKgHandler = this.changemaxTakeoffWeightKgHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            AircraftVariantService.getAircraftVariantById(this.state.id).then( (res) =>{
                let aircraftVariant = res.data;
                this.setState({
                    variantCode: aircraftVariant.variantCode,
                    rangeNm: aircraftVariant.rangeNm,
                    maxTakeoffWeightKg: aircraftVariant.maxTakeoffWeightKg
                });
            });
        }        
    }
    saveOrUpdateAircraftVariant = (e) => {
        e.preventDefault();
        let aircraftVariant = {
                aircraftVariantId: this.state.id,
                variantCode: this.state.variantCode,
                rangeNm: this.state.rangeNm,
                maxTakeoffWeightKg: this.state.maxTakeoffWeightKg
            };
        console.log('aircraftVariant => ' + JSON.stringify(aircraftVariant));

        // step 5
        if(this.state.id === '_add'){
            aircraftVariant.aircraftVariantId=''
            AircraftVariantService.createAircraftVariant(aircraftVariant).then(res =>{
                this.props.history.push('/aircraftVariants');
            });
        }else{
            AircraftVariantService.updateAircraftVariant(aircraftVariant).then( res => {
                this.props.history.push('/aircraftVariants');
            });
        }
    }
    
    changevariantCodeHandler= (event) => {
        this.setState({variantCode: event.target.value});
    }
    changerangeNmHandler= (event) => {
        this.setState({rangeNm: event.target.value});
    }
    changemaxTakeoffWeightKgHandler= (event) => {
        this.setState({maxTakeoffWeightKg: event.target.value});
    }

    cancel(){
        this.props.history.push('/aircraftVariants');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add AircraftVariant</h3>
        }else{
            return <h3 className="text-center">Update AircraftVariant</h3>
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
                                            <label> variantCode:&emsp; </label>
                                                <input placeholder="variantCode" name="variantCode" className="form-control" value={this.state.variantCode} onChange={this.changevariantCodeHandler}/>

                                            <label> rangeNm:&emsp; </label>
                                                <input type="number" placeholder="rangeNm" name="rangeNm" className="form-control" value={this.state.rangeNm} onChange={this.changerangeNmHandler}/>

                                            <label> maxTakeoffWeightKg:&emsp; </label>
                                                <input placeholder="maxTakeoffWeightKg" name="maxTakeoffWeightKg" className="form-control" value={this.state.maxTakeoffWeightKg} onChange={this.changemaxTakeoffWeightKgHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateAircraftVariant}>Save</button>
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

export default CreateAircraftVariantComponent
