import React, { Component } from 'react'
import PromotionService from '../services/PromotionService'

class ListPromotionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                promotions: []
        }
        this.addPromotion = this.addPromotion.bind(this);
        this.editPromotion = this.editPromotion.bind(this);
        this.deletePromotion = this.deletePromotion.bind(this);
    }

    deletePromotion(id){
        PromotionService.deletePromotion(id).then( res => {
            this.setState({promotions: this.state.promotions.filter(promotion => promotion.promotionId !== id)});
        });
    }
    viewPromotion(id){
        this.props.history.push(`/view-promotion/${id}`);
    }
    editPromotion(id){
        this.props.history.push(`/add-promotion/${id}`);
    }

    componentDidMount(){
        PromotionService.getPromotions().then((res) => {
            this.setState({ promotions: res.data});
        });
    }

    addPromotion(){
        this.props.history.push('/add-promotion/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Promotion List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addPromotion}> Add Promotion</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Code </th>
                                    <th> Value </th>
                                    <th> StartDate </th>
                                    <th> EndDate </th>
                                    <th> AsStackable </th>
                                    <th> MaxRedemptions </th>
                                    <th> PromotionType </th>
                                    <th> DiscountType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.promotions.map(
                                        promotion => 
                                        <tr key = {promotion.promotionId}>
                                             <td> { promotion.name } </td>
                                             <td> { promotion.code } </td>
                                             <td> { promotion.value } </td>
                                             <td> { promotion.startDate } </td>
                                             <td> { promotion.endDate } </td>
                                             <td> { promotion.asStackable } </td>
                                             <td> { promotion.maxRedemptions } </td>
                                             <td> { promotion.promotionType } </td>
                                             <td> { promotion.discountType } </td>
                                             <td>
                                                 <button onClick={ () => this.editPromotion(promotion.promotionId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deletePromotion(promotion.promotionId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewPromotion(promotion.promotionId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListPromotionComponent
