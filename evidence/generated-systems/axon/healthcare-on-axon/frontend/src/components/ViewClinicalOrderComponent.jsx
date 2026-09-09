import React, { Component } from 'react'
import ClinicalOrderService from '../services/ClinicalOrderService'

class ViewClinicalOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            clinicalOrder: {}
        }
    }

    componentDidMount(){
        ClinicalOrderService.getClinicalOrderById(this.state.id).then( res => {
            this.setState({clinicalOrder: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View ClinicalOrder Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> orderNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.clinicalOrder.orderNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.clinicalOrder.status }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> OrderType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.clinicalOrder.orderType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Priority:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.clinicalOrder.priority }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewClinicalOrderComponent
