import React, { Component } from 'react'
import GoodsReceiptLineService from '../services/GoodsReceiptLineService'

class ListGoodsReceiptLineComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                goodsReceiptLines: []
        }
        this.addGoodsReceiptLine = this.addGoodsReceiptLine.bind(this);
        this.editGoodsReceiptLine = this.editGoodsReceiptLine.bind(this);
        this.deleteGoodsReceiptLine = this.deleteGoodsReceiptLine.bind(this);
    }

    deleteGoodsReceiptLine(id){
        GoodsReceiptLineService.deleteGoodsReceiptLine(id).then( res => {
            this.setState({goodsReceiptLines: this.state.goodsReceiptLines.filter(goodsReceiptLine => goodsReceiptLine.goodsReceiptLineId !== id)});
        });
    }
    viewGoodsReceiptLine(id){
        this.props.history.push(`/view-goodsReceiptLine/${id}`);
    }
    editGoodsReceiptLine(id){
        this.props.history.push(`/add-goodsReceiptLine/${id}`);
    }

    componentDidMount(){
        GoodsReceiptLineService.getGoodsReceiptLines().then((res) => {
            this.setState({ goodsReceiptLines: res.data});
        });
    }

    addGoodsReceiptLine(){
        this.props.history.push('/add-goodsReceiptLine/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">GoodsReceiptLine List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addGoodsReceiptLine}> Add GoodsReceiptLine</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> LineNumber </th>
                                    <th> ReceivedQuantity </th>
                                    <th> AcceptedQuantity </th>
                                    <th> RejectedQuantity </th>
                                    <th> Lot </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.goodsReceiptLines.map(
                                        goodsReceiptLine => 
                                        <tr key = {goodsReceiptLine.goodsReceiptLineId}>
                                             <td> { goodsReceiptLine.lineNumber } </td>
                                             <td> { goodsReceiptLine.receivedQuantity } </td>
                                             <td> { goodsReceiptLine.acceptedQuantity } </td>
                                             <td> { goodsReceiptLine.rejectedQuantity } </td>
                                             <td> { goodsReceiptLine.lot } </td>
                                             <td>
                                                 <button onClick={ () => this.editGoodsReceiptLine(goodsReceiptLine.goodsReceiptLineId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteGoodsReceiptLine(goodsReceiptLine.goodsReceiptLineId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewGoodsReceiptLine(goodsReceiptLine.goodsReceiptLineId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListGoodsReceiptLineComponent
