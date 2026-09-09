import React, { Component } from 'react'
import ProductService from '../services/ProductService';

class CreateProductComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                slug: '',
                asActive: '',
                productType: '',
                defaultTaxClass: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeslugHandler = this.changeslugHandler.bind(this);
        this.changeasActiveHandler = this.changeasActiveHandler.bind(this);
        this.changeProductTypeHandler = this.changeProductTypeHandler.bind(this);
        this.changeDefaultTaxClassHandler = this.changeDefaultTaxClassHandler.bind(this);
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
                    name: product.name,
                    slug: product.slug,
                    asActive: product.asActive,
                    productType: product.productType,
                    defaultTaxClass: product.defaultTaxClass
                });
            });
        }        
    }
    saveOrUpdateProduct = (e) => {
        e.preventDefault();
        let product = {
                productId: this.state.id,
                name: this.state.name,
                slug: this.state.slug,
                asActive: this.state.asActive,
                productType: this.state.productType,
                defaultTaxClass: this.state.defaultTaxClass
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
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeslugHandler= (event) => {
        this.setState({slug: event.target.value});
    }
    changeasActiveHandler= (event) => {
        this.setState({asActive: event.target.value});
    }
    changeProductTypeHandler= (event) => {
        this.setState({productType: event.target.value});
    }
    changeDefaultTaxClassHandler= (event) => {
        this.setState({defaultTaxClass: event.target.value});
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
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> slug:&emsp; </label>
                                                <input placeholder="slug" name="slug" className="form-control" value={this.state.slug} onChange={this.changeslugHandler}/>

                                            <label> asActive:&emsp; </label>
                                                <input type="checkbox" placeholder="asActive" name="asActive" className="form-control" value={this.state.asActive} onChange={this.changeasActiveHandler}/>


                                            <label> ProductType:&emsp; </label>
                                                <select value={this.state.productType} onChange={this.changeProductTypeHandler}>
                      <option name="ProductType" className="form-control" >
                          Physical
                      </option>
                      <option name="ProductType" className="form-control" >
                          Digital
                      </option>
                      <option name="ProductType" className="form-control" >
                          Service
                      </option>
                      <option name="ProductType" className="form-control" >
                          Bundle
                      </option>
                      <option name="ProductType" className="form-control" >
                          Subscription
                      </option>
                    </select>

                                            <label> DefaultTaxClass:&emsp; </label>
                                                <select value={this.state.defaultTaxClass} onChange={this.changeDefaultTaxClassHandler}>
                      <option name="DefaultTaxClass" className="form-control" >
                          Standard
                      </option>
                      <option name="DefaultTaxClass" className="form-control" >
                          Reduced
                      </option>
                      <option name="DefaultTaxClass" className="form-control" >
                          Zero
                      </option>
                      <option name="DefaultTaxClass" className="form-control" >
                          Exempt
                      </option>
                      <option name="DefaultTaxClass" className="form-control" >
                          DigitalServices
                      </option>
                      <option name="DefaultTaxClass" className="form-control" >
                          Food
                      </option>
                      <option name="DefaultTaxClass" className="form-control" >
                          Clothing
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
