import React, { Component } from 'react'
import ReviewService from '../services/ReviewService'

class ListReviewComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                reviews: []
        }
        this.addReview = this.addReview.bind(this);
        this.editReview = this.editReview.bind(this);
        this.deleteReview = this.deleteReview.bind(this);
    }

    deleteReview(id){
        ReviewService.deleteReview(id).then( res => {
            this.setState({reviews: this.state.reviews.filter(review => review.reviewId !== id)});
        });
    }
    viewReview(id){
        this.props.history.push(`/view-review/${id}`);
    }
    editReview(id){
        this.props.history.push(`/add-review/${id}`);
    }

    componentDidMount(){
        ReviewService.getReviews().then((res) => {
            this.setState({ reviews: res.data});
        });
    }

    addReview(){
        this.props.history.push('/add-review/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Review List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addReview}> Add Review</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Rating </th>
                                    <th> Title </th>
                                    <th> Content </th>
                                    <th> CreatedAt </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.reviews.map(
                                        review => 
                                        <tr key = {review.reviewId}>
                                             <td> { review.rating } </td>
                                             <td> { review.title } </td>
                                             <td> { review.content } </td>
                                             <td> { review.createdAt } </td>
                                             <td> { review.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editReview(review.reviewId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteReview(review.reviewId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewReview(review.reviewId)} className="btn btn-outline-info btn-sm">View </button>
                                             </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>

                 </div>

            </div>
        )
    }
}

export default ListReviewComponent
