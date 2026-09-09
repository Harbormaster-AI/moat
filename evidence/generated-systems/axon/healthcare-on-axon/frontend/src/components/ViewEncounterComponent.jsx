import React, { Component } from 'react'
import EncounterService from '../services/EncounterService'

class ViewEncounterComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            encounter: {}
        }
    }

    componentDidMount(){
        EncounterService.getEncounterById(this.state.id).then( res => {
            this.setState({encounter: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Encounter Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> encounterNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.encounter.encounterNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> startDateTime:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.encounter.startDateTime }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> endDateTime:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.encounter.endDateTime }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.encounter.status }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> EncounterType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.encounter.encounterType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewEncounterComponent
