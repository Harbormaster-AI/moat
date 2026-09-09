import React, { Component } from 'react'
import DepartmentService from '../services/DepartmentService';

class CreateDepartmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                departmentType: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeDepartmentTypeHandler = this.changeDepartmentTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            DepartmentService.getDepartmentById(this.state.id).then( (res) =>{
                let department = res.data;
                this.setState({
                    name: department.name,
                    departmentType: department.departmentType
                });
            });
        }        
    }
    saveOrUpdateDepartment = (e) => {
        e.preventDefault();
        let department = {
                departmentId: this.state.id,
                name: this.state.name,
                departmentType: this.state.departmentType
            };
        console.log('department => ' + JSON.stringify(department));

        // step 5
        if(this.state.id === '_add'){
            department.departmentId=''
            DepartmentService.createDepartment(department).then(res =>{
                this.props.history.push('/departments');
            });
        }else{
            DepartmentService.updateDepartment(department).then( res => {
                this.props.history.push('/departments');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Department</h3>
        }else{
            return <h3 className="text-center">Update Department</h3>
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
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> DepartmentType:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateDepartment}>Save</button>
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

export default CreateDepartmentComponent
