import React, { Component } from 'react'
import ComplianceProgramService from '../services/ComplianceProgramService';

class UpdateComplianceProgramComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                framework: '',
                status: ''
        }
        this.updateComplianceProgram = this.updateComplianceProgram.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeframeworkHandler = this.changeframeworkHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        ComplianceProgramService.getComplianceProgramById(this.state.id).then( (res) =>{
            let complianceProgram = res.data;
            this.setState({
                name: complianceProgram.name,
                framework: complianceProgram.framework,
                status: complianceProgram.status
            });
        });
    }

    updateComplianceProgram = (e) => {
        e.preventDefault();
        let complianceProgram = {
            complianceProgramId: this.state.id,
            name: this.state.name,
            framework: this.state.framework,
            status: this.state.status
        };
        console.log('complianceProgram => ' + JSON.stringify(complianceProgram));
        console.log('id => ' + JSON.stringify(this.state.id));
        ComplianceProgramService.updateComplianceProgram(complianceProgram).then( res => {
            this.props.history.push('/compliancePrograms');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update ComplianceProgram</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> framework: </label>
                                                <input placeholder="framework" name="framework" className="form-control" value={this.state.framework} onChange={this.changeframeworkHandler}/>

                                            <label> Status: </label>
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
                                        <button className="btn btn-success" onClick={this.updateComplianceProgram}>Save</button>
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

export default UpdateComplianceProgramComponent
