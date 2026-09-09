import React, { Component } from 'react'
import LegalHoldService from '../services/LegalHoldService';

class UpdateLegalHoldComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                reason: '',
                issuedDate: '',
                releaseDate: '',
                holdStatus: ''
        }
        this.updateLegalHold = this.updateLegalHold.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changereasonHandler = this.changereasonHandler.bind(this);
        this.changeissuedDateHandler = this.changeissuedDateHandler.bind(this);
        this.changereleaseDateHandler = this.changereleaseDateHandler.bind(this);
        this.changeHoldStatusHandler = this.changeHoldStatusHandler.bind(this);
    }

    componentDidMount(){
        LegalHoldService.getLegalHoldById(this.state.id).then( (res) =>{
            let legalHold = res.data;
            this.setState({
                name: legalHold.name,
                reason: legalHold.reason,
                issuedDate: legalHold.issuedDate,
                releaseDate: legalHold.releaseDate,
                holdStatus: legalHold.holdStatus
            });
        });
    }

    updateLegalHold = (e) => {
        e.preventDefault();
        let legalHold = {
            legalHoldId: this.state.id,
            name: this.state.name,
            reason: this.state.reason,
            issuedDate: this.state.issuedDate,
            releaseDate: this.state.releaseDate,
            holdStatus: this.state.holdStatus
        };
        console.log('legalHold => ' + JSON.stringify(legalHold));
        console.log('id => ' + JSON.stringify(this.state.id));
        LegalHoldService.updateLegalHold(legalHold).then( res => {
            this.props.history.push('/legalHolds');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changereasonHandler= (event) => {
        this.setState({reason: event.target.value});
    }
    changeissuedDateHandler= (event) => {
        this.setState({issuedDate: event.target.value});
    }
    changereleaseDateHandler= (event) => {
        this.setState({releaseDate: event.target.value});
    }
    changeHoldStatusHandler= (event) => {
        this.setState({holdStatus: event.target.value});
    }

    cancel(){
        this.props.history.push('/legalHolds');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update LegalHold</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> reason: </label>
                                                <input placeholder="reason" name="reason" className="form-control" value={this.state.reason} onChange={this.changereasonHandler}/>

                                            <label> issuedDate: </label>
                                                <input type="date" placeholder="issuedDate" name="issuedDate" className="form-control" value={this.state.issuedDate} onChange={this.changeissuedDateHandler}/>

                                            <label> releaseDate: </label>
                                                <input type="date" placeholder="releaseDate" name="releaseDate" className="form-control" value={this.state.releaseDate} onChange={this.changereleaseDateHandler}/>

                                            <label> HoldStatus: </label>
                                                <select value={this.state.holdStatus} onChange={this.changeHoldStatusHandler}>
                      <option name="HoldStatus" className="form-control" >
                          Active
                      </option>
                      <option name="HoldStatus" className="form-control" >
                          Released
                      </option>
                      <option name="HoldStatus" className="form-control" >
                          Superseded
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateLegalHold}>Save</button>
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

export default UpdateLegalHoldComponent
