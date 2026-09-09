import React, { Component } from 'react'
import ProductionCertificateService from '../services/ProductionCertificateService';

class UpdateProductionCertificateComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                certificateNumber: '',
                authority: ''
        }
        this.updateProductionCertificate = this.updateProductionCertificate.bind(this);

        this.changecertificateNumberHandler = this.changecertificateNumberHandler.bind(this);
        this.changeauthorityHandler = this.changeauthorityHandler.bind(this);
    }

    componentDidMount(){
        ProductionCertificateService.getProductionCertificateById(this.state.id).then( (res) =>{
            let productionCertificate = res.data;
            this.setState({
                certificateNumber: productionCertificate.certificateNumber,
                authority: productionCertificate.authority
            });
        });
    }

    updateProductionCertificate = (e) => {
        e.preventDefault();
        let productionCertificate = {
            productionCertificateId: this.state.id,
            certificateNumber: this.state.certificateNumber,
            authority: this.state.authority
        };
        console.log('productionCertificate => ' + JSON.stringify(productionCertificate));
        console.log('id => ' + JSON.stringify(this.state.id));
        ProductionCertificateService.updateProductionCertificate(productionCertificate).then( res => {
            this.props.history.push('/productionCertificates');
        });
    }

    changecertificateNumberHandler= (event) => {
        this.setState({certificateNumber: event.target.value});
    }
    changeauthorityHandler= (event) => {
        this.setState({authority: event.target.value});
    }

    cancel(){
        this.props.history.push('/productionCertificates');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update ProductionCertificate</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> certificateNumber: </label>
                                                <input placeholder="certificateNumber" name="certificateNumber" className="form-control" value={this.state.certificateNumber} onChange={this.changecertificateNumberHandler}/>

                                            <label> authority: </label>
                                                <input placeholder="authority" name="authority" className="form-control" value={this.state.authority} onChange={this.changeauthorityHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateProductionCertificate}>Save</button>
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

export default UpdateProductionCertificateComponent
