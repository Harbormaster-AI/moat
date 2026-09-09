import React, { Component } from 'react'
import OpportunityStageHistoryService from '../services/OpportunityStageHistoryService'

class ViewOpportunityStageHistoryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            opportunityStageHistory: {}
        }
    }

    componentDidMount(){
        OpportunityStageHistoryService.getOpportunityStageHistoryById(this.state.id).then( res => {
            this.setState({opportunityStageHistory: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View OpportunityStageHistory Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> changedAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.opportunityStageHistory.changedAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> comment:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.opportunityStageHistory.comment }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> FromStage:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.opportunityStageHistory.fromStage }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ToStage:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.opportunityStageHistory.toStage }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewOpportunityStageHistoryComponent
