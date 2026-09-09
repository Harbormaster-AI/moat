import React, { Component } from 'react'
import RetentionScheduleService from '../services/RetentionScheduleService';

class CreateRetentionScheduleComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                retentionPeriodMonths: '',
                retentionTrigger: '',
                dispositionAction: '',
                status: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeretentionPeriodMonthsHandler = this.changeretentionPeriodMonthsHandler.bind(this);
        this.changeRetentionTriggerHandler = this.changeRetentionTriggerHandler.bind(this);
        this.changeDispositionActionHandler = this.changeDispositionActionHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            RetentionScheduleService.getRetentionScheduleById(this.state.id).then( (res) =>{
                let retentionSchedule = res.data;
                this.setState({
                    name: retentionSchedule.name,
                    retentionPeriodMonths: retentionSchedule.retentionPeriodMonths,
                    retentionTrigger: retentionSchedule.retentionTrigger,
                    dispositionAction: retentionSchedule.dispositionAction,
                    status: retentionSchedule.status
                });
            });
        }        
    }
    saveOrUpdateRetentionSchedule = (e) => {
        e.preventDefault();
        let retentionSchedule = {
                retentionScheduleId: this.state.id,
                name: this.state.name,
                retentionPeriodMonths: this.state.retentionPeriodMonths,
                retentionTrigger: this.state.retentionTrigger,
                dispositionAction: this.state.dispositionAction,
                status: this.state.status
            };
        console.log('retentionSchedule => ' + JSON.stringify(retentionSchedule));

        // step 5
        if(this.state.id === '_add'){
            retentionSchedule.retentionScheduleId=''
            RetentionScheduleService.createRetentionSchedule(retentionSchedule).then(res =>{
                this.props.history.push('/retentionSchedules');
            });
        }else{
            RetentionScheduleService.updateRetentionSchedule(retentionSchedule).then( res => {
                this.props.history.push('/retentionSchedules');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeretentionPeriodMonthsHandler= (event) => {
        this.setState({retentionPeriodMonths: event.target.value});
    }
    changeRetentionTriggerHandler= (event) => {
        this.setState({retentionTrigger: event.target.value});
    }
    changeDispositionActionHandler= (event) => {
        this.setState({dispositionAction: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/retentionSchedules');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add RetentionSchedule</h3>
        }else{
            return <h3 className="text-center">Update RetentionSchedule</h3>
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

                                            <label> retentionPeriodMonths:&emsp; </label>
                                                <input type="number" placeholder="retentionPeriodMonths" name="retentionPeriodMonths" className="form-control" value={this.state.retentionPeriodMonths} onChange={this.changeretentionPeriodMonthsHandler}/>

                                            <label> RetentionTrigger:&emsp; </label>
                                                <select value={this.state.retentionTrigger} onChange={this.changeRetentionTriggerHandler}>
                      <option name="RetentionTrigger" className="form-control" >
                          CreationDate
                      </option>
                      <option name="RetentionTrigger" className="form-control" >
                          LastModified
                      </option>
                      <option name="RetentionTrigger" className="form-control" >
                          Termination
                      </option>
                      <option name="RetentionTrigger" className="form-control" >
                          ContractEnd
                      </option>
                      <option name="RetentionTrigger" className="form-control" >
                          EventCompletion
                      </option>
                      <option name="RetentionTrigger" className="form-control" >
                          FiscalYearEnd
                      </option>
                    </select>

                                            <label> DispositionAction:&emsp; </label>
                                                <select value={this.state.dispositionAction} onChange={this.changeDispositionActionHandler}>
                      <option name="DispositionAction" className="form-control" >
                          Destroy
                      </option>
                      <option name="DispositionAction" className="form-control" >
                          TransferToArchive
                      </option>
                      <option name="DispositionAction" className="form-control" >
                          Review
                      </option>
                      <option name="DispositionAction" className="form-control" >
                          SecureDelete
                      </option>
                      <option name="DispositionAction" className="form-control" >
                          ReturnToOwner
                      </option>
                    </select>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Approved
                      </option>
                      <option name="Status" className="form-control" >
                          InEffect
                      </option>
                      <option name="Status" className="form-control" >
                          Suspended
                      </option>
                      <option name="Status" className="form-control" >
                          Retired
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateRetentionSchedule}>Save</button>
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

export default CreateRetentionScheduleComponent
