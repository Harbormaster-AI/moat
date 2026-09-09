import React, { Component } from 'react'
import AnomalyService from '../services/AnomalyService';

class CreateAnomalyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                occurredAt: '',
                details: '',
                anomalyType: '',
                severity: ''
        }
        this.changeoccurredAtHandler = this.changeoccurredAtHandler.bind(this);
        this.changedetailsHandler = this.changedetailsHandler.bind(this);
        this.changeAnomalyTypeHandler = this.changeAnomalyTypeHandler.bind(this);
        this.changeSeverityHandler = this.changeSeverityHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            AnomalyService.getAnomalyById(this.state.id).then( (res) =>{
                let anomaly = res.data;
                this.setState({
                    occurredAt: anomaly.occurredAt,
                    details: anomaly.details,
                    anomalyType: anomaly.anomalyType,
                    severity: anomaly.severity
                });
            });
        }        
    }
    saveOrUpdateAnomaly = (e) => {
        e.preventDefault();
        let anomaly = {
                anomalyId: this.state.id,
                occurredAt: this.state.occurredAt,
                details: this.state.details,
                anomalyType: this.state.anomalyType,
                severity: this.state.severity
            };
        console.log('anomaly => ' + JSON.stringify(anomaly));

        // step 5
        if(this.state.id === '_add'){
            anomaly.anomalyId=''
            AnomalyService.createAnomaly(anomaly).then(res =>{
                this.props.history.push('/anomalys');
            });
        }else{
            AnomalyService.updateAnomaly(anomaly).then( res => {
                this.props.history.push('/anomalys');
            });
        }
    }
    
    changeoccurredAtHandler= (event) => {
        this.setState({occurredAt: event.target.value});
    }
    changedetailsHandler= (event) => {
        this.setState({details: event.target.value});
    }
    changeAnomalyTypeHandler= (event) => {
        this.setState({anomalyType: event.target.value});
    }
    changeSeverityHandler= (event) => {
        this.setState({severity: event.target.value});
    }

    cancel(){
        this.props.history.push('/anomalys');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Anomaly</h3>
        }else{
            return <h3 className="text-center">Update Anomaly</h3>
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
                                            <label> occurredAt:&emsp; </label>
                                                <input type="date" placeholder="occurredAt" name="occurredAt" className="form-control" value={this.state.occurredAt} onChange={this.changeoccurredAtHandler}/>

                                            <label> details:&emsp; </label>
                                                <input placeholder="details" name="details" className="form-control" value={this.state.details} onChange={this.changedetailsHandler}/>

                                            <label> AnomalyType:&emsp; </label>
                                                <select value={this.state.anomalyType} onChange={this.changeAnomalyTypeHandler}>
                      <option name="AnomalyType" className="form-control" >
                          Spike
                      </option>
                      <option name="AnomalyType" className="form-control" >
                          Drop
                      </option>
                      <option name="AnomalyType" className="form-control" >
                          Drift
                      </option>
                      <option name="AnomalyType" className="form-control" >
                          Seasonal
                      </option>
                      <option name="AnomalyType" className="form-control" >
                          LevelShift
                      </option>
                    </select>

                                            <label> Severity:&emsp; </label>
                                                <select value={this.state.severity} onChange={this.changeSeverityHandler}>
                      <option name="Severity" className="form-control" >
                          Info
                      </option>
                      <option name="Severity" className="form-control" >
                          Warning
                      </option>
                      <option name="Severity" className="form-control" >
                          Critical
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateAnomaly}>Save</button>
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

export default CreateAnomalyComponent
