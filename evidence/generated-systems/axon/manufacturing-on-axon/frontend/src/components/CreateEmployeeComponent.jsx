import React, { Component } from 'react'
import EmployeeService from '../services/EmployeeService';

class CreateEmployeeComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                firstName: '',
                lastName: '',
                role: '',
                skillLevel: ''
        }
        this.changefirstNameHandler = this.changefirstNameHandler.bind(this);
        this.changelastNameHandler = this.changelastNameHandler.bind(this);
        this.changeRoleHandler = this.changeRoleHandler.bind(this);
        this.changeSkillLevelHandler = this.changeSkillLevelHandler.bind(this);
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
                    firstName: employee.firstName,
                    lastName: employee.lastName,
                    role: employee.role,
                    skillLevel: employee.skillLevel
                });
            });
        }        
    }
    saveOrUpdateEmployee = (e) => {
        e.preventDefault();
        let employee = {
                employeeId: this.state.id,
                firstName: this.state.firstName,
                lastName: this.state.lastName,
                role: this.state.role,
                skillLevel: this.state.skillLevel
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
    
    changefirstNameHandler= (event) => {
        this.setState({firstName: event.target.value});
    }
    changelastNameHandler= (event) => {
        this.setState({lastName: event.target.value});
    }
    changeRoleHandler= (event) => {
        this.setState({role: event.target.value});
    }
    changeSkillLevelHandler= (event) => {
        this.setState({skillLevel: event.target.value});
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
                                            <label> firstName:&emsp; </label>
                                                <input placeholder="firstName" name="firstName" className="form-control" value={this.state.firstName} onChange={this.changefirstNameHandler}/>

                                            <label> lastName:&emsp; </label>
                                                <input placeholder="lastName" name="lastName" className="form-control" value={this.state.lastName} onChange={this.changelastNameHandler}/>

                                            <label> Role:&emsp; </label>
                                                <select value={this.state.role} onChange={this.changeRoleHandler}>
                      <option name="Role" className="form-control" >
                          Operator
                      </option>
                      <option name="Role" className="form-control" >
                          Technician
                      </option>
                      <option name="Role" className="form-control" >
                          Supervisor
                      </option>
                      <option name="Role" className="form-control" >
                          Planner
                      </option>
                      <option name="Role" className="form-control" >
                          QualityEngineer
                      </option>
                      <option name="Role" className="form-control" >
                          Buyer
                      </option>
                    </select>

                                            <label> SkillLevel:&emsp; </label>
                                                <select value={this.state.skillLevel} onChange={this.changeSkillLevelHandler}>
                      <option name="SkillLevel" className="form-control" >
                          Novice
                      </option>
                      <option name="SkillLevel" className="form-control" >
                          Competent
                      </option>
                      <option name="SkillLevel" className="form-control" >
                          Proficient
                      </option>
                      <option name="SkillLevel" className="form-control" >
                          Expert
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
