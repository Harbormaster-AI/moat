import React, { Component } from 'react'
import AircraftModelService from '../services/AircraftModelService'

class ViewAircraftModelComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            aircraftModel: {}
        }
    }

    componentDidMount(){
        AircraftModelService.getAircraftModelById(this.state.id).then( res => {
            this.setState({aircraftModel: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View AircraftModel Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.aircraftModel.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> modelDesignation:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.aircraftModel.modelDesignation }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> AircraftType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.aircraftModel.aircraftType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewAircraftModelComponent
