import React, { Component } from 'react'
import InferenceEndpointService from '../services/InferenceEndpointService'

class ViewInferenceEndpointComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            inferenceEndpoint: {}
        }
    }

    componentDidMount(){
        InferenceEndpointService.getInferenceEndpointById(this.state.id).then( res => {
            this.setState({inferenceEndpoint: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View InferenceEndpoint Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inferenceEndpoint.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> endpointUrl:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inferenceEndpoint.endpointUrl }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> trafficShare:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inferenceEndpoint.trafficShare }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Mode:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.inferenceEndpoint.mode }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewInferenceEndpointComponent
