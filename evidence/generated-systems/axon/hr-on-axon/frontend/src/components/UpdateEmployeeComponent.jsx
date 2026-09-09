import React, { Component } from 'react'
import EmployeeService from '../services/EmployeeService';

class UpdateEmployeeComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                employeeNumber: '',
                name: '',
                workEmail: '',
                workPhone: '',
                dateOfHire: '',
                nationalId: '',
                status: ''
        }
        this.updateEmployee = this.updateEmployee.bind(this);

        this.changeemployeeNumberHandler = this.changeemployeeNumberHandler.bind(this);
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeworkEmailHandler = this.changeworkEmailHandler.bind(this);
        this.changeworkPhoneHandler = this.changeworkPhoneHandler.bind(this);
        this.changedateOfHireHandler = this.changedateOfHireHandler.bind(this);
        this.changenationalIdHandler = this.changenationalIdHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        EmployeeService.getEmployeeById(this.state.id).then( (res) =>{
            let employee = res.data;
            this.setState({
                employeeNumber: employee.employeeNumber,
                name: employee.name,
                workEmail: employee.workEmail,
                workPhone: employee.workPhone,
                dateOfHire: employee.dateOfHire,
                nationalId: employee.nationalId,
                status: employee.status
            });
        });
    }

    updateEmployee = (e) => {
        e.preventDefault();
        let employee = {
            employeeId: this.state.id,
            employeeNumber: this.state.employeeNumber,
            name: this.state.name,
            workEmail: this.state.workEmail,
            workPhone: this.state.workPhone,
            dateOfHire: this.state.dateOfHire,
            nationalId: this.state.nationalId,
            status: this.state.status
        };
        console.log('employee => ' + JSON.stringify(employee));
        console.log('id => ' + JSON.stringify(this.state.id));
        EmployeeService.updateEmployee(employee).then( res => {
            this.props.history.push('/employees');
        });
    }

    changeemployeeNumberHandler= (event) => {
        this.setState({employeeNumber: event.target.value});
    }
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeworkEmailHandler= (event) => {
        this.setState({workEmail: event.target.value});
    }
    changeworkPhoneHandler= (event) => {
        this.setState({workPhone: event.target.value});
    }
    changedateOfHireHandler= (event) => {
        this.setState({dateOfHire: event.target.value});
    }
    changenationalIdHandler= (event) => {
        this.setState({nationalId: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/employees');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Employee</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> employeeNumber: </label>
                                                <input placeholder="employeeNumber" name="employeeNumber" className="form-control" value={this.state.employeeNumber} onChange={this.changeemployeeNumberHandler}/>

                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> workEmail: </label>
                                                <input placeholder="workEmail" name="workEmail" className="form-control" value={this.state.workEmail} onChange={this.changeworkEmailHandler}/>

                                            <label> workPhone: </label>
                                                <input placeholder="workPhone" name="workPhone" className="form-control" value={this.state.workPhone} onChange={this.changeworkPhoneHandler}/>

                                            <label> dateOfHire: </label>
                                                <input type="date" placeholder="dateOfHire" name="dateOfHire" className="form-control" value={this.state.dateOfHire} onChange={this.changedateOfHireHandler}/>

                                            <label> nationalId: </label>
                                                <input placeholder="nationalId" name="nationalId" className="form-control" value={this.state.nationalId} onChange={this.changenationalIdHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          OnLeave
                      </option>
                      <option name="Status" className="form-control" >
                          Suspended
                      </option>
                      <option name="Status" className="form-control" >
                          Terminated
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateEmployee}>Save</button>
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

export default UpdateEmployeeComponent
