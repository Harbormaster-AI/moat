import React, { Component } from 'react'
import GiftCardService from '../services/GiftCardService'

class ListGiftCardComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                giftCards: []
        }
        this.addGiftCard = this.addGiftCard.bind(this);
        this.editGiftCard = this.editGiftCard.bind(this);
        this.deleteGiftCard = this.deleteGiftCard.bind(this);
    }

    deleteGiftCard(id){
        GiftCardService.deleteGiftCard(id).then( res => {
            this.setState({giftCards: this.state.giftCards.filter(giftCard => giftCard.giftCardId !== id)});
        });
    }
    viewGiftCard(id){
        this.props.history.push(`/view-giftCard/${id}`);
    }
    editGiftCard(id){
        this.props.history.push(`/add-giftCard/${id}`);
    }

    componentDidMount(){
        GiftCardService.getGiftCards().then((res) => {
            this.setState({ giftCards: res.data});
        });
    }

    addGiftCard(){
        this.props.history.push('/add-giftCard/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">GiftCard List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addGiftCard}> Add GiftCard</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Code </th>
                                    <th> Balance </th>
                                    <th> ExpirationDate </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.giftCards.map(
                                        giftCard => 
                                        <tr key = {giftCard.giftCardId}>
                                             <td> { giftCard.code } </td>
                                             <td> { giftCard.balance } </td>
                                             <td> { giftCard.expirationDate } </td>
                                             <td> { giftCard.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editGiftCard(giftCard.giftCardId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteGiftCard(giftCard.giftCardId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewGiftCard(giftCard.giftCardId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListGiftCardComponent
