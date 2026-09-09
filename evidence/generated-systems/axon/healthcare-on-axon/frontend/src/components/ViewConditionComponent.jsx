import React, { Component } from 'react'
import ConditionService from '../services/ConditionService'

class ViewConditionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            condition: {}
        }
    }

    componentDidMount(){
        ConditionService.getConditionById(this.state.id).then( res => {
            this.setState({condition: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Condition Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> code:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.condition.code }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> onsetDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.condition.onsetDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> abatementDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.condition.abatementDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ClinicalStatus:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.condition.clinicalStatus }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> VerificationStatus:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.condition.verificationStatus }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewConditionComponent
