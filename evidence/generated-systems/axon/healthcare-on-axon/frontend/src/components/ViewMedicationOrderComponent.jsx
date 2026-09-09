import React, { Component } from 'react'
import MedicationOrderService from '../services/MedicationOrderService'

class ViewMedicationOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            medicationOrder: {}
        }
    }

    componentDidMount(){
        MedicationOrderService.getMedicationOrderById(this.state.id).then( res => {
            this.setState({medicationOrder: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View MedicationOrder Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> medicationCode:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.medicationOrder.medicationCode }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> dose:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.medicationOrder.dose }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> frequency:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.medicationOrder.frequency }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> duration:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.medicationOrder.duration }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Route:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.medicationOrder.route }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewMedicationOrderComponent
