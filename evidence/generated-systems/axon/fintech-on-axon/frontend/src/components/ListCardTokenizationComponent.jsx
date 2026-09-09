import React, { Component } from 'react'
import CardTokenizationService from '../services/CardTokenizationService'

class ListCardTokenizationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                cardTokenizations: []
        }
        this.addCardTokenization = this.addCardTokenization.bind(this);
        this.editCardTokenization = this.editCardTokenization.bind(this);
        this.deleteCardTokenization = this.deleteCardTokenization.bind(this);
    }

    deleteCardTokenization(id){
        CardTokenizationService.deleteCardTokenization(id).then( res => {
            this.setState({cardTokenizations: this.state.cardTokenizations.filter(cardTokenization => cardTokenization.cardTokenizationId !== id)});
        });
    }
    viewCardTokenization(id){
        this.props.history.push(`/view-cardTokenization/${id}`);
    }
    editCardTokenization(id){
        this.props.history.push(`/add-cardTokenization/${id}`);
    }

    componentDidMount(){
        CardTokenizationService.getCardTokenizations().then((res) => {
            this.setState({ cardTokenizations: res.data});
        });
    }

    addCardTokenization(){
        this.props.history.push('/add-cardTokenization/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">CardTokenization List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCardTokenization}> Add CardTokenization</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> TokenReference </th>
                                    <th> CreatedAt </th>
                                    <th> WalletProvider </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.cardTokenizations.map(
                                        cardTokenization => 
                                        <tr key = {cardTokenization.cardTokenizationId}>
                                             <td> { cardTokenization.tokenReference } </td>
                                             <td> { cardTokenization.createdAt } </td>
                                             <td> { cardTokenization.walletProvider } </td>
                                             <td> { cardTokenization.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editCardTokenization(cardTokenization.cardTokenizationId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCardTokenization(cardTokenization.cardTokenizationId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCardTokenization(cardTokenization.cardTokenizationId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListCardTokenizationComponent
