import React, { Component } from 'react'
import ForecastService from '../services/ForecastService'

class ListForecastComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                forecasts: []
        }
        this.addForecast = this.addForecast.bind(this);
        this.editForecast = this.editForecast.bind(this);
        this.deleteForecast = this.deleteForecast.bind(this);
    }

    deleteForecast(id){
        ForecastService.deleteForecast(id).then( res => {
            this.setState({forecasts: this.state.forecasts.filter(forecast => forecast.forecastId !== id)});
        });
    }
    viewForecast(id){
        this.props.history.push(`/view-forecast/${id}`);
    }
    editForecast(id){
        this.props.history.push(`/add-forecast/${id}`);
    }

    componentDidMount(){
        ForecastService.getForecasts().then((res) => {
            this.setState({ forecasts: res.data});
        });
    }

    addForecast(){
        this.props.history.push('/add-forecast/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Forecast List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addForecast}> Add Forecast</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ForecastNumber </th>
                                    <th> ForecastHorizonStart </th>
                                    <th> ForecastHorizonEnd </th>
                                    <th> Method </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.forecasts.map(
                                        forecast => 
                                        <tr key = {forecast.forecastId}>
                                             <td> { forecast.forecastNumber } </td>
                                             <td> { forecast.forecastHorizonStart } </td>
                                             <td> { forecast.forecastHorizonEnd } </td>
                                             <td> { forecast.method } </td>
                                             <td>
                                                 <button onClick={ () => this.editForecast(forecast.forecastId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteForecast(forecast.forecastId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewForecast(forecast.forecastId)} className="btn btn-outline-info btn-sm">View </button>
                                             </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>

                 </div>

            </div>
        )
    }
}

export default ListForecastComponent
