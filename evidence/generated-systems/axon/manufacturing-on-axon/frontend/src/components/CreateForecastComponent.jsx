import React, { Component } from 'react'
import ForecastService from '../services/ForecastService';

class CreateForecastComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                forecastNumber: '',
                forecastHorizonStart: '',
                forecastHorizonEnd: '',
                method: ''
        }
        this.changeforecastNumberHandler = this.changeforecastNumberHandler.bind(this);
        this.changeforecastHorizonStartHandler = this.changeforecastHorizonStartHandler.bind(this);
        this.changeforecastHorizonEndHandler = this.changeforecastHorizonEndHandler.bind(this);
        this.changeMethodHandler = this.changeMethodHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateForecast = (e) => {
        e.preventDefault();
        let forecast = {
                forecastId: this.state.id,
                forecastNumber: this.state.forecastNumber,
                forecastHorizonStart: this.state.forecastHorizonStart,
                forecastHorizonEnd: this.state.forecastHorizonEnd,
                method: this.state.method
            };
        console.log('forecast => ' + JSON.stringify(forecast));

        // step 5
        if(this.state.id === '_add'){
            forecast.forecastId=''
            ForecastService.createForecast(forecast).then(res =>{
                this.props.history.push('/forecasts');
            });
        }else{
            ForecastService.updateForecast(forecast).then( res => {
                this.props.history.push('/forecasts');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Forecast</h3>
        }else{
            return <h3 className="text-center">Update Forecast</h3>
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
                                            <label> forecastNumber:&emsp; </label>
                                                <input placeholder="forecastNumber" name="forecastNumber" className="form-control" value={this.state.forecastNumber} onChange={this.changeforecastNumberHandler}/>

                                            <label> forecastHorizonStart:&emsp; </label>
                                                <input type="date" placeholder="forecastHorizonStart" name="forecastHorizonStart" className="form-control" value={this.state.forecastHorizonStart} onChange={this.changeforecastHorizonStartHandler}/>

                                            <label> forecastHorizonEnd:&emsp; </label>
                                                <input type="date" placeholder="forecastHorizonEnd" name="forecastHorizonEnd" className="form-control" value={this.state.forecastHorizonEnd} onChange={this.changeforecastHorizonEndHandler}/>

                                            <label> Method:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateForecast}>Save</button>
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

export default CreateForecastComponent
