import React, { Component } from 'react'
import RefundService from '../services/RefundService'

class ListRefundComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                refunds: []
        }
        this.addRefund = this.addRefund.bind(this);
        this.editRefund = this.editRefund.bind(this);
        this.deleteRefund = this.deleteRefund.bind(this);
    }

    deleteRefund(id){
        RefundService.deleteRefund(id).then( res => {
            this.setState({refunds: this.state.refunds.filter(refund => refund.refundId !== id)});
        });
    }
    viewRefund(id){
        this.props.history.push(`/view-refund/${id}`);
    }
    editRefund(id){
        this.props.history.push(`/add-refund/${id}`);
    }

    componentDidMount(){
        RefundService.getRefunds().then((res) => {
            this.setState({ refunds: res.data});
        });
    }

    addRefund(){
        this.props.history.push('/add-refund/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Refund List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addRefund}> Add Refund</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> RefundNumber </th>
                                    <th> Amount </th>
                                    <th> Reason </th>
                                    <th> CreatedAt </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.refunds.map(
                                        refund => 
                                        <tr key = {refund.refundId}>
                                             <td> { refund.refundNumber } </td>
                                             <td> { refund.amount } </td>
                                             <td> { refund.reason } </td>
                                             <td> { refund.createdAt } </td>
                                             <td> { refund.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editRefund(refund.refundId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteRefund(refund.refundId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewRefund(refund.refundId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListRefundComponent
