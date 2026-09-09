import React, { Component } from 'react'
import StockAdjustmentLineService from '../services/StockAdjustmentLineService'

class ViewStockAdjustmentLineComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            stockAdjustmentLine: {}
        }
    }

    componentDidMount(){
        StockAdjustmentLineService.getStockAdjustmentLineById(this.state.id).then( res => {
            this.setState({stockAdjustmentLine: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View StockAdjustmentLine Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> lineNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.stockAdjustmentLine.lineNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> quantity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.stockAdjustmentLine.quantity }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> UnitOfMeasure:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.stockAdjustmentLine.unitOfMeasure }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> StockStatus:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.stockAdjustmentLine.stockStatus }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewStockAdjustmentLineComponent
