import React, { Component } from 'react'
import GoodsReceiptService from '../services/GoodsReceiptService'

class ViewGoodsReceiptComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            goodsReceipt: {}
        }
    }

    componentDidMount(){
        GoodsReceiptService.getGoodsReceiptById(this.state.id).then( res => {
            this.setState({goodsReceipt: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View GoodsReceipt Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> receiptNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.goodsReceipt.receiptNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> receiptDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.goodsReceipt.receiptDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.goodsReceipt.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewGoodsReceiptComponent
