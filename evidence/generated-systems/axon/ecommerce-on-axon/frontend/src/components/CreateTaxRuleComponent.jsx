import React, { Component } from 'react'
import TaxRuleService from '../services/TaxRuleService';

class CreateTaxRuleComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                country: '',
                region: '',
                rate: '',
                taxInclusive: '',
                taxClass: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changecountryHandler = this.changecountryHandler.bind(this);
        this.changeregionHandler = this.changeregionHandler.bind(this);
        this.changerateHandler = this.changerateHandler.bind(this);
        this.changetaxInclusiveHandler = this.changetaxInclusiveHandler.bind(this);
        this.changeTaxClassHandler = this.changeTaxClassHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            TaxRuleService.getTaxRuleById(this.state.id).then( (res) =>{
                let taxRule = res.data;
                this.setState({
                    name: taxRule.name,
                    country: taxRule.country,
                    region: taxRule.region,
                    rate: taxRule.rate,
                    taxInclusive: taxRule.taxInclusive,
                    taxClass: taxRule.taxClass
                });
            });
        }        
    }
    saveOrUpdateTaxRule = (e) => {
        e.preventDefault();
        let taxRule = {
                taxRuleId: this.state.id,
                name: this.state.name,
                country: this.state.country,
                region: this.state.region,
                rate: this.state.rate,
                taxInclusive: this.state.taxInclusive,
                taxClass: this.state.taxClass
            };
        console.log('taxRule => ' + JSON.stringify(taxRule));

        // step 5
        if(this.state.id === '_add'){
            taxRule.taxRuleId=''
            TaxRuleService.createTaxRule(taxRule).then(res =>{
                this.props.history.push('/taxRules');
            });
        }else{
            TaxRuleService.updateTaxRule(taxRule).then( res => {
                this.props.history.push('/taxRules');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changecountryHandler= (event) => {
        this.setState({country: event.target.value});
    }
    changeregionHandler= (event) => {
        this.setState({region: event.target.value});
    }
    changerateHandler= (event) => {
        this.setState({rate: event.target.value});
    }
    changetaxInclusiveHandler= (event) => {
        this.setState({taxInclusive: event.target.value});
    }
    changeTaxClassHandler= (event) => {
        this.setState({taxClass: event.target.value});
    }

    cancel(){
        this.props.history.push('/taxRules');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add TaxRule</h3>
        }else{
            return <h3 className="text-center">Update TaxRule</h3>
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

                                            <label> country:&emsp; </label>
                                                <input placeholder="country" name="country" className="form-control" value={this.state.country} onChange={this.changecountryHandler}/>

                                            <label> region:&emsp; </label>
                                                <input placeholder="region" name="region" className="form-control" value={this.state.region} onChange={this.changeregionHandler}/>

                                            <label> rate:&emsp; </label>
                                                <input placeholder="rate" name="rate" className="form-control" value={this.state.rate} onChange={this.changerateHandler}/>

                                            <label> taxInclusive:&emsp; </label>
                                                <input type="checkbox" placeholder="taxInclusive" name="taxInclusive" className="form-control" value={this.state.taxInclusive} onChange={this.changetaxInclusiveHandler}/>


                                            <label> TaxClass:&emsp; </label>
                                                <select value={this.state.taxClass} onChange={this.changeTaxClassHandler}>
                      <option name="TaxClass" className="form-control" >
                          Standard
                      </option>
                      <option name="TaxClass" className="form-control" >
                          Reduced
                      </option>
                      <option name="TaxClass" className="form-control" >
                          Zero
                      </option>
                      <option name="TaxClass" className="form-control" >
                          Exempt
                      </option>
                      <option name="TaxClass" className="form-control" >
                          DigitalServices
                      </option>
                      <option name="TaxClass" className="form-control" >
                          Food
                      </option>
                      <option name="TaxClass" className="form-control" >
                          Clothing
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateTaxRule}>Save</button>
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

export default CreateTaxRuleComponent
