import React, { Component } from 'react'
import RateCardService from '../services/RateCardService';

class UpdateRateCardComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                effectiveDate: '',
                currency: ''
        }
        this.updateRateCard = this.updateRateCard.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeeffectiveDateHandler = this.changeeffectiveDateHandler.bind(this);
        this.changecurrencyHandler = this.changecurrencyHandler.bind(this);
    }

    componentDidMount(){
        RateCardService.getRateCardById(this.state.id).then( (res) =>{
            let rateCard = res.data;
            this.setState({
                name: rateCard.name,
                effectiveDate: rateCard.effectiveDate,
                currency: rateCard.currency
            });
        });
    }

    updateRateCard = (e) => {
        e.preventDefault();
        let rateCard = {
            rateCardId: this.state.id,
            name: this.state.name,
            effectiveDate: this.state.effectiveDate,
            currency: this.state.currency
        };
        console.log('rateCard => ' + JSON.stringify(rateCard));
        console.log('id => ' + JSON.stringify(this.state.id));
        RateCardService.updateRateCard(rateCard).then( res => {
            this.props.history.push('/rateCards');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update RateCard</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> effectiveDate: </label>
                                                <input type="date" placeholder="effectiveDate" name="effectiveDate" className="form-control" value={this.state.effectiveDate} onChange={this.changeeffectiveDateHandler}/>

                                            <label> currency: </label>
                                                <input placeholder="currency" name="currency" className="form-control" value={this.state.currency} onChange={this.changecurrencyHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateRateCard}>Save</button>
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

export default UpdateRateCardComponent
