import React, { Component } from 'react'
import ReinsuranceAgreementService from '../services/ReinsuranceAgreementService';

class CreateReinsuranceAgreementComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                agreementNumber: '',
                effectivePeriod: '',
                retention: '',
                limit: '',
                cessionPercentage: '',
                reinsuranceType: '',
                treatyType: ''
        }
        this.changeagreementNumberHandler = this.changeagreementNumberHandler.bind(this);
        this.changeeffectivePeriodHandler = this.changeeffectivePeriodHandler.bind(this);
        this.changeretentionHandler = this.changeretentionHandler.bind(this);
        this.changelimitHandler = this.changelimitHandler.bind(this);
        this.changecessionPercentageHandler = this.changecessionPercentageHandler.bind(this);
        this.changeReinsuranceTypeHandler = this.changeReinsuranceTypeHandler.bind(this);
        this.changeTreatyTypeHandler = this.changeTreatyTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ReinsuranceAgreementService.getReinsuranceAgreementById(this.state.id).then( (res) =>{
                let reinsuranceAgreement = res.data;
                this.setState({
                    agreementNumber: reinsuranceAgreement.agreementNumber,
                    effectivePeriod: reinsuranceAgreement.effectivePeriod,
                    retention: reinsuranceAgreement.retention,
                    limit: reinsuranceAgreement.limit,
                    cessionPercentage: reinsuranceAgreement.cessionPercentage,
                    reinsuranceType: reinsuranceAgreement.reinsuranceType,
                    treatyType: reinsuranceAgreement.treatyType
                });
            });
        }        
    }
    saveOrUpdateReinsuranceAgreement = (e) => {
        e.preventDefault();
        let reinsuranceAgreement = {
                reinsuranceAgreementId: this.state.id,
                agreementNumber: this.state.agreementNumber,
                effectivePeriod: this.state.effectivePeriod,
                retention: this.state.retention,
                limit: this.state.limit,
                cessionPercentage: this.state.cessionPercentage,
                reinsuranceType: this.state.reinsuranceType,
                treatyType: this.state.treatyType
            };
        console.log('reinsuranceAgreement => ' + JSON.stringify(reinsuranceAgreement));

        // step 5
        if(this.state.id === '_add'){
            reinsuranceAgreement.reinsuranceAgreementId=''
            ReinsuranceAgreementService.createReinsuranceAgreement(reinsuranceAgreement).then(res =>{
                this.props.history.push('/reinsuranceAgreements');
            });
        }else{
            ReinsuranceAgreementService.updateReinsuranceAgreement(reinsuranceAgreement).then( res => {
                this.props.history.push('/reinsuranceAgreements');
            });
        }
    }
    
    changeagreementNumberHandler= (event) => {
        this.setState({agreementNumber: event.target.value});
    }
    changeeffectivePeriodHandler= (event) => {
        this.setState({effectivePeriod: event.target.value});
    }
    changeretentionHandler= (event) => {
        this.setState({retention: event.target.value});
    }
    changelimitHandler= (event) => {
        this.setState({limit: event.target.value});
    }
    changecessionPercentageHandler= (event) => {
        this.setState({cessionPercentage: event.target.value});
    }
    changeReinsuranceTypeHandler= (event) => {
        this.setState({reinsuranceType: event.target.value});
    }
    changeTreatyTypeHandler= (event) => {
        this.setState({treatyType: event.target.value});
    }

    cancel(){
        this.props.history.push('/reinsuranceAgreements');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add ReinsuranceAgreement</h3>
        }else{
            return <h3 className="text-center">Update ReinsuranceAgreement</h3>
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
                                            <label> agreementNumber:&emsp; </label>
                                                <input placeholder="agreementNumber" name="agreementNumber" className="form-control" value={this.state.agreementNumber} onChange={this.changeagreementNumberHandler}/>

                                            <label> effectivePeriod:&emsp; </label>
                                                <input placeholder="effectivePeriod" name="effectivePeriod" className="form-control" value={this.state.effectivePeriod} onChange={this.changeeffectivePeriodHandler}/>

                                            <label> retention:&emsp; </label>
                                                <input placeholder="retention" name="retention" className="form-control" value={this.state.retention} onChange={this.changeretentionHandler}/>

                                            <label> limit:&emsp; </label>
                                                <input placeholder="limit" name="limit" className="form-control" value={this.state.limit} onChange={this.changelimitHandler}/>

                                            <label> cessionPercentage:&emsp; </label>
                                                <input placeholder="cessionPercentage" name="cessionPercentage" className="form-control" value={this.state.cessionPercentage} onChange={this.changecessionPercentageHandler}/>

                                            <label> ReinsuranceType:&emsp; </label>
                                                <select value={this.state.reinsuranceType} onChange={this.changeReinsuranceTypeHandler}>
                      <option name="ReinsuranceType" className="form-control" >
                          Treaty
                      </option>
                      <option name="ReinsuranceType" className="form-control" >
                          Facultative
                      </option>
                    </select>

                                            <label> TreatyType:&emsp; </label>
                                                <select value={this.state.treatyType} onChange={this.changeTreatyTypeHandler}>
                      <option name="TreatyType" className="form-control" >
                          QuotaShare
                      </option>
                      <option name="TreatyType" className="form-control" >
                          Surplus
                      </option>
                      <option name="TreatyType" className="form-control" >
                          ExcessOfLoss
                      </option>
                      <option name="TreatyType" className="form-control" >
                          StopLoss
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateReinsuranceAgreement}>Save</button>
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

export default CreateReinsuranceAgreementComponent
