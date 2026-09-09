import React, { Component } from 'react'
import FraudSignalService from '../services/FraudSignalService'

class ViewFraudSignalComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            fraudSignal: {}
        }
    }

    componentDidMount(){
        FraudSignalService.getFraudSignalById(this.state.id).then( res => {
            this.setState({fraudSignal: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View FraudSignal Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.fraudSignal.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ruleLogic:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.fraudSignal.ruleLogic }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> SignalType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.fraudSignal.signalType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewFraudSignalComponent
