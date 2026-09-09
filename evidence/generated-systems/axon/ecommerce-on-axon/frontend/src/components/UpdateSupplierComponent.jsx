import React, { Component } from 'react'
import SupplierService from '../services/SupplierService';

class UpdateSupplierComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                contactEmail: '',
                website: '',
                status: ''
        }
        this.updateSupplier = this.updateSupplier.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changecontactEmailHandler = this.changecontactEmailHandler.bind(this);
        this.changewebsiteHandler = this.changewebsiteHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        SupplierService.getSupplierById(this.state.id).then( (res) =>{
            let supplier = res.data;
            this.setState({
                name: supplier.name,
                contactEmail: supplier.contactEmail,
                website: supplier.website,
                status: supplier.status
            });
        });
    }

    updateSupplier = (e) => {
        e.preventDefault();
        let supplier = {
            supplierId: this.state.id,
            name: this.state.name,
            contactEmail: this.state.contactEmail,
            website: this.state.website,
            status: this.state.status
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
    changecontactEmailHandler= (event) => {
        this.setState({contactEmail: event.target.value});
    }
    changewebsiteHandler= (event) => {
        this.setState({website: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
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

                                            <label> contactEmail: </label>
                                                <input placeholder="contactEmail" name="contactEmail" className="form-control" value={this.state.contactEmail} onChange={this.changecontactEmailHandler}/>

                                            <label> website: </label>
                                                <input placeholder="website" name="website" className="form-control" value={this.state.website} onChange={this.changewebsiteHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Inactive
                      </option>
                      <option name="Status" className="form-control" >
                          Onboarding
                      </option>
                      <option name="Status" className="form-control" >
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
