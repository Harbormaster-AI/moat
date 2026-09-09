import React, { Component } from 'react'
import AuditProgramService from '../services/AuditProgramService';

class UpdateAuditProgramComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                scope: '',
                cycle: '',
                status: ''
        }
        this.updateAuditProgram = this.updateAuditProgram.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changescopeHandler = this.changescopeHandler.bind(this);
        this.changeCycleHandler = this.changeCycleHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        AuditProgramService.getAuditProgramById(this.state.id).then( (res) =>{
            let auditProgram = res.data;
            this.setState({
                name: auditProgram.name,
                scope: auditProgram.scope,
                cycle: auditProgram.cycle,
                status: auditProgram.status
            });
        });
    }

    updateAuditProgram = (e) => {
        e.preventDefault();
        let auditProgram = {
            auditProgramId: this.state.id,
            name: this.state.name,
            scope: this.state.scope,
            cycle: this.state.cycle,
            status: this.state.status
        };
        console.log('auditProgram => ' + JSON.stringify(auditProgram));
        console.log('id => ' + JSON.stringify(this.state.id));
        AuditProgramService.updateAuditProgram(auditProgram).then( res => {
            this.props.history.push('/auditPrograms');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changescopeHandler= (event) => {
        this.setState({scope: event.target.value});
    }
    changeCycleHandler= (event) => {
        this.setState({cycle: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/auditPrograms');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update AuditProgram</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> scope: </label>
                                                <input placeholder="scope" name="scope" className="form-control" value={this.state.scope} onChange={this.changescopeHandler}/>

                                            <label> Cycle: </label>
                                                <select value={this.state.cycle} onChange={this.changeCycleHandler}>
                      <option name="Cycle" className="form-control" >
                          Annual
                      </option>
                      <option name="Cycle" className="form-control" >
                          SemiAnnual
                      </option>
                      <option name="Cycle" className="form-control" >
                          Quarterly
                      </option>
                      <option name="Cycle" className="form-control" >
                          Continuous
                      </option>
                      <option name="Cycle" className="form-control" >
                          OneTime
                      </option>
                    </select>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Planned
                      </option>
                      <option name="Status" className="form-control" >
                          Fieldwork
                      </option>
                      <option name="Status" className="form-control" >
                          Reporting
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
                      </option>
                      <option name="Status" className="form-control" >
                          OnHold
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateAuditProgram}>Save</button>
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

export default UpdateAuditProgramComponent
