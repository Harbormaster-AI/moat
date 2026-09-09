import React, { Component } from 'react'
import DispositionReviewService from '../services/DispositionReviewService'

class ListDispositionReviewComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                dispositionReviews: []
        }
        this.addDispositionReview = this.addDispositionReview.bind(this);
        this.editDispositionReview = this.editDispositionReview.bind(this);
        this.deleteDispositionReview = this.deleteDispositionReview.bind(this);
    }

    deleteDispositionReview(id){
        DispositionReviewService.deleteDispositionReview(id).then( res => {
            this.setState({dispositionReviews: this.state.dispositionReviews.filter(dispositionReview => dispositionReview.dispositionReviewId !== id)});
        });
    }
    viewDispositionReview(id){
        this.props.history.push(`/view-dispositionReview/${id}`);
    }
    editDispositionReview(id){
        this.props.history.push(`/add-dispositionReview/${id}`);
    }

    componentDidMount(){
        DispositionReviewService.getDispositionReviews().then((res) => {
            this.setState({ dispositionReviews: res.data});
        });
    }

    addDispositionReview(){
        this.props.history.push('/add-dispositionReview/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">DispositionReview List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addDispositionReview}> Add DispositionReview</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ReviewDate </th>
                                    <th> Reviewer </th>
                                    <th> Notes </th>
                                    <th> Outcome </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.dispositionReviews.map(
                                        dispositionReview => 
                                        <tr key = {dispositionReview.dispositionReviewId}>
                                             <td> { dispositionReview.reviewDate } </td>
                                             <td> { dispositionReview.reviewer } </td>
                                             <td> { dispositionReview.notes } </td>
                                             <td> { dispositionReview.outcome } </td>
                                             <td>
                                                 <button onClick={ () => this.editDispositionReview(dispositionReview.dispositionReviewId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteDispositionReview(dispositionReview.dispositionReviewId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewDispositionReview(dispositionReview.dispositionReviewId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListDispositionReviewComponent
