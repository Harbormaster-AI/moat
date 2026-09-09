import React, { Component } from 'react'
import ReviewService from '../services/ReviewService';

class CreateReviewComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                rating: '',
                title: '',
                content: '',
                createdAt: '',
                status: ''
        }
        this.changeratingHandler = this.changeratingHandler.bind(this);
        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changecontentHandler = this.changecontentHandler.bind(this);
        this.changecreatedAtHandler = this.changecreatedAtHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateReview = (e) => {
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

        // step 5
        if(this.state.id === '_add'){
            review.reviewId=''
            ReviewService.createReview(review).then(res =>{
                this.props.history.push('/reviews');
            });
        }else{
            ReviewService.updateReview(review).then( res => {
                this.props.history.push('/reviews');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Review</h3>
        }else{
            return <h3 className="text-center">Update Review</h3>
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
                                            <label> rating:&emsp; </label>
                                                <input type="number" placeholder="rating" name="rating" className="form-control" value={this.state.rating} onChange={this.changeratingHandler}/>

                                            <label> title:&emsp; </label>
                                                <input placeholder="title" name="title" className="form-control" value={this.state.title} onChange={this.changetitleHandler}/>

                                            <label> content:&emsp; </label>
                                                <input placeholder="content" name="content" className="form-control" value={this.state.content} onChange={this.changecontentHandler}/>

                                            <label> createdAt:&emsp; </label>
                                                <input type="date" placeholder="createdAt" name="createdAt" className="form-control" value={this.state.createdAt} onChange={this.changecreatedAtHandler}/>

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateReview}>Save</button>
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

export default CreateReviewComponent
