import React, { Component } from 'react'
import RecommendationScenarioService from '../services/RecommendationScenarioService'

class ListRecommendationScenarioComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                recommendationScenarios: []
        }
        this.addRecommendationScenario = this.addRecommendationScenario.bind(this);
        this.editRecommendationScenario = this.editRecommendationScenario.bind(this);
        this.deleteRecommendationScenario = this.deleteRecommendationScenario.bind(this);
    }

    deleteRecommendationScenario(id){
        RecommendationScenarioService.deleteRecommendationScenario(id).then( res => {
            this.setState({recommendationScenarios: this.state.recommendationScenarios.filter(recommendationScenario => recommendationScenario.recommendationScenarioId !== id)});
        });
    }
    viewRecommendationScenario(id){
        this.props.history.push(`/view-recommendationScenario/${id}`);
    }
    editRecommendationScenario(id){
        this.props.history.push(`/add-recommendationScenario/${id}`);
    }

    componentDidMount(){
        RecommendationScenarioService.getRecommendationScenarios().then((res) => {
            this.setState({ recommendationScenarios: res.data});
        });
    }

    addRecommendationScenario(){
        this.props.history.push('/add-recommendationScenario/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">RecommendationScenario List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addRecommendationScenario}> Add RecommendationScenario</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Objective </th>
                                    <th> RecommendationType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.recommendationScenarios.map(
                                        recommendationScenario => 
                                        <tr key = {recommendationScenario.recommendationScenarioId}>
                                             <td> { recommendationScenario.name } </td>
                                             <td> { recommendationScenario.objective } </td>
                                             <td> { recommendationScenario.recommendationType } </td>
                                             <td>
                                                 <button onClick={ () => this.editRecommendationScenario(recommendationScenario.recommendationScenarioId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteRecommendationScenario(recommendationScenario.recommendationScenarioId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewRecommendationScenario(recommendationScenario.recommendationScenarioId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListRecommendationScenarioComponent
