import React, { Component } from 'react'
import EmploymentContractService from '../services/EmploymentContractService';

class UpdateEmploymentContractComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                contractNumber: '',
                startDate: '',
                endDate: '',
                workHoursPerWeek: '',
                employmentType: '',
                status: '',
                payFrequency: ''
        }
        this.updateEmploymentContract = this.updateEmploymentContract.bind(this);

        this.changecontractNumberHandler = this.changecontractNumberHandler.bind(this);
        this.changestartDateHandler = this.changestartDateHandler.bind(this);
        this.changeendDateHandler = this.changeendDateHandler.bind(this);
        this.changeworkHoursPerWeekHandler = this.changeworkHoursPerWeekHandler.bind(this);
        this.changeEmploymentTypeHandler = this.changeEmploymentTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
        this.changePayFrequencyHandler = this.changePayFrequencyHandler.bind(this);
    }

    componentDidMount(){
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

    updateEmploymentContract = (e) => {
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
        console.log('id => ' + JSON.stringify(this.state.id));
        EmploymentContractService.updateEmploymentContract(employmentContract).then( res => {
            this.props.history.push('/employmentContracts');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update EmploymentContract</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> contractNumber: </label>
                                                <input placeholder="contractNumber" name="contractNumber" className="form-control" value={this.state.contractNumber} onChange={this.changecontractNumberHandler}/>

                                            <label> startDate: </label>
                                                <input type="date" placeholder="startDate" name="startDate" className="form-control" value={this.state.startDate} onChange={this.changestartDateHandler}/>

                                            <label> endDate: </label>
                                                <input type="date" placeholder="endDate" name="endDate" className="form-control" value={this.state.endDate} onChange={this.changeendDateHandler}/>

                                            <label> workHoursPerWeek: </label>
                                                <input placeholder="workHoursPerWeek" name="workHoursPerWeek" className="form-control" value={this.state.workHoursPerWeek} onChange={this.changeworkHoursPerWeekHandler}/>

                                            <label> EmploymentType: </label>
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

                                            <label> Status: </label>
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

                                            <label> PayFrequency: </label>
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
                                        <button className="btn btn-success" onClick={this.updateEmploymentContract}>Save</button>
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

export default UpdateEmploymentContractComponent
