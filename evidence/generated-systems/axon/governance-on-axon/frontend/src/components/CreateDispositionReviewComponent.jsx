import React, { Component } from 'react'
import DispositionReviewService from '../services/DispositionReviewService';

class CreateDispositionReviewComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                reviewDate: '',
                reviewer: '',
                notes: '',
                outcome: ''
        }
        this.changereviewDateHandler = this.changereviewDateHandler.bind(this);
        this.changereviewerHandler = this.changereviewerHandler.bind(this);
        this.changenotesHandler = this.changenotesHandler.bind(this);
        this.changeOutcomeHandler = this.changeOutcomeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateDispositionReview = (e) => {
        e.preventDefault();
        let dispositionReview = {
                dispositionReviewId: this.state.id,
                reviewDate: this.state.reviewDate,
                reviewer: this.state.reviewer,
                notes: this.state.notes,
                outcome: this.state.outcome
            };
        console.log('dispositionReview => ' + JSON.stringify(dispositionReview));

        // step 5
        if(this.state.id === '_add'){
            dispositionReview.dispositionReviewId=''
            DispositionReviewService.createDispositionReview(dispositionReview).then(res =>{
                this.props.history.push('/dispositionReviews');
            });
        }else{
            DispositionReviewService.updateDispositionReview(dispositionReview).then( res => {
                this.props.history.push('/dispositionReviews');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add DispositionReview</h3>
        }else{
            return <h3 className="text-center">Update DispositionReview</h3>
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
                                            <label> reviewDate:&emsp; </label>
                                                <input type="date" placeholder="reviewDate" name="reviewDate" className="form-control" value={this.state.reviewDate} onChange={this.changereviewDateHandler}/>

                                            <label> reviewer:&emsp; </label>
                                                <input placeholder="reviewer" name="reviewer" className="form-control" value={this.state.reviewer} onChange={this.changereviewerHandler}/>

                                            <label> notes:&emsp; </label>
                                                <input placeholder="notes" name="notes" className="form-control" value={this.state.notes} onChange={this.changenotesHandler}/>

                                            <label> Outcome:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateDispositionReview}>Save</button>
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

export default CreateDispositionReviewComponent
