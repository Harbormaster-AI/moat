import React, { Component } from 'react'
import DispositionReviewService from '../services/DispositionReviewService';

class UpdateDispositionReviewComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                reviewDate: '',
                reviewer: '',
                notes: '',
                outcome: ''
        }
        this.updateDispositionReview = this.updateDispositionReview.bind(this);

        this.changereviewDateHandler = this.changereviewDateHandler.bind(this);
        this.changereviewerHandler = this.changereviewerHandler.bind(this);
        this.changenotesHandler = this.changenotesHandler.bind(this);
        this.changeOutcomeHandler = this.changeOutcomeHandler.bind(this);
    }

    componentDidMount(){
        DispositionReviewService.getDispositionReviewById(this.state.id).then( (res) =>{
            let dispositionReview = res.data;
            this.setState({
                reviewDate: dispositionReview.reviewDate,
                reviewer: dispositionReview.reviewer,
                notes: dispositionReview.notes,
                outcome: dispositionReview.outcome
            });
        });
    }

    updateDispositionReview = (e) => {
        e.preventDefault();
        let dispositionReview = {
            dispositionReviewId: this.state.id,
            reviewDate: this.state.reviewDate,
            reviewer: this.state.reviewer,
            notes: this.state.notes,
            outcome: this.state.outcome
        };
        console.log('dispositionReview => ' + JSON.stringify(dispositionReview));
        console.log('id => ' + JSON.stringify(this.state.id));
        DispositionReviewService.updateDispositionReview(dispositionReview).then( res => {
            this.props.history.push('/dispositionReviews');
        });
    }

    changereviewDateHandler= (event) => {
        this.setState({reviewDate: event.target.value});
    }
    changereviewerHandler= (event) => {
        this.setState({reviewer: event.target.value});
    }
    changenotesHandler= (event) => {
        this.setState({notes: event.target.value});
    }
    changeOutcomeHandler= (event) => {
        this.setState({outcome: event.target.value});
    }

    cancel(){
        this.props.history.push('/dispositionReviews');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update DispositionReview</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> reviewDate: </label>
                                                <input type="date" placeholder="reviewDate" name="reviewDate" className="form-control" value={this.state.reviewDate} onChange={this.changereviewDateHandler}/>

                                            <label> reviewer: </label>
                                                <input placeholder="reviewer" name="reviewer" className="form-control" value={this.state.reviewer} onChange={this.changereviewerHandler}/>

                                            <label> notes: </label>
                                                <input placeholder="notes" name="notes" className="form-control" value={this.state.notes} onChange={this.changenotesHandler}/>

                                            <label> Outcome: </label>
                                                <select value={this.state.outcome} onChange={this.changeOutcomeHandler}>
                      <option name="Outcome" className="form-control" >
                          Approved
                      </option>
                      <option name="Outcome" className="form-control" >
                          Deferred
                      </option>
                      <option name="Outcome" className="form-control" >
                          Rejected
                      </option>
                      <option name="Outcome" className="form-control" >
                          Executed
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateDispositionReview}>Save</button>
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

export default UpdateDispositionReviewComponent
