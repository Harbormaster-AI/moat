import React, { Component } from 'react'
import LabResultService from '../services/LabResultService';

class CreateLabResultComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                resultCode: '',
                issuedDate: '',
                status: ''
        }
        this.changeresultCodeHandler = this.changeresultCodeHandler.bind(this);
        this.changeissuedDateHandler = this.changeissuedDateHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            LabResultService.getLabResultById(this.state.id).then( (res) =>{
                let labResult = res.data;
                this.setState({
                    resultCode: labResult.resultCode,
                    issuedDate: labResult.issuedDate,
                    status: labResult.status
                });
            });
        }        
    }
    saveOrUpdateLabResult = (e) => {
        e.preventDefault();
        let labResult = {
                labResultId: this.state.id,
                resultCode: this.state.resultCode,
                issuedDate: this.state.issuedDate,
                status: this.state.status
            };
        console.log('labResult => ' + JSON.stringify(labResult));

        // step 5
        if(this.state.id === '_add'){
            labResult.labResultId=''
            LabResultService.createLabResult(labResult).then(res =>{
                this.props.history.push('/labResults');
            });
        }else{
            LabResultService.updateLabResult(labResult).then( res => {
                this.props.history.push('/labResults');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add LabResult</h3>
        }else{
            return <h3 className="text-center">Update LabResult</h3>
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
                                            <label> resultCode:&emsp; </label>
                                                <input placeholder="resultCode" name="resultCode" className="form-control" value={this.state.resultCode} onChange={this.changeresultCodeHandler}/>

                                            <label> issuedDate:&emsp; </label>
                                                <input type="time" placeholder="issuedDate" name="issuedDate" className="form-control" value={this.state.issuedDate} onChange={this.changeissuedDateHandler}/>

                                            <label> Status:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateLabResult}>Save</button>
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

export default CreateLabResultComponent
