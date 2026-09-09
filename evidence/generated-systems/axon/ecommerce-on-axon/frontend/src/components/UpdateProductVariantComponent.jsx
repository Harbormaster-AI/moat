import React, { Component } from 'react'
import ProductVariantService from '../services/ProductVariantService';

class UpdateProductVariantComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                sku: '',
                barcode: '',
                title: '',
                weight: '',
                requiresShipping: '',
                weightUnit: ''
        }
        this.updateProductVariant = this.updateProductVariant.bind(this);

        this.changeskuHandler = this.changeskuHandler.bind(this);
        this.changebarcodeHandler = this.changebarcodeHandler.bind(this);
        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changeweightHandler = this.changeweightHandler.bind(this);
        this.changerequiresShippingHandler = this.changerequiresShippingHandler.bind(this);
        this.changeWeightUnitHandler = this.changeWeightUnitHandler.bind(this);
    }

    componentDidMount(){
        ProductVariantService.getProductVariantById(this.state.id).then( (res) =>{
            let productVariant = res.data;
            this.setState({
                sku: productVariant.sku,
                barcode: productVariant.barcode,
                title: productVariant.title,
                weight: productVariant.weight,
                requiresShipping: productVariant.requiresShipping,
                weightUnit: productVariant.weightUnit
            });
        });
    }

    updateProductVariant = (e) => {
        e.preventDefault();
        let productVariant = {
            productVariantId: this.state.id,
            sku: this.state.sku,
            barcode: this.state.barcode,
            title: this.state.title,
            weight: this.state.weight,
            requiresShipping: this.state.requiresShipping,
            weightUnit: this.state.weightUnit
        };
        console.log('productVariant => ' + JSON.stringify(productVariant));
        console.log('id => ' + JSON.stringify(this.state.id));
        ProductVariantService.updateProductVariant(productVariant).then( res => {
            this.props.history.push('/productVariants');
        });
    }

    changeskuHandler= (event) => {
        this.setState({sku: event.target.value});
    }
    changebarcodeHandler= (event) => {
        this.setState({barcode: event.target.value});
    }
    changetitleHandler= (event) => {
        this.setState({title: event.target.value});
    }
    changeweightHandler= (event) => {
        this.setState({weight: event.target.value});
    }
    changerequiresShippingHandler= (event) => {
        this.setState({requiresShipping: event.target.value});
    }
    changeWeightUnitHandler= (event) => {
        this.setState({weightUnit: event.target.value});
    }

    cancel(){
        this.props.history.push('/productVariants');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update ProductVariant</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> sku: </label>
                                                <input placeholder="sku" name="sku" className="form-control" value={this.state.sku} onChange={this.changeskuHandler}/>

                                            <label> barcode: </label>
                                                <input placeholder="barcode" name="barcode" className="form-control" value={this.state.barcode} onChange={this.changebarcodeHandler}/>

                                            <label> title: </label>
                                                <input placeholder="title" name="title" className="form-control" value={this.state.title} onChange={this.changetitleHandler}/>

                                            <label> weight: </label>
                                                <input placeholder="weight" name="weight" className="form-control" value={this.state.weight} onChange={this.changeweightHandler}/>

                                            <label> requiresShipping: </label>
                                                <input type="checkbox" placeholder="requiresShipping" name="requiresShipping" className="form-control" value={this.state.requiresShipping} onChange={this.changerequiresShippingHandler}/>


                                            <label> WeightUnit: </label>
                                                <select value={this.state.weightUnit} onChange={this.changeWeightUnitHandler}>
                      <option name="WeightUnit" className="form-control" >
                          Gram
                      </option>
                      <option name="WeightUnit" className="form-control" >
                          Kilogram
                      </option>
                      <option name="WeightUnit" className="form-control" >
                          Ounce
                      </option>
                      <option name="WeightUnit" className="form-control" >
                          Pound
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateProductVariant}>Save</button>
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

export default UpdateProductVariantComponent
