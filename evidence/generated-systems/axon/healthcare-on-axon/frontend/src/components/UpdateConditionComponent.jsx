import React, { Component } from 'react'
import ConditionService from '../services/ConditionService';

class UpdateConditionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                code: '',
                onsetDate: '',
                abatementDate: '',
                clinicalStatus: '',
                verificationStatus: ''
        }
        this.updateCondition = this.updateCondition.bind(this);

        this.changecodeHandler = this.changecodeHandler.bind(this);
        this.changeonsetDateHandler = this.changeonsetDateHandler.bind(this);
        this.changeabatementDateHandler = this.changeabatementDateHandler.bind(this);
        this.changeClinicalStatusHandler = this.changeClinicalStatusHandler.bind(this);
        this.changeVerificationStatusHandler = this.changeVerificationStatusHandler.bind(this);
    }

    componentDidMount(){
        ConditionService.getConditionById(this.state.id).then( (res) =>{
            let condition = res.data;
            this.setState({
                code: condition.code,
                onsetDate: condition.onsetDate,
                abatementDate: condition.abatementDate,
                clinicalStatus: condition.clinicalStatus,
                verificationStatus: condition.verificationStatus
            });
        });
    }

    updateCondition = (e) => {
        e.preventDefault();
        let condition = {
            conditionId: this.state.id,
            code: this.state.code,
            onsetDate: this.state.onsetDate,
            abatementDate: this.state.abatementDate,
            clinicalStatus: this.state.clinicalStatus,
            verificationStatus: this.state.verificationStatus
        };
        console.log('condition => ' + JSON.stringify(condition));
        console.log('id => ' + JSON.stringify(this.state.id));
        ConditionService.updateCondition(condition).then( res => {
            this.props.history.push('/conditions');
        });
    }

    changecodeHandler= (event) => {
        this.setState({code: event.target.value});
    }
    changeonsetDateHandler= (event) => {
        this.setState({onsetDate: event.target.value});
    }
    changeabatementDateHandler= (event) => {
        this.setState({abatementDate: event.target.value});
    }
    changeClinicalStatusHandler= (event) => {
        this.setState({clinicalStatus: event.target.value});
    }
    changeVerificationStatusHandler= (event) => {
        this.setState({verificationStatus: event.target.value});
    }

    cancel(){
        this.props.history.push('/conditions');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Condition</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> code: </label>
                                                <input placeholder="code" name="code" className="form-control" value={this.state.code} onChange={this.changecodeHandler}/>

                                            <label> onsetDate: </label>
                                                <input type="date" placeholder="onsetDate" name="onsetDate" className="form-control" value={this.state.onsetDate} onChange={this.changeonsetDateHandler}/>

                                            <label> abatementDate: </label>
                                                <input type="date" placeholder="abatementDate" name="abatementDate" className="form-control" value={this.state.abatementDate} onChange={this.changeabatementDateHandler}/>

                                            <label> ClinicalStatus: </label>
                                                <select value={this.state.clinicalStatus} onChange={this.changeClinicalStatusHandler}>
                      <option name="ClinicalStatus" className="form-control" >
                          Active
                      </option>
                      <option name="ClinicalStatus" className="form-control" >
                          Recurrence
                      </option>
                      <option name="ClinicalStatus" className="form-control" >
                          Relapse
                      </option>
                      <option name="ClinicalStatus" className="form-control" >
                          Remission
                      </option>
                      <option name="ClinicalStatus" className="form-control" >
                          Resolved
                      </option>
                    </select>

                                            <label> VerificationStatus: </label>
                                                <select value={this.state.verificationStatus} onChange={this.changeVerificationStatusHandler}>
                      <option name="VerificationStatus" className="form-control" >
                          Suspected
                      </option>
                      <option name="VerificationStatus" className="form-control" >
                          Presumptive
                      </option>
                      <option name="VerificationStatus" className="form-control" >
                          Confirmed
                      </option>
                      <option name="VerificationStatus" className="form-control" >
                          RuledOut
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateCondition}>Save</button>
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

export default UpdateConditionComponent
