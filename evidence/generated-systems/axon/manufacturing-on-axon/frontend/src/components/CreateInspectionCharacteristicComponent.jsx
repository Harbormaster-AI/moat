import React, { Component } from 'react'
import InspectionCharacteristicService from '../services/InspectionCharacteristicService';

class CreateInspectionCharacteristicComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                characteristicCode: '',
                name: '',
                lowerSpecLimit: '',
                upperSpecLimit: '',
                target: '',
                measurementType: ''
        }
        this.changecharacteristicCodeHandler = this.changecharacteristicCodeHandler.bind(this);
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changelowerSpecLimitHandler = this.changelowerSpecLimitHandler.bind(this);
        this.changeupperSpecLimitHandler = this.changeupperSpecLimitHandler.bind(this);
        this.changetargetHandler = this.changetargetHandler.bind(this);
        this.changeMeasurementTypeHandler = this.changeMeasurementTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateInspectionCharacteristic = (e) => {
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

        // step 5
        if(this.state.id === '_add'){
            inspectionCharacteristic.inspectionCharacteristicId=''
            InspectionCharacteristicService.createInspectionCharacteristic(inspectionCharacteristic).then(res =>{
                this.props.history.push('/inspectionCharacteristics');
            });
        }else{
            InspectionCharacteristicService.updateInspectionCharacteristic(inspectionCharacteristic).then( res => {
                this.props.history.push('/inspectionCharacteristics');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add InspectionCharacteristic</h3>
        }else{
            return <h3 className="text-center">Update InspectionCharacteristic</h3>
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
                                            <label> characteristicCode:&emsp; </label>
                                                <input placeholder="characteristicCode" name="characteristicCode" className="form-control" value={this.state.characteristicCode} onChange={this.changecharacteristicCodeHandler}/>

                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> lowerSpecLimit:&emsp; </label>
                                                <input placeholder="lowerSpecLimit" name="lowerSpecLimit" className="form-control" value={this.state.lowerSpecLimit} onChange={this.changelowerSpecLimitHandler}/>

                                            <label> upperSpecLimit:&emsp; </label>
                                                <input placeholder="upperSpecLimit" name="upperSpecLimit" className="form-control" value={this.state.upperSpecLimit} onChange={this.changeupperSpecLimitHandler}/>

                                            <label> target:&emsp; </label>
                                                <input placeholder="target" name="target" className="form-control" value={this.state.target} onChange={this.changetargetHandler}/>

                                            <label> MeasurementType:&emsp; </label>
                                                <select value={this.state.measurementType} onChange={this.changeMeasurementTypeHandler}>
                      <option name="MeasurementType" className="form-control" >
                          Attribute
                      </option>
                      <option name="MeasurementType" className="form-control" >
                          Variable
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateInspectionCharacteristic}>Save</button>
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

export default CreateInspectionCharacteristicComponent
