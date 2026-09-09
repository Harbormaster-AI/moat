import React, { Component } from 'react'
import CycleCountService from '../services/CycleCountService';

class UpdateCycleCountComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                countNumber: '',
                scheduledDate: '',
                performedDate: '',
                approvedBy: '',
                status: ''
        }
        this.updateCycleCount = this.updateCycleCount.bind(this);

        this.changecountNumberHandler = this.changecountNumberHandler.bind(this);
        this.changescheduledDateHandler = this.changescheduledDateHandler.bind(this);
        this.changeperformedDateHandler = this.changeperformedDateHandler.bind(this);
        this.changeapprovedByHandler = this.changeapprovedByHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        CycleCountService.getCycleCountById(this.state.id).then( (res) =>{
            let cycleCount = res.data;
            this.setState({
                countNumber: cycleCount.countNumber,
                scheduledDate: cycleCount.scheduledDate,
                performedDate: cycleCount.performedDate,
                approvedBy: cycleCount.approvedBy,
                status: cycleCount.status
            });
        });
    }

    updateCycleCount = (e) => {
        e.preventDefault();
        let cycleCount = {
            cycleCountId: this.state.id,
            countNumber: this.state.countNumber,
            scheduledDate: this.state.scheduledDate,
            performedDate: this.state.performedDate,
            approvedBy: this.state.approvedBy,
            status: this.state.status
        };
        console.log('cycleCount => ' + JSON.stringify(cycleCount));
        console.log('id => ' + JSON.stringify(this.state.id));
        CycleCountService.updateCycleCount(cycleCount).then( res => {
            this.props.history.push('/cycleCounts');
        });
    }

    changecountNumberHandler= (event) => {
        this.setState({countNumber: event.target.value});
    }
    changescheduledDateHandler= (event) => {
        this.setState({scheduledDate: event.target.value});
    }
    changeperformedDateHandler= (event) => {
        this.setState({performedDate: event.target.value});
    }
    changeapprovedByHandler= (event) => {
        this.setState({approvedBy: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/cycleCounts');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update CycleCount</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> countNumber: </label>
                                                <input placeholder="countNumber" name="countNumber" className="form-control" value={this.state.countNumber} onChange={this.changecountNumberHandler}/>

                                            <label> scheduledDate: </label>
                                                <input type="date" placeholder="scheduledDate" name="scheduledDate" className="form-control" value={this.state.scheduledDate} onChange={this.changescheduledDateHandler}/>

                                            <label> performedDate: </label>
                                                <input type="date" placeholder="performedDate" name="performedDate" className="form-control" value={this.state.performedDate} onChange={this.changeperformedDateHandler}/>

                                            <label> approvedBy: </label>
                                                <input placeholder="approvedBy" name="approvedBy" className="form-control" value={this.state.approvedBy} onChange={this.changeapprovedByHandler}/>

                                            <label> Status: </label>
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
                          Posted
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateCycleCount}>Save</button>
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

export default UpdateCycleCountComponent
