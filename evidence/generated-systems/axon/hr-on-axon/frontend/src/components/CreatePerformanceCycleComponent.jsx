import React, { Component } from 'react'
import PerformanceCycleService from '../services/PerformanceCycleService';

class CreatePerformanceCycleComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                startDate: '',
                endDate: '',
                status: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changestartDateHandler = this.changestartDateHandler.bind(this);
        this.changeendDateHandler = this.changeendDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdatePerformanceCycle = (e) => {
        e.preventDefault();
        let performanceCycle = {
                performanceCycleId: this.state.id,
                name: this.state.name,
                startDate: this.state.startDate,
                endDate: this.state.endDate,
                status: this.state.status
            };
        console.log('performanceCycle => ' + JSON.stringify(performanceCycle));

        // step 5
        if(this.state.id === '_add'){
            performanceCycle.performanceCycleId=''
            PerformanceCycleService.createPerformanceCycle(performanceCycle).then(res =>{
                this.props.history.push('/performanceCycles');
            });
        }else{
            PerformanceCycleService.updatePerformanceCycle(performanceCycle).then( res => {
                this.props.history.push('/performanceCycles');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add PerformanceCycle</h3>
        }else{
            return <h3 className="text-center">Update PerformanceCycle</h3>
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

                                            <label> startDate:&emsp; </label>
                                                <input type="date" placeholder="startDate" name="startDate" className="form-control" value={this.state.startDate} onChange={this.changestartDateHandler}/>

                                            <label> endDate:&emsp; </label>
                                                <input type="date" placeholder="endDate" name="endDate" className="form-control" value={this.state.endDate} onChange={this.changeendDateHandler}/>

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdatePerformanceCycle}>Save</button>
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

export default CreatePerformanceCycleComponent
