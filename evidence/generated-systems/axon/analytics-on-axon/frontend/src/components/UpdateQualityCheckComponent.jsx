import React, { Component } from 'react'
import QualityCheckService from '../services/QualityCheckService';

class UpdateQualityCheckComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                checkedAt: '',
                observedValue: '',
                sampleSize: '',
                status: ''
        }
        this.updateQualityCheck = this.updateQualityCheck.bind(this);

        this.changecheckedAtHandler = this.changecheckedAtHandler.bind(this);
        this.changeobservedValueHandler = this.changeobservedValueHandler.bind(this);
        this.changesampleSizeHandler = this.changesampleSizeHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        QualityCheckService.getQualityCheckById(this.state.id).then( (res) =>{
            let qualityCheck = res.data;
            this.setState({
                checkedAt: qualityCheck.checkedAt,
                observedValue: qualityCheck.observedValue,
                sampleSize: qualityCheck.sampleSize,
                status: qualityCheck.status
            });
        });
    }

    updateQualityCheck = (e) => {
        e.preventDefault();
        let qualityCheck = {
            qualityCheckId: this.state.id,
            checkedAt: this.state.checkedAt,
            observedValue: this.state.observedValue,
            sampleSize: this.state.sampleSize,
            status: this.state.status
        };
        console.log('qualityCheck => ' + JSON.stringify(qualityCheck));
        console.log('id => ' + JSON.stringify(this.state.id));
        QualityCheckService.updateQualityCheck(qualityCheck).then( res => {
            this.props.history.push('/qualityChecks');
        });
    }

    changecheckedAtHandler= (event) => {
        this.setState({checkedAt: event.target.value});
    }
    changeobservedValueHandler= (event) => {
        this.setState({observedValue: event.target.value});
    }
    changesampleSizeHandler= (event) => {
        this.setState({sampleSize: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/qualityChecks');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update QualityCheck</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> checkedAt: </label>
                                                <input type="date" placeholder="checkedAt" name="checkedAt" className="form-control" value={this.state.checkedAt} onChange={this.changecheckedAtHandler}/>

                                            <label> observedValue: </label>
                                                <input placeholder="observedValue" name="observedValue" className="form-control" value={this.state.observedValue} onChange={this.changeobservedValueHandler}/>

                                            <label> sampleSize: </label>
                                                <input type="number" placeholder="sampleSize" name="sampleSize" className="form-control" value={this.state.sampleSize} onChange={this.changesampleSizeHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Passed
                      </option>
                      <option name="Status" className="form-control" >
                          Failed
                      </option>
                      <option name="Status" className="form-control" >
                          Warning
                      </option>
                      <option name="Status" className="form-control" >
                          Skipped
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateQualityCheck}>Save</button>
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

export default UpdateQualityCheckComponent
