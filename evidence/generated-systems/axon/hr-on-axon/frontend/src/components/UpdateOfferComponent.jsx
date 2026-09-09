import React, { Component } from 'react'
import OfferService from '../services/OfferService';

class UpdateOfferComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                offerNumber: '',
                proposedStartDate: '',
                baseSalary: '',
                signOnBonus: '',
                status: ''
        }
        this.updateOffer = this.updateOffer.bind(this);

        this.changeofferNumberHandler = this.changeofferNumberHandler.bind(this);
        this.changeproposedStartDateHandler = this.changeproposedStartDateHandler.bind(this);
        this.changebaseSalaryHandler = this.changebaseSalaryHandler.bind(this);
        this.changesignOnBonusHandler = this.changesignOnBonusHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
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

    updateOffer = (e) => {
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
        console.log('id => ' + JSON.stringify(this.state.id));
        OfferService.updateOffer(offer).then( res => {
            this.props.history.push('/offers');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Offer</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> offerNumber: </label>
                                                <input placeholder="offerNumber" name="offerNumber" className="form-control" value={this.state.offerNumber} onChange={this.changeofferNumberHandler}/>

                                            <label> proposedStartDate: </label>
                                                <input type="date" placeholder="proposedStartDate" name="proposedStartDate" className="form-control" value={this.state.proposedStartDate} onChange={this.changeproposedStartDateHandler}/>

                                            <label> baseSalary: </label>
                                                <input placeholder="baseSalary" name="baseSalary" className="form-control" value={this.state.baseSalary} onChange={this.changebaseSalaryHandler}/>

                                            <label> signOnBonus: </label>
                                                <input placeholder="signOnBonus" name="signOnBonus" className="form-control" value={this.state.signOnBonus} onChange={this.changesignOnBonusHandler}/>

                                            <label> Status: </label>
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
                                        <button className="btn btn-success" onClick={this.updateOffer}>Save</button>
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

export default UpdateOfferComponent
