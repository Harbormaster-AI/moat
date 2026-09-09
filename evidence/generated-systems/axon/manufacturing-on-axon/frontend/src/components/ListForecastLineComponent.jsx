import React, { Component } from 'react'
import ForecastLineService from '../services/ForecastLineService'

class ListForecastLineComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                forecastLines: []
        }
        this.addForecastLine = this.addForecastLine.bind(this);
        this.editForecastLine = this.editForecastLine.bind(this);
        this.deleteForecastLine = this.deleteForecastLine.bind(this);
    }

    deleteForecastLine(id){
        ForecastLineService.deleteForecastLine(id).then( res => {
            this.setState({forecastLines: this.state.forecastLines.filter(forecastLine => forecastLine.forecastLineId !== id)});
        });
    }
    viewForecastLine(id){
        this.props.history.push(`/view-forecastLine/${id}`);
    }
    editForecastLine(id){
        this.props.history.push(`/add-forecastLine/${id}`);
    }

    componentDidMount(){
        ForecastLineService.getForecastLines().then((res) => {
            this.setState({ forecastLines: res.data});
        });
    }

    addForecastLine(){
        this.props.history.push('/add-forecastLine/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ForecastLine List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addForecastLine}> Add ForecastLine</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Period </th>
                                    <th> Quantity </th>
                                    <th> Confidence </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.forecastLines.map(
                                        forecastLine => 
                                        <tr key = {forecastLine.forecastLineId}>
                                             <td> { forecastLine.period } </td>
                                             <td> { forecastLine.quantity } </td>
                                             <td> { forecastLine.confidence } </td>
                                             <td>
                                                 <button onClick={ () => this.editForecastLine(forecastLine.forecastLineId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteForecastLine(forecastLine.forecastLineId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewForecastLine(forecastLine.forecastLineId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListForecastLineComponent
