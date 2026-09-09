import React, { Component } from 'react'
import ObservationService from '../services/ObservationService'

class ViewObservationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            observation: {}
        }
    }

    componentDidMount(){
        ObservationService.getObservationById(this.state.id).then( res => {
            this.setState({observation: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Observation Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> code:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.observation.code }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> value:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.observation.value }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> unit:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.observation.unit }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> effectiveDateTime:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.observation.effectiveDateTime }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Interpretation:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.observation.interpretation }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewObservationComponent
