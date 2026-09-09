import React, { Component } from 'react'
import ProductionCertificateService from '../services/ProductionCertificateService';

class CreateProductionCertificateComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                certificateNumber: '',
                authority: ''
        }
        this.changecertificateNumberHandler = this.changecertificateNumberHandler.bind(this);
        this.changeauthorityHandler = this.changeauthorityHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ProductionCertificateService.getProductionCertificateById(this.state.id).then( (res) =>{
                let productionCertificate = res.data;
                this.setState({
                    certificateNumber: productionCertificate.certificateNumber,
                    authority: productionCertificate.authority
                });
            });
        }        
    }
    saveOrUpdateProductionCertificate = (e) => {
        e.preventDefault();
        let productionCertificate = {
                productionCertificateId: this.state.id,
                certificateNumber: this.state.certificateNumber,
                authority: this.state.authority
            };
        console.log('productionCertificate => ' + JSON.stringify(productionCertificate));

        // step 5
        if(this.state.id === '_add'){
            productionCertificate.productionCertificateId=''
            ProductionCertificateService.createProductionCertificate(productionCertificate).then(res =>{
                this.props.history.push('/productionCertificates');
            });
        }else{
            ProductionCertificateService.updateProductionCertificate(productionCertificate).then( res => {
                this.props.history.push('/productionCertificates');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add ProductionCertificate</h3>
        }else{
            return <h3 className="text-center">Update ProductionCertificate</h3>
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
                                            <label> certificateNumber:&emsp; </label>
                                                <input placeholder="certificateNumber" name="certificateNumber" className="form-control" value={this.state.certificateNumber} onChange={this.changecertificateNumberHandler}/>

                                            <label> authority:&emsp; </label>
                                                <input placeholder="authority" name="authority" className="form-control" value={this.state.authority} onChange={this.changeauthorityHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateProductionCertificate}>Save</button>
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

export default CreateProductionCertificateComponent
