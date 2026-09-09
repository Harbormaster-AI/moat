import React, { Component } from 'react'
import DealService from '../services/DealService'

class ListDealComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                deals: []
        }
        this.addDeal = this.addDeal.bind(this);
        this.editDeal = this.editDeal.bind(this);
        this.deleteDeal = this.deleteDeal.bind(this);
    }

    deleteDeal(id){
        DealService.deleteDeal(id).then( res => {
            this.setState({deals: this.state.deals.filter(deal => deal.dealId !== id)});
        });
    }
    viewDeal(id){
        this.props.history.push(`/view-deal/${id}`);
    }
    editDeal(id){
        this.props.history.push(`/add-deal/${id}`);
    }

    componentDidMount(){
        DealService.getDeals().then((res) => {
            this.setState({ deals: res.data});
        });
    }

    addDeal(){
        this.props.history.push('/add-deal/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Deal List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addDeal}> Add Deal</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> FloorPrice </th>
                                    <th> DealType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.deals.map(
                                        deal => 
                                        <tr key = {deal.dealId}>
                                             <td> { deal.floorPrice } </td>
                                             <td> { deal.dealType } </td>
                                             <td>
                                                 <button onClick={ () => this.editDeal(deal.dealId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteDeal(deal.dealId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewDeal(deal.dealId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListDealComponent
