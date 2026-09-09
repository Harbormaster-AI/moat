import React, { Component } from 'react'
import FeeScheduleService from '../services/FeeScheduleService'

class ViewFeeScheduleComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            feeSchedule: {}
        }
    }

    componentDidMount(){
        FeeScheduleService.getFeeScheduleById(this.state.id).then( res => {
            this.setState({feeSchedule: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View FeeSchedule Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.feeSchedule.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> amount:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.feeSchedule.amount }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> percentage:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.feeSchedule.percentage }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> minimum:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.feeSchedule.minimum }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> maximum:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.feeSchedule.maximum }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> FeeType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.feeSchedule.feeType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> CalculationMethod:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.feeSchedule.calculationMethod }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewFeeScheduleComponent
