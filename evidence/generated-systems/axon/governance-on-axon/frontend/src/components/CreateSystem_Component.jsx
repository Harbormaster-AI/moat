import React, { Component } from 'react'
import System_Service from '../services/System_Service';

class CreateSystem_Component extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                ownerDepartment: '',
                systemType: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeownerDepartmentHandler = this.changeownerDepartmentHandler.bind(this);
        this.changeSystemTypeHandler = this.changeSystemTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            System_Service.getSystem_ById(this.state.id).then( (res) =>{
                let system_ = res.data;
                this.setState({
                    name: system_.name,
                    ownerDepartment: system_.ownerDepartment,
                    systemType: system_.systemType
                });
            });
        }        
    }
    saveOrUpdateSystem_ = (e) => {
        e.preventDefault();
        let system_ = {
                system_Id: this.state.id,
                name: this.state.name,
                ownerDepartment: this.state.ownerDepartment,
                systemType: this.state.systemType
            };
        console.log('system_ => ' + JSON.stringify(system_));

        // step 5
        if(this.state.id === '_add'){
            system_.system_Id=''
            System_Service.createSystem_(system_).then(res =>{
                this.props.history.push('/system_s');
            });
        }else{
            System_Service.updateSystem_(system_).then( res => {
                this.props.history.push('/system_s');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add System_</h3>
        }else{
            return <h3 className="text-center">Update System_</h3>
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

                                            <label> ownerDepartment:&emsp; </label>
                                                <input placeholder="ownerDepartment" name="ownerDepartment" className="form-control" value={this.state.ownerDepartment} onChange={this.changeownerDepartmentHandler}/>

                                            <label> SystemType:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateSystem_}>Save</button>
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

export default CreateSystem_Component
