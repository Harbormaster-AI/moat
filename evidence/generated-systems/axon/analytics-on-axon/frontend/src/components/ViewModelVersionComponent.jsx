import React, { Component } from 'react'
import ModelVersionService from '../services/ModelVersionService'

class ViewModelVersionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            modelVersion: {}
        }
    }

    componentDidMount(){
        ModelVersionService.getModelVersionById(this.state.id).then( res => {
            this.setState({modelVersion: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View ModelVersion Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> version:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.modelVersion.version }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Lifecycle:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.modelVersion.lifecycle }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> TrainingStatus:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.modelVersion.trainingStatus }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewModelVersionComponent
