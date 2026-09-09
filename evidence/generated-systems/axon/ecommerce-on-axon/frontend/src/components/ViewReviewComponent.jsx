import React, { Component } from 'react'
import ReviewService from '../services/ReviewService'

class ViewReviewComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            review: {}
        }
    }

    componentDidMount(){
        ReviewService.getReviewById(this.state.id).then( res => {
            this.setState({review: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Review Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> rating:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.review.rating }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> title:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.review.title }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> content:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.review.content }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> createdAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.review.createdAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.review.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewReviewComponent
