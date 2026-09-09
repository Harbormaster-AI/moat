import React, { Component } from 'react'
import SellerService from '../services/SellerService';

class CreateSellerComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                sellerCode: '',
                contactEmail: '',
                status: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changesellerCodeHandler = this.changesellerCodeHandler.bind(this);
        this.changecontactEmailHandler = this.changecontactEmailHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateSeller = (e) => {
        e.preventDefault();
        let seller = {
                sellerId: this.state.id,
                name: this.state.name,
                sellerCode: this.state.sellerCode,
                contactEmail: this.state.contactEmail,
                status: this.state.status
            };
        console.log('seller => ' + JSON.stringify(seller));

        // step 5
        if(this.state.id === '_add'){
            seller.sellerId=''
            SellerService.createSeller(seller).then(res =>{
                this.props.history.push('/sellers');
            });
        }else{
            SellerService.updateSeller(seller).then( res => {
                this.props.history.push('/sellers');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Seller</h3>
        }else{
            return <h3 className="text-center">Update Seller</h3>
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

                                            <label> sellerCode:&emsp; </label>
                                                <input placeholder="sellerCode" name="sellerCode" className="form-control" value={this.state.sellerCode} onChange={this.changesellerCodeHandler}/>

                                            <label> contactEmail:&emsp; </label>
                                                <input placeholder="contactEmail" name="contactEmail" className="form-control" value={this.state.contactEmail} onChange={this.changecontactEmailHandler}/>

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateSeller}>Save</button>
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

export default CreateSellerComponent
