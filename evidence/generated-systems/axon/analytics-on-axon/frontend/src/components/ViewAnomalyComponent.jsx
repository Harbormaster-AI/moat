import React, { Component } from 'react'
import AnomalyService from '../services/AnomalyService'

class ViewAnomalyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            anomaly: {}
        }
    }

    componentDidMount(){
        AnomalyService.getAnomalyById(this.state.id).then( res => {
            this.setState({anomaly: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Anomaly Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> occurredAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.anomaly.occurredAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> details:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.anomaly.details }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> AnomalyType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.anomaly.anomalyType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Severity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.anomaly.severity }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewAnomalyComponent
