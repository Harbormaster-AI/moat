import React, { Component } from 'react'
import ProductPricingService from '../services/ProductPricingService';

class UpdateProductPricingComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                listPrice: '',
                salePrice: '',
                validFrom: '',
                validTo: ''
        }
        this.updateProductPricing = this.updateProductPricing.bind(this);

        this.changelistPriceHandler = this.changelistPriceHandler.bind(this);
        this.changesalePriceHandler = this.changesalePriceHandler.bind(this);
        this.changevalidFromHandler = this.changevalidFromHandler.bind(this);
        this.changevalidToHandler = this.changevalidToHandler.bind(this);
    }

    componentDidMount(){
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

    updateProductPricing = (e) => {
        e.preventDefault();
        let productPricing = {
            productPricingId: this.state.id,
            listPrice: this.state.listPrice,
            salePrice: this.state.salePrice,
            validFrom: this.state.validFrom,
            validTo: this.state.validTo
        };
        console.log('productPricing => ' + JSON.stringify(productPricing));
        console.log('id => ' + JSON.stringify(this.state.id));
        ProductPricingService.updateProductPricing(productPricing).then( res => {
            this.props.history.push('/productPricings');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update ProductPricing</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> listPrice: </label>
                                                <input placeholder="listPrice" name="listPrice" className="form-control" value={this.state.listPrice} onChange={this.changelistPriceHandler}/>

                                            <label> salePrice: </label>
                                                <input placeholder="salePrice" name="salePrice" className="form-control" value={this.state.salePrice} onChange={this.changesalePriceHandler}/>

                                            <label> validFrom: </label>
                                                <input type="date" placeholder="validFrom" name="validFrom" className="form-control" value={this.state.validFrom} onChange={this.changevalidFromHandler}/>

                                            <label> validTo: </label>
                                                <input type="date" placeholder="validTo" name="validTo" className="form-control" value={this.state.validTo} onChange={this.changevalidToHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateProductPricing}>Save</button>
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

export default UpdateProductPricingComponent
