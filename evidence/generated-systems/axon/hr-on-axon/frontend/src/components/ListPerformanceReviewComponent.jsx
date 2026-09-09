import React, { Component } from 'react'
import PerformanceReviewService from '../services/PerformanceReviewService'

class ListPerformanceReviewComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                performanceReviews: []
        }
        this.addPerformanceReview = this.addPerformanceReview.bind(this);
        this.editPerformanceReview = this.editPerformanceReview.bind(this);
        this.deletePerformanceReview = this.deletePerformanceReview.bind(this);
    }

    deletePerformanceReview(id){
        PerformanceReviewService.deletePerformanceReview(id).then( res => {
            this.setState({performanceReviews: this.state.performanceReviews.filter(performanceReview => performanceReview.performanceReviewId !== id)});
        });
    }
    viewPerformanceReview(id){
        this.props.history.push(`/view-performanceReview/${id}`);
    }
    editPerformanceReview(id){
        this.props.history.push(`/add-performanceReview/${id}`);
    }

    componentDidMount(){
        PerformanceReviewService.getPerformanceReviews().then((res) => {
            this.setState({ performanceReviews: res.data});
        });
    }

    addPerformanceReview(){
        this.props.history.push('/add-performanceReview/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">PerformanceReview List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addPerformanceReview}> Add PerformanceReview</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ReviewNumber </th>
                                    <th> ReviewDate </th>
                                    <th> ReviewerComments </th>
                                    <th> Rating </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.performanceReviews.map(
                                        performanceReview => 
                                        <tr key = {performanceReview.performanceReviewId}>
                                             <td> { performanceReview.reviewNumber } </td>
                                             <td> { performanceReview.reviewDate } </td>
                                             <td> { performanceReview.reviewerComments } </td>
                                             <td> { performanceReview.rating } </td>
                                             <td> { performanceReview.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editPerformanceReview(performanceReview.performanceReviewId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deletePerformanceReview(performanceReview.performanceReviewId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewPerformanceReview(performanceReview.performanceReviewId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListPerformanceReviewComponent
