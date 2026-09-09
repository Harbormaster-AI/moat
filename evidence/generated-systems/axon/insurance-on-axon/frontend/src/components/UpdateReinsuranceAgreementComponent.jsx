import React, { Component } from 'react'
import ReinsuranceAgreementService from '../services/ReinsuranceAgreementService';

class UpdateReinsuranceAgreementComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                agreementNumber: '',
                effectivePeriod: '',
                retention: '',
                limit: '',
                cessionPercentage: '',
                reinsuranceType: '',
                treatyType: ''
        }
        this.updateReinsuranceAgreement = this.updateReinsuranceAgreement.bind(this);

        this.changeagreementNumberHandler = this.changeagreementNumberHandler.bind(this);
        this.changeeffectivePeriodHandler = this.changeeffectivePeriodHandler.bind(this);
        this.changeretentionHandler = this.changeretentionHandler.bind(this);
        this.changelimitHandler = this.changelimitHandler.bind(this);
        this.changecessionPercentageHandler = this.changecessionPercentageHandler.bind(this);
        this.changeReinsuranceTypeHandler = this.changeReinsuranceTypeHandler.bind(this);
        this.changeTreatyTypeHandler = this.changeTreatyTypeHandler.bind(this);
    }

    componentDidMount(){
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

    updateReinsuranceAgreement = (e) => {
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
        console.log('id => ' + JSON.stringify(this.state.id));
        ReinsuranceAgreementService.updateReinsuranceAgreement(reinsuranceAgreement).then( res => {
            this.props.history.push('/reinsuranceAgreements');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update ReinsuranceAgreement</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> agreementNumber: </label>
                                                <input placeholder="agreementNumber" name="agreementNumber" className="form-control" value={this.state.agreementNumber} onChange={this.changeagreementNumberHandler}/>

                                            <label> effectivePeriod: </label>
                                                <input placeholder="effectivePeriod" name="effectivePeriod" className="form-control" value={this.state.effectivePeriod} onChange={this.changeeffectivePeriodHandler}/>

                                            <label> retention: </label>
                                                <input placeholder="retention" name="retention" className="form-control" value={this.state.retention} onChange={this.changeretentionHandler}/>

                                            <label> limit: </label>
                                                <input placeholder="limit" name="limit" className="form-control" value={this.state.limit} onChange={this.changelimitHandler}/>

                                            <label> cessionPercentage: </label>
                                                <input placeholder="cessionPercentage" name="cessionPercentage" className="form-control" value={this.state.cessionPercentage} onChange={this.changecessionPercentageHandler}/>

                                            <label> ReinsuranceType: </label>
                                                <select value={this.state.reinsuranceType} onChange={this.changeReinsuranceTypeHandler}>
                      <option name="ReinsuranceType" className="form-control" >
                          Treaty
                      </option>
                      <option name="ReinsuranceType" className="form-control" >
                          Facultative
                      </option>
                    </select>

                                            <label> TreatyType: </label>
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
                                        <button className="btn btn-success" onClick={this.updateReinsuranceAgreement}>Save</button>
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

export default UpdateReinsuranceAgreementComponent
