import React, { Component } from 'react'
import PaymentProviderService from '../services/PaymentProviderService';

class UpdatePaymentProviderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                enabled: '',
                merchantAccountId: '',
                providerType: ''
        }
        this.updatePaymentProvider = this.updatePaymentProvider.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeenabledHandler = this.changeenabledHandler.bind(this);
        this.changemerchantAccountIdHandler = this.changemerchantAccountIdHandler.bind(this);
        this.changeProviderTypeHandler = this.changeProviderTypeHandler.bind(this);
    }

    componentDidMount(){
        PaymentProviderService.getPaymentProviderById(this.state.id).then( (res) =>{
            let paymentProvider = res.data;
            this.setState({
                name: paymentProvider.name,
                enabled: paymentProvider.enabled,
                merchantAccountId: paymentProvider.merchantAccountId,
                providerType: paymentProvider.providerType
            });
        });
    }

    updatePaymentProvider = (e) => {
        e.preventDefault();
        let paymentProvider = {
            paymentProviderId: this.state.id,
            name: this.state.name,
            enabled: this.state.enabled,
            merchantAccountId: this.state.merchantAccountId,
            providerType: this.state.providerType
        };
        console.log('paymentProvider => ' + JSON.stringify(paymentProvider));
        console.log('id => ' + JSON.stringify(this.state.id));
        PaymentProviderService.updatePaymentProvider(paymentProvider).then( res => {
            this.props.history.push('/paymentProviders');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeenabledHandler= (event) => {
        this.setState({enabled: event.target.value});
    }
    changemerchantAccountIdHandler= (event) => {
        this.setState({merchantAccountId: event.target.value});
    }
    changeProviderTypeHandler= (event) => {
        this.setState({providerType: event.target.value});
    }

    cancel(){
        this.props.history.push('/paymentProviders');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update PaymentProvider</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> enabled: </label>
                                                <input type="checkbox" placeholder="enabled" name="enabled" className="form-control" value={this.state.enabled} onChange={this.changeenabledHandler}/>


                                            <label> merchantAccountId: </label>
                                                <input placeholder="merchantAccountId" name="merchantAccountId" className="form-control" value={this.state.merchantAccountId} onChange={this.changemerchantAccountIdHandler}/>

                                            <label> ProviderType: </label>
                                                <select value={this.state.providerType} onChange={this.changeProviderTypeHandler}>
                      <option name="ProviderType" className="form-control" >
                          PSP
                      </option>
                      <option name="ProviderType" className="form-control" >
                          Gateway
                      </option>
                      <option name="ProviderType" className="form-control" >
                          Aggregator
                      </option>
                      <option name="ProviderType" className="form-control" >
                          Manual
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updatePaymentProvider}>Save</button>
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

export default UpdatePaymentProviderComponent
