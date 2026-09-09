import React, { Component } from 'react'
import MeasureService from '../services/MeasureService'

class ViewMeasureComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            measure: {}
        }
    }

    componentDidMount(){
        MeasureService.getMeasureById(this.state.id).then( res => {
            this.setState({measure: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Measure Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.measure.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> format:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.measure.format }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Aggregation:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.measure.aggregation }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewMeasureComponent
