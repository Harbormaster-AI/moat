import React, { Component } from 'react'
import SecurityService from '../services/SecurityService';

class CreateSecurityComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                symbol: '',
                isin: '',
                cusip: '',
                currency: '',
                securityType: ''
        }
        this.changesymbolHandler = this.changesymbolHandler.bind(this);
        this.changeisinHandler = this.changeisinHandler.bind(this);
        this.changecusipHandler = this.changecusipHandler.bind(this);
        this.changecurrencyHandler = this.changecurrencyHandler.bind(this);
        this.changeSecurityTypeHandler = this.changeSecurityTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            SecurityService.getSecurityById(this.state.id).then( (res) =>{
                let security = res.data;
                this.setState({
                    symbol: security.symbol,
                    isin: security.isin,
                    cusip: security.cusip,
                    currency: security.currency,
                    securityType: security.securityType
                });
            });
        }        
    }
    saveOrUpdateSecurity = (e) => {
        e.preventDefault();
        let security = {
                securityId: this.state.id,
                symbol: this.state.symbol,
                isin: this.state.isin,
                cusip: this.state.cusip,
                currency: this.state.currency,
                securityType: this.state.securityType
            };
        console.log('security => ' + JSON.stringify(security));

        // step 5
        if(this.state.id === '_add'){
            security.securityId=''
            SecurityService.createSecurity(security).then(res =>{
                this.props.history.push('/securitys');
            });
        }else{
            SecurityService.updateSecurity(security).then( res => {
                this.props.history.push('/securitys');
            });
        }
    }
    
    changesymbolHandler= (event) => {
        this.setState({symbol: event.target.value});
    }
    changeisinHandler= (event) => {
        this.setState({isin: event.target.value});
    }
    changecusipHandler= (event) => {
        this.setState({cusip: event.target.value});
    }
    changecurrencyHandler= (event) => {
        this.setState({currency: event.target.value});
    }
    changeSecurityTypeHandler= (event) => {
        this.setState({securityType: event.target.value});
    }

    cancel(){
        this.props.history.push('/securitys');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Security</h3>
        }else{
            return <h3 className="text-center">Update Security</h3>
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
                                            <label> symbol:&emsp; </label>
                                                <input placeholder="symbol" name="symbol" className="form-control" value={this.state.symbol} onChange={this.changesymbolHandler}/>

                                            <label> isin:&emsp; </label>
                                                <input placeholder="isin" name="isin" className="form-control" value={this.state.isin} onChange={this.changeisinHandler}/>

                                            <label> cusip:&emsp; </label>
                                                <input placeholder="cusip" name="cusip" className="form-control" value={this.state.cusip} onChange={this.changecusipHandler}/>

                                            <label> currency:&emsp; </label>
                                                <input placeholder="currency" name="currency" className="form-control" value={this.state.currency} onChange={this.changecurrencyHandler}/>

                                            <label> SecurityType:&emsp; </label>
                                                <select value={this.state.securityType} onChange={this.changeSecurityTypeHandler}>
                      <option name="SecurityType" className="form-control" >
                          Equity
                      </option>
                      <option name="SecurityType" className="form-control" >
                          Bond
                      </option>
                      <option name="SecurityType" className="form-control" >
                          ETF
                      </option>
                      <option name="SecurityType" className="form-control" >
                          MutualFund
                      </option>
                      <option name="SecurityType" className="form-control" >
                          Derivative
                      </option>
                      <option name="SecurityType" className="form-control" >
                          Crypto
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateSecurity}>Save</button>
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

export default CreateSecurityComponent
