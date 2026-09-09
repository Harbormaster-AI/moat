import React, { Component } from 'react'
import ProductOfferingService from '../services/ProductOfferingService';

class CreateProductOfferingComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                productCode: '',
                category: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeproductCodeHandler = this.changeproductCodeHandler.bind(this);
        this.changeCategoryHandler = this.changeCategoryHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ProductOfferingService.getProductOfferingById(this.state.id).then( (res) =>{
                let productOffering = res.data;
                this.setState({
                    name: productOffering.name,
                    productCode: productOffering.productCode,
                    category: productOffering.category
                });
            });
        }        
    }
    saveOrUpdateProductOffering = (e) => {
        e.preventDefault();
        let productOffering = {
                productOfferingId: this.state.id,
                name: this.state.name,
                productCode: this.state.productCode,
                category: this.state.category
            };
        console.log('productOffering => ' + JSON.stringify(productOffering));

        // step 5
        if(this.state.id === '_add'){
            productOffering.productOfferingId=''
            ProductOfferingService.createProductOffering(productOffering).then(res =>{
                this.props.history.push('/productOfferings');
            });
        }else{
            ProductOfferingService.updateProductOffering(productOffering).then( res => {
                this.props.history.push('/productOfferings');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add ProductOffering</h3>
        }else{
            return <h3 className="text-center">Update ProductOffering</h3>
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

                                            <label> productCode:&emsp; </label>
                                                <input placeholder="productCode" name="productCode" className="form-control" value={this.state.productCode} onChange={this.changeproductCodeHandler}/>

                                            <label> Category:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateProductOffering}>Save</button>
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

export default CreateProductOfferingComponent
