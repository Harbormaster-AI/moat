import React, { Component } from 'react'
import TimeSeriesService from '../services/TimeSeriesService'

class ListTimeSeriesComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                timeSeriess: []
        }
        this.addTimeSeries = this.addTimeSeries.bind(this);
        this.editTimeSeries = this.editTimeSeries.bind(this);
        this.deleteTimeSeries = this.deleteTimeSeries.bind(this);
    }

    deleteTimeSeries(id){
        TimeSeriesService.deleteTimeSeries(id).then( res => {
            this.setState({timeSeriess: this.state.timeSeriess.filter(timeSeries => timeSeries.timeSeriesId !== id)});
        });
    }
    viewTimeSeries(id){
        this.props.history.push(`/view-timeSeries/${id}`);
    }
    editTimeSeries(id){
        this.props.history.push(`/add-timeSeries/${id}`);
    }

    componentDidMount(){
        TimeSeriesService.getTimeSeriess().then((res) => {
            this.setState({ timeSeriess: res.data});
        });
    }

    addTimeSeries(){
        this.props.history.push('/add-timeSeries/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">TimeSeries List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addTimeSeries}> Add TimeSeries</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Timezone </th>
                                    <th> Granularity </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.timeSeriess.map(
                                        timeSeries => 
                                        <tr key = {timeSeries.timeSeriesId}>
                                             <td> { timeSeries.name } </td>
                                             <td> { timeSeries.timezone } </td>
                                             <td> { timeSeries.granularity } </td>
                                             <td>
                                                 <button onClick={ () => this.editTimeSeries(timeSeries.timeSeriesId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteTimeSeries(timeSeries.timeSeriesId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewTimeSeries(timeSeries.timeSeriesId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListTimeSeriesComponent
