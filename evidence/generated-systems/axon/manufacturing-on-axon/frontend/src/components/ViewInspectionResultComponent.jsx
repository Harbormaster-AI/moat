import React, { Component } from 'react'
import InspectionResultService from '../services/InspectionResultService'

class ViewInspectionResultComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            inspectionResult: {}
        }
    }

    componentDidMount(){
        InspectionResultService.getInspectionResultById(this.state.id).then( res => {
            this.setState({inspectionResult: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View InspectionResult Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> resultValue:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inspectionResult.resultValue }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> recordedOn:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inspectionResult.recordedOn }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> notes:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inspectionResult.notes }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ResultStatus:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inspectionResult.resultStatus }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewInspectionResultComponent
