import React, { Component } from 'react'
import LabResultService from '../services/LabResultService';

class UpdateLabResultComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                resultCode: '',
                issuedDate: '',
                status: ''
        }
        this.updateLabResult = this.updateLabResult.bind(this);

        this.changeresultCodeHandler = this.changeresultCodeHandler.bind(this);
        this.changeissuedDateHandler = this.changeissuedDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        LabResultService.getLabResultById(this.state.id).then( (res) =>{
            let labResult = res.data;
            this.setState({
                resultCode: labResult.resultCode,
                issuedDate: labResult.issuedDate,
                status: labResult.status
            });
        });
    }

    updateLabResult = (e) => {
        e.preventDefault();
        let labResult = {
            labResultId: this.state.id,
            resultCode: this.state.resultCode,
            issuedDate: this.state.issuedDate,
            status: this.state.status
        };
        console.log('labResult => ' + JSON.stringify(labResult));
        console.log('id => ' + JSON.stringify(this.state.id));
        LabResultService.updateLabResult(labResult).then( res => {
            this.props.history.push('/labResults');
        });
    }

    changeresultCodeHandler= (event) => {
        this.setState({resultCode: event.target.value});
    }
    changeissuedDateHandler= (event) => {
        this.setState({issuedDate: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/labResults');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update LabResult</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> resultCode: </label>
                                                <input placeholder="resultCode" name="resultCode" className="form-control" value={this.state.resultCode} onChange={this.changeresultCodeHandler}/>

                                            <label> issuedDate: </label>
                                                <input type="time" placeholder="issuedDate" name="issuedDate" className="form-control" value={this.state.issuedDate} onChange={this.changeissuedDateHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Registered
                      </option>
                      <option name="Status" className="form-control" >
                          Partial
                      </option>
                      <option name="Status" className="form-control" >
                          Final
                      </option>
                      <option name="Status" className="form-control" >
                          Corrected
                      </option>
                      <option name="Status" className="form-control" >
                          Cancelled
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateLabResult}>Save</button>
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

export default UpdateLabResultComponent
