import React, { Component } from 'react'
import CycleCountService from '../services/CycleCountService'

class ViewCycleCountComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            cycleCount: {}
        }
    }

    componentDidMount(){
        CycleCountService.getCycleCountById(this.state.id).then( res => {
            this.setState({cycleCount: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View CycleCount Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> countNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.cycleCount.countNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> scheduledDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.cycleCount.scheduledDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> performedDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.cycleCount.performedDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> approvedBy:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.cycleCount.approvedBy }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.cycleCount.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewCycleCountComponent
