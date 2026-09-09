import React, { Component } from 'react'
import AircraftVariantService from '../services/AircraftVariantService'

class ViewAircraftVariantComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            aircraftVariant: {}
        }
    }

    componentDidMount(){
        AircraftVariantService.getAircraftVariantById(this.state.id).then( res => {
            this.setState({aircraftVariant: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View AircraftVariant Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> variantCode:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.aircraftVariant.variantCode }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> rangeNm:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.aircraftVariant.rangeNm }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> maxTakeoffWeightKg:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.aircraftVariant.maxTakeoffWeightKg }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewAircraftVariantComponent
