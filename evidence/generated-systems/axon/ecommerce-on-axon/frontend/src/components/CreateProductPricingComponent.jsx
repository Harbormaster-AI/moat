import React, { Component } from 'react'
import ProductPricingService from '../services/ProductPricingService';

class CreateProductPricingComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                listPrice: '',
                salePrice: '',
                validFrom: '',
                validTo: ''
        }
        this.changelistPriceHandler = this.changelistPriceHandler.bind(this);
        this.changesalePriceHandler = this.changesalePriceHandler.bind(this);
        this.changevalidFromHandler = this.changevalidFromHandler.bind(this);
        this.changevalidToHandler = this.changevalidToHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ProductPricingService.getProductPricingById(this.state.id).then( (res) =>{
                let productPricing = res.data;
                this.setState({
                    listPrice: productPricing.listPrice,
                    salePrice: productPricing.salePrice,
                    validFrom: productPricing.validFrom,
                    validTo: productPricing.validTo
                });
            });
        }        
    }
    saveOrUpdateProductPricing = (e) => {
        e.preventDefault();
        let productPricing = {
                productPricingId: this.state.id,
                listPrice: this.state.listPrice,
                salePrice: this.state.salePrice,
                validFrom: this.state.validFrom,
                validTo: this.state.validTo
            };
        console.log('productPricing => ' + JSON.stringify(productPricing));

        // step 5
        if(this.state.id === '_add'){
            productPricing.productPricingId=''
            ProductPricingService.createProductPricing(productPricing).then(res =>{
                this.props.history.push('/productPricings');
            });
        }else{
            ProductPricingService.updateProductPricing(productPricing).then( res => {
                this.props.history.push('/productPricings');
            });
        }
    }
    
    changelistPriceHandler= (event) => {
        this.setState({listPrice: event.target.value});
    }
    changesalePriceHandler= (event) => {
        this.setState({salePrice: event.target.value});
    }
    changevalidFromHandler= (event) => {
        this.setState({validFrom: event.target.value});
    }
    changevalidToHandler= (event) => {
        this.setState({validTo: event.target.value});
    }

    cancel(){
        this.props.history.push('/productPricings');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add ProductPricing</h3>
        }else{
            return <h3 className="text-center">Update ProductPricing</h3>
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
                                            <label> listPrice:&emsp; </label>
                                                <input placeholder="listPrice" name="listPrice" className="form-control" value={this.state.listPrice} onChange={this.changelistPriceHandler}/>

                                            <label> salePrice:&emsp; </label>
                                                <input placeholder="salePrice" name="salePrice" className="form-control" value={this.state.salePrice} onChange={this.changesalePriceHandler}/>

                                            <label> validFrom:&emsp; </label>
                                                <input type="date" placeholder="validFrom" name="validFrom" className="form-control" value={this.state.validFrom} onChange={this.changevalidFromHandler}/>

                                            <label> validTo:&emsp; </label>
                                                <input type="date" placeholder="validTo" name="validTo" className="form-control" value={this.state.validTo} onChange={this.changevalidToHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateProductPricing}>Save</button>
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

export default CreateProductPricingComponent
