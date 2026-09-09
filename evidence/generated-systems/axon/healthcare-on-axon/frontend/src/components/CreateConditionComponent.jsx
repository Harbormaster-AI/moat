import React, { Component } from 'react'
import ConditionService from '../services/ConditionService';

class CreateConditionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                code: '',
                onsetDate: '',
                abatementDate: '',
                clinicalStatus: '',
                verificationStatus: ''
        }
        this.changecodeHandler = this.changecodeHandler.bind(this);
        this.changeonsetDateHandler = this.changeonsetDateHandler.bind(this);
        this.changeabatementDateHandler = this.changeabatementDateHandler.bind(this);
        this.changeClinicalStatusHandler = this.changeClinicalStatusHandler.bind(this);
        this.changeVerificationStatusHandler = this.changeVerificationStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateCondition = (e) => {
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

        // step 5
        if(this.state.id === '_add'){
            condition.conditionId=''
            ConditionService.createCondition(condition).then(res =>{
                this.props.history.push('/conditions');
            });
        }else{
            ConditionService.updateCondition(condition).then( res => {
                this.props.history.push('/conditions');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Condition</h3>
        }else{
            return <h3 className="text-center">Update Condition</h3>
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
                                            <label> code:&emsp; </label>
                                                <input placeholder="code" name="code" className="form-control" value={this.state.code} onChange={this.changecodeHandler}/>

                                            <label> onsetDate:&emsp; </label>
                                                <input type="date" placeholder="onsetDate" name="onsetDate" className="form-control" value={this.state.onsetDate} onChange={this.changeonsetDateHandler}/>

                                            <label> abatementDate:&emsp; </label>
                                                <input type="date" placeholder="abatementDate" name="abatementDate" className="form-control" value={this.state.abatementDate} onChange={this.changeabatementDateHandler}/>

                                            <label> ClinicalStatus:&emsp; </label>
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

                                            <label> VerificationStatus:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateCondition}>Save</button>
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

export default CreateConditionComponent
