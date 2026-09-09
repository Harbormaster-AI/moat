import React, { Component } from 'react'
import TimeSeriesService from '../services/TimeSeriesService';

class CreateTimeSeriesComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                timezone: '',
                granularity: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changetimezoneHandler = this.changetimezoneHandler.bind(this);
        this.changeGranularityHandler = this.changeGranularityHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            TimeSeriesService.getTimeSeriesById(this.state.id).then( (res) =>{
                let timeSeries = res.data;
                this.setState({
                    name: timeSeries.name,
                    timezone: timeSeries.timezone,
                    granularity: timeSeries.granularity
                });
            });
        }        
    }
    saveOrUpdateTimeSeries = (e) => {
        e.preventDefault();
        let timeSeries = {
                timeSeriesId: this.state.id,
                name: this.state.name,
                timezone: this.state.timezone,
                granularity: this.state.granularity
            };
        console.log('timeSeries => ' + JSON.stringify(timeSeries));

        // step 5
        if(this.state.id === '_add'){
            timeSeries.timeSeriesId=''
            TimeSeriesService.createTimeSeries(timeSeries).then(res =>{
                this.props.history.push('/timeSeriess');
            });
        }else{
            TimeSeriesService.updateTimeSeries(timeSeries).then( res => {
                this.props.history.push('/timeSeriess');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changetimezoneHandler= (event) => {
        this.setState({timezone: event.target.value});
    }
    changeGranularityHandler= (event) => {
        this.setState({granularity: event.target.value});
    }

    cancel(){
        this.props.history.push('/timeSeriess');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add TimeSeries</h3>
        }else{
            return <h3 className="text-center">Update TimeSeries</h3>
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

                                            <label> timezone:&emsp; </label>
                                                <input placeholder="timezone" name="timezone" className="form-control" value={this.state.timezone} onChange={this.changetimezoneHandler}/>

                                            <label> Granularity:&emsp; </label>
                                                <select value={this.state.granularity} onChange={this.changeGranularityHandler}>
                      <option name="Granularity" className="form-control" >
                          Minute
                      </option>
                      <option name="Granularity" className="form-control" >
                          Hour
                      </option>
                      <option name="Granularity" className="form-control" >
                          Day
                      </option>
                      <option name="Granularity" className="form-control" >
                          Week
                      </option>
                      <option name="Granularity" className="form-control" >
                          Month
                      </option>
                      <option name="Granularity" className="form-control" >
                          Quarter
                      </option>
                      <option name="Granularity" className="form-control" >
                          Year
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateTimeSeries}>Save</button>
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

export default CreateTimeSeriesComponent
