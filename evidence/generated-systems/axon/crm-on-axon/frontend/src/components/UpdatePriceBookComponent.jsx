import React, { Component } from 'react'
import PriceBookService from '../services/PriceBookService';

class UpdatePriceBookComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                asActive: '',
                description: ''
        }
        this.updatePriceBook = this.updatePriceBook.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeasActiveHandler = this.changeasActiveHandler.bind(this);
        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
    }

    componentDidMount(){
        PriceBookService.getPriceBookById(this.state.id).then( (res) =>{
            let priceBook = res.data;
            this.setState({
                name: priceBook.name,
                asActive: priceBook.asActive,
                description: priceBook.description
            });
        });
    }

    updatePriceBook = (e) => {
        e.preventDefault();
        let priceBook = {
            priceBookId: this.state.id,
            name: this.state.name,
            asActive: this.state.asActive,
            description: this.state.description
        };
        console.log('priceBook => ' + JSON.stringify(priceBook));
        console.log('id => ' + JSON.stringify(this.state.id));
        PriceBookService.updatePriceBook(priceBook).then( res => {
            this.props.history.push('/priceBooks');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeasActiveHandler= (event) => {
        this.setState({asActive: event.target.value});
    }
    changedescriptionHandler= (event) => {
        this.setState({description: event.target.value});
    }

    cancel(){
        this.props.history.push('/priceBooks');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update PriceBook</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> asActive: </label>
                                                <input type="checkbox" placeholder="asActive" name="asActive" className="form-control" value={this.state.asActive} onChange={this.changeasActiveHandler}/>


                                            <label> description: </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updatePriceBook}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>

                   </div>
            </div>
        )
    }
}

export default UpdatePriceBookComponent
