import React, { Component } from 'react'
import AnomalyService from '../services/AnomalyService';

class UpdateAnomalyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                occurredAt: '',
                details: '',
                anomalyType: '',
                severity: ''
        }
        this.updateAnomaly = this.updateAnomaly.bind(this);

        this.changeoccurredAtHandler = this.changeoccurredAtHandler.bind(this);
        this.changedetailsHandler = this.changedetailsHandler.bind(this);
        this.changeAnomalyTypeHandler = this.changeAnomalyTypeHandler.bind(this);
        this.changeSeverityHandler = this.changeSeverityHandler.bind(this);
    }

    componentDidMount(){
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

    updateAnomaly = (e) => {
        e.preventDefault();
        let anomaly = {
            anomalyId: this.state.id,
            occurredAt: this.state.occurredAt,
            details: this.state.details,
            anomalyType: this.state.anomalyType,
            severity: this.state.severity
        };
        console.log('anomaly => ' + JSON.stringify(anomaly));
        console.log('id => ' + JSON.stringify(this.state.id));
        AnomalyService.updateAnomaly(anomaly).then( res => {
            this.props.history.push('/anomalys');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Anomaly</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> occurredAt: </label>
                                                <input type="date" placeholder="occurredAt" name="occurredAt" className="form-control" value={this.state.occurredAt} onChange={this.changeoccurredAtHandler}/>

                                            <label> details: </label>
                                                <input placeholder="details" name="details" className="form-control" value={this.state.details} onChange={this.changedetailsHandler}/>

                                            <label> AnomalyType: </label>
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

                                            <label> Severity: </label>
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
                                        <button className="btn btn-success" onClick={this.updateAnomaly}>Save</button>
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

export default UpdateAnomalyComponent
