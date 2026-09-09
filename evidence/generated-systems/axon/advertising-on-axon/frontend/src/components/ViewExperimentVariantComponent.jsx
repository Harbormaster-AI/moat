import React, { Component } from 'react'
import ExperimentVariantService from '../services/ExperimentVariantService'

class ViewExperimentVariantComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            experimentVariant: {}
        }
    }

    componentDidMount(){
        ExperimentVariantService.getExperimentVariantById(this.state.id).then( res => {
            this.setState({experimentVariant: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View ExperimentVariant Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.experimentVariant.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> allocation:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.experimentVariant.allocation }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewExperimentVariantComponent
