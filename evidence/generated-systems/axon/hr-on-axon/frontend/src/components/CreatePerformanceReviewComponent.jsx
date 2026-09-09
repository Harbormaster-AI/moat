import React, { Component } from 'react'
import PerformanceReviewService from '../services/PerformanceReviewService';

class CreatePerformanceReviewComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                reviewNumber: '',
                reviewDate: '',
                reviewerComments: '',
                rating: '',
                status: ''
        }
        this.changereviewNumberHandler = this.changereviewNumberHandler.bind(this);
        this.changereviewDateHandler = this.changereviewDateHandler.bind(this);
        this.changereviewerCommentsHandler = this.changereviewerCommentsHandler.bind(this);
        this.changeRatingHandler = this.changeRatingHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            PerformanceReviewService.getPerformanceReviewById(this.state.id).then( (res) =>{
                let performanceReview = res.data;
                this.setState({
                    reviewNumber: performanceReview.reviewNumber,
                    reviewDate: performanceReview.reviewDate,
                    reviewerComments: performanceReview.reviewerComments,
                    rating: performanceReview.rating,
                    status: performanceReview.status
                });
            });
        }        
    }
    saveOrUpdatePerformanceReview = (e) => {
        e.preventDefault();
        let performanceReview = {
                performanceReviewId: this.state.id,
                reviewNumber: this.state.reviewNumber,
                reviewDate: this.state.reviewDate,
                reviewerComments: this.state.reviewerComments,
                rating: this.state.rating,
                status: this.state.status
            };
        console.log('performanceReview => ' + JSON.stringify(performanceReview));

        // step 5
        if(this.state.id === '_add'){
            performanceReview.performanceReviewId=''
            PerformanceReviewService.createPerformanceReview(performanceReview).then(res =>{
                this.props.history.push('/performanceReviews');
            });
        }else{
            PerformanceReviewService.updatePerformanceReview(performanceReview).then( res => {
                this.props.history.push('/performanceReviews');
            });
        }
    }
    
    changereviewNumberHandler= (event) => {
        this.setState({reviewNumber: event.target.value});
    }
    changereviewDateHandler= (event) => {
        this.setState({reviewDate: event.target.value});
    }
    changereviewerCommentsHandler= (event) => {
        this.setState({reviewerComments: event.target.value});
    }
    changeRatingHandler= (event) => {
        this.setState({rating: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/performanceReviews');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add PerformanceReview</h3>
        }else{
            return <h3 className="text-center">Update PerformanceReview</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> reviewNumber:&emsp; </label>
                                                <input placeholder="reviewNumber" name="reviewNumber" className="form-control" value={this.state.reviewNumber} onChange={this.changereviewNumberHandler}/>

                                            <label> reviewDate:&emsp; </label>
                                                <input type="date" placeholder="reviewDate" name="reviewDate" className="form-control" value={this.state.reviewDate} onChange={this.changereviewDateHandler}/>

                                            <label> reviewerComments:&emsp; </label>
                                                <input placeholder="reviewerComments" name="reviewerComments" className="form-control" value={this.state.reviewerComments} onChange={this.changereviewerCommentsHandler}/>

                                            <label> Rating:&emsp; </label>
                                                <select value={this.state.rating} onChange={this.changeRatingHandler}>
                      <option name="Rating" className="form-control" >
                          Unsatisfactory
                      </option>
                      <option name="Rating" className="form-control" >
                          NeedsImprovement
                      </option>
                      <option name="Rating" className="form-control" >
                          MeetsExpectations
                      </option>
                      <option name="Rating" className="form-control" >
                          ExceedsExpectations
                      </option>
                      <option name="Rating" className="form-control" >
                          Outstanding
                      </option>
                    </select>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          NotStarted
                      </option>
                      <option name="Status" className="form-control" >
                          InProgress
                      </option>
                      <option name="Status" className="form-control" >
                          Finalized
                      </option>
                      <option name="Status" className="form-control" >
                          Acknowledged
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdatePerformanceReview}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>
                   </div>
            </div>
        )
    }
}

export default CreatePerformanceReviewComponent
