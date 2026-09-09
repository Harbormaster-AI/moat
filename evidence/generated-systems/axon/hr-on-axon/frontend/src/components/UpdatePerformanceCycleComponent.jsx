import React, { Component } from 'react'
import PerformanceCycleService from '../services/PerformanceCycleService';

class UpdatePerformanceCycleComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                startDate: '',
                endDate: '',
                status: ''
        }
        this.updatePerformanceCycle = this.updatePerformanceCycle.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changestartDateHandler = this.changestartDateHandler.bind(this);
        this.changeendDateHandler = this.changeendDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        PerformanceCycleService.getPerformanceCycleById(this.state.id).then( (res) =>{
            let performanceCycle = res.data;
            this.setState({
                name: performanceCycle.name,
                startDate: performanceCycle.startDate,
                endDate: performanceCycle.endDate,
                status: performanceCycle.status
            });
        });
    }

    updatePerformanceCycle = (e) => {
        e.preventDefault();
        let performanceCycle = {
            performanceCycleId: this.state.id,
            name: this.state.name,
            startDate: this.state.startDate,
            endDate: this.state.endDate,
            status: this.state.status
        };
        console.log('performanceCycle => ' + JSON.stringify(performanceCycle));
        console.log('id => ' + JSON.stringify(this.state.id));
        PerformanceCycleService.updatePerformanceCycle(performanceCycle).then( res => {
            this.props.history.push('/performanceCycles');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changestartDateHandler= (event) => {
        this.setState({startDate: event.target.value});
    }
    changeendDateHandler= (event) => {
        this.setState({endDate: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/performanceCycles');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update PerformanceCycle</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> startDate: </label>
                                                <input type="date" placeholder="startDate" name="startDate" className="form-control" value={this.state.startDate} onChange={this.changestartDateHandler}/>

                                            <label> endDate: </label>
                                                <input type="date" placeholder="endDate" name="endDate" className="form-control" value={this.state.endDate} onChange={this.changeendDateHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Planned
                      </option>
                      <option name="Status" className="form-control" >
                          Open
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updatePerformanceCycle}>Save</button>
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

export default UpdatePerformanceCycleComponent
