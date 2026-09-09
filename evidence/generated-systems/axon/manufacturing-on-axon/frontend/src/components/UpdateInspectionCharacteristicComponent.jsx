import React, { Component } from 'react'
import InspectionCharacteristicService from '../services/InspectionCharacteristicService';

class UpdateInspectionCharacteristicComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                characteristicCode: '',
                name: '',
                lowerSpecLimit: '',
                upperSpecLimit: '',
                target: '',
                measurementType: ''
        }
        this.updateInspectionCharacteristic = this.updateInspectionCharacteristic.bind(this);

        this.changecharacteristicCodeHandler = this.changecharacteristicCodeHandler.bind(this);
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changelowerSpecLimitHandler = this.changelowerSpecLimitHandler.bind(this);
        this.changeupperSpecLimitHandler = this.changeupperSpecLimitHandler.bind(this);
        this.changetargetHandler = this.changetargetHandler.bind(this);
        this.changeMeasurementTypeHandler = this.changeMeasurementTypeHandler.bind(this);
    }

    componentDidMount(){
        InspectionCharacteristicService.getInspectionCharacteristicById(this.state.id).then( (res) =>{
            let inspectionCharacteristic = res.data;
            this.setState({
                characteristicCode: inspectionCharacteristic.characteristicCode,
                name: inspectionCharacteristic.name,
                lowerSpecLimit: inspectionCharacteristic.lowerSpecLimit,
                upperSpecLimit: inspectionCharacteristic.upperSpecLimit,
                target: inspectionCharacteristic.target,
                measurementType: inspectionCharacteristic.measurementType
            });
        });
    }

    updateInspectionCharacteristic = (e) => {
        e.preventDefault();
        let inspectionCharacteristic = {
            inspectionCharacteristicId: this.state.id,
            characteristicCode: this.state.characteristicCode,
            name: this.state.name,
            lowerSpecLimit: this.state.lowerSpecLimit,
            upperSpecLimit: this.state.upperSpecLimit,
            target: this.state.target,
            measurementType: this.state.measurementType
        };
        console.log('inspectionCharacteristic => ' + JSON.stringify(inspectionCharacteristic));
        console.log('id => ' + JSON.stringify(this.state.id));
        InspectionCharacteristicService.updateInspectionCharacteristic(inspectionCharacteristic).then( res => {
            this.props.history.push('/inspectionCharacteristics');
        });
    }

    changecharacteristicCodeHandler= (event) => {
        this.setState({characteristicCode: event.target.value});
    }
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changelowerSpecLimitHandler= (event) => {
        this.setState({lowerSpecLimit: event.target.value});
    }
    changeupperSpecLimitHandler= (event) => {
        this.setState({upperSpecLimit: event.target.value});
    }
    changetargetHandler= (event) => {
        this.setState({target: event.target.value});
    }
    changeMeasurementTypeHandler= (event) => {
        this.setState({measurementType: event.target.value});
    }

    cancel(){
        this.props.history.push('/inspectionCharacteristics');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update InspectionCharacteristic</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> characteristicCode: </label>
                                                <input placeholder="characteristicCode" name="characteristicCode" className="form-control" value={this.state.characteristicCode} onChange={this.changecharacteristicCodeHandler}/>

                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> lowerSpecLimit: </label>
                                                <input placeholder="lowerSpecLimit" name="lowerSpecLimit" className="form-control" value={this.state.lowerSpecLimit} onChange={this.changelowerSpecLimitHandler}/>

                                            <label> upperSpecLimit: </label>
                                                <input placeholder="upperSpecLimit" name="upperSpecLimit" className="form-control" value={this.state.upperSpecLimit} onChange={this.changeupperSpecLimitHandler}/>

                                            <label> target: </label>
                                                <input placeholder="target" name="target" className="form-control" value={this.state.target} onChange={this.changetargetHandler}/>

                                            <label> MeasurementType: </label>
                                                <select value={this.state.measurementType} onChange={this.changeMeasurementTypeHandler}>
                      <option name="MeasurementType" className="form-control" >
                          Attribute
                      </option>
                      <option name="MeasurementType" className="form-control" >
                          Variable
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateInspectionCharacteristic}>Save</button>
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

export default UpdateInspectionCharacteristicComponent
