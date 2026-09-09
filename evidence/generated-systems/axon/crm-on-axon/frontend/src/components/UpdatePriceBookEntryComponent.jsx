import React, { Component } from 'react'
import PriceBookEntryService from '../services/PriceBookEntryService';

class UpdatePriceBookEntryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                unitPrice: '',
                effectiveDate: '',
                expirationDate: '',
                asActive: ''
        }
        this.updatePriceBookEntry = this.updatePriceBookEntry.bind(this);

        this.changeunitPriceHandler = this.changeunitPriceHandler.bind(this);
        this.changeeffectiveDateHandler = this.changeeffectiveDateHandler.bind(this);
        this.changeexpirationDateHandler = this.changeexpirationDateHandler.bind(this);
        this.changeasActiveHandler = this.changeasActiveHandler.bind(this);
    }

    componentDidMount(){
        PriceBookEntryService.getPriceBookEntryById(this.state.id).then( (res) =>{
            let priceBookEntry = res.data;
            this.setState({
                unitPrice: priceBookEntry.unitPrice,
                effectiveDate: priceBookEntry.effectiveDate,
                expirationDate: priceBookEntry.expirationDate,
                asActive: priceBookEntry.asActive
            });
        });
    }

    updatePriceBookEntry = (e) => {
        e.preventDefault();
        let priceBookEntry = {
            priceBookEntryId: this.state.id,
            unitPrice: this.state.unitPrice,
            effectiveDate: this.state.effectiveDate,
            expirationDate: this.state.expirationDate,
            asActive: this.state.asActive
        };
        console.log('priceBookEntry => ' + JSON.stringify(priceBookEntry));
        console.log('id => ' + JSON.stringify(this.state.id));
        PriceBookEntryService.updatePriceBookEntry(priceBookEntry).then( res => {
            this.props.history.push('/priceBookEntrys');
        });
    }

    changeunitPriceHandler= (event) => {
        this.setState({unitPrice: event.target.value});
    }
    changeeffectiveDateHandler= (event) => {
        this.setState({effectiveDate: event.target.value});
    }
    changeexpirationDateHandler= (event) => {
        this.setState({expirationDate: event.target.value});
    }
    changeasActiveHandler= (event) => {
        this.setState({asActive: event.target.value});
    }

    cancel(){
        this.props.history.push('/priceBookEntrys');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update PriceBookEntry</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> unitPrice: </label>
                                                <input placeholder="unitPrice" name="unitPrice" className="form-control" value={this.state.unitPrice} onChange={this.changeunitPriceHandler}/>

                                            <label> effectiveDate: </label>
                                                <input type="date" placeholder="effectiveDate" name="effectiveDate" className="form-control" value={this.state.effectiveDate} onChange={this.changeeffectiveDateHandler}/>

                                            <label> expirationDate: </label>
                                                <input type="date" placeholder="expirationDate" name="expirationDate" className="form-control" value={this.state.expirationDate} onChange={this.changeexpirationDateHandler}/>

                                            <label> asActive: </label>
                                                <input type="checkbox" placeholder="asActive" name="asActive" className="form-control" value={this.state.asActive} onChange={this.changeasActiveHandler}/>


                                        </div>
                                        <button className="btn btn-success" onClick={this.updatePriceBookEntry}>Save</button>
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

export default UpdatePriceBookEntryComponent
