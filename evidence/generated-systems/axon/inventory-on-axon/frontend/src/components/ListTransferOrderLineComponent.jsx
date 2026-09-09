import React, { Component } from 'react'
import TransferOrderLineService from '../services/TransferOrderLineService'

class ListTransferOrderLineComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                transferOrderLines: []
        }
        this.addTransferOrderLine = this.addTransferOrderLine.bind(this);
        this.editTransferOrderLine = this.editTransferOrderLine.bind(this);
        this.deleteTransferOrderLine = this.deleteTransferOrderLine.bind(this);
    }

    deleteTransferOrderLine(id){
        TransferOrderLineService.deleteTransferOrderLine(id).then( res => {
            this.setState({transferOrderLines: this.state.transferOrderLines.filter(transferOrderLine => transferOrderLine.transferOrderLineId !== id)});
        });
    }
    viewTransferOrderLine(id){
        this.props.history.push(`/view-transferOrderLine/${id}`);
    }
    editTransferOrderLine(id){
        this.props.history.push(`/add-transferOrderLine/${id}`);
    }

    componentDidMount(){
        TransferOrderLineService.getTransferOrderLines().then((res) => {
            this.setState({ transferOrderLines: res.data});
        });
    }

    addTransferOrderLine(){
        this.props.history.push('/add-transferOrderLine/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">TransferOrderLine List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addTransferOrderLine}> Add TransferOrderLine</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> LineNumber </th>
                                    <th> Quantity </th>
                                    <th> UnitOfMeasure </th>
                                    <th> StockStatus </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.transferOrderLines.map(
                                        transferOrderLine => 
                                        <tr key = {transferOrderLine.transferOrderLineId}>
                                             <td> { transferOrderLine.lineNumber } </td>
                                             <td> { transferOrderLine.quantity } </td>
                                             <td> { transferOrderLine.unitOfMeasure } </td>
                                             <td> { transferOrderLine.stockStatus } </td>
                                             <td>
                                                 <button onClick={ () => this.editTransferOrderLine(transferOrderLine.transferOrderLineId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteTransferOrderLine(transferOrderLine.transferOrderLineId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewTransferOrderLine(transferOrderLine.transferOrderLineId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListTransferOrderLineComponent
