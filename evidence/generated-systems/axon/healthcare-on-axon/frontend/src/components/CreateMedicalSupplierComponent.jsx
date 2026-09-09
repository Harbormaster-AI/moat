import React, { Component } from 'react'
import MedicalSupplierService from '../services/MedicalSupplierService';

class CreateMedicalSupplierComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                website: '',
                supplierTier: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changewebsiteHandler = this.changewebsiteHandler.bind(this);
        this.changeSupplierTierHandler = this.changeSupplierTierHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            MedicalSupplierService.getMedicalSupplierById(this.state.id).then( (res) =>{
                let medicalSupplier = res.data;
                this.setState({
                    name: medicalSupplier.name,
                    website: medicalSupplier.website,
                    supplierTier: medicalSupplier.supplierTier
                });
            });
        }        
    }
    saveOrUpdateMedicalSupplier = (e) => {
        e.preventDefault();
        let medicalSupplier = {
                medicalSupplierId: this.state.id,
                name: this.state.name,
                website: this.state.website,
                supplierTier: this.state.supplierTier
            };
        console.log('medicalSupplier => ' + JSON.stringify(medicalSupplier));

        // step 5
        if(this.state.id === '_add'){
            medicalSupplier.medicalSupplierId=''
            MedicalSupplierService.createMedicalSupplier(medicalSupplier).then(res =>{
                this.props.history.push('/medicalSuppliers');
            });
        }else{
            MedicalSupplierService.updateMedicalSupplier(medicalSupplier).then( res => {
                this.props.history.push('/medicalSuppliers');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add MedicalSupplier</h3>
        }else{
            return <h3 className="text-center">Update MedicalSupplier</h3>
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

                                            <label> website:&emsp; </label>
                                                <input placeholder="website" name="website" className="form-control" value={this.state.website} onChange={this.changewebsiteHandler}/>

                                            <label> SupplierTier:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateMedicalSupplier}>Save</button>
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

export default CreateMedicalSupplierComponent
