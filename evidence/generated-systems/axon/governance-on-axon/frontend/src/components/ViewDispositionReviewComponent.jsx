import React, { Component } from 'react'
import DispositionReviewService from '../services/DispositionReviewService'

class ViewDispositionReviewComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            dispositionReview: {}
        }
    }

    componentDidMount(){
        DispositionReviewService.getDispositionReviewById(this.state.id).then( res => {
            this.setState({dispositionReview: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View DispositionReview Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> reviewDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dispositionReview.reviewDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> reviewer:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dispositionReview.reviewer }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> notes:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dispositionReview.notes }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Outcome:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dispositionReview.outcome }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewDispositionReviewComponent
