import React, { Component } from 'react'
import ScreeningService from '../services/ScreeningService';

class UpdateScreeningComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                score: '',
                screenedAt: '',
                screeningType: '',
                status: ''
        }
        this.updateScreening = this.updateScreening.bind(this);

        this.changescoreHandler = this.changescoreHandler.bind(this);
        this.changescreenedAtHandler = this.changescreenedAtHandler.bind(this);
        this.changeScreeningTypeHandler = this.changeScreeningTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        ScreeningService.getScreeningById(this.state.id).then( (res) =>{
            let screening = res.data;
            this.setState({
                score: screening.score,
                screenedAt: screening.screenedAt,
                screeningType: screening.screeningType,
                status: screening.status
            });
        });
    }

    updateScreening = (e) => {
        e.preventDefault();
        let screening = {
            screeningId: this.state.id,
            score: this.state.score,
            screenedAt: this.state.screenedAt,
            screeningType: this.state.screeningType,
            status: this.state.status
        };
        console.log('screening => ' + JSON.stringify(screening));
        console.log('id => ' + JSON.stringify(this.state.id));
        ScreeningService.updateScreening(screening).then( res => {
            this.props.history.push('/screenings');
        });
    }

    changescoreHandler= (event) => {
        this.setState({score: event.target.value});
    }
    changescreenedAtHandler= (event) => {
        this.setState({screenedAt: event.target.value});
    }
    changeScreeningTypeHandler= (event) => {
        this.setState({screeningType: event.target.value});
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
                                            <label> score: </label>
                                                <input placeholder="score" name="score" className="form-control" value={this.state.score} onChange={this.changescoreHandler}/>

                                            <label> screenedAt: </label>
                                                <input type="time" placeholder="screenedAt" name="screenedAt" className="form-control" value={this.state.screenedAt} onChange={this.changescreenedAtHandler}/>

                                            <label> ScreeningType: </label>
                                                <select value={this.state.screeningType} onChange={this.changeScreeningTypeHandler}>
                      <option name="ScreeningType" className="form-control" >
                          Sanctions
                      </option>
                      <option name="ScreeningType" className="form-control" >
                          PEP
                      </option>
                      <option name="ScreeningType" className="form-control" >
                          AdverseMedia
                      </option>
                    </select>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Clear
                      </option>
                      <option name="Status" className="form-control" >
                          Review
                      </option>
                      <option name="Status" className="form-control" >
                          Match
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
