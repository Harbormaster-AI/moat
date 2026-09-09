import React, { Component } from 'react'
import MetricService from '../services/MetricService'

class ListMetricComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                metrics: []
        }
        this.addMetric = this.addMetric.bind(this);
        this.editMetric = this.editMetric.bind(this);
        this.deleteMetric = this.deleteMetric.bind(this);
    }

    deleteMetric(id){
        MetricService.deleteMetric(id).then( res => {
            this.setState({metrics: this.state.metrics.filter(metric => metric.metricId !== id)});
        });
    }
    viewMetric(id){
        this.props.history.push(`/view-metric/${id}`);
    }
    editMetric(id){
        this.props.history.push(`/add-metric/${id}`);
    }

    componentDidMount(){
        MetricService.getMetrics().then((res) => {
            this.setState({ metrics: res.data});
        });
    }

    addMetric(){
        this.props.history.push('/add-metric/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Metric List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addMetric}> Add Metric</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Expression </th>
                                    <th> Unit </th>
                                    <th> MetricType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.metrics.map(
                                        metric => 
                                        <tr key = {metric.metricId}>
                                             <td> { metric.name } </td>
                                             <td> { metric.expression } </td>
                                             <td> { metric.unit } </td>
                                             <td> { metric.metricType } </td>
                                             <td>
                                                 <button onClick={ () => this.editMetric(metric.metricId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteMetric(metric.metricId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewMetric(metric.metricId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListMetricComponent
