import React, { Component } from 'react'
import GiftCardRedemptionService from '../services/GiftCardRedemptionService'

class ListGiftCardRedemptionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                giftCardRedemptions: []
        }
        this.addGiftCardRedemption = this.addGiftCardRedemption.bind(this);
        this.editGiftCardRedemption = this.editGiftCardRedemption.bind(this);
        this.deleteGiftCardRedemption = this.deleteGiftCardRedemption.bind(this);
    }

    deleteGiftCardRedemption(id){
        GiftCardRedemptionService.deleteGiftCardRedemption(id).then( res => {
            this.setState({giftCardRedemptions: this.state.giftCardRedemptions.filter(giftCardRedemption => giftCardRedemption.giftCardRedemptionId !== id)});
        });
    }
    viewGiftCardRedemption(id){
        this.props.history.push(`/view-giftCardRedemption/${id}`);
    }
    editGiftCardRedemption(id){
        this.props.history.push(`/add-giftCardRedemption/${id}`);
    }

    componentDidMount(){
        GiftCardRedemptionService.getGiftCardRedemptions().then((res) => {
            this.setState({ giftCardRedemptions: res.data});
        });
    }

    addGiftCardRedemption(){
        this.props.history.push('/add-giftCardRedemption/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">GiftCardRedemption List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addGiftCardRedemption}> Add GiftCardRedemption</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> RedeemedAt </th>
                                    <th> Amount </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.giftCardRedemptions.map(
                                        giftCardRedemption => 
                                        <tr key = {giftCardRedemption.giftCardRedemptionId}>
                                             <td> { giftCardRedemption.redeemedAt } </td>
                                             <td> { giftCardRedemption.amount } </td>
                                             <td>
                                                 <button onClick={ () => this.editGiftCardRedemption(giftCardRedemption.giftCardRedemptionId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteGiftCardRedemption(giftCardRedemption.giftCardRedemptionId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewGiftCardRedemption(giftCardRedemption.giftCardRedemptionId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListGiftCardRedemptionComponent
