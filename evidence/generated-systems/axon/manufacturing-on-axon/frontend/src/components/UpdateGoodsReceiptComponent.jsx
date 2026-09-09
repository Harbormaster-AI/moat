import React, { Component } from 'react'
import GoodsReceiptService from '../services/GoodsReceiptService';

class UpdateGoodsReceiptComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                receiptNumber: '',
                receiptDate: '',
                status: ''
        }
        this.updateGoodsReceipt = this.updateGoodsReceipt.bind(this);

        this.changereceiptNumberHandler = this.changereceiptNumberHandler.bind(this);
        this.changereceiptDateHandler = this.changereceiptDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        GoodsReceiptService.getGoodsReceiptById(this.state.id).then( (res) =>{
            let goodsReceipt = res.data;
            this.setState({
                receiptNumber: goodsReceipt.receiptNumber,
                receiptDate: goodsReceipt.receiptDate,
                status: goodsReceipt.status
            });
        });
    }

    updateGoodsReceipt = (e) => {
        e.preventDefault();
        let goodsReceipt = {
            goodsReceiptId: this.state.id,
            receiptNumber: this.state.receiptNumber,
            receiptDate: this.state.receiptDate,
            status: this.state.status
        };
        console.log('goodsReceipt => ' + JSON.stringify(goodsReceipt));
        console.log('id => ' + JSON.stringify(this.state.id));
        GoodsReceiptService.updateGoodsReceipt(goodsReceipt).then( res => {
            this.props.history.push('/goodsReceipts');
        });
    }

    changereceiptNumberHandler= (event) => {
        this.setState({receiptNumber: event.target.value});
    }
    changereceiptDateHandler= (event) => {
        this.setState({receiptDate: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/goodsReceipts');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update GoodsReceipt</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> receiptNumber: </label>
                                                <input placeholder="receiptNumber" name="receiptNumber" className="form-control" value={this.state.receiptNumber} onChange={this.changereceiptNumberHandler}/>

                                            <label> receiptDate: </label>
                                                <input type="date" placeholder="receiptDate" name="receiptDate" className="form-control" value={this.state.receiptDate} onChange={this.changereceiptDateHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Open
                      </option>
                      <option name="Status" className="form-control" >
                          PartiallyProcessed
                      </option>
                      <option name="Status" className="form-control" >
                          Completed
                      </option>
                      <option name="Status" className="form-control" >
                          Rejected
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateGoodsReceipt}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>

                   </div>
            </div>
        )
    }
}

export default UpdateGoodsReceiptComponent
