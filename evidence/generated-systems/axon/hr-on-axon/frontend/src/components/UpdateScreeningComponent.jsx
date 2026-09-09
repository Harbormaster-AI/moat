import React, { Component } from 'react'
import ScreeningService from '../services/ScreeningService';

class UpdateScreeningComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                completedDate: '',
                status: ''
        }
        this.updateScreening = this.updateScreening.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changecompletedDateHandler = this.changecompletedDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        ScreeningService.getScreeningById(this.state.id).then( (res) =>{
            let screening = res.data;
            this.setState({
                name: screening.name,
                completedDate: screening.completedDate,
                status: screening.status
            });
        });
    }

    updateScreening = (e) => {
        e.preventDefault();
        let screening = {
            screeningId: this.state.id,
            name: this.state.name,
            completedDate: this.state.completedDate,
            status: this.state.status
        };
        console.log('screening => ' + JSON.stringify(screening));
        console.log('id => ' + JSON.stringify(this.state.id));
        ScreeningService.updateScreening(screening).then( res => {
            this.props.history.push('/screenings');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changecompletedDateHandler= (event) => {
        this.setState({completedDate: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/screenings');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Screening</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> completedDate: </label>
                                                <input type="date" placeholder="completedDate" name="completedDate" className="form-control" value={this.state.completedDate} onChange={this.changecompletedDateHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Ordered
                      </option>
                      <option name="Status" className="form-control" >
                          InProgress
                      </option>
                      <option name="Status" className="form-control" >
                          Clear
                      </option>
                      <option name="Status" className="form-control" >
                          Adverse
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateScreening}>Save</button>
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

export default UpdateScreeningComponent
