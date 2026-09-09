import React, { Component } from 'react'
import InspectionCharacteristicService from '../services/InspectionCharacteristicService'

class ViewInspectionCharacteristicComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            inspectionCharacteristic: {}
        }
    }

    componentDidMount(){
        InspectionCharacteristicService.getInspectionCharacteristicById(this.state.id).then( res => {
            this.setState({inspectionCharacteristic: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View InspectionCharacteristic Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> characteristicCode:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inspectionCharacteristic.characteristicCode }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inspectionCharacteristic.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> lowerSpecLimit:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inspectionCharacteristic.lowerSpecLimit }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> upperSpecLimit:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inspectionCharacteristic.upperSpecLimit }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> target:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inspectionCharacteristic.target }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> MeasurementType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inspectionCharacteristic.measurementType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewInspectionCharacteristicComponent
