import React, { Component } from 'react'
import OfferService from '../services/OfferService'

class ListOfferComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                offers: []
        }
        this.addOffer = this.addOffer.bind(this);
        this.editOffer = this.editOffer.bind(this);
        this.deleteOffer = this.deleteOffer.bind(this);
    }

    deleteOffer(id){
        OfferService.deleteOffer(id).then( res => {
            this.setState({offers: this.state.offers.filter(offer => offer.offerId !== id)});
        });
    }
    viewOffer(id){
        this.props.history.push(`/view-offer/${id}`);
    }
    editOffer(id){
        this.props.history.push(`/add-offer/${id}`);
    }

    componentDidMount(){
        OfferService.getOffers().then((res) => {
            this.setState({ offers: res.data});
        });
    }

    addOffer(){
        this.props.history.push('/add-offer/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Offer List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addOffer}> Add Offer</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> OfferNumber </th>
                                    <th> ProposedStartDate </th>
                                    <th> BaseSalary </th>
                                    <th> SignOnBonus </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.offers.map(
                                        offer => 
                                        <tr key = {offer.offerId}>
                                             <td> { offer.offerNumber } </td>
                                             <td> { offer.proposedStartDate } </td>
                                             <td> { offer.baseSalary } </td>
                                             <td> { offer.signOnBonus } </td>
                                             <td> { offer.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editOffer(offer.offerId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteOffer(offer.offerId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewOffer(offer.offerId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListOfferComponent
