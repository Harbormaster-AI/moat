import React, { Component } from 'react'
import ChargebackService from '../services/ChargebackService'

class ListChargebackComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                chargebacks: []
        }
        this.addChargeback = this.addChargeback.bind(this);
        this.editChargeback = this.editChargeback.bind(this);
        this.deleteChargeback = this.deleteChargeback.bind(this);
    }

    deleteChargeback(id){
        ChargebackService.deleteChargeback(id).then( res => {
            this.setState({chargebacks: this.state.chargebacks.filter(chargeback => chargeback.chargebackId !== id)});
        });
    }
    viewChargeback(id){
        this.props.history.push(`/view-chargeback/${id}`);
    }
    editChargeback(id){
        this.props.history.push(`/add-chargeback/${id}`);
    }

    componentDidMount(){
        ChargebackService.getChargebacks().then((res) => {
            this.setState({ chargebacks: res.data});
        });
    }

    addChargeback(){
        this.props.history.push('/add-chargeback/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Chargeback List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addChargeback}> Add Chargeback</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ChargebackReference </th>
                                    <th> Amount </th>
                                    <th> PostedAt </th>
                                    <th> Stage </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.chargebacks.map(
                                        chargeback => 
                                        <tr key = {chargeback.chargebackId}>
                                             <td> { chargeback.chargebackReference } </td>
                                             <td> { chargeback.amount } </td>
                                             <td> { chargeback.postedAt } </td>
                                             <td> { chargeback.stage } </td>
                                             <td> { chargeback.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editChargeback(chargeback.chargebackId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteChargeback(chargeback.chargebackId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewChargeback(chargeback.chargebackId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListChargebackComponent
