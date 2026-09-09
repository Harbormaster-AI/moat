import React, { Component } from 'react'
import SettlementBatchService from '../services/SettlementBatchService'

class ListSettlementBatchComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                settlementBatchs: []
        }
        this.addSettlementBatch = this.addSettlementBatch.bind(this);
        this.editSettlementBatch = this.editSettlementBatch.bind(this);
        this.deleteSettlementBatch = this.deleteSettlementBatch.bind(this);
    }

    deleteSettlementBatch(id){
        SettlementBatchService.deleteSettlementBatch(id).then( res => {
            this.setState({settlementBatchs: this.state.settlementBatchs.filter(settlementBatch => settlementBatch.settlementBatchId !== id)});
        });
    }
    viewSettlementBatch(id){
        this.props.history.push(`/view-settlementBatch/${id}`);
    }
    editSettlementBatch(id){
        this.props.history.push(`/add-settlementBatch/${id}`);
    }

    componentDidMount(){
        SettlementBatchService.getSettlementBatchs().then((res) => {
            this.setState({ settlementBatchs: res.data});
        });
    }

    addSettlementBatch(){
        this.props.history.push('/add-settlementBatch/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">SettlementBatch List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addSettlementBatch}> Add SettlementBatch</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> BatchId </th>
                                    <th> PeriodStart </th>
                                    <th> PeriodEnd </th>
                                    <th> TotalVolume </th>
                                    <th> TotalCount </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.settlementBatchs.map(
                                        settlementBatch => 
                                        <tr key = {settlementBatch.settlementBatchId}>
                                             <td> { settlementBatch.batchId } </td>
                                             <td> { settlementBatch.periodStart } </td>
                                             <td> { settlementBatch.periodEnd } </td>
                                             <td> { settlementBatch.totalVolume } </td>
                                             <td> { settlementBatch.totalCount } </td>
                                             <td> { settlementBatch.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editSettlementBatch(settlementBatch.settlementBatchId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteSettlementBatch(settlementBatch.settlementBatchId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewSettlementBatch(settlementBatch.settlementBatchId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListSettlementBatchComponent
