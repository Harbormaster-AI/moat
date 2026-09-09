import React, { Component } from 'react'
import TaxWithholdingService from '../services/TaxWithholdingService';

class CreateTaxWithholdingComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                taxId: '',
                allowances: '',
                additionalAmount: '',
                filingStatus: ''
        }
        this.changetaxIdHandler = this.changetaxIdHandler.bind(this);
        this.changeallowancesHandler = this.changeallowancesHandler.bind(this);
        this.changeadditionalAmountHandler = this.changeadditionalAmountHandler.bind(this);
        this.changeFilingStatusHandler = this.changeFilingStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            TaxWithholdingService.getTaxWithholdingById(this.state.id).then( (res) =>{
                let taxWithholding = res.data;
                this.setState({
                    taxId: taxWithholding.taxId,
                    allowances: taxWithholding.allowances,
                    additionalAmount: taxWithholding.additionalAmount,
                    filingStatus: taxWithholding.filingStatus
                });
            });
        }        
    }
    saveOrUpdateTaxWithholding = (e) => {
        e.preventDefault();
        let taxWithholding = {
                taxWithholdingId: this.state.id,
                taxId: this.state.taxId,
                allowances: this.state.allowances,
                additionalAmount: this.state.additionalAmount,
                filingStatus: this.state.filingStatus
            };
        console.log('taxWithholding => ' + JSON.stringify(taxWithholding));

        // step 5
        if(this.state.id === '_add'){
            taxWithholding.taxWithholdingId=''
            TaxWithholdingService.createTaxWithholding(taxWithholding).then(res =>{
                this.props.history.push('/taxWithholdings');
            });
        }else{
            TaxWithholdingService.updateTaxWithholding(taxWithholding).then( res => {
                this.props.history.push('/taxWithholdings');
            });
        }
    }
    
    changetaxIdHandler= (event) => {
        this.setState({taxId: event.target.value});
    }
    changeallowancesHandler= (event) => {
        this.setState({allowances: event.target.value});
    }
    changeadditionalAmountHandler= (event) => {
        this.setState({additionalAmount: event.target.value});
    }
    changeFilingStatusHandler= (event) => {
        this.setState({filingStatus: event.target.value});
    }

    cancel(){
        this.props.history.push('/taxWithholdings');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add TaxWithholding</h3>
        }else{
            return <h3 className="text-center">Update TaxWithholding</h3>
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
                                            <label> taxId:&emsp; </label>
                                                <input placeholder="taxId" name="taxId" className="form-control" value={this.state.taxId} onChange={this.changetaxIdHandler}/>

                                            <label> allowances:&emsp; </label>
                                                <input type="number" placeholder="allowances" name="allowances" className="form-control" value={this.state.allowances} onChange={this.changeallowancesHandler}/>

                                            <label> additionalAmount:&emsp; </label>
                                                <input placeholder="additionalAmount" name="additionalAmount" className="form-control" value={this.state.additionalAmount} onChange={this.changeadditionalAmountHandler}/>

                                            <label> FilingStatus:&emsp; </label>
                                                <select value={this.state.filingStatus} onChange={this.changeFilingStatusHandler}>
                      <option name="FilingStatus" className="form-control" >
                          Single
                      </option>
                      <option name="FilingStatus" className="form-control" >
                          MarriedFilingJointly
                      </option>
                      <option name="FilingStatus" className="form-control" >
                          MarriedFilingSeparately
                      </option>
                      <option name="FilingStatus" className="form-control" >
                          HeadOfHousehold
                      </option>
                      <option name="FilingStatus" className="form-control" >
                          QualifyingWidowEr
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateTaxWithholding}>Save</button>
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

export default CreateTaxWithholdingComponent
