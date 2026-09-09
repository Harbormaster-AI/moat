import React, { Component } from 'react'
import System_Service from '../services/System_Service';

class UpdateSystem_Component extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                ownerDepartment: '',
                systemType: ''
        }
        this.updateSystem_ = this.updateSystem_.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeownerDepartmentHandler = this.changeownerDepartmentHandler.bind(this);
        this.changeSystemTypeHandler = this.changeSystemTypeHandler.bind(this);
    }

    componentDidMount(){
        System_Service.getSystem_ById(this.state.id).then( (res) =>{
            let system_ = res.data;
            this.setState({
                name: system_.name,
                ownerDepartment: system_.ownerDepartment,
                systemType: system_.systemType
            });
        });
    }

    updateSystem_ = (e) => {
        e.preventDefault();
        let system_ = {
            system_Id: this.state.id,
            name: this.state.name,
            ownerDepartment: this.state.ownerDepartment,
            systemType: this.state.systemType
        };
        console.log('system_ => ' + JSON.stringify(system_));
        console.log('id => ' + JSON.stringify(this.state.id));
        System_Service.updateSystem_(system_).then( res => {
            this.props.history.push('/system_s');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeownerDepartmentHandler= (event) => {
        this.setState({ownerDepartment: event.target.value});
    }
    changeSystemTypeHandler= (event) => {
        this.setState({systemType: event.target.value});
    }

    cancel(){
        this.props.history.push('/system_s');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update System_</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> ownerDepartment: </label>
                                                <input placeholder="ownerDepartment" name="ownerDepartment" className="form-control" value={this.state.ownerDepartment} onChange={this.changeownerDepartmentHandler}/>

                                            <label> SystemType: </label>
                                                <select value={this.state.systemType} onChange={this.changeSystemTypeHandler}>
                      <option name="SystemType" className="form-control" >
                          Application
                      </option>
                      <option name="SystemType" className="form-control" >
                          Database
                      </option>
                      <option name="SystemType" className="form-control" >
                          DataWarehouse
                      </option>
                      <option name="SystemType" className="form-control" >
                          SaaS
                      </option>
                      <option name="SystemType" className="form-control" >
                          Infrastructure
                      </option>
                      <option name="SystemType" className="form-control" >
                          Endpoint
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateSystem_}>Save</button>
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

export default UpdateSystem_Component
