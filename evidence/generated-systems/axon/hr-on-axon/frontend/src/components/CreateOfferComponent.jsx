import React, { Component } from 'react'
import OfferService from '../services/OfferService';

class CreateOfferComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                offerNumber: '',
                proposedStartDate: '',
                baseSalary: '',
                signOnBonus: '',
                status: ''
        }
        this.changeofferNumberHandler = this.changeofferNumberHandler.bind(this);
        this.changeproposedStartDateHandler = this.changeproposedStartDateHandler.bind(this);
        this.changebaseSalaryHandler = this.changebaseSalaryHandler.bind(this);
        this.changesignOnBonusHandler = this.changesignOnBonusHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            OfferService.getOfferById(this.state.id).then( (res) =>{
                let offer = res.data;
                this.setState({
                    offerNumber: offer.offerNumber,
                    proposedStartDate: offer.proposedStartDate,
                    baseSalary: offer.baseSalary,
                    signOnBonus: offer.signOnBonus,
                    status: offer.status
                });
            });
        }        
    }
    saveOrUpdateOffer = (e) => {
        e.preventDefault();
        let offer = {
                offerId: this.state.id,
                offerNumber: this.state.offerNumber,
                proposedStartDate: this.state.proposedStartDate,
                baseSalary: this.state.baseSalary,
                signOnBonus: this.state.signOnBonus,
                status: this.state.status
            };
        console.log('offer => ' + JSON.stringify(offer));

        // step 5
        if(this.state.id === '_add'){
            offer.offerId=''
            OfferService.createOffer(offer).then(res =>{
                this.props.history.push('/offers');
            });
        }else{
            OfferService.updateOffer(offer).then( res => {
                this.props.history.push('/offers');
            });
        }
    }
    
    changeofferNumberHandler= (event) => {
        this.setState({offerNumber: event.target.value});
    }
    changeproposedStartDateHandler= (event) => {
        this.setState({proposedStartDate: event.target.value});
    }
    changebaseSalaryHandler= (event) => {
        this.setState({baseSalary: event.target.value});
    }
    changesignOnBonusHandler= (event) => {
        this.setState({signOnBonus: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/offers');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Offer</h3>
        }else{
            return <h3 className="text-center">Update Offer</h3>
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
                                            <label> offerNumber:&emsp; </label>
                                                <input placeholder="offerNumber" name="offerNumber" className="form-control" value={this.state.offerNumber} onChange={this.changeofferNumberHandler}/>

                                            <label> proposedStartDate:&emsp; </label>
                                                <input type="date" placeholder="proposedStartDate" name="proposedStartDate" className="form-control" value={this.state.proposedStartDate} onChange={this.changeproposedStartDateHandler}/>

                                            <label> baseSalary:&emsp; </label>
                                                <input placeholder="baseSalary" name="baseSalary" className="form-control" value={this.state.baseSalary} onChange={this.changebaseSalaryHandler}/>

                                            <label> signOnBonus:&emsp; </label>
                                                <input placeholder="signOnBonus" name="signOnBonus" className="form-control" value={this.state.signOnBonus} onChange={this.changesignOnBonusHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Sent
                      </option>
                      <option name="Status" className="form-control" >
                          Accepted
                      </option>
                      <option name="Status" className="form-control" >
                          Declined
                      </option>
                      <option name="Status" className="form-control" >
                          Withdrawn
                      </option>
                      <option name="Status" className="form-control" >
                          Expired
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateOffer}>Save</button>
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

export default CreateOfferComponent
