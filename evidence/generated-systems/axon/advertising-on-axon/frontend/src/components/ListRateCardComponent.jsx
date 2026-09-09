import React, { Component } from 'react'
import RateCardService from '../services/RateCardService'

class ListRateCardComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                rateCards: []
        }
        this.addRateCard = this.addRateCard.bind(this);
        this.editRateCard = this.editRateCard.bind(this);
        this.deleteRateCard = this.deleteRateCard.bind(this);
    }

    deleteRateCard(id){
        RateCardService.deleteRateCard(id).then( res => {
            this.setState({rateCards: this.state.rateCards.filter(rateCard => rateCard.rateCardId !== id)});
        });
    }
    viewRateCard(id){
        this.props.history.push(`/view-rateCard/${id}`);
    }
    editRateCard(id){
        this.props.history.push(`/add-rateCard/${id}`);
    }

    componentDidMount(){
        RateCardService.getRateCards().then((res) => {
            this.setState({ rateCards: res.data});
        });
    }

    addRateCard(){
        this.props.history.push('/add-rateCard/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">RateCard List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addRateCard}> Add RateCard</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> EffectiveDate </th>
                                    <th> Currency </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.rateCards.map(
                                        rateCard => 
                                        <tr key = {rateCard.rateCardId}>
                                             <td> { rateCard.name } </td>
                                             <td> { rateCard.effectiveDate } </td>
                                             <td> { rateCard.currency } </td>
                                             <td>
                                                 <button onClick={ () => this.editRateCard(rateCard.rateCardId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteRateCard(rateCard.rateCardId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewRateCard(rateCard.rateCardId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListRateCardComponent
