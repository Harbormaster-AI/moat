import React, { Component } from 'react'
import RunMetricService from '../services/RunMetricService'

class ListRunMetricComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                runMetrics: []
        }
        this.addRunMetric = this.addRunMetric.bind(this);
        this.editRunMetric = this.editRunMetric.bind(this);
        this.deleteRunMetric = this.deleteRunMetric.bind(this);
    }

    deleteRunMetric(id){
        RunMetricService.deleteRunMetric(id).then( res => {
            this.setState({runMetrics: this.state.runMetrics.filter(runMetric => runMetric.runMetricId !== id)});
        });
    }
    viewRunMetric(id){
        this.props.history.push(`/view-runMetric/${id}`);
    }
    editRunMetric(id){
        this.props.history.push(`/add-runMetric/${id}`);
    }

    componentDidMount(){
        RunMetricService.getRunMetrics().then((res) => {
            this.setState({ runMetrics: res.data});
        });
    }

    addRunMetric(){
        this.props.history.push('/add-runMetric/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">RunMetric List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addRunMetric}> Add RunMetric</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Value </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.runMetrics.map(
                                        runMetric => 
                                        <tr key = {runMetric.runMetricId}>
                                             <td> { runMetric.name } </td>
                                             <td> { runMetric.value } </td>
                                             <td>
                                                 <button onClick={ () => this.editRunMetric(runMetric.runMetricId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteRunMetric(runMetric.runMetricId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewRunMetric(runMetric.runMetricId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListRunMetricComponent
