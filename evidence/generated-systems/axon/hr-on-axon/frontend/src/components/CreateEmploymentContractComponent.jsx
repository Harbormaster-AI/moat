import React, { Component } from 'react'
import EmploymentContractService from '../services/EmploymentContractService';

class CreateEmploymentContractComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                contractNumber: '',
                startDate: '',
                endDate: '',
                workHoursPerWeek: '',
                employmentType: '',
                status: '',
                payFrequency: ''
        }
        this.changecontractNumberHandler = this.changecontractNumberHandler.bind(this);
        this.changestartDateHandler = this.changestartDateHandler.bind(this);
        this.changeendDateHandler = this.changeendDateHandler.bind(this);
        this.changeworkHoursPerWeekHandler = this.changeworkHoursPerWeekHandler.bind(this);
        this.changeEmploymentTypeHandler = this.changeEmploymentTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
        this.changePayFrequencyHandler = this.changePayFrequencyHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            EmploymentContractService.getEmploymentContractById(this.state.id).then( (res) =>{
                let employmentContract = res.data;
                this.setState({
                    contractNumber: employmentContract.contractNumber,
                    startDate: employmentContract.startDate,
                    endDate: employmentContract.endDate,
                    workHoursPerWeek: employmentContract.workHoursPerWeek,
                    employmentType: employmentContract.employmentType,
                    status: employmentContract.status,
                    payFrequency: employmentContract.payFrequency
                });
            });
        }        
    }
    saveOrUpdateEmploymentContract = (e) => {
        e.preventDefault();
        let employmentContract = {
                employmentContractId: this.state.id,
                contractNumber: this.state.contractNumber,
                startDate: this.state.startDate,
                endDate: this.state.endDate,
                workHoursPerWeek: this.state.workHoursPerWeek,
                employmentType: this.state.employmentType,
                status: this.state.status,
                payFrequency: this.state.payFrequency
            };
        console.log('employmentContract => ' + JSON.stringify(employmentContract));

        // step 5
        if(this.state.id === '_add'){
            employmentContract.employmentContractId=''
            EmploymentContractService.createEmploymentContract(employmentContract).then(res =>{
                this.props.history.push('/employmentContracts');
            });
        }else{
            EmploymentContractService.updateEmploymentContract(employmentContract).then( res => {
                this.props.history.push('/employmentContracts');
            });
        }
    }
    
    changecontractNumberHandler= (event) => {
        this.setState({contractNumber: event.target.value});
    }
    changestartDateHandler= (event) => {
        this.setState({startDate: event.target.value});
    }
    changeendDateHandler= (event) => {
        this.setState({endDate: event.target.value});
    }
    changeworkHoursPerWeekHandler= (event) => {
        this.setState({workHoursPerWeek: event.target.value});
    }
    changeEmploymentTypeHandler= (event) => {
        this.setState({employmentType: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }
    changePayFrequencyHandler= (event) => {
        this.setState({payFrequency: event.target.value});
    }

    cancel(){
        this.props.history.push('/employmentContracts');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add EmploymentContract</h3>
        }else{
            return <h3 className="text-center">Update EmploymentContract</h3>
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
                                            <label> contractNumber:&emsp; </label>
                                                <input placeholder="contractNumber" name="contractNumber" className="form-control" value={this.state.contractNumber} onChange={this.changecontractNumberHandler}/>

                                            <label> startDate:&emsp; </label>
                                                <input type="date" placeholder="startDate" name="startDate" className="form-control" value={this.state.startDate} onChange={this.changestartDateHandler}/>

                                            <label> endDate:&emsp; </label>
                                                <input type="date" placeholder="endDate" name="endDate" className="form-control" value={this.state.endDate} onChange={this.changeendDateHandler}/>

                                            <label> workHoursPerWeek:&emsp; </label>
                                                <input placeholder="workHoursPerWeek" name="workHoursPerWeek" className="form-control" value={this.state.workHoursPerWeek} onChange={this.changeworkHoursPerWeekHandler}/>

                                            <label> EmploymentType:&emsp; </label>
                                                <select value={this.state.employmentType} onChange={this.changeEmploymentTypeHandler}>
                      <option name="EmploymentType" className="form-control" >
                          FullTime
                      </option>
                      <option name="EmploymentType" className="form-control" >
                          PartTime
                      </option>
                      <option name="EmploymentType" className="form-control" >
                          Temporary
                      </option>
                      <option name="EmploymentType" className="form-control" >
                          Intern
                      </option>
                      <option name="EmploymentType" className="form-control" >
                          Contractor
                      </option>
                      <option name="EmploymentType" className="form-control" >
                          Seasonal
                      </option>
                    </select>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Suspended
                      </option>
                      <option name="Status" className="form-control" >
                          Expired
                      </option>
                      <option name="Status" className="form-control" >
                          Terminated
                      </option>
                    </select>

                                            <label> PayFrequency:&emsp; </label>
                                                <select value={this.state.payFrequency} onChange={this.changePayFrequencyHandler}>
                      <option name="PayFrequency" className="form-control" >
                          Weekly
                      </option>
                      <option name="PayFrequency" className="form-control" >
                          Biweekly
                      </option>
                      <option name="PayFrequency" className="form-control" >
                          Semimonthly
                      </option>
                      <option name="PayFrequency" className="form-control" >
                          Monthly
                      </option>
                      <option name="PayFrequency" className="form-control" >
                          Quarterly
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateEmploymentContract}>Save</button>
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

export default CreateEmploymentContractComponent
