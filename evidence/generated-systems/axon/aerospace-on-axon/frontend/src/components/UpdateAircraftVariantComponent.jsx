import React, { Component } from 'react'
import AircraftVariantService from '../services/AircraftVariantService';

class UpdateAircraftVariantComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                variantCode: '',
                rangeNm: '',
                maxTakeoffWeightKg: ''
        }
        this.updateAircraftVariant = this.updateAircraftVariant.bind(this);

        this.changevariantCodeHandler = this.changevariantCodeHandler.bind(this);
        this.changerangeNmHandler = this.changerangeNmHandler.bind(this);
        this.changemaxTakeoffWeightKgHandler = this.changemaxTakeoffWeightKgHandler.bind(this);
    }

    componentDidMount(){
        AircraftVariantService.getAircraftVariantById(this.state.id).then( (res) =>{
            let aircraftVariant = res.data;
            this.setState({
                variantCode: aircraftVariant.variantCode,
                rangeNm: aircraftVariant.rangeNm,
                maxTakeoffWeightKg: aircraftVariant.maxTakeoffWeightKg
            });
        });
    }

    updateAircraftVariant = (e) => {
        e.preventDefault();
        let aircraftVariant = {
            aircraftVariantId: this.state.id,
            variantCode: this.state.variantCode,
            rangeNm: this.state.rangeNm,
            maxTakeoffWeightKg: this.state.maxTakeoffWeightKg
        };
        console.log('aircraftVariant => ' + JSON.stringify(aircraftVariant));
        console.log('id => ' + JSON.stringify(this.state.id));
        AircraftVariantService.updateAircraftVariant(aircraftVariant).then( res => {
            this.props.history.push('/aircraftVariants');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update AircraftVariant</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> variantCode: </label>
                                                <input placeholder="variantCode" name="variantCode" className="form-control" value={this.state.variantCode} onChange={this.changevariantCodeHandler}/>

                                            <label> rangeNm: </label>
                                                <input type="number" placeholder="rangeNm" name="rangeNm" className="form-control" value={this.state.rangeNm} onChange={this.changerangeNmHandler}/>

                                            <label> maxTakeoffWeightKg: </label>
                                                <input placeholder="maxTakeoffWeightKg" name="maxTakeoffWeightKg" className="form-control" value={this.state.maxTakeoffWeightKg} onChange={this.changemaxTakeoffWeightKgHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateAircraftVariant}>Save</button>
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

export default UpdateAircraftVariantComponent
