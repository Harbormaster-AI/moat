import React, { Component } from 'react'
import EvaluationMetricService from '../services/EvaluationMetricService'

class ListEvaluationMetricComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                evaluationMetrics: []
        }
        this.addEvaluationMetric = this.addEvaluationMetric.bind(this);
        this.editEvaluationMetric = this.editEvaluationMetric.bind(this);
        this.deleteEvaluationMetric = this.deleteEvaluationMetric.bind(this);
    }

    deleteEvaluationMetric(id){
        EvaluationMetricService.deleteEvaluationMetric(id).then( res => {
            this.setState({evaluationMetrics: this.state.evaluationMetrics.filter(evaluationMetric => evaluationMetric.evaluationMetricId !== id)});
        });
    }
    viewEvaluationMetric(id){
        this.props.history.push(`/view-evaluationMetric/${id}`);
    }
    editEvaluationMetric(id){
        this.props.history.push(`/add-evaluationMetric/${id}`);
    }

    componentDidMount(){
        EvaluationMetricService.getEvaluationMetrics().then((res) => {
            this.setState({ evaluationMetrics: res.data});
        });
    }

    addEvaluationMetric(){
        this.props.history.push('/add-evaluationMetric/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">EvaluationMetric List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addEvaluationMetric}> Add EvaluationMetric</button>
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
                                    this.state.evaluationMetrics.map(
                                        evaluationMetric => 
                                        <tr key = {evaluationMetric.evaluationMetricId}>
                                             <td> { evaluationMetric.name } </td>
                                             <td> { evaluationMetric.value } </td>
                                             <td>
                                                 <button onClick={ () => this.editEvaluationMetric(evaluationMetric.evaluationMetricId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteEvaluationMetric(evaluationMetric.evaluationMetricId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewEvaluationMetric(evaluationMetric.evaluationMetricId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListEvaluationMetricComponent
