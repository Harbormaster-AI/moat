import React, { Component } from 'react'
import SupplierService from '../services/SupplierService';

class CreateSupplierComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                supplierType: '',
                approvalStatus: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeSupplierTypeHandler = this.changeSupplierTypeHandler.bind(this);
        this.changeApprovalStatusHandler = this.changeApprovalStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            SupplierService.getSupplierById(this.state.id).then( (res) =>{
                let supplier = res.data;
                this.setState({
                    name: supplier.name,
                    supplierType: supplier.supplierType,
                    approvalStatus: supplier.approvalStatus
                });
            });
        }        
    }
    saveOrUpdateSupplier = (e) => {
        e.preventDefault();
        let supplier = {
                supplierId: this.state.id,
                name: this.state.name,
                supplierType: this.state.supplierType,
                approvalStatus: this.state.approvalStatus
            };
        console.log('supplier => ' + JSON.stringify(supplier));

        // step 5
        if(this.state.id === '_add'){
            supplier.supplierId=''
            SupplierService.createSupplier(supplier).then(res =>{
                this.props.history.push('/suppliers');
            });
        }else{
            SupplierService.updateSupplier(supplier).then( res => {
                this.props.history.push('/suppliers');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeSupplierTypeHandler= (event) => {
        this.setState({supplierType: event.target.value});
    }
    changeApprovalStatusHandler= (event) => {
        this.setState({approvalStatus: event.target.value});
    }

    cancel(){
        this.props.history.push('/suppliers');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Supplier</h3>
        }else{
            return <h3 className="text-center">Update Supplier</h3>
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

                                            <label> SupplierType:&emsp; </label>
                                                <select value={this.state.supplierType} onChange={this.changeSupplierTypeHandler}>
                      <option name="SupplierType" className="form-control" >
                          Airframe
                      </option>
                      <option name="SupplierType" className="form-control" >
                          Engine
                      </option>
                      <option name="SupplierType" className="form-control" >
                          Avionics
                      </option>
                      <option name="SupplierType" className="form-control" >
                          Systems
                      </option>
                      <option name="SupplierType" className="form-control" >
                          Materials
                      </option>
                      <option name="SupplierType" className="form-control" >
                          MRO
                      </option>
                      <option name="SupplierType" className="form-control" >
                          Testing
                      </option>
                    </select>

                                            <label> ApprovalStatus:&emsp; </label>
                                                <select value={this.state.approvalStatus} onChange={this.changeApprovalStatusHandler}>
                      <option name="ApprovalStatus" className="form-control" >
                          Applied
                      </option>
                      <option name="ApprovalStatus" className="form-control" >
                          Approved
                      </option>
                      <option name="ApprovalStatus" className="form-control" >
                          OnHold
                      </option>
                      <option name="ApprovalStatus" className="form-control" >
                          Suspended
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateSupplier}>Save</button>
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

export default CreateSupplierComponent
