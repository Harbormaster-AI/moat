import React, { Component } from 'react'
import ProductService from '../services/ProductService';

class CreateProductComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                sku: '',
                name: '',
                asActive: '',
                standardPrice: '',
                description: '',
                productType: '',
                uom: ''
        }
        this.changeskuHandler = this.changeskuHandler.bind(this);
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeasActiveHandler = this.changeasActiveHandler.bind(this);
        this.changestandardPriceHandler = this.changestandardPriceHandler.bind(this);
        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
        this.changeProductTypeHandler = this.changeProductTypeHandler.bind(this);
        this.changeUomHandler = this.changeUomHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ProductService.getProductById(this.state.id).then( (res) =>{
                let product = res.data;
                this.setState({
                    sku: product.sku,
                    name: product.name,
                    asActive: product.asActive,
                    standardPrice: product.standardPrice,
                    description: product.description,
                    productType: product.productType,
                    uom: product.uom
                });
            });
        }        
    }
    saveOrUpdateProduct = (e) => {
        e.preventDefault();
        let product = {
                productId: this.state.id,
                sku: this.state.sku,
                name: this.state.name,
                asActive: this.state.asActive,
                standardPrice: this.state.standardPrice,
                description: this.state.description,
                productType: this.state.productType,
                uom: this.state.uom
            };
        console.log('product => ' + JSON.stringify(product));

        // step 5
        if(this.state.id === '_add'){
            product.productId=''
            ProductService.createProduct(product).then(res =>{
                this.props.history.push('/products');
            });
        }else{
            ProductService.updateProduct(product).then( res => {
                this.props.history.push('/products');
            });
        }
    }
    
    changeskuHandler= (event) => {
        this.setState({sku: event.target.value});
    }
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeasActiveHandler= (event) => {
        this.setState({asActive: event.target.value});
    }
    changestandardPriceHandler= (event) => {
        this.setState({standardPrice: event.target.value});
    }
    changedescriptionHandler= (event) => {
        this.setState({description: event.target.value});
    }
    changeProductTypeHandler= (event) => {
        this.setState({productType: event.target.value});
    }
    changeUomHandler= (event) => {
        this.setState({uom: event.target.value});
    }

    cancel(){
        this.props.history.push('/products');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Product</h3>
        }else{
            return <h3 className="text-center">Update Product</h3>
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
                                            <label> sku:&emsp; </label>
                                                <input placeholder="sku" name="sku" className="form-control" value={this.state.sku} onChange={this.changeskuHandler}/>

                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> asActive:&emsp; </label>
                                                <input type="checkbox" placeholder="asActive" name="asActive" className="form-control" value={this.state.asActive} onChange={this.changeasActiveHandler}/>


                                            <label> standardPrice:&emsp; </label>
                                                <input placeholder="standardPrice" name="standardPrice" className="form-control" value={this.state.standardPrice} onChange={this.changestandardPriceHandler}/>

                                            <label> description:&emsp; </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

                                            <label> ProductType:&emsp; </label>
                                                <select value={this.state.productType} onChange={this.changeProductTypeHandler}>
                      <option name="ProductType" className="form-control" >
                          Good
                      </option>
                      <option name="ProductType" className="form-control" >
                          Service
                      </option>
                      <option name="ProductType" className="form-control" >
                          Subscription
                      </option>
                      <option name="ProductType" className="form-control" >
                          Bundle
                      </option>
                    </select>

                                            <label> Uom:&emsp; </label>
                                                <select value={this.state.uom} onChange={this.changeUomHandler}>
                      <option name="Uom" className="form-control" >
                          Each
                      </option>
                      <option name="Uom" className="form-control" >
                          Hour
                      </option>
                      <option name="Uom" className="form-control" >
                          Day
                      </option>
                      <option name="Uom" className="form-control" >
                          Month
                      </option>
                      <option name="Uom" className="form-control" >
                          User
                      </option>
                      <option name="Uom" className="form-control" >
                          Package
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateProduct}>Save</button>
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

export default CreateProductComponent
