import React, { Component } from 'react'
import GoodsReceiptLineService from '../services/GoodsReceiptLineService';

class CreateGoodsReceiptLineComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                lineNumber: '',
                receivedQuantity: '',
                acceptedQuantity: '',
                rejectedQuantity: '',
                lot: ''
        }
        this.changelineNumberHandler = this.changelineNumberHandler.bind(this);
        this.changereceivedQuantityHandler = this.changereceivedQuantityHandler.bind(this);
        this.changeacceptedQuantityHandler = this.changeacceptedQuantityHandler.bind(this);
        this.changerejectedQuantityHandler = this.changerejectedQuantityHandler.bind(this);
        this.changelotHandler = this.changelotHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateGoodsReceiptLine = (e) => {
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

        // step 5
        if(this.state.id === '_add'){
            goodsReceiptLine.goodsReceiptLineId=''
            GoodsReceiptLineService.createGoodsReceiptLine(goodsReceiptLine).then(res =>{
                this.props.history.push('/goodsReceiptLines');
            });
        }else{
            GoodsReceiptLineService.updateGoodsReceiptLine(goodsReceiptLine).then( res => {
                this.props.history.push('/goodsReceiptLines');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add GoodsReceiptLine</h3>
        }else{
            return <h3 className="text-center">Update GoodsReceiptLine</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> lineNumber:&emsp; </label>
                                                <input type="number" placeholder="lineNumber" name="lineNumber" className="form-control" value={this.state.lineNumber} onChange={this.changelineNumberHandler}/>

                                            <label> receivedQuantity:&emsp; </label>
                                                <input placeholder="receivedQuantity" name="receivedQuantity" className="form-control" value={this.state.receivedQuantity} onChange={this.changereceivedQuantityHandler}/>

                                            <label> acceptedQuantity:&emsp; </label>
                                                <input placeholder="acceptedQuantity" name="acceptedQuantity" className="form-control" value={this.state.acceptedQuantity} onChange={this.changeacceptedQuantityHandler}/>

                                            <label> rejectedQuantity:&emsp; </label>
                                                <input placeholder="rejectedQuantity" name="rejectedQuantity" className="form-control" value={this.state.rejectedQuantity} onChange={this.changerejectedQuantityHandler}/>

                                            <label> lot:&emsp; </label>
                                                <input placeholder="lot" name="lot" className="form-control" value={this.state.lot} onChange={this.changelotHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateGoodsReceiptLine}>Save</button>
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

export default CreateGoodsReceiptLineComponent
