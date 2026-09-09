import React, { Component } from 'react'
import VisualizationService from '../services/VisualizationService'

class ViewVisualizationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            visualization: {}
        }
    }

    componentDidMount(){
        VisualizationService.getVisualizationById(this.state.id).then( res => {
            this.setState({visualization: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Visualization Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> title:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.visualization.title }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> options:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.visualization.options }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ChartType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.visualization.chartType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewVisualizationComponent
