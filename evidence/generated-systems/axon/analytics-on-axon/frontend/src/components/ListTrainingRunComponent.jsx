import React, { Component } from 'react'
import TrainingRunService from '../services/TrainingRunService'

class ListTrainingRunComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                trainingRuns: []
        }
        this.addTrainingRun = this.addTrainingRun.bind(this);
        this.editTrainingRun = this.editTrainingRun.bind(this);
        this.deleteTrainingRun = this.deleteTrainingRun.bind(this);
    }

    deleteTrainingRun(id){
        TrainingRunService.deleteTrainingRun(id).then( res => {
            this.setState({trainingRuns: this.state.trainingRuns.filter(trainingRun => trainingRun.trainingRunId !== id)});
        });
    }
    viewTrainingRun(id){
        this.props.history.push(`/view-trainingRun/${id}`);
    }
    editTrainingRun(id){
        this.props.history.push(`/add-trainingRun/${id}`);
    }

    componentDidMount(){
        TrainingRunService.getTrainingRuns().then((res) => {
            this.setState({ trainingRuns: res.data});
        });
    }

    addTrainingRun(){
        this.props.history.push('/add-trainingRun/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">TrainingRun List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addTrainingRun}> Add TrainingRun</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> RunLabel </th>
                                    <th> StartedAt </th>
                                    <th> CompletedAt </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.trainingRuns.map(
                                        trainingRun => 
                                        <tr key = {trainingRun.trainingRunId}>
                                             <td> { trainingRun.runLabel } </td>
                                             <td> { trainingRun.startedAt } </td>
                                             <td> { trainingRun.completedAt } </td>
                                             <td> { trainingRun.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editTrainingRun(trainingRun.trainingRunId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteTrainingRun(trainingRun.trainingRunId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewTrainingRun(trainingRun.trainingRunId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListTrainingRunComponent
