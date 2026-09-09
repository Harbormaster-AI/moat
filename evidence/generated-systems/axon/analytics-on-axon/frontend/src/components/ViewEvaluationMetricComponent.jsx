import React, { Component } from 'react'
import EvaluationMetricService from '../services/EvaluationMetricService'

class ViewEvaluationMetricComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            evaluationMetric: {}
        }
    }

    componentDidMount(){
        EvaluationMetricService.getEvaluationMetricById(this.state.id).then( res => {
            this.setState({evaluationMetric: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View EvaluationMetric Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.evaluationMetric.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> value:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.evaluationMetric.value }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewEvaluationMetricComponent
