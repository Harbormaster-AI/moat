import React, { Component } from 'react'
import PriceBookService from '../services/PriceBookService';

class CreatePriceBookComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                asActive: '',
                description: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeasActiveHandler = this.changeasActiveHandler.bind(this);
        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            PriceBookService.getPriceBookById(this.state.id).then( (res) =>{
                let priceBook = res.data;
                this.setState({
                    name: priceBook.name,
                    asActive: priceBook.asActive,
                    description: priceBook.description
                });
            });
        }        
    }
    saveOrUpdatePriceBook = (e) => {
        e.preventDefault();
        let priceBook = {
                priceBookId: this.state.id,
                name: this.state.name,
                asActive: this.state.asActive,
                description: this.state.description
            };
        console.log('priceBook => ' + JSON.stringify(priceBook));

        // step 5
        if(this.state.id === '_add'){
            priceBook.priceBookId=''
            PriceBookService.createPriceBook(priceBook).then(res =>{
                this.props.history.push('/priceBooks');
            });
        }else{
            PriceBookService.updatePriceBook(priceBook).then( res => {
                this.props.history.push('/priceBooks');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add PriceBook</h3>
        }else{
            return <h3 className="text-center">Update PriceBook</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> asActive:&emsp; </label>
                                                <input type="checkbox" placeholder="asActive" name="asActive" className="form-control" value={this.state.asActive} onChange={this.changeasActiveHandler}/>


                                            <label> description:&emsp; </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdatePriceBook}>Save</button>
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

export default CreatePriceBookComponent
