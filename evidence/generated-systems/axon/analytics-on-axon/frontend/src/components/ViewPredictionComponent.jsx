import React, { Component } from 'react'
import PredictionService from '../services/PredictionService'

class ViewPredictionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            prediction: {}
        }
    }

    componentDidMount(){
        PredictionService.getPredictionById(this.state.id).then( res => {
            this.setState({prediction: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Prediction Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> referenceKey:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.prediction.referenceKey }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> predictedAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.prediction.predictedAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> score:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.prediction.score }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewPredictionComponent
