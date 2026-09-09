import React, { Component } from 'react'
import DepartmentService from '../services/DepartmentService';

class UpdateDepartmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                code: ''
        }
        this.updateDepartment = this.updateDepartment.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changecodeHandler = this.changecodeHandler.bind(this);
    }

    componentDidMount(){
        DepartmentService.getDepartmentById(this.state.id).then( (res) =>{
            let department = res.data;
            this.setState({
                name: department.name,
                code: department.code
            });
        });
    }

    updateDepartment = (e) => {
        e.preventDefault();
        let department = {
            departmentId: this.state.id,
            name: this.state.name,
            code: this.state.code
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
    changecodeHandler= (event) => {
        this.setState({code: event.target.value});
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

                                            <label> code: </label>
                                                <input placeholder="code" name="code" className="form-control" value={this.state.code} onChange={this.changecodeHandler}/>

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
