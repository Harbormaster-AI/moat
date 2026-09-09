import React, { Component } from 'react'
import ExperimentService from '../services/ExperimentService'

class ListExperimentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                experiments: []
        }
        this.addExperiment = this.addExperiment.bind(this);
        this.editExperiment = this.editExperiment.bind(this);
        this.deleteExperiment = this.deleteExperiment.bind(this);
    }

    deleteExperiment(id){
        ExperimentService.deleteExperiment(id).then( res => {
            this.setState({experiments: this.state.experiments.filter(experiment => experiment.experimentId !== id)});
        });
    }
    viewExperiment(id){
        this.props.history.push(`/view-experiment/${id}`);
    }
    editExperiment(id){
        this.props.history.push(`/add-experiment/${id}`);
    }

    componentDidMount(){
        ExperimentService.getExperiments().then((res) => {
            this.setState({ experiments: res.data});
        });
    }

    addExperiment(){
        this.props.history.push('/add-experiment/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Experiment List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addExperiment}> Add Experiment</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Hypothesis </th>
                                    <th> StartDate </th>
                                    <th> EndDate </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.experiments.map(
                                        experiment => 
                                        <tr key = {experiment.experimentId}>
                                             <td> { experiment.name } </td>
                                             <td> { experiment.hypothesis } </td>
                                             <td> { experiment.startDate } </td>
                                             <td> { experiment.endDate } </td>
                                             <td> { experiment.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editExperiment(experiment.experimentId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteExperiment(experiment.experimentId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewExperiment(experiment.experimentId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListExperimentComponent
