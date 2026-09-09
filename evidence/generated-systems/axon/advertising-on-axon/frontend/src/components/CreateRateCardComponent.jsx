import React, { Component } from 'react'
import RateCardService from '../services/RateCardService';

class CreateRateCardComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                effectiveDate: '',
                currency: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeeffectiveDateHandler = this.changeeffectiveDateHandler.bind(this);
        this.changecurrencyHandler = this.changecurrencyHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            RateCardService.getRateCardById(this.state.id).then( (res) =>{
                let rateCard = res.data;
                this.setState({
                    name: rateCard.name,
                    effectiveDate: rateCard.effectiveDate,
                    currency: rateCard.currency
                });
            });
        }        
    }
    saveOrUpdateRateCard = (e) => {
        e.preventDefault();
        let rateCard = {
                rateCardId: this.state.id,
                name: this.state.name,
                effectiveDate: this.state.effectiveDate,
                currency: this.state.currency
            };
        console.log('rateCard => ' + JSON.stringify(rateCard));

        // step 5
        if(this.state.id === '_add'){
            rateCard.rateCardId=''
            RateCardService.createRateCard(rateCard).then(res =>{
                this.props.history.push('/rateCards');
            });
        }else{
            RateCardService.updateRateCard(rateCard).then( res => {
                this.props.history.push('/rateCards');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeeffectiveDateHandler= (event) => {
        this.setState({effectiveDate: event.target.value});
    }
    changecurrencyHandler= (event) => {
        this.setState({currency: event.target.value});
    }

    cancel(){
        this.props.history.push('/rateCards');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add RateCard</h3>
        }else{
            return <h3 className="text-center">Update RateCard</h3>
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

                                            <label> effectiveDate:&emsp; </label>
                                                <input type="date" placeholder="effectiveDate" name="effectiveDate" className="form-control" value={this.state.effectiveDate} onChange={this.changeeffectiveDateHandler}/>

                                            <label> currency:&emsp; </label>
                                                <input placeholder="currency" name="currency" className="form-control" value={this.state.currency} onChange={this.changecurrencyHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateRateCard}>Save</button>
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

export default CreateRateCardComponent
