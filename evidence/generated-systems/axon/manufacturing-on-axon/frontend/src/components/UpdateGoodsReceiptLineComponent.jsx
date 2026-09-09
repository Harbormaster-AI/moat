import React, { Component } from 'react'
import GoodsReceiptLineService from '../services/GoodsReceiptLineService';

class UpdateGoodsReceiptLineComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                lineNumber: '',
                receivedQuantity: '',
                acceptedQuantity: '',
                rejectedQuantity: '',
                lot: ''
        }
        this.updateGoodsReceiptLine = this.updateGoodsReceiptLine.bind(this);

        this.changelineNumberHandler = this.changelineNumberHandler.bind(this);
        this.changereceivedQuantityHandler = this.changereceivedQuantityHandler.bind(this);
        this.changeacceptedQuantityHandler = this.changeacceptedQuantityHandler.bind(this);
        this.changerejectedQuantityHandler = this.changerejectedQuantityHandler.bind(this);
        this.changelotHandler = this.changelotHandler.bind(this);
    }

    componentDidMount(){
        GoodsReceiptLineService.getGoodsReceiptLineById(this.state.id).then( (res) =>{
            let goodsReceiptLine = res.data;
            this.setState({
                lineNumber: goodsReceiptLine.lineNumber,
                receivedQuantity: goodsReceiptLine.receivedQuantity,
                acceptedQuantity: goodsReceiptLine.acceptedQuantity,
                rejectedQuantity: goodsReceiptLine.rejectedQuantity,
                lot: goodsReceiptLine.lot
            });
        });
    }

    updateGoodsReceiptLine = (e) => {
        e.preventDefault();
        let goodsReceiptLine = {
            goodsReceiptLineId: this.state.id,
            lineNumber: this.state.lineNumber,
            receivedQuantity: this.state.receivedQuantity,
            acceptedQuantity: this.state.acceptedQuantity,
            rejectedQuantity: this.state.rejectedQuantity,
            lot: this.state.lot
        };
        console.log('goodsReceiptLine => ' + JSON.stringify(goodsReceiptLine));
        console.log('id => ' + JSON.stringify(this.state.id));
        GoodsReceiptLineService.updateGoodsReceiptLine(goodsReceiptLine).then( res => {
            this.props.history.push('/goodsReceiptLines');
        });
    }

    changelineNumberHandler= (event) => {
        this.setState({lineNumber: event.target.value});
    }
    changereceivedQuantityHandler= (event) => {
        this.setState({receivedQuantity: event.target.value});
    }
    changeacceptedQuantityHandler= (event) => {
        this.setState({acceptedQuantity: event.target.value});
    }
    changerejectedQuantityHandler= (event) => {
        this.setState({rejectedQuantity: event.target.value});
    }
    changelotHandler= (event) => {
        this.setState({lot: event.target.value});
    }

    cancel(){
        this.props.history.push('/goodsReceiptLines');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update GoodsReceiptLine</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> lineNumber: </label>
                                                <input type="number" placeholder="lineNumber" name="lineNumber" className="form-control" value={this.state.lineNumber} onChange={this.changelineNumberHandler}/>

                                            <label> receivedQuantity: </label>
                                                <input placeholder="receivedQuantity" name="receivedQuantity" className="form-control" value={this.state.receivedQuantity} onChange={this.changereceivedQuantityHandler}/>

                                            <label> acceptedQuantity: </label>
                                                <input placeholder="acceptedQuantity" name="acceptedQuantity" className="form-control" value={this.state.acceptedQuantity} onChange={this.changeacceptedQuantityHandler}/>

                                            <label> rejectedQuantity: </label>
                                                <input placeholder="rejectedQuantity" name="rejectedQuantity" className="form-control" value={this.state.rejectedQuantity} onChange={this.changerejectedQuantityHandler}/>

                                            <label> lot: </label>
                                                <input placeholder="lot" name="lot" className="form-control" value={this.state.lot} onChange={this.changelotHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateGoodsReceiptLine}>Save</button>
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

export default UpdateGoodsReceiptLineComponent
