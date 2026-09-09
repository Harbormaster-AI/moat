import React, { Component } from 'react'
import TimeSeriesService from '../services/TimeSeriesService';

class UpdateTimeSeriesComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                timezone: '',
                granularity: ''
        }
        this.updateTimeSeries = this.updateTimeSeries.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changetimezoneHandler = this.changetimezoneHandler.bind(this);
        this.changeGranularityHandler = this.changeGranularityHandler.bind(this);
    }

    componentDidMount(){
        TimeSeriesService.getTimeSeriesById(this.state.id).then( (res) =>{
            let timeSeries = res.data;
            this.setState({
                name: timeSeries.name,
                timezone: timeSeries.timezone,
                granularity: timeSeries.granularity
            });
        });
    }

    updateTimeSeries = (e) => {
        e.preventDefault();
        let timeSeries = {
            timeSeriesId: this.state.id,
            name: this.state.name,
            timezone: this.state.timezone,
            granularity: this.state.granularity
        };
        console.log('timeSeries => ' + JSON.stringify(timeSeries));
        console.log('id => ' + JSON.stringify(this.state.id));
        TimeSeriesService.updateTimeSeries(timeSeries).then( res => {
            this.props.history.push('/timeSeriess');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update TimeSeries</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> timezone: </label>
                                                <input placeholder="timezone" name="timezone" className="form-control" value={this.state.timezone} onChange={this.changetimezoneHandler}/>

                                            <label> Granularity: </label>
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
                                        <button className="btn btn-success" onClick={this.updateTimeSeries}>Save</button>
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

export default UpdateTimeSeriesComponent
