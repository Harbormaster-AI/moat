import React, { Component } from 'react'
import CompetencyRatingService from '../services/CompetencyRatingService';

class UpdateCompetencyRatingComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                comment: '',
                rating: ''
        }
        this.updateCompetencyRating = this.updateCompetencyRating.bind(this);

        this.changecommentHandler = this.changecommentHandler.bind(this);
        this.changeRatingHandler = this.changeRatingHandler.bind(this);
    }

    componentDidMount(){
        CompetencyRatingService.getCompetencyRatingById(this.state.id).then( (res) =>{
            let competencyRating = res.data;
            this.setState({
                comment: competencyRating.comment,
                rating: competencyRating.rating
            });
        });
    }

    updateCompetencyRating = (e) => {
        e.preventDefault();
        let competencyRating = {
            competencyRatingId: this.state.id,
            comment: this.state.comment,
            rating: this.state.rating
        };
        console.log('competencyRating => ' + JSON.stringify(competencyRating));
        console.log('id => ' + JSON.stringify(this.state.id));
        CompetencyRatingService.updateCompetencyRating(competencyRating).then( res => {
            this.props.history.push('/competencyRatings');
        });
    }

    changecommentHandler= (event) => {
        this.setState({comment: event.target.value});
    }
    changeRatingHandler= (event) => {
        this.setState({rating: event.target.value});
    }

    cancel(){
        this.props.history.push('/competencyRatings');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update CompetencyRating</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> comment: </label>
                                                <input placeholder="comment" name="comment" className="form-control" value={this.state.comment} onChange={this.changecommentHandler}/>

                                            <label> Rating: </label>
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

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateCompetencyRating}>Save</button>
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

export default UpdateCompetencyRatingComponent
