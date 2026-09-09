import React, { Component } from 'react'
import PredictionService from '../services/PredictionService';

class UpdatePredictionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                referenceKey: '',
                predictedAt: '',
                score: ''
        }
        this.updatePrediction = this.updatePrediction.bind(this);

        this.changereferenceKeyHandler = this.changereferenceKeyHandler.bind(this);
        this.changepredictedAtHandler = this.changepredictedAtHandler.bind(this);
        this.changescoreHandler = this.changescoreHandler.bind(this);
    }

    componentDidMount(){
        PredictionService.getPredictionById(this.state.id).then( (res) =>{
            let prediction = res.data;
            this.setState({
                referenceKey: prediction.referenceKey,
                predictedAt: prediction.predictedAt,
                score: prediction.score
            });
        });
    }

    updatePrediction = (e) => {
        e.preventDefault();
        let prediction = {
            predictionId: this.state.id,
            referenceKey: this.state.referenceKey,
            predictedAt: this.state.predictedAt,
            score: this.state.score
        };
        console.log('prediction => ' + JSON.stringify(prediction));
        console.log('id => ' + JSON.stringify(this.state.id));
        PredictionService.updatePrediction(prediction).then( res => {
            this.props.history.push('/predictions');
        });
    }

    changereferenceKeyHandler= (event) => {
        this.setState({referenceKey: event.target.value});
    }
    changepredictedAtHandler= (event) => {
        this.setState({predictedAt: event.target.value});
    }
    changescoreHandler= (event) => {
        this.setState({score: event.target.value});
    }

    cancel(){
        this.props.history.push('/predictions');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Prediction</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> referenceKey: </label>
                                                <input placeholder="referenceKey" name="referenceKey" className="form-control" value={this.state.referenceKey} onChange={this.changereferenceKeyHandler}/>

                                            <label> predictedAt: </label>
                                                <input type="date" placeholder="predictedAt" name="predictedAt" className="form-control" value={this.state.predictedAt} onChange={this.changepredictedAtHandler}/>

                                            <label> score: </label>
                                                <input placeholder="score" name="score" className="form-control" value={this.state.score} onChange={this.changescoreHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updatePrediction}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>

                   </div>
            </div>
        )
    }
}

export default UpdatePredictionComponent
