import React, { Component } from 'react'
import ForecastService from '../services/ForecastService';

class UpdateForecastComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                forecastNumber: '',
                forecastHorizonStart: '',
                forecastHorizonEnd: '',
                method: ''
        }
        this.updateForecast = this.updateForecast.bind(this);

        this.changeforecastNumberHandler = this.changeforecastNumberHandler.bind(this);
        this.changeforecastHorizonStartHandler = this.changeforecastHorizonStartHandler.bind(this);
        this.changeforecastHorizonEndHandler = this.changeforecastHorizonEndHandler.bind(this);
        this.changeMethodHandler = this.changeMethodHandler.bind(this);
    }

    componentDidMount(){
        ForecastService.getForecastById(this.state.id).then( (res) =>{
            let forecast = res.data;
            this.setState({
                forecastNumber: forecast.forecastNumber,
                forecastHorizonStart: forecast.forecastHorizonStart,
                forecastHorizonEnd: forecast.forecastHorizonEnd,
                method: forecast.method
            });
        });
    }

    updateForecast = (e) => {
        e.preventDefault();
        let forecast = {
            forecastId: this.state.id,
            forecastNumber: this.state.forecastNumber,
            forecastHorizonStart: this.state.forecastHorizonStart,
            forecastHorizonEnd: this.state.forecastHorizonEnd,
            method: this.state.method
        };
        console.log('forecast => ' + JSON.stringify(forecast));
        console.log('id => ' + JSON.stringify(this.state.id));
        ForecastService.updateForecast(forecast).then( res => {
            this.props.history.push('/forecasts');
        });
    }

    changeforecastNumberHandler= (event) => {
        this.setState({forecastNumber: event.target.value});
    }
    changeforecastHorizonStartHandler= (event) => {
        this.setState({forecastHorizonStart: event.target.value});
    }
    changeforecastHorizonEndHandler= (event) => {
        this.setState({forecastHorizonEnd: event.target.value});
    }
    changeMethodHandler= (event) => {
        this.setState({method: event.target.value});
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
                                            <label> forecastNumber: </label>
                                                <input placeholder="forecastNumber" name="forecastNumber" className="form-control" value={this.state.forecastNumber} onChange={this.changeforecastNumberHandler}/>

                                            <label> forecastHorizonStart: </label>
                                                <input type="date" placeholder="forecastHorizonStart" name="forecastHorizonStart" className="form-control" value={this.state.forecastHorizonStart} onChange={this.changeforecastHorizonStartHandler}/>

                                            <label> forecastHorizonEnd: </label>
                                                <input type="date" placeholder="forecastHorizonEnd" name="forecastHorizonEnd" className="form-control" value={this.state.forecastHorizonEnd} onChange={this.changeforecastHorizonEndHandler}/>

                                            <label> Method: </label>
                                                <select value={this.state.method} onChange={this.changeMethodHandler}>
                      <option name="Method" className="form-control" >
                          MovingAverage
                      </option>
                      <option name="Method" className="form-control" >
                          ExponentialSmoothing
                      </option>
                      <option name="Method" className="form-control" >
                          Croston
                      </option>
                      <option name="Method" className="form-control" >
                          ARIMA
                      </option>
                      <option name="Method" className="form-control" >
                          Manual
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
