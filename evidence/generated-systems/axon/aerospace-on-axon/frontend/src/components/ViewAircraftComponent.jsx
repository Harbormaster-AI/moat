import React, { Component } from 'react'
import AircraftService from '../services/AircraftService'

class ViewAircraftComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            aircraft: {}
        }
    }

    componentDidMount(){
        AircraftService.getAircraftById(this.state.id).then( res => {
            this.setState({aircraft: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Aircraft Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> msn:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.aircraft.msn }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> deliveryDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.aircraft.deliveryDate }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewAircraftComponent
