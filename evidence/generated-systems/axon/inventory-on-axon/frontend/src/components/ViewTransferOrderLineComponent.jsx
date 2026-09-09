import React, { Component } from 'react'
import TransferOrderLineService from '../services/TransferOrderLineService'

class ViewTransferOrderLineComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            transferOrderLine: {}
        }
    }

    componentDidMount(){
        TransferOrderLineService.getTransferOrderLineById(this.state.id).then( res => {
            this.setState({transferOrderLine: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View TransferOrderLine Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> lineNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.transferOrderLine.lineNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> quantity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.transferOrderLine.quantity }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> UnitOfMeasure:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.transferOrderLine.unitOfMeasure }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> StockStatus:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.transferOrderLine.stockStatus }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewTransferOrderLineComponent
