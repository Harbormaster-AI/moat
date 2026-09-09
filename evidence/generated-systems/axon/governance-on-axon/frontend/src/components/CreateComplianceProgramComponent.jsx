import React, { Component } from 'react'
import ComplianceProgramService from '../services/ComplianceProgramService';

class CreateComplianceProgramComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                framework: '',
                status: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeframeworkHandler = this.changeframeworkHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ComplianceProgramService.getComplianceProgramById(this.state.id).then( (res) =>{
                let complianceProgram = res.data;
                this.setState({
                    name: complianceProgram.name,
                    framework: complianceProgram.framework,
                    status: complianceProgram.status
                });
            });
        }        
    }
    saveOrUpdateComplianceProgram = (e) => {
        e.preventDefault();
        let complianceProgram = {
                complianceProgramId: this.state.id,
                name: this.state.name,
                framework: this.state.framework,
                status: this.state.status
            };
        console.log('complianceProgram => ' + JSON.stringify(complianceProgram));

        // step 5
        if(this.state.id === '_add'){
            complianceProgram.complianceProgramId=''
            ComplianceProgramService.createComplianceProgram(complianceProgram).then(res =>{
                this.props.history.push('/compliancePrograms');
            });
        }else{
            ComplianceProgramService.updateComplianceProgram(complianceProgram).then( res => {
                this.props.history.push('/compliancePrograms');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeframeworkHandler= (event) => {
        this.setState({framework: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/compliancePrograms');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add ComplianceProgram</h3>
        }else{
            return <h3 className="text-center">Update ComplianceProgram</h3>
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

                                            <label> framework:&emsp; </label>
                                                <input placeholder="framework" name="framework" className="form-control" value={this.state.framework} onChange={this.changeframeworkHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          NotStarted
                      </option>
                      <option name="Status" className="form-control" >
                          InProgress
                      </option>
                      <option name="Status" className="form-control" >
                          Compliant
                      </option>
                      <option name="Status" className="form-control" >
                          NonCompliant
                      </option>
                      <option name="Status" className="form-control" >
                          Waived
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateComplianceProgram}>Save</button>
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

export default CreateComplianceProgramComponent
