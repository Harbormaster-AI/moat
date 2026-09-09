import React, { Component } from 'react'
import InboundShipmentLineService from '../services/InboundShipmentLineService'

class ListInboundShipmentLineComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                inboundShipmentLines: []
        }
        this.addInboundShipmentLine = this.addInboundShipmentLine.bind(this);
        this.editInboundShipmentLine = this.editInboundShipmentLine.bind(this);
        this.deleteInboundShipmentLine = this.deleteInboundShipmentLine.bind(this);
    }

    deleteInboundShipmentLine(id){
        InboundShipmentLineService.deleteInboundShipmentLine(id).then( res => {
            this.setState({inboundShipmentLines: this.state.inboundShipmentLines.filter(inboundShipmentLine => inboundShipmentLine.inboundShipmentLineId !== id)});
        });
    }
    viewInboundShipmentLine(id){
        this.props.history.push(`/view-inboundShipmentLine/${id}`);
    }
    editInboundShipmentLine(id){
        this.props.history.push(`/add-inboundShipmentLine/${id}`);
    }

    componentDidMount(){
        InboundShipmentLineService.getInboundShipmentLines().then((res) => {
            this.setState({ inboundShipmentLines: res.data});
        });
    }

    addInboundShipmentLine(){
        this.props.history.push('/add-inboundShipmentLine/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">InboundShipmentLine List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addInboundShipmentLine}> Add InboundShipmentLine</button>
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
                                    this.state.inboundShipmentLines.map(
                                        inboundShipmentLine => 
                                        <tr key = {inboundShipmentLine.inboundShipmentLineId}>
                                             <td> { inboundShipmentLine.lineNumber } </td>
                                             <td> { inboundShipmentLine.quantity } </td>
                                             <td> { inboundShipmentLine.unitOfMeasure } </td>
                                             <td> { inboundShipmentLine.stockStatus } </td>
                                             <td>
                                                 <button onClick={ () => this.editInboundShipmentLine(inboundShipmentLine.inboundShipmentLineId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteInboundShipmentLine(inboundShipmentLine.inboundShipmentLineId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewInboundShipmentLine(inboundShipmentLine.inboundShipmentLineId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListInboundShipmentLineComponent
