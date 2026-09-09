import React, { Component } from 'react'
import CompetencyRatingService from '../services/CompetencyRatingService'

class ViewCompetencyRatingComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            competencyRating: {}
        }
    }

    componentDidMount(){
        CompetencyRatingService.getCompetencyRatingById(this.state.id).then( res => {
            this.setState({competencyRating: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View CompetencyRating Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> comment:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.competencyRating.comment }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Rating:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.competencyRating.rating }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewCompetencyRatingComponent
