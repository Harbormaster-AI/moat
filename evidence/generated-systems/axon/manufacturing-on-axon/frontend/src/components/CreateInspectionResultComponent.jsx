import React, { Component } from 'react'
import InspectionResultService from '../services/InspectionResultService';

class CreateInspectionResultComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                resultValue: '',
                recordedOn: '',
                notes: '',
                resultStatus: ''
        }
        this.changeresultValueHandler = this.changeresultValueHandler.bind(this);
        this.changerecordedOnHandler = this.changerecordedOnHandler.bind(this);
        this.changenotesHandler = this.changenotesHandler.bind(this);
        this.changeResultStatusHandler = this.changeResultStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            InspectionResultService.getInspectionResultById(this.state.id).then( (res) =>{
                let inspectionResult = res.data;
                this.setState({
                    resultValue: inspectionResult.resultValue,
                    recordedOn: inspectionResult.recordedOn,
                    notes: inspectionResult.notes,
                    resultStatus: inspectionResult.resultStatus
                });
            });
        }        
    }
    saveOrUpdateInspectionResult = (e) => {
        e.preventDefault();
        let inspectionResult = {
                inspectionResultId: this.state.id,
                resultValue: this.state.resultValue,
                recordedOn: this.state.recordedOn,
                notes: this.state.notes,
                resultStatus: this.state.resultStatus
            };
        console.log('inspectionResult => ' + JSON.stringify(inspectionResult));

        // step 5
        if(this.state.id === '_add'){
            inspectionResult.inspectionResultId=''
            InspectionResultService.createInspectionResult(inspectionResult).then(res =>{
                this.props.history.push('/inspectionResults');
            });
        }else{
            InspectionResultService.updateInspectionResult(inspectionResult).then( res => {
                this.props.history.push('/inspectionResults');
            });
        }
    }
    
    changeresultValueHandler= (event) => {
        this.setState({resultValue: event.target.value});
    }
    changerecordedOnHandler= (event) => {
        this.setState({recordedOn: event.target.value});
    }
    changenotesHandler= (event) => {
        this.setState({notes: event.target.value});
    }
    changeResultStatusHandler= (event) => {
        this.setState({resultStatus: event.target.value});
    }

    cancel(){
        this.props.history.push('/inspectionResults');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add InspectionResult</h3>
        }else{
            return <h3 className="text-center">Update InspectionResult</h3>
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
                                            <label> resultValue:&emsp; </label>
                                                <input placeholder="resultValue" name="resultValue" className="form-control" value={this.state.resultValue} onChange={this.changeresultValueHandler}/>

                                            <label> recordedOn:&emsp; </label>
                                                <input type="time" placeholder="recordedOn" name="recordedOn" className="form-control" value={this.state.recordedOn} onChange={this.changerecordedOnHandler}/>

                                            <label> notes:&emsp; </label>
                                                <input placeholder="notes" name="notes" className="form-control" value={this.state.notes} onChange={this.changenotesHandler}/>

                                            <label> ResultStatus:&emsp; </label>
                                                <select value={this.state.resultStatus} onChange={this.changeResultStatusHandler}>
                      <option name="ResultStatus" className="form-control" >
                          Pass
                      </option>
                      <option name="ResultStatus" className="form-control" >
                          Fail
                      </option>
                      <option name="ResultStatus" className="form-control" >
                          Rework
                      </option>
                      <option name="ResultStatus" className="form-control" >
                          Scrap
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateInspectionResult}>Save</button>
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

export default CreateInspectionResultComponent
