import React, { Component } from 'react'
import ChargebackService from '../services/ChargebackService'

class ViewChargebackComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            chargeback: {}
        }
    }

    componentDidMount(){
        ChargebackService.getChargebackById(this.state.id).then( res => {
            this.setState({chargeback: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Chargeback Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> chargebackReference:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.chargeback.chargebackReference }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> amount:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.chargeback.amount }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> postedAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.chargeback.postedAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Stage:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.chargeback.stage }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.chargeback.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewChargebackComponent
