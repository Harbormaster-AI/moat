import React, { Component } from 'react'
import CompetencyRatingService from '../services/CompetencyRatingService'

class ListCompetencyRatingComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                competencyRatings: []
        }
        this.addCompetencyRating = this.addCompetencyRating.bind(this);
        this.editCompetencyRating = this.editCompetencyRating.bind(this);
        this.deleteCompetencyRating = this.deleteCompetencyRating.bind(this);
    }

    deleteCompetencyRating(id){
        CompetencyRatingService.deleteCompetencyRating(id).then( res => {
            this.setState({competencyRatings: this.state.competencyRatings.filter(competencyRating => competencyRating.competencyRatingId !== id)});
        });
    }
    viewCompetencyRating(id){
        this.props.history.push(`/view-competencyRating/${id}`);
    }
    editCompetencyRating(id){
        this.props.history.push(`/add-competencyRating/${id}`);
    }

    componentDidMount(){
        CompetencyRatingService.getCompetencyRatings().then((res) => {
            this.setState({ competencyRatings: res.data});
        });
    }

    addCompetencyRating(){
        this.props.history.push('/add-competencyRating/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">CompetencyRating List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCompetencyRating}> Add CompetencyRating</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Comment </th>
                                    <th> Rating </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.competencyRatings.map(
                                        competencyRating => 
                                        <tr key = {competencyRating.competencyRatingId}>
                                             <td> { competencyRating.comment } </td>
                                             <td> { competencyRating.rating } </td>
                                             <td>
                                                 <button onClick={ () => this.editCompetencyRating(competencyRating.competencyRatingId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCompetencyRating(competencyRating.competencyRatingId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCompetencyRating(competencyRating.competencyRatingId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListCompetencyRatingComponent
