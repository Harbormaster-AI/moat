import React, { Component } from 'react'
import OpportunityService from '../services/OpportunityService'

class ViewOpportunityComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            opportunity: {}
        }
    }

    componentDidMount(){
        OpportunityService.getOpportunityById(this.state.id).then( res => {
            this.setState({opportunity: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Opportunity Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.opportunity.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> amount:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.opportunity.amount }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> closeDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.opportunity.closeDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> probability:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.opportunity.probability }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> description:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.opportunity.description }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Stage:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.opportunity.stage }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Type:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.opportunity.type }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ForecastCategory:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.opportunity.forecastCategory }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewOpportunityComponent
