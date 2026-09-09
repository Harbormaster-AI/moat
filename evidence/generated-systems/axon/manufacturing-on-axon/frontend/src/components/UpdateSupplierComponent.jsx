import React, { Component } from 'react'
import SupplierService from '../services/SupplierService';

class UpdateSupplierComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                supplierCode: '',
                address: '',
                supplierTier: '',
                paymentTerms: ''
        }
        this.updateSupplier = this.updateSupplier.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changesupplierCodeHandler = this.changesupplierCodeHandler.bind(this);
        this.changeaddressHandler = this.changeaddressHandler.bind(this);
        this.changeSupplierTierHandler = this.changeSupplierTierHandler.bind(this);
        this.changePaymentTermsHandler = this.changePaymentTermsHandler.bind(this);
    }

    componentDidMount(){
        SupplierService.getSupplierById(this.state.id).then( (res) =>{
            let supplier = res.data;
            this.setState({
                name: supplier.name,
                supplierCode: supplier.supplierCode,
                address: supplier.address,
                supplierTier: supplier.supplierTier,
                paymentTerms: supplier.paymentTerms
            });
        });
    }

    updateSupplier = (e) => {
        e.preventDefault();
        let supplier = {
            supplierId: this.state.id,
            name: this.state.name,
            supplierCode: this.state.supplierCode,
            address: this.state.address,
            supplierTier: this.state.supplierTier,
            paymentTerms: this.state.paymentTerms
        };
        console.log('supplier => ' + JSON.stringify(supplier));
        console.log('id => ' + JSON.stringify(this.state.id));
        SupplierService.updateSupplier(supplier).then( res => {
            this.props.history.push('/suppliers');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changesupplierCodeHandler= (event) => {
        this.setState({supplierCode: event.target.value});
    }
    changeaddressHandler= (event) => {
        this.setState({address: event.target.value});
    }
    changeSupplierTierHandler= (event) => {
        this.setState({supplierTier: event.target.value});
    }
    changePaymentTermsHandler= (event) => {
        this.setState({paymentTerms: event.target.value});
    }

    cancel(){
        this.props.history.push('/suppliers');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Supplier</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> supplierCode: </label>
                                                <input placeholder="supplierCode" name="supplierCode" className="form-control" value={this.state.supplierCode} onChange={this.changesupplierCodeHandler}/>

                                            <label> address: </label>
                                                <input placeholder="address" name="address" className="form-control" value={this.state.address} onChange={this.changeaddressHandler}/>

                                            <label> SupplierTier: </label>
                                                <select value={this.state.supplierTier} onChange={this.changeSupplierTierHandler}>
                      <option name="SupplierTier" className="form-control" >
                          Tier1
                      </option>
                      <option name="SupplierTier" className="form-control" >
                          Tier2
                      </option>
                      <option name="SupplierTier" className="form-control" >
                          Tier3
                      </option>
                    </select>

                                            <label> PaymentTerms: </label>
                                                <select value={this.state.paymentTerms} onChange={this.changePaymentTermsHandler}>
                      <option name="PaymentTerms" className="form-control" >
                          Net30
                      </option>
                      <option name="PaymentTerms" className="form-control" >
                          Net45
                      </option>
                      <option name="PaymentTerms" className="form-control" >
                          Net60
                      </option>
                      <option name="PaymentTerms" className="form-control" >
                          Prepaid
                      </option>
                      <option name="PaymentTerms" className="form-control" >
                          COD
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateSupplier}>Save</button>
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

export default UpdateSupplierComponent
