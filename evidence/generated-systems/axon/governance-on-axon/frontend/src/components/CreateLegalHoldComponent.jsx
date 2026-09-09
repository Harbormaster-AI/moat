import React, { Component } from 'react'
import LegalHoldService from '../services/LegalHoldService';

class CreateLegalHoldComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                reason: '',
                issuedDate: '',
                releaseDate: '',
                holdStatus: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changereasonHandler = this.changereasonHandler.bind(this);
        this.changeissuedDateHandler = this.changeissuedDateHandler.bind(this);
        this.changereleaseDateHandler = this.changereleaseDateHandler.bind(this);
        this.changeHoldStatusHandler = this.changeHoldStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateLegalHold = (e) => {
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

        // step 5
        if(this.state.id === '_add'){
            legalHold.legalHoldId=''
            LegalHoldService.createLegalHold(legalHold).then(res =>{
                this.props.history.push('/legalHolds');
            });
        }else{
            LegalHoldService.updateLegalHold(legalHold).then( res => {
                this.props.history.push('/legalHolds');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add LegalHold</h3>
        }else{
            return <h3 className="text-center">Update LegalHold</h3>
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

                                            <label> reason:&emsp; </label>
                                                <input placeholder="reason" name="reason" className="form-control" value={this.state.reason} onChange={this.changereasonHandler}/>

                                            <label> issuedDate:&emsp; </label>
                                                <input type="date" placeholder="issuedDate" name="issuedDate" className="form-control" value={this.state.issuedDate} onChange={this.changeissuedDateHandler}/>

                                            <label> releaseDate:&emsp; </label>
                                                <input type="date" placeholder="releaseDate" name="releaseDate" className="form-control" value={this.state.releaseDate} onChange={this.changereleaseDateHandler}/>

                                            <label> HoldStatus:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateLegalHold}>Save</button>
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

export default CreateLegalHoldComponent
