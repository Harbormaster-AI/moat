import React, { Component } from 'react'
import JobRequisitionService from '../services/JobRequisitionService'

class ViewJobRequisitionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            jobRequisition: {}
        }
    }

    componentDidMount(){
        JobRequisitionService.getJobRequisitionById(this.state.id).then( res => {
            this.setState({jobRequisition: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View JobRequisition Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> requisitionNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.jobRequisition.requisitionNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> title:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.jobRequisition.title }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> openings:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.jobRequisition.openings }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> targetStartDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.jobRequisition.targetStartDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.jobRequisition.status }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Priority:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.jobRequisition.priority }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewJobRequisitionComponent
