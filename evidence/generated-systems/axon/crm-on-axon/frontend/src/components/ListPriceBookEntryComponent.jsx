import React, { Component } from 'react'
import PriceBookEntryService from '../services/PriceBookEntryService'

class ListPriceBookEntryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                priceBookEntrys: []
        }
        this.addPriceBookEntry = this.addPriceBookEntry.bind(this);
        this.editPriceBookEntry = this.editPriceBookEntry.bind(this);
        this.deletePriceBookEntry = this.deletePriceBookEntry.bind(this);
    }

    deletePriceBookEntry(id){
        PriceBookEntryService.deletePriceBookEntry(id).then( res => {
            this.setState({priceBookEntrys: this.state.priceBookEntrys.filter(priceBookEntry => priceBookEntry.priceBookEntryId !== id)});
        });
    }
    viewPriceBookEntry(id){
        this.props.history.push(`/view-priceBookEntry/${id}`);
    }
    editPriceBookEntry(id){
        this.props.history.push(`/add-priceBookEntry/${id}`);
    }

    componentDidMount(){
        PriceBookEntryService.getPriceBookEntrys().then((res) => {
            this.setState({ priceBookEntrys: res.data});
        });
    }

    addPriceBookEntry(){
        this.props.history.push('/add-priceBookEntry/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">PriceBookEntry List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addPriceBookEntry}> Add PriceBookEntry</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> UnitPrice </th>
                                    <th> EffectiveDate </th>
                                    <th> ExpirationDate </th>
                                    <th> AsActive </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.priceBookEntrys.map(
                                        priceBookEntry => 
                                        <tr key = {priceBookEntry.priceBookEntryId}>
                                             <td> { priceBookEntry.unitPrice } </td>
                                             <td> { priceBookEntry.effectiveDate } </td>
                                             <td> { priceBookEntry.expirationDate } </td>
                                             <td> { priceBookEntry.asActive } </td>
                                             <td>
                                                 <button onClick={ () => this.editPriceBookEntry(priceBookEntry.priceBookEntryId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deletePriceBookEntry(priceBookEntry.priceBookEntryId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewPriceBookEntry(priceBookEntry.priceBookEntryId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListPriceBookEntryComponent
