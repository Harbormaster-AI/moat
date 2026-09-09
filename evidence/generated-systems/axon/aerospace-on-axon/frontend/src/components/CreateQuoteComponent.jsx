import React, { Component } from 'react'
import QuoteService from '../services/QuoteService';

class CreateQuoteComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                quoteNumber: '',
                totalAmount: ''
        }
        this.changequoteNumberHandler = this.changequoteNumberHandler.bind(this);
        this.changetotalAmountHandler = this.changetotalAmountHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            QuoteService.getQuoteById(this.state.id).then( (res) =>{
                let quote = res.data;
                this.setState({
                    quoteNumber: quote.quoteNumber,
                    totalAmount: quote.totalAmount
                });
            });
        }        
    }
    saveOrUpdateQuote = (e) => {
        e.preventDefault();
        let quote = {
                quoteId: this.state.id,
                quoteNumber: this.state.quoteNumber,
                totalAmount: this.state.totalAmount
            };
        console.log('quote => ' + JSON.stringify(quote));

        // step 5
        if(this.state.id === '_add'){
            quote.quoteId=''
            QuoteService.createQuote(quote).then(res =>{
                this.props.history.push('/quotes');
            });
        }else{
            QuoteService.updateQuote(quote).then( res => {
                this.props.history.push('/quotes');
            });
        }
    }
    
    changequoteNumberHandler= (event) => {
        this.setState({quoteNumber: event.target.value});
    }
    changetotalAmountHandler= (event) => {
        this.setState({totalAmount: event.target.value});
    }

    cancel(){
        this.props.history.push('/quotes');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Quote</h3>
        }else{
            return <h3 className="text-center">Update Quote</h3>
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
                                            <label> quoteNumber:&emsp; </label>
                                                <input placeholder="quoteNumber" name="quoteNumber" className="form-control" value={this.state.quoteNumber} onChange={this.changequoteNumberHandler}/>

                                            <label> totalAmount:&emsp; </label>
                                                <input placeholder="totalAmount" name="totalAmount" className="form-control" value={this.state.totalAmount} onChange={this.changetotalAmountHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateQuote}>Save</button>
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

export default CreateQuoteComponent
