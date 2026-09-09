import React, { Component } from 'react'
import SemanticModelService from '../services/SemanticModelService'

class ViewSemanticModelComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            semanticModel: {}
        }
    }

    componentDidMount(){
        SemanticModelService.getSemanticModelById(this.state.id).then( res => {
            this.setState({semanticModel: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View SemanticModel Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.semanticModel.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> version:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.semanticModel.version }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> grain:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.semanticModel.grain }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewSemanticModelComponent
