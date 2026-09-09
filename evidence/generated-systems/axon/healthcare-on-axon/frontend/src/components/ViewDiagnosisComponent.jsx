import React, { Component } from 'react'
import DiagnosisService from '../services/DiagnosisService'

class ViewDiagnosisComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            diagnosis: {}
        }
    }

    componentDidMount(){
        DiagnosisService.getDiagnosisById(this.state.id).then( res => {
            this.setState({diagnosis: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Diagnosis Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> code:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.diagnosis.code }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> description:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.diagnosis.description }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> onsetDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.diagnosis.onsetDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Certainty:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.diagnosis.certainty }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewDiagnosisComponent
