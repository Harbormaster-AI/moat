import React, { Component } from 'react'
import AircraftProgramService from '../services/AircraftProgramService'

class ViewAircraftProgramComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            aircraftProgram: {}
        }
    }

    componentDidMount(){
        AircraftProgramService.getAircraftProgramById(this.state.id).then( res => {
            this.setState({aircraftProgram: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View AircraftProgram Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.aircraftProgram.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> programCode:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.aircraftProgram.programCode }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> entryIntoServiceYear:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.aircraftProgram.entryIntoServiceYear }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.aircraftProgram.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewAircraftProgramComponent
