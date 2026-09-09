import React, { Component } from 'react'
import StockAdjustmentService from '../services/StockAdjustmentService'

class ViewStockAdjustmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            stockAdjustment: {}
        }
    }

    componentDidMount(){
        StockAdjustmentService.getStockAdjustmentById(this.state.id).then( res => {
            this.setState({stockAdjustment: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View StockAdjustment Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> adjustmentNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.stockAdjustment.adjustmentNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> reason:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.stockAdjustment.reason }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> adjustmentDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.stockAdjustment.adjustmentDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> AdjustmentType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.stockAdjustment.adjustmentType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.stockAdjustment.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewStockAdjustmentComponent
