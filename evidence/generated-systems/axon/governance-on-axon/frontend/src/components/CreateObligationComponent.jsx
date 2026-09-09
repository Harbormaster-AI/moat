import React, { Component } from 'react'
import ObligationService from '../services/ObligationService';

class CreateObligationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                referenceNumber: '',
                descriptionText: '',
                obligationType: '',
                reviewFrequency: ''
        }
        this.changereferenceNumberHandler = this.changereferenceNumberHandler.bind(this);
        this.changedescriptionTextHandler = this.changedescriptionTextHandler.bind(this);
        this.changeObligationTypeHandler = this.changeObligationTypeHandler.bind(this);
        this.changeReviewFrequencyHandler = this.changeReviewFrequencyHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ObligationService.getObligationById(this.state.id).then( (res) =>{
                let obligation = res.data;
                this.setState({
                    referenceNumber: obligation.referenceNumber,
                    descriptionText: obligation.descriptionText,
                    obligationType: obligation.obligationType,
                    reviewFrequency: obligation.reviewFrequency
                });
            });
        }        
    }
    saveOrUpdateObligation = (e) => {
        e.preventDefault();
        let obligation = {
                obligationId: this.state.id,
                referenceNumber: this.state.referenceNumber,
                descriptionText: this.state.descriptionText,
                obligationType: this.state.obligationType,
                reviewFrequency: this.state.reviewFrequency
            };
        console.log('obligation => ' + JSON.stringify(obligation));

        // step 5
        if(this.state.id === '_add'){
            obligation.obligationId=''
            ObligationService.createObligation(obligation).then(res =>{
                this.props.history.push('/obligations');
            });
        }else{
            ObligationService.updateObligation(obligation).then( res => {
                this.props.history.push('/obligations');
            });
        }
    }
    
    changereferenceNumberHandler= (event) => {
        this.setState({referenceNumber: event.target.value});
    }
    changedescriptionTextHandler= (event) => {
        this.setState({descriptionText: event.target.value});
    }
    changeObligationTypeHandler= (event) => {
        this.setState({obligationType: event.target.value});
    }
    changeReviewFrequencyHandler= (event) => {
        this.setState({reviewFrequency: event.target.value});
    }

    cancel(){
        this.props.history.push('/obligations');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Obligation</h3>
        }else{
            return <h3 className="text-center">Update Obligation</h3>
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
                                            <label> referenceNumber:&emsp; </label>
                                                <input placeholder="referenceNumber" name="referenceNumber" className="form-control" value={this.state.referenceNumber} onChange={this.changereferenceNumberHandler}/>

                                            <label> descriptionText:&emsp; </label>
                                                <input placeholder="descriptionText" name="descriptionText" className="form-control" value={this.state.descriptionText} onChange={this.changedescriptionTextHandler}/>

                                            <label> ObligationType:&emsp; </label>
                                                <select value={this.state.obligationType} onChange={this.changeObligationTypeHandler}>
                      <option name="ObligationType" className="form-control" >
                          Regulatory
                      </option>
                      <option name="ObligationType" className="form-control" >
                          Contractual
                      </option>
                      <option name="ObligationType" className="form-control" >
                          PolicyDerived
                      </option>
                      <option name="ObligationType" className="form-control" >
                          IndustryStandard
                      </option>
                    </select>

                                            <label> ReviewFrequency:&emsp; </label>
                                                <select value={this.state.reviewFrequency} onChange={this.changeReviewFrequencyHandler}>
                      <option name="ReviewFrequency" className="form-control" >
                          Continuous
                      </option>
                      <option name="ReviewFrequency" className="form-control" >
                          Daily
                      </option>
                      <option name="ReviewFrequency" className="form-control" >
                          Weekly
                      </option>
                      <option name="ReviewFrequency" className="form-control" >
                          Monthly
                      </option>
                      <option name="ReviewFrequency" className="form-control" >
                          Quarterly
                      </option>
                      <option name="ReviewFrequency" className="form-control" >
                          Annually
                      </option>
                      <option name="ReviewFrequency" className="form-control" >
                          AdHoc
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateObligation}>Save</button>
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

export default CreateObligationComponent
