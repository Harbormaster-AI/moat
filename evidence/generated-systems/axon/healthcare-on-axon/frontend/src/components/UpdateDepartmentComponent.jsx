import React, { Component } from 'react'
import DepartmentService from '../services/DepartmentService';

class UpdateDepartmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                departmentType: ''
        }
        this.updateDepartment = this.updateDepartment.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeDepartmentTypeHandler = this.changeDepartmentTypeHandler.bind(this);
    }

    componentDidMount(){
        DepartmentService.getDepartmentById(this.state.id).then( (res) =>{
            let department = res.data;
            this.setState({
                name: department.name,
                departmentType: department.departmentType
            });
        });
    }

    updateDepartment = (e) => {
        e.preventDefault();
        let department = {
            departmentId: this.state.id,
            name: this.state.name,
            departmentType: this.state.departmentType
        };
        console.log('department => ' + JSON.stringify(department));
        console.log('id => ' + JSON.stringify(this.state.id));
        DepartmentService.updateDepartment(department).then( res => {
            this.props.history.push('/departments');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeDepartmentTypeHandler= (event) => {
        this.setState({departmentType: event.target.value});
    }

    cancel(){
        this.props.history.push('/departments');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Department</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> DepartmentType: </label>
                                                <select value={this.state.departmentType} onChange={this.changeDepartmentTypeHandler}>
                      <option name="DepartmentType" className="form-control" >
                          Emergency
                      </option>
                      <option name="DepartmentType" className="form-control" >
                          Cardiology
                      </option>
                      <option name="DepartmentType" className="form-control" >
                          Oncology
                      </option>
                      <option name="DepartmentType" className="form-control" >
                          Orthopedics
                      </option>
                      <option name="DepartmentType" className="form-control" >
                          Pediatrics
                      </option>
                      <option name="DepartmentType" className="form-control" >
                          Radiology
                      </option>
                      <option name="DepartmentType" className="form-control" >
                          Pathology
                      </option>
                      <option name="DepartmentType" className="form-control" >
                          Pharmacy
                      </option>
                      <option name="DepartmentType" className="form-control" >
                          IntensiveCare
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateDepartment}>Save</button>
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

export default UpdateDepartmentComponent
