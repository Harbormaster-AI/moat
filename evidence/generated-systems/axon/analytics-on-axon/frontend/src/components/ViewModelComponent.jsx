import React, { Component } from 'react'
import ModelService from '../services/ModelService'

class ViewModelComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            model: {}
        }
    }

    componentDidMount(){
        ModelService.getModelById(this.state.id).then( res => {
            this.setState({model: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Model Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.model.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> taskDescription:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.model.taskDescription }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ModelType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.model.modelType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewModelComponent
