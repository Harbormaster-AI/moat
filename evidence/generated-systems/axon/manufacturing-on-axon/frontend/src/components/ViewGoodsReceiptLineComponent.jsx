import React, { Component } from 'react'
import GoodsReceiptLineService from '../services/GoodsReceiptLineService'

class ViewGoodsReceiptLineComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            goodsReceiptLine: {}
        }
    }

    componentDidMount(){
        GoodsReceiptLineService.getGoodsReceiptLineById(this.state.id).then( res => {
            this.setState({goodsReceiptLine: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View GoodsReceiptLine Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> lineNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.goodsReceiptLine.lineNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> receivedQuantity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.goodsReceiptLine.receivedQuantity }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> acceptedQuantity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.goodsReceiptLine.acceptedQuantity }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> rejectedQuantity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.goodsReceiptLine.rejectedQuantity }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> lot:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.goodsReceiptLine.lot }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewGoodsReceiptLineComponent
