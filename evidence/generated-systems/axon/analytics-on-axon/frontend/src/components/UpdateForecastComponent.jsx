import React, { Component } from 'react'
import ForecastService from '../services/ForecastService';

class UpdateForecastComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                horizon: '',
                granularity: ''
        }
        this.updateForecast = this.updateForecast.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changehorizonHandler = this.changehorizonHandler.bind(this);
        this.changeGranularityHandler = this.changeGranularityHandler.bind(this);
    }

    componentDidMount(){
        ForecastService.getForecastById(this.state.id).then( (res) =>{
            let forecast = res.data;
            this.setState({
                name: forecast.name,
                horizon: forecast.horizon,
                granularity: forecast.granularity
            });
        });
    }

    updateForecast = (e) => {
        e.preventDefault();
        let forecast = {
            forecastId: this.state.id,
            name: this.state.name,
            horizon: this.state.horizon,
            granularity: this.state.granularity
        };
        console.log('forecast => ' + JSON.stringify(forecast));
        console.log('id => ' + JSON.stringify(this.state.id));
        ForecastService.updateForecast(forecast).then( res => {
            this.props.history.push('/forecasts');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changehorizonHandler= (event) => {
        this.setState({horizon: event.target.value});
    }
    changeGranularityHandler= (event) => {
        this.setState({granularity: event.target.value});
    }

    cancel(){
        this.props.history.push('/forecasts');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Forecast</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> horizon: </label>
                                                <input type="number" placeholder="horizon" name="horizon" className="form-control" value={this.state.horizon} onChange={this.changehorizonHandler}/>

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
                                        <button className="btn btn-success" onClick={this.updateForecast}>Save</button>
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

export default UpdateForecastComponent
