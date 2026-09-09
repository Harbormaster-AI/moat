import React, { Component } from 'react'
import AccessPolicyService from '../services/AccessPolicyService';

class UpdateAccessPolicyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                subjectName: '',
                accessLevel: '',
                subjectType: ''
        }
        this.updateAccessPolicy = this.updateAccessPolicy.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changesubjectNameHandler = this.changesubjectNameHandler.bind(this);
        this.changeAccessLevelHandler = this.changeAccessLevelHandler.bind(this);
        this.changeSubjectTypeHandler = this.changeSubjectTypeHandler.bind(this);
    }

    componentDidMount(){
        AccessPolicyService.getAccessPolicyById(this.state.id).then( (res) =>{
            let accessPolicy = res.data;
            this.setState({
                name: accessPolicy.name,
                subjectName: accessPolicy.subjectName,
                accessLevel: accessPolicy.accessLevel,
                subjectType: accessPolicy.subjectType
            });
        });
    }

    updateAccessPolicy = (e) => {
        e.preventDefault();
        let accessPolicy = {
            accessPolicyId: this.state.id,
            name: this.state.name,
            subjectName: this.state.subjectName,
            accessLevel: this.state.accessLevel,
            subjectType: this.state.subjectType
        };
        console.log('accessPolicy => ' + JSON.stringify(accessPolicy));
        console.log('id => ' + JSON.stringify(this.state.id));
        AccessPolicyService.updateAccessPolicy(accessPolicy).then( res => {
            this.props.history.push('/accessPolicys');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changesubjectNameHandler= (event) => {
        this.setState({subjectName: event.target.value});
    }
    changeAccessLevelHandler= (event) => {
        this.setState({accessLevel: event.target.value});
    }
    changeSubjectTypeHandler= (event) => {
        this.setState({subjectType: event.target.value});
    }

    cancel(){
        this.props.history.push('/accessPolicys');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update AccessPolicy</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> subjectName: </label>
                                                <input placeholder="subjectName" name="subjectName" className="form-control" value={this.state.subjectName} onChange={this.changesubjectNameHandler}/>

                                            <label> AccessLevel: </label>
                                                <select value={this.state.accessLevel} onChange={this.changeAccessLevelHandler}>
                      <option name="AccessLevel" className="form-control" >
                          View
                      </option>
                      <option name="AccessLevel" className="form-control" >
                          Query
                      </option>
                      <option name="AccessLevel" className="form-control" >
                          Modify
                      </option>
                      <option name="AccessLevel" className="form-control" >
                          Admin
                      </option>
                    </select>

                                            <label> SubjectType: </label>
                                                <select value={this.state.subjectType} onChange={this.changeSubjectTypeHandler}>
                      <option name="SubjectType" className="form-control" >
                          User
                      </option>
                      <option name="SubjectType" className="form-control" >
                          Group
                      </option>
                      <option name="SubjectType" className="form-control" >
                          Service
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateAccessPolicy}>Save</button>
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

export default UpdateAccessPolicyComponent
