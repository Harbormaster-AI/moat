import React, { Component } from 'react'
import EmployeeService from '../services/EmployeeService';

class CreateEmployeeComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                employeeNumber: '',
                name: '',
                workEmail: '',
                workPhone: '',
                dateOfHire: '',
                nationalId: '',
                status: ''
        }
        this.changeemployeeNumberHandler = this.changeemployeeNumberHandler.bind(this);
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeworkEmailHandler = this.changeworkEmailHandler.bind(this);
        this.changeworkPhoneHandler = this.changeworkPhoneHandler.bind(this);
        this.changedateOfHireHandler = this.changedateOfHireHandler.bind(this);
        this.changenationalIdHandler = this.changenationalIdHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateEmployee = (e) => {
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

        // step 5
        if(this.state.id === '_add'){
            employee.employeeId=''
            EmployeeService.createEmployee(employee).then(res =>{
                this.props.history.push('/employees');
            });
        }else{
            EmployeeService.updateEmployee(employee).then( res => {
                this.props.history.push('/employees');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Employee</h3>
        }else{
            return <h3 className="text-center">Update Employee</h3>
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
                                            <label> employeeNumber:&emsp; </label>
                                                <input placeholder="employeeNumber" name="employeeNumber" className="form-control" value={this.state.employeeNumber} onChange={this.changeemployeeNumberHandler}/>

                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> workEmail:&emsp; </label>
                                                <input placeholder="workEmail" name="workEmail" className="form-control" value={this.state.workEmail} onChange={this.changeworkEmailHandler}/>

                                            <label> workPhone:&emsp; </label>
                                                <input placeholder="workPhone" name="workPhone" className="form-control" value={this.state.workPhone} onChange={this.changeworkPhoneHandler}/>

                                            <label> dateOfHire:&emsp; </label>
                                                <input type="date" placeholder="dateOfHire" name="dateOfHire" className="form-control" value={this.state.dateOfHire} onChange={this.changedateOfHireHandler}/>

                                            <label> nationalId:&emsp; </label>
                                                <input placeholder="nationalId" name="nationalId" className="form-control" value={this.state.nationalId} onChange={this.changenationalIdHandler}/>

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateEmployee}>Save</button>
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

export default CreateEmployeeComponent
