import React, { Component } from 'react'
import LandingGearService from '../services/LandingGearService'

class ViewLandingGearComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            landingGear: {}
        }
    }

    componentDidMount(){
        LandingGearService.getLandingGearById(this.state.id).then( res => {
            this.setState({landingGear: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View LandingGear Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> supplierPartNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.landingGear.supplierPartNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> GearType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.landingGear.gearType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewLandingGearComponent
