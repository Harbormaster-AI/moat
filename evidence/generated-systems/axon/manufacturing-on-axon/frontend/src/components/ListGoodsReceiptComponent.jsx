import React, { Component } from 'react'
import GoodsReceiptService from '../services/GoodsReceiptService'

class ListGoodsReceiptComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                goodsReceipts: []
        }
        this.addGoodsReceipt = this.addGoodsReceipt.bind(this);
        this.editGoodsReceipt = this.editGoodsReceipt.bind(this);
        this.deleteGoodsReceipt = this.deleteGoodsReceipt.bind(this);
    }

    deleteGoodsReceipt(id){
        GoodsReceiptService.deleteGoodsReceipt(id).then( res => {
            this.setState({goodsReceipts: this.state.goodsReceipts.filter(goodsReceipt => goodsReceipt.goodsReceiptId !== id)});
        });
    }
    viewGoodsReceipt(id){
        this.props.history.push(`/view-goodsReceipt/${id}`);
    }
    editGoodsReceipt(id){
        this.props.history.push(`/add-goodsReceipt/${id}`);
    }

    componentDidMount(){
        GoodsReceiptService.getGoodsReceipts().then((res) => {
            this.setState({ goodsReceipts: res.data});
        });
    }

    addGoodsReceipt(){
        this.props.history.push('/add-goodsReceipt/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">GoodsReceipt List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addGoodsReceipt}> Add GoodsReceipt</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ReceiptNumber </th>
                                    <th> ReceiptDate </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.goodsReceipts.map(
                                        goodsReceipt => 
                                        <tr key = {goodsReceipt.goodsReceiptId}>
                                             <td> { goodsReceipt.receiptNumber } </td>
                                             <td> { goodsReceipt.receiptDate } </td>
                                             <td> { goodsReceipt.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editGoodsReceipt(goodsReceipt.goodsReceiptId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteGoodsReceipt(goodsReceipt.goodsReceiptId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewGoodsReceipt(goodsReceipt.goodsReceiptId)} className="btn btn-outline-info btn-sm">View </button>
                                             </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>

                 </div>

            </div>
        )
    }
}

export default ListGoodsReceiptComponent
