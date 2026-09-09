import React, { Component } from 'react'
import PlacementService from '../services/PlacementService';

class UpdatePlacementComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                flight: '',
                goalImpressions: ''
        }
        this.updatePlacement = this.updatePlacement.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeflightHandler = this.changeflightHandler.bind(this);
        this.changegoalImpressionsHandler = this.changegoalImpressionsHandler.bind(this);
    }

    componentDidMount(){
        PlacementService.getPlacementById(this.state.id).then( (res) =>{
            let placement = res.data;
            this.setState({
                name: placement.name,
                flight: placement.flight,
                goalImpressions: placement.goalImpressions
            });
        });
    }

    updatePlacement = (e) => {
        e.preventDefault();
        let placement = {
            placementId: this.state.id,
            name: this.state.name,
            flight: this.state.flight,
            goalImpressions: this.state.goalImpressions
        };
        console.log('placement => ' + JSON.stringify(placement));
        console.log('id => ' + JSON.stringify(this.state.id));
        PlacementService.updatePlacement(placement).then( res => {
            this.props.history.push('/placements');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeflightHandler= (event) => {
        this.setState({flight: event.target.value});
    }
    changegoalImpressionsHandler= (event) => {
        this.setState({goalImpressions: event.target.value});
    }

    cancel(){
        this.props.history.push('/placements');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Placement</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> flight: </label>
                                                <input placeholder="flight" name="flight" className="form-control" value={this.state.flight} onChange={this.changeflightHandler}/>

                                            <label> goalImpressions: </label>
                                                <input type="number" placeholder="goalImpressions" name="goalImpressions" className="form-control" value={this.state.goalImpressions} onChange={this.changegoalImpressionsHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updatePlacement}>Save</button>
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

export default UpdatePlacementComponent
