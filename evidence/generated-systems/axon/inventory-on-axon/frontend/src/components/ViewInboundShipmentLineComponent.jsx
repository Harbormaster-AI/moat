import React, { Component } from 'react'
import InboundShipmentLineService from '../services/InboundShipmentLineService'

class ViewInboundShipmentLineComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            inboundShipmentLine: {}
        }
    }

    componentDidMount(){
        InboundShipmentLineService.getInboundShipmentLineById(this.state.id).then( res => {
            this.setState({inboundShipmentLine: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View InboundShipmentLine Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> lineNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inboundShipmentLine.lineNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> quantity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inboundShipmentLine.quantity }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> UnitOfMeasure:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inboundShipmentLine.unitOfMeasure }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> StockStatus:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inboundShipmentLine.stockStatus }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewInboundShipmentLineComponent
