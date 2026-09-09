import React, { Component } from 'react'
import AircraftPackageService from '../services/AircraftPackageService'

class ViewAircraftPackageComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            aircraftPackage: {}
        }
    }

    componentDidMount(){
        AircraftPackageService.getAircraftPackageById(this.state.id).then( res => {
            this.setState({aircraftPackage: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View AircraftPackage Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.aircraftPackage.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> PackageType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.aircraftPackage.packageType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewAircraftPackageComponent
