import React, { Component } from 'react'
import IncidentService from '../services/IncidentService'

class ViewIncidentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            incident: {}
        }
    }

    componentDidMount(){
        IncidentService.getIncidentById(this.state.id).then( res => {
            this.setState({incident: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Incident Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> location:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.incident.location }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> description:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.incident.description }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> IncidentType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.incident.incidentType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewIncidentComponent
