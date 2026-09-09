import React, { Component } from 'react'
import ScreeningService from '../services/ScreeningService';

class CreateScreeningComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                score: '',
                screenedAt: '',
                screeningType: '',
                status: ''
        }
        this.changescoreHandler = this.changescoreHandler.bind(this);
        this.changescreenedAtHandler = this.changescreenedAtHandler.bind(this);
        this.changeScreeningTypeHandler = this.changeScreeningTypeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateScreening = (e) => {
        e.preventDefault();
        let screening = {
                screeningId: this.state.id,
                score: this.state.score,
                screenedAt: this.state.screenedAt,
                screeningType: this.state.screeningType,
                status: this.state.status
            };
        console.log('screening => ' + JSON.stringify(screening));

        // step 5
        if(this.state.id === '_add'){
            screening.screeningId=''
            ScreeningService.createScreening(screening).then(res =>{
                this.props.history.push('/screenings');
            });
        }else{
            ScreeningService.updateScreening(screening).then( res => {
                this.props.history.push('/screenings');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Screening</h3>
        }else{
            return <h3 className="text-center">Update Screening</h3>
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
                                            <label> score:&emsp; </label>
                                                <input placeholder="score" name="score" className="form-control" value={this.state.score} onChange={this.changescoreHandler}/>

                                            <label> screenedAt:&emsp; </label>
                                                <input type="time" placeholder="screenedAt" name="screenedAt" className="form-control" value={this.state.screenedAt} onChange={this.changescreenedAtHandler}/>

                                            <label> ScreeningType:&emsp; </label>
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

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateScreening}>Save</button>
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

export default CreateScreeningComponent
