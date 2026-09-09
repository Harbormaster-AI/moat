import React, { Component } from 'react'
import AircraftFamilyService from '../services/AircraftFamilyService'

class ViewAircraftFamilyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            aircraftFamily: {}
        }
    }

    componentDidMount(){
        AircraftFamilyService.getAircraftFamilyById(this.state.id).then( res => {
            this.setState({aircraftFamily: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View AircraftFamily Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.aircraftFamily.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> familyCode:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.aircraftFamily.familyCode }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewAircraftFamilyComponent
