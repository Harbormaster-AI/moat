import React, { Component } from 'react'
import PerformanceReviewService from '../services/PerformanceReviewService'

class ViewPerformanceReviewComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            performanceReview: {}
        }
    }

    componentDidMount(){
        PerformanceReviewService.getPerformanceReviewById(this.state.id).then( res => {
            this.setState({performanceReview: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View PerformanceReview Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> reviewNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.performanceReview.reviewNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> reviewDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.performanceReview.reviewDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> reviewerComments:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.performanceReview.reviewerComments }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Rating:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.performanceReview.rating }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.performanceReview.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewPerformanceReviewComponent
