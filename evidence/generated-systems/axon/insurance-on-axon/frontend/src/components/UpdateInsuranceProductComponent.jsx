import React, { Component } from 'react'
import InsuranceProductService from '../services/InsuranceProductService';

class UpdateInsuranceProductComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                productCode: '',
                lineOfBusiness: ''
        }
        this.updateInsuranceProduct = this.updateInsuranceProduct.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeproductCodeHandler = this.changeproductCodeHandler.bind(this);
        this.changeLineOfBusinessHandler = this.changeLineOfBusinessHandler.bind(this);
    }

    componentDidMount(){
        InsuranceProductService.getInsuranceProductById(this.state.id).then( (res) =>{
            let insuranceProduct = res.data;
            this.setState({
                name: insuranceProduct.name,
                productCode: insuranceProduct.productCode,
                lineOfBusiness: insuranceProduct.lineOfBusiness
            });
        });
    }

    updateInsuranceProduct = (e) => {
        e.preventDefault();
        let insuranceProduct = {
            insuranceProductId: this.state.id,
            name: this.state.name,
            productCode: this.state.productCode,
            lineOfBusiness: this.state.lineOfBusiness
        };
        console.log('insuranceProduct => ' + JSON.stringify(insuranceProduct));
        console.log('id => ' + JSON.stringify(this.state.id));
        InsuranceProductService.updateInsuranceProduct(insuranceProduct).then( res => {
            this.props.history.push('/insuranceProducts');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeproductCodeHandler= (event) => {
        this.setState({productCode: event.target.value});
    }
    changeLineOfBusinessHandler= (event) => {
        this.setState({lineOfBusiness: event.target.value});
    }

    cancel(){
        this.props.history.push('/insuranceProducts');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update InsuranceProduct</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> productCode: </label>
                                                <input placeholder="productCode" name="productCode" className="form-control" value={this.state.productCode} onChange={this.changeproductCodeHandler}/>

                                            <label> LineOfBusiness: </label>
                                                <select value={this.state.lineOfBusiness} onChange={this.changeLineOfBusinessHandler}>
                      <option name="LineOfBusiness" className="form-control" >
                          PersonalAuto
                      </option>
                      <option name="LineOfBusiness" className="form-control" >
                          Homeowners
                      </option>
                      <option name="LineOfBusiness" className="form-control" >
                          Renters
                      </option>
                      <option name="LineOfBusiness" className="form-control" >
                          TermLife
                      </option>
                      <option name="LineOfBusiness" className="form-control" >
                          WholeLife
                      </option>
                      <option name="LineOfBusiness" className="form-control" >
                          CommercialProperty
                      </option>
                      <option name="LineOfBusiness" className="form-control" >
                          GeneralLiability
                      </option>
                      <option name="LineOfBusiness" className="form-control" >
                          WorkersCompensation
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateInsuranceProduct}>Save</button>
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

export default UpdateInsuranceProductComponent
