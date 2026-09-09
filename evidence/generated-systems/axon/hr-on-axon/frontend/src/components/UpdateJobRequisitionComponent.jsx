import React, { Component } from 'react'
import JobRequisitionService from '../services/JobRequisitionService';

class UpdateJobRequisitionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                requisitionNumber: '',
                title: '',
                openings: '',
                targetStartDate: '',
                status: '',
                priority: ''
        }
        this.updateJobRequisition = this.updateJobRequisition.bind(this);

        this.changerequisitionNumberHandler = this.changerequisitionNumberHandler.bind(this);
        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changeopeningsHandler = this.changeopeningsHandler.bind(this);
        this.changetargetStartDateHandler = this.changetargetStartDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
        this.changePriorityHandler = this.changePriorityHandler.bind(this);
    }

    componentDidMount(){
        JobRequisitionService.getJobRequisitionById(this.state.id).then( (res) =>{
            let jobRequisition = res.data;
            this.setState({
                requisitionNumber: jobRequisition.requisitionNumber,
                title: jobRequisition.title,
                openings: jobRequisition.openings,
                targetStartDate: jobRequisition.targetStartDate,
                status: jobRequisition.status,
                priority: jobRequisition.priority
            });
        });
    }

    updateJobRequisition = (e) => {
        e.preventDefault();
        let jobRequisition = {
            jobRequisitionId: this.state.id,
            requisitionNumber: this.state.requisitionNumber,
            title: this.state.title,
            openings: this.state.openings,
            targetStartDate: this.state.targetStartDate,
            status: this.state.status,
            priority: this.state.priority
        };
        console.log('jobRequisition => ' + JSON.stringify(jobRequisition));
        console.log('id => ' + JSON.stringify(this.state.id));
        JobRequisitionService.updateJobRequisition(jobRequisition).then( res => {
            this.props.history.push('/jobRequisitions');
        });
    }

    changerequisitionNumberHandler= (event) => {
        this.setState({requisitionNumber: event.target.value});
    }
    changetitleHandler= (event) => {
        this.setState({title: event.target.value});
    }
    changeopeningsHandler= (event) => {
        this.setState({openings: event.target.value});
    }
    changetargetStartDateHandler= (event) => {
        this.setState({targetStartDate: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }
    changePriorityHandler= (event) => {
        this.setState({priority: event.target.value});
    }

    cancel(){
        this.props.history.push('/jobRequisitions');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update JobRequisition</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> requisitionNumber: </label>
                                                <input placeholder="requisitionNumber" name="requisitionNumber" className="form-control" value={this.state.requisitionNumber} onChange={this.changerequisitionNumberHandler}/>

                                            <label> title: </label>
                                                <input placeholder="title" name="title" className="form-control" value={this.state.title} onChange={this.changetitleHandler}/>

                                            <label> openings: </label>
                                                <input type="number" placeholder="openings" name="openings" className="form-control" value={this.state.openings} onChange={this.changeopeningsHandler}/>

                                            <label> targetStartDate: </label>
                                                <input type="date" placeholder="targetStartDate" name="targetStartDate" className="form-control" value={this.state.targetStartDate} onChange={this.changetargetStartDateHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Open
                      </option>
                      <option name="Status" className="form-control" >
                          OnHold
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                    </select>

                                            <label> Priority: </label>
                                                <select value={this.state.priority} onChange={this.changePriorityHandler}>
                      <option name="Priority" className="form-control" >
                          Low
                      </option>
                      <option name="Priority" className="form-control" >
                          Medium
                      </option>
                      <option name="Priority" className="form-control" >
                          High
                      </option>
                      <option name="Priority" className="form-control" >
                          Critical
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateJobRequisition}>Save</button>
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

export default UpdateJobRequisitionComponent
