import React, { Component } from 'react'
import PriceBookEntryService from '../services/PriceBookEntryService';

class CreatePriceBookEntryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                unitPrice: '',
                effectiveDate: '',
                expirationDate: '',
                asActive: ''
        }
        this.changeunitPriceHandler = this.changeunitPriceHandler.bind(this);
        this.changeeffectiveDateHandler = this.changeeffectiveDateHandler.bind(this);
        this.changeexpirationDateHandler = this.changeexpirationDateHandler.bind(this);
        this.changeasActiveHandler = this.changeasActiveHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdatePriceBookEntry = (e) => {
        e.preventDefault();
        let priceBookEntry = {
                priceBookEntryId: this.state.id,
                unitPrice: this.state.unitPrice,
                effectiveDate: this.state.effectiveDate,
                expirationDate: this.state.expirationDate,
                asActive: this.state.asActive
            };
        console.log('priceBookEntry => ' + JSON.stringify(priceBookEntry));

        // step 5
        if(this.state.id === '_add'){
            priceBookEntry.priceBookEntryId=''
            PriceBookEntryService.createPriceBookEntry(priceBookEntry).then(res =>{
                this.props.history.push('/priceBookEntrys');
            });
        }else{
            PriceBookEntryService.updatePriceBookEntry(priceBookEntry).then( res => {
                this.props.history.push('/priceBookEntrys');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add PriceBookEntry</h3>
        }else{
            return <h3 className="text-center">Update PriceBookEntry</h3>
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
                                            <label> unitPrice:&emsp; </label>
                                                <input placeholder="unitPrice" name="unitPrice" className="form-control" value={this.state.unitPrice} onChange={this.changeunitPriceHandler}/>

                                            <label> effectiveDate:&emsp; </label>
                                                <input type="date" placeholder="effectiveDate" name="effectiveDate" className="form-control" value={this.state.effectiveDate} onChange={this.changeeffectiveDateHandler}/>

                                            <label> expirationDate:&emsp; </label>
                                                <input type="date" placeholder="expirationDate" name="expirationDate" className="form-control" value={this.state.expirationDate} onChange={this.changeexpirationDateHandler}/>

                                            <label> asActive:&emsp; </label>
                                                <input type="checkbox" placeholder="asActive" name="asActive" className="form-control" value={this.state.asActive} onChange={this.changeasActiveHandler}/>


                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdatePriceBookEntry}>Save</button>
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

export default CreatePriceBookEntryComponent
