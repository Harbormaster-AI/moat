import React, { Component } from 'react'
import PriceBookService from '../services/PriceBookService'

class ListPriceBookComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                priceBooks: []
        }
        this.addPriceBook = this.addPriceBook.bind(this);
        this.editPriceBook = this.editPriceBook.bind(this);
        this.deletePriceBook = this.deletePriceBook.bind(this);
    }

    deletePriceBook(id){
        PriceBookService.deletePriceBook(id).then( res => {
            this.setState({priceBooks: this.state.priceBooks.filter(priceBook => priceBook.priceBookId !== id)});
        });
    }
    viewPriceBook(id){
        this.props.history.push(`/view-priceBook/${id}`);
    }
    editPriceBook(id){
        this.props.history.push(`/add-priceBook/${id}`);
    }

    componentDidMount(){
        PriceBookService.getPriceBooks().then((res) => {
            this.setState({ priceBooks: res.data});
        });
    }

    addPriceBook(){
        this.props.history.push('/add-priceBook/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">PriceBook List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addPriceBook}> Add PriceBook</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> AsActive </th>
                                    <th> Description </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.priceBooks.map(
                                        priceBook => 
                                        <tr key = {priceBook.priceBookId}>
                                             <td> { priceBook.name } </td>
                                             <td> { priceBook.asActive } </td>
                                             <td> { priceBook.description } </td>
                                             <td>
                                                 <button onClick={ () => this.editPriceBook(priceBook.priceBookId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deletePriceBook(priceBook.priceBookId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewPriceBook(priceBook.priceBookId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListPriceBookComponent
