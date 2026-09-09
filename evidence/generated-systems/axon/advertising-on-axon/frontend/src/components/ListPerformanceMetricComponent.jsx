import React, { Component } from 'react'
import PerformanceMetricService from '../services/PerformanceMetricService'

class ListPerformanceMetricComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                performanceMetrics: []
        }
        this.addPerformanceMetric = this.addPerformanceMetric.bind(this);
        this.editPerformanceMetric = this.editPerformanceMetric.bind(this);
        this.deletePerformanceMetric = this.deletePerformanceMetric.bind(this);
    }

    deletePerformanceMetric(id){
        PerformanceMetricService.deletePerformanceMetric(id).then( res => {
            this.setState({performanceMetrics: this.state.performanceMetrics.filter(performanceMetric => performanceMetric.performanceMetricId !== id)});
        });
    }
    viewPerformanceMetric(id){
        this.props.history.push(`/view-performanceMetric/${id}`);
    }
    editPerformanceMetric(id){
        this.props.history.push(`/add-performanceMetric/${id}`);
    }

    componentDidMount(){
        PerformanceMetricService.getPerformanceMetrics().then((res) => {
            this.setState({ performanceMetrics: res.data});
        });
    }

    addPerformanceMetric(){
        this.props.history.push('/add-performanceMetric/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">PerformanceMetric List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addPerformanceMetric}> Add PerformanceMetric</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Date </th>
                                    <th> Value </th>
                                    <th> MetricType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.performanceMetrics.map(
                                        performanceMetric => 
                                        <tr key = {performanceMetric.performanceMetricId}>
                                             <td> { performanceMetric.date } </td>
                                             <td> { performanceMetric.value } </td>
                                             <td> { performanceMetric.metricType } </td>
                                             <td>
                                                 <button onClick={ () => this.editPerformanceMetric(performanceMetric.performanceMetricId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deletePerformanceMetric(performanceMetric.performanceMetricId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewPerformanceMetric(performanceMetric.performanceMetricId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListPerformanceMetricComponent
