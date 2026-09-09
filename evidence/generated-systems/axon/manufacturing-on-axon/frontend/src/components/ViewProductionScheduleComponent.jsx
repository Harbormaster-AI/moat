import React, { Component } from 'react'
import ProductionScheduleService from '../services/ProductionScheduleService'

class ViewProductionScheduleComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            productionSchedule: {}
        }
    }

    componentDidMount(){
        ProductionScheduleService.getProductionScheduleById(this.state.id).then( res => {
            this.setState({productionSchedule: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View ProductionSchedule Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> scheduleNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.productionSchedule.scheduleNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> horizonStart:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.productionSchedule.horizonStart }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> horizonEnd:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.productionSchedule.horizonEnd }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.productionSchedule.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewProductionScheduleComponent
