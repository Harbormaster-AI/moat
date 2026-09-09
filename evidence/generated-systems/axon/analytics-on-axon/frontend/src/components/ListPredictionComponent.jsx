import React, { Component } from 'react'
import PredictionService from '../services/PredictionService'

class ListPredictionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                predictions: []
        }
        this.addPrediction = this.addPrediction.bind(this);
        this.editPrediction = this.editPrediction.bind(this);
        this.deletePrediction = this.deletePrediction.bind(this);
    }

    deletePrediction(id){
        PredictionService.deletePrediction(id).then( res => {
            this.setState({predictions: this.state.predictions.filter(prediction => prediction.predictionId !== id)});
        });
    }
    viewPrediction(id){
        this.props.history.push(`/view-prediction/${id}`);
    }
    editPrediction(id){
        this.props.history.push(`/add-prediction/${id}`);
    }

    componentDidMount(){
        PredictionService.getPredictions().then((res) => {
            this.setState({ predictions: res.data});
        });
    }

    addPrediction(){
        this.props.history.push('/add-prediction/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Prediction List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addPrediction}> Add Prediction</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ReferenceKey </th>
                                    <th> PredictedAt </th>
                                    <th> Score </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.predictions.map(
                                        prediction => 
                                        <tr key = {prediction.predictionId}>
                                             <td> { prediction.referenceKey } </td>
                                             <td> { prediction.predictedAt } </td>
                                             <td> { prediction.score } </td>
                                             <td>
                                                 <button onClick={ () => this.editPrediction(prediction.predictionId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deletePrediction(prediction.predictionId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewPrediction(prediction.predictionId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListPredictionComponent
