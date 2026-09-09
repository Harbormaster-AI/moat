import React, { Component } from 'react'
import ProcedureService from '../services/ProcedureService';

class CreateProcedureComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                procedureCode: '',
                startDateTime: '',
                endDateTime: '',
                status: ''
        }
        this.changeprocedureCodeHandler = this.changeprocedureCodeHandler.bind(this);
        this.changestartDateTimeHandler = this.changestartDateTimeHandler.bind(this);
        this.changeendDateTimeHandler = this.changeendDateTimeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ProcedureService.getProcedureById(this.state.id).then( (res) =>{
                let procedure = res.data;
                this.setState({
                    procedureCode: procedure.procedureCode,
                    startDateTime: procedure.startDateTime,
                    endDateTime: procedure.endDateTime,
                    status: procedure.status
                });
            });
        }        
    }
    saveOrUpdateProcedure = (e) => {
        e.preventDefault();
        let procedure = {
                procedureId: this.state.id,
                procedureCode: this.state.procedureCode,
                startDateTime: this.state.startDateTime,
                endDateTime: this.state.endDateTime,
                status: this.state.status
            };
        console.log('procedure => ' + JSON.stringify(procedure));

        // step 5
        if(this.state.id === '_add'){
            procedure.procedureId=''
            ProcedureService.createProcedure(procedure).then(res =>{
                this.props.history.push('/procedures');
            });
        }else{
            ProcedureService.updateProcedure(procedure).then( res => {
                this.props.history.push('/procedures');
            });
        }
    }
    
    changeprocedureCodeHandler= (event) => {
        this.setState({procedureCode: event.target.value});
    }
    changestartDateTimeHandler= (event) => {
        this.setState({startDateTime: event.target.value});
    }
    changeendDateTimeHandler= (event) => {
        this.setState({endDateTime: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/procedures');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Procedure</h3>
        }else{
            return <h3 className="text-center">Update Procedure</h3>
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
                                            <label> procedureCode:&emsp; </label>
                                                <input placeholder="procedureCode" name="procedureCode" className="form-control" value={this.state.procedureCode} onChange={this.changeprocedureCodeHandler}/>

                                            <label> startDateTime:&emsp; </label>
                                                <input type="time" placeholder="startDateTime" name="startDateTime" className="form-control" value={this.state.startDateTime} onChange={this.changestartDateTimeHandler}/>

                                            <label> endDateTime:&emsp; </label>
                                                <input type="time" placeholder="endDateTime" name="endDateTime" className="form-control" value={this.state.endDateTime} onChange={this.changeendDateTimeHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Planned
                      </option>
                      <option name="Status" className="form-control" >
                          InProgress
                      </option>
                      <option name="Status" className="form-control" >
                          Completed
                      </option>
                      <option name="Status" className="form-control" >
                          Aborted
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateProcedure}>Save</button>
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

export default CreateProcedureComponent
