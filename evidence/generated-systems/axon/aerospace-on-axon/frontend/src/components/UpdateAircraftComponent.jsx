import React, { Component } from 'react'
import AircraftService from '../services/AircraftService';

class UpdateAircraftComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                msn: '',
                deliveryDate: ''
        }
        this.updateAircraft = this.updateAircraft.bind(this);

        this.changemsnHandler = this.changemsnHandler.bind(this);
        this.changedeliveryDateHandler = this.changedeliveryDateHandler.bind(this);
    }

    componentDidMount(){
        AircraftService.getAircraftById(this.state.id).then( (res) =>{
            let aircraft = res.data;
            this.setState({
                msn: aircraft.msn,
                deliveryDate: aircraft.deliveryDate
            });
        });
    }

    updateAircraft = (e) => {
        e.preventDefault();
        let aircraft = {
            aircraftId: this.state.id,
            msn: this.state.msn,
            deliveryDate: this.state.deliveryDate
        };
        console.log('aircraft => ' + JSON.stringify(aircraft));
        console.log('id => ' + JSON.stringify(this.state.id));
        AircraftService.updateAircraft(aircraft).then( res => {
            this.props.history.push('/aircrafts');
        });
    }

    changemsnHandler= (event) => {
        this.setState({msn: event.target.value});
    }
    changedeliveryDateHandler= (event) => {
        this.setState({deliveryDate: event.target.value});
    }

    cancel(){
        this.props.history.push('/aircrafts');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Aircraft</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> msn: </label>
                                                <input placeholder="msn" name="msn" className="form-control" value={this.state.msn} onChange={this.changemsnHandler}/>

                                            <label> deliveryDate: </label>
                                                <input type="date" placeholder="deliveryDate" name="deliveryDate" className="form-control" value={this.state.deliveryDate} onChange={this.changedeliveryDateHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateAircraft}>Save</button>
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

export default UpdateAircraftComponent
