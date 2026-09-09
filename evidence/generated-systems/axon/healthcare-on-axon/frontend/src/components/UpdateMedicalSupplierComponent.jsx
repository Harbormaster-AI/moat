import React, { Component } from 'react'
import MedicalSupplierService from '../services/MedicalSupplierService';

class UpdateMedicalSupplierComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                website: '',
                supplierTier: ''
        }
        this.updateMedicalSupplier = this.updateMedicalSupplier.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changewebsiteHandler = this.changewebsiteHandler.bind(this);
        this.changeSupplierTierHandler = this.changeSupplierTierHandler.bind(this);
    }

    componentDidMount(){
        MedicalSupplierService.getMedicalSupplierById(this.state.id).then( (res) =>{
            let medicalSupplier = res.data;
            this.setState({
                name: medicalSupplier.name,
                website: medicalSupplier.website,
                supplierTier: medicalSupplier.supplierTier
            });
        });
    }

    updateMedicalSupplier = (e) => {
        e.preventDefault();
        let medicalSupplier = {
            medicalSupplierId: this.state.id,
            name: this.state.name,
            website: this.state.website,
            supplierTier: this.state.supplierTier
        };
        console.log('medicalSupplier => ' + JSON.stringify(medicalSupplier));
        console.log('id => ' + JSON.stringify(this.state.id));
        MedicalSupplierService.updateMedicalSupplier(medicalSupplier).then( res => {
            this.props.history.push('/medicalSuppliers');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changewebsiteHandler= (event) => {
        this.setState({website: event.target.value});
    }
    changeSupplierTierHandler= (event) => {
        this.setState({supplierTier: event.target.value});
    }

    cancel(){
        this.props.history.push('/medicalSuppliers');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update MedicalSupplier</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> website: </label>
                                                <input placeholder="website" name="website" className="form-control" value={this.state.website} onChange={this.changewebsiteHandler}/>

                                            <label> SupplierTier: </label>
                                                <select value={this.state.supplierTier} onChange={this.changeSupplierTierHandler}>
                      <option name="SupplierTier" className="form-control" >
                          Primary
                      </option>
                      <option name="SupplierTier" className="form-control" >
                          Secondary
                      </option>
                      <option name="SupplierTier" className="form-control" >
                          Distributor
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateMedicalSupplier}>Save</button>
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

export default UpdateMedicalSupplierComponent
