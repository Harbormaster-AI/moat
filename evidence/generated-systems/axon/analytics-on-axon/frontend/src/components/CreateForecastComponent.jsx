import React, { Component } from 'react'
import ForecastService from '../services/ForecastService';

class CreateForecastComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                horizon: '',
                granularity: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changehorizonHandler = this.changehorizonHandler.bind(this);
        this.changeGranularityHandler = this.changeGranularityHandler.bind(this);
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
                    name: forecast.name,
                    horizon: forecast.horizon,
                    granularity: forecast.granularity
                });
            });
        }        
    }
    saveOrUpdateForecast = (e) => {
        e.preventDefault();
        let forecast = {
                forecastId: this.state.id,
                name: this.state.name,
                horizon: this.state.horizon,
                granularity: this.state.granularity
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
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> horizon:&emsp; </label>
                                                <input type="number" placeholder="horizon" name="horizon" className="form-control" value={this.state.horizon} onChange={this.changehorizonHandler}/>

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
