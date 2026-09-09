import React, { Component } from 'react'
import FraudScenarioService from '../services/FraudScenarioService'

class ViewFraudScenarioComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            fraudScenario: {}
        }
    }

    componentDidMount(){
        FraudScenarioService.getFraudScenarioById(this.state.id).then( res => {
            this.setState({fraudScenario: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View FraudScenario Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.fraudScenario.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> riskAppetite:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.fraudScenario.riskAppetite }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> DetectionType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.fraudScenario.detectionType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewFraudScenarioComponent
