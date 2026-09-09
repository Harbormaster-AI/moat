import React, { Component } from 'react'
import SellerService from '../services/SellerService';

class UpdateSellerComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                sellerCode: '',
                contactEmail: '',
                status: ''
        }
        this.updateSeller = this.updateSeller.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changesellerCodeHandler = this.changesellerCodeHandler.bind(this);
        this.changecontactEmailHandler = this.changecontactEmailHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        SellerService.getSellerById(this.state.id).then( (res) =>{
            let seller = res.data;
            this.setState({
                name: seller.name,
                sellerCode: seller.sellerCode,
                contactEmail: seller.contactEmail,
                status: seller.status
            });
        });
    }

    updateSeller = (e) => {
        e.preventDefault();
        let seller = {
            sellerId: this.state.id,
            name: this.state.name,
            sellerCode: this.state.sellerCode,
            contactEmail: this.state.contactEmail,
            status: this.state.status
        };
        console.log('seller => ' + JSON.stringify(seller));
        console.log('id => ' + JSON.stringify(this.state.id));
        SellerService.updateSeller(seller).then( res => {
            this.props.history.push('/sellers');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changesellerCodeHandler= (event) => {
        this.setState({sellerCode: event.target.value});
    }
    changecontactEmailHandler= (event) => {
        this.setState({contactEmail: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/sellers');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Seller</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> sellerCode: </label>
                                                <input placeholder="sellerCode" name="sellerCode" className="form-control" value={this.state.sellerCode} onChange={this.changesellerCodeHandler}/>

                                            <label> contactEmail: </label>
                                                <input placeholder="contactEmail" name="contactEmail" className="form-control" value={this.state.contactEmail} onChange={this.changecontactEmailHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Inactive
                      </option>
                      <option name="Status" className="form-control" >
                          Suspended
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateSeller}>Save</button>
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

export default UpdateSellerComponent
