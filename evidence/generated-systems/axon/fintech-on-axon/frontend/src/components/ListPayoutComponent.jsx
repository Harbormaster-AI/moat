import React, { Component } from 'react'
import PayoutService from '../services/PayoutService'

class ListPayoutComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                payouts: []
        }
        this.addPayout = this.addPayout.bind(this);
        this.editPayout = this.editPayout.bind(this);
        this.deletePayout = this.deletePayout.bind(this);
    }

    deletePayout(id){
        PayoutService.deletePayout(id).then( res => {
            this.setState({payouts: this.state.payouts.filter(payout => payout.payoutId !== id)});
        });
    }
    viewPayout(id){
        this.props.history.push(`/view-payout/${id}`);
    }
    editPayout(id){
        this.props.history.push(`/add-payout/${id}`);
    }

    componentDidMount(){
        PayoutService.getPayouts().then((res) => {
            this.setState({ payouts: res.data});
        });
    }

    addPayout(){
        this.props.history.push('/add-payout/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Payout List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addPayout}> Add Payout</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> PayoutReference </th>
                                    <th> Amount </th>
                                    <th> Currency </th>
                                    <th> ScheduledDate </th>
                                    <th> PaidDate </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.payouts.map(
                                        payout => 
                                        <tr key = {payout.payoutId}>
                                             <td> { payout.payoutReference } </td>
                                             <td> { payout.amount } </td>
                                             <td> { payout.currency } </td>
                                             <td> { payout.scheduledDate } </td>
                                             <td> { payout.paidDate } </td>
                                             <td> { payout.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editPayout(payout.payoutId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deletePayout(payout.payoutId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewPayout(payout.payoutId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListPayoutComponent
