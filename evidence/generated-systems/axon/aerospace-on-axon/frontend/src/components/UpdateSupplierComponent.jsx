import React, { Component } from 'react'
import SupplierService from '../services/SupplierService';

class UpdateSupplierComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                supplierType: '',
                approvalStatus: ''
        }
        this.updateSupplier = this.updateSupplier.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeSupplierTypeHandler = this.changeSupplierTypeHandler.bind(this);
        this.changeApprovalStatusHandler = this.changeApprovalStatusHandler.bind(this);
    }

    componentDidMount(){
        SupplierService.getSupplierById(this.state.id).then( (res) =>{
            let supplier = res.data;
            this.setState({
                name: supplier.name,
                supplierType: supplier.supplierType,
                approvalStatus: supplier.approvalStatus
            });
        });
    }

    updateSupplier = (e) => {
        e.preventDefault();
        let supplier = {
            supplierId: this.state.id,
            name: this.state.name,
            supplierType: this.state.supplierType,
            approvalStatus: this.state.approvalStatus
        };
        console.log('supplier => ' + JSON.stringify(supplier));
        console.log('id => ' + JSON.stringify(this.state.id));
        SupplierService.updateSupplier(supplier).then( res => {
            this.props.history.push('/suppliers');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Supplier</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> SupplierType: </label>
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

                                            <label> ApprovalStatus: </label>
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
                                        <button className="btn btn-success" onClick={this.updateSupplier}>Save</button>
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

export default UpdateSupplierComponent
