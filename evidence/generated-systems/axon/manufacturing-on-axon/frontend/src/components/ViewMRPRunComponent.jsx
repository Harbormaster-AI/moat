import React, { Component } from 'react'
import MRPRunService from '../services/MRPRunService'

class ViewMRPRunComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            mRPRun: {}
        }
    }

    componentDidMount(){
        MRPRunService.getMRPRunById(this.state.id).then( res => {
            this.setState({mRPRun: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View MRPRun Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> runNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.mRPRun.runNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> runDateTime:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.mRPRun.runDateTime }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> planningHorizonDays:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.mRPRun.planningHorizonDays }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.mRPRun.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewMRPRunComponent
