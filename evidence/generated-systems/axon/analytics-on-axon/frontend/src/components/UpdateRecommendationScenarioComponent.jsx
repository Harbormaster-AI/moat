import React, { Component } from 'react'
import RecommendationScenarioService from '../services/RecommendationScenarioService';

class UpdateRecommendationScenarioComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                objective: '',
                recommendationType: ''
        }
        this.updateRecommendationScenario = this.updateRecommendationScenario.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeobjectiveHandler = this.changeobjectiveHandler.bind(this);
        this.changeRecommendationTypeHandler = this.changeRecommendationTypeHandler.bind(this);
    }

    componentDidMount(){
        RecommendationScenarioService.getRecommendationScenarioById(this.state.id).then( (res) =>{
            let recommendationScenario = res.data;
            this.setState({
                name: recommendationScenario.name,
                objective: recommendationScenario.objective,
                recommendationType: recommendationScenario.recommendationType
            });
        });
    }

    updateRecommendationScenario = (e) => {
        e.preventDefault();
        let recommendationScenario = {
            recommendationScenarioId: this.state.id,
            name: this.state.name,
            objective: this.state.objective,
            recommendationType: this.state.recommendationType
        };
        console.log('recommendationScenario => ' + JSON.stringify(recommendationScenario));
        console.log('id => ' + JSON.stringify(this.state.id));
        RecommendationScenarioService.updateRecommendationScenario(recommendationScenario).then( res => {
            this.props.history.push('/recommendationScenarios');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeobjectiveHandler= (event) => {
        this.setState({objective: event.target.value});
    }
    changeRecommendationTypeHandler= (event) => {
        this.setState({recommendationType: event.target.value});
    }

    cancel(){
        this.props.history.push('/recommendationScenarios');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update RecommendationScenario</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> objective: </label>
                                                <input placeholder="objective" name="objective" className="form-control" value={this.state.objective} onChange={this.changeobjectiveHandler}/>

                                            <label> RecommendationType: </label>
                                                <select value={this.state.recommendationType} onChange={this.changeRecommendationTypeHandler}>
                      <option name="RecommendationType" className="form-control" >
                          Personalized
                      </option>
                      <option name="RecommendationType" className="form-control" >
                          Trending
                      </option>
                      <option name="RecommendationType" className="form-control" >
                          SimilarItems
                      </option>
                      <option name="RecommendationType" className="form-control" >
                          FrequentlyBoughtTogether
                      </option>
                      <option name="RecommendationType" className="form-control" >
                          ContentBased
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateRecommendationScenario}>Save</button>
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

export default UpdateRecommendationScenarioComponent
