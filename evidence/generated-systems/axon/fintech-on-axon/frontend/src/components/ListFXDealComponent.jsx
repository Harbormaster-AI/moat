import React, { Component } from 'react'
import FXDealService from '../services/FXDealService'

class ListFXDealComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                fXDeals: []
        }
        this.addFXDeal = this.addFXDeal.bind(this);
        this.editFXDeal = this.editFXDeal.bind(this);
        this.deleteFXDeal = this.deleteFXDeal.bind(this);
    }

    deleteFXDeal(id){
        FXDealService.deleteFXDeal(id).then( res => {
            this.setState({fXDeals: this.state.fXDeals.filter(fXDeal => fXDeal.fXDealId !== id)});
        });
    }
    viewFXDeal(id){
        this.props.history.push(`/view-fXDeal/${id}`);
    }
    editFXDeal(id){
        this.props.history.push(`/add-fXDeal/${id}`);
    }

    componentDidMount(){
        FXDealService.getFXDeals().then((res) => {
            this.setState({ fXDeals: res.data});
        });
    }

    addFXDeal(){
        this.props.history.push('/add-fXDeal/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">FXDeal List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addFXDeal}> Add FXDeal</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> DealReference </th>
                                    <th> BaseCurrency </th>
                                    <th> QuoteCurrency </th>
                                    <th> Rate </th>
                                    <th> Amount </th>
                                    <th> SettlementDate </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.fXDeals.map(
                                        fXDeal => 
                                        <tr key = {fXDeal.fXDealId}>
                                             <td> { fXDeal.dealReference } </td>
                                             <td> { fXDeal.baseCurrency } </td>
                                             <td> { fXDeal.quoteCurrency } </td>
                                             <td> { fXDeal.rate } </td>
                                             <td> { fXDeal.amount } </td>
                                             <td> { fXDeal.settlementDate } </td>
                                             <td> { fXDeal.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editFXDeal(fXDeal.fXDealId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteFXDeal(fXDeal.fXDealId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewFXDeal(fXDeal.fXDealId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListFXDealComponent
