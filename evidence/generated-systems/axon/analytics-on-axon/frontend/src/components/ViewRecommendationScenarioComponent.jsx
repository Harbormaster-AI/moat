import React, { Component } from 'react'
import RecommendationScenarioService from '../services/RecommendationScenarioService'

class ViewRecommendationScenarioComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            recommendationScenario: {}
        }
    }

    componentDidMount(){
        RecommendationScenarioService.getRecommendationScenarioById(this.state.id).then( res => {
            this.setState({recommendationScenario: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View RecommendationScenario Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.recommendationScenario.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> objective:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.recommendationScenario.objective }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> RecommendationType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.recommendationScenario.recommendationType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewRecommendationScenarioComponent
