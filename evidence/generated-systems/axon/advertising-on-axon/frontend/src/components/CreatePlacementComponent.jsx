import React, { Component } from 'react'
import PlacementService from '../services/PlacementService';

class CreatePlacementComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                flight: '',
                goalImpressions: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeflightHandler = this.changeflightHandler.bind(this);
        this.changegoalImpressionsHandler = this.changegoalImpressionsHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            PlacementService.getPlacementById(this.state.id).then( (res) =>{
                let placement = res.data;
                this.setState({
                    name: placement.name,
                    flight: placement.flight,
                    goalImpressions: placement.goalImpressions
                });
            });
        }        
    }
    saveOrUpdatePlacement = (e) => {
        e.preventDefault();
        let placement = {
                placementId: this.state.id,
                name: this.state.name,
                flight: this.state.flight,
                goalImpressions: this.state.goalImpressions
            };
        console.log('placement => ' + JSON.stringify(placement));

        // step 5
        if(this.state.id === '_add'){
            placement.placementId=''
            PlacementService.createPlacement(placement).then(res =>{
                this.props.history.push('/placements');
            });
        }else{
            PlacementService.updatePlacement(placement).then( res => {
                this.props.history.push('/placements');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Placement</h3>
        }else{
            return <h3 className="text-center">Update Placement</h3>
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
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> flight:&emsp; </label>
                                                <input placeholder="flight" name="flight" className="form-control" value={this.state.flight} onChange={this.changeflightHandler}/>

                                            <label> goalImpressions:&emsp; </label>
                                                <input type="number" placeholder="goalImpressions" name="goalImpressions" className="form-control" value={this.state.goalImpressions} onChange={this.changegoalImpressionsHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdatePlacement}>Save</button>
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

export default CreatePlacementComponent
