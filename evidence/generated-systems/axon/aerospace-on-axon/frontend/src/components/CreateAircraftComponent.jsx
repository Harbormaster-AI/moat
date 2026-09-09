import React, { Component } from 'react'
import AircraftService from '../services/AircraftService';

class CreateAircraftComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                msn: '',
                deliveryDate: ''
        }
        this.changemsnHandler = this.changemsnHandler.bind(this);
        this.changedeliveryDateHandler = this.changedeliveryDateHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            AircraftService.getAircraftById(this.state.id).then( (res) =>{
                let aircraft = res.data;
                this.setState({
                    msn: aircraft.msn,
                    deliveryDate: aircraft.deliveryDate
                });
            });
        }        
    }
    saveOrUpdateAircraft = (e) => {
        e.preventDefault();
        let aircraft = {
                aircraftId: this.state.id,
                msn: this.state.msn,
                deliveryDate: this.state.deliveryDate
            };
        console.log('aircraft => ' + JSON.stringify(aircraft));

        // step 5
        if(this.state.id === '_add'){
            aircraft.aircraftId=''
            AircraftService.createAircraft(aircraft).then(res =>{
                this.props.history.push('/aircrafts');
            });
        }else{
            AircraftService.updateAircraft(aircraft).then( res => {
                this.props.history.push('/aircrafts');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Aircraft</h3>
        }else{
            return <h3 className="text-center">Update Aircraft</h3>
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
                                            <label> msn:&emsp; </label>
                                                <input placeholder="msn" name="msn" className="form-control" value={this.state.msn} onChange={this.changemsnHandler}/>

                                            <label> deliveryDate:&emsp; </label>
                                                <input type="date" placeholder="deliveryDate" name="deliveryDate" className="form-control" value={this.state.deliveryDate} onChange={this.changedeliveryDateHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateAircraft}>Save</button>
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

export default CreateAircraftComponent
