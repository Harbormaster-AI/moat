import React, { Component } from 'react'
import FacilityService from '../services/FacilityService'

class ViewFacilityComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            facility: {}
        }
    }

    componentDidMount(){
        FacilityService.getFacilityById(this.state.id).then( res => {
            this.setState({facility: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Facility Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.facility.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> facilityCode:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.facility.facilityCode }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> address:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.facility.address }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> FacilityType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.facility.facilityType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewFacilityComponent
