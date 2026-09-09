import React, { Component } from 'react'
import CompetencyRatingService from '../services/CompetencyRatingService';

class CreateCompetencyRatingComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                comment: '',
                rating: ''
        }
        this.changecommentHandler = this.changecommentHandler.bind(this);
        this.changeRatingHandler = this.changeRatingHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            CompetencyRatingService.getCompetencyRatingById(this.state.id).then( (res) =>{
                let competencyRating = res.data;
                this.setState({
                    comment: competencyRating.comment,
                    rating: competencyRating.rating
                });
            });
        }        
    }
    saveOrUpdateCompetencyRating = (e) => {
        e.preventDefault();
        let competencyRating = {
                competencyRatingId: this.state.id,
                comment: this.state.comment,
                rating: this.state.rating
            };
        console.log('competencyRating => ' + JSON.stringify(competencyRating));

        // step 5
        if(this.state.id === '_add'){
            competencyRating.competencyRatingId=''
            CompetencyRatingService.createCompetencyRating(competencyRating).then(res =>{
                this.props.history.push('/competencyRatings');
            });
        }else{
            CompetencyRatingService.updateCompetencyRating(competencyRating).then( res => {
                this.props.history.push('/competencyRatings');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add CompetencyRating</h3>
        }else{
            return <h3 className="text-center">Update CompetencyRating</h3>
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
                                            <label> comment:&emsp; </label>
                                                <input placeholder="comment" name="comment" className="form-control" value={this.state.comment} onChange={this.changecommentHandler}/>

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

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateCompetencyRating}>Save</button>
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

export default CreateCompetencyRatingComponent
