import React, { Component } from 'react'
import EvidenceService from '../services/EvidenceService'

class ViewEvidenceComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            evidence: {}
        }
    }

    componentDidMount(){
        EvidenceService.getEvidenceById(this.state.id).then( res => {
            this.setState({evidence: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Evidence Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> title:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.evidence.title }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> locationUrl:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.evidence.locationUrl }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> receivedDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.evidence.receivedDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> EvidenceType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.evidence.evidenceType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewEvidenceComponent
