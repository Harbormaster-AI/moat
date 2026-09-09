import React, { Component } from 'react'
import ReviewService from '../services/ReviewService';

class UpdateReviewComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                rating: '',
                title: '',
                content: '',
                createdAt: '',
                status: ''
        }
        this.updateReview = this.updateReview.bind(this);

        this.changeratingHandler = this.changeratingHandler.bind(this);
        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changecontentHandler = this.changecontentHandler.bind(this);
        this.changecreatedAtHandler = this.changecreatedAtHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        ReviewService.getReviewById(this.state.id).then( (res) =>{
            let review = res.data;
            this.setState({
                rating: review.rating,
                title: review.title,
                content: review.content,
                createdAt: review.createdAt,
                status: review.status
            });
        });
    }

    updateReview = (e) => {
        e.preventDefault();
        let review = {
            reviewId: this.state.id,
            rating: this.state.rating,
            title: this.state.title,
            content: this.state.content,
            createdAt: this.state.createdAt,
            status: this.state.status
        };
        console.log('review => ' + JSON.stringify(review));
        console.log('id => ' + JSON.stringify(this.state.id));
        ReviewService.updateReview(review).then( res => {
            this.props.history.push('/reviews');
        });
    }

    changeratingHandler= (event) => {
        this.setState({rating: event.target.value});
    }
    changetitleHandler= (event) => {
        this.setState({title: event.target.value});
    }
    changecontentHandler= (event) => {
        this.setState({content: event.target.value});
    }
    changecreatedAtHandler= (event) => {
        this.setState({createdAt: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/reviews');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Review</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> rating: </label>
                                                <input type="number" placeholder="rating" name="rating" className="form-control" value={this.state.rating} onChange={this.changeratingHandler}/>

                                            <label> title: </label>
                                                <input placeholder="title" name="title" className="form-control" value={this.state.title} onChange={this.changetitleHandler}/>

                                            <label> content: </label>
                                                <input placeholder="content" name="content" className="form-control" value={this.state.content} onChange={this.changecontentHandler}/>

                                            <label> createdAt: </label>
                                                <input type="date" placeholder="createdAt" name="createdAt" className="form-control" value={this.state.createdAt} onChange={this.changecreatedAtHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Pending
                      </option>
                      <option name="Status" className="form-control" >
                          Approved
                      </option>
                      <option name="Status" className="form-control" >
                          Rejected
                      </option>
                      <option name="Status" className="form-control" >
                          Flagged
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateReview}>Save</button>
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

export default UpdateReviewComponent
