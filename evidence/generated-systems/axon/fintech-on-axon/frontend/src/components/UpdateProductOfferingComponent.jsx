import React, { Component } from 'react'
import ProductOfferingService from '../services/ProductOfferingService';

class UpdateProductOfferingComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                productCode: '',
                category: ''
        }
        this.updateProductOffering = this.updateProductOffering.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeproductCodeHandler = this.changeproductCodeHandler.bind(this);
        this.changeCategoryHandler = this.changeCategoryHandler.bind(this);
    }

    componentDidMount(){
        ProductOfferingService.getProductOfferingById(this.state.id).then( (res) =>{
            let productOffering = res.data;
            this.setState({
                name: productOffering.name,
                productCode: productOffering.productCode,
                category: productOffering.category
            });
        });
    }

    updateProductOffering = (e) => {
        e.preventDefault();
        let productOffering = {
            productOfferingId: this.state.id,
            name: this.state.name,
            productCode: this.state.productCode,
            category: this.state.category
        };
        console.log('productOffering => ' + JSON.stringify(productOffering));
        console.log('id => ' + JSON.stringify(this.state.id));
        ProductOfferingService.updateProductOffering(productOffering).then( res => {
            this.props.history.push('/productOfferings');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeproductCodeHandler= (event) => {
        this.setState({productCode: event.target.value});
    }
    changeCategoryHandler= (event) => {
        this.setState({category: event.target.value});
    }

    cancel(){
        this.props.history.push('/productOfferings');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update ProductOffering</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> productCode: </label>
                                                <input placeholder="productCode" name="productCode" className="form-control" value={this.state.productCode} onChange={this.changeproductCodeHandler}/>

                                            <label> Category: </label>
                                                <select value={this.state.category} onChange={this.changeCategoryHandler}>
                      <option name="Category" className="form-control" >
                          Checking
                      </option>
                      <option name="Category" className="form-control" >
                          Savings
                      </option>
                      <option name="Category" className="form-control" >
                          CreditCard
                      </option>
                      <option name="Category" className="form-control" >
                          Loan
                      </option>
                      <option name="Category" className="form-control" >
                          Investment
                      </option>
                      <option name="Category" className="form-control" >
                          Insurance
                      </option>
                      <option name="Category" className="form-control" >
                          Payments
                      </option>
                      <option name="Category" className="form-control" >
                          FX
                      </option>
                      <option name="Category" className="form-control" >
                          Wallet
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateProductOffering}>Save</button>
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

export default UpdateProductOfferingComponent
