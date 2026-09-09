import React, { Component } from 'react'
import KPIService from '../services/KPIService';

class CreateKPIComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                targetValue: '',
                metricType: ''
        }
        this.changetargetValueHandler = this.changetargetValueHandler.bind(this);
        this.changeMetricTypeHandler = this.changeMetricTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            KPIService.getKPIById(this.state.id).then( (res) =>{
                let kPI = res.data;
                this.setState({
                    targetValue: kPI.targetValue,
                    metricType: kPI.metricType
                });
            });
        }        
    }
    saveOrUpdateKPI = (e) => {
        e.preventDefault();
        let kPI = {
                kPIId: this.state.id,
                targetValue: this.state.targetValue,
                metricType: this.state.metricType
            };
        console.log('kPI => ' + JSON.stringify(kPI));

        // step 5
        if(this.state.id === '_add'){
            kPI.kPIId=''
            KPIService.createKPI(kPI).then(res =>{
                this.props.history.push('/kPIs');
            });
        }else{
            KPIService.updateKPI(kPI).then( res => {
                this.props.history.push('/kPIs');
            });
        }
    }
    
    changetargetValueHandler= (event) => {
        this.setState({targetValue: event.target.value});
    }
    changeMetricTypeHandler= (event) => {
        this.setState({metricType: event.target.value});
    }

    cancel(){
        this.props.history.push('/kPIs');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add KPI</h3>
        }else{
            return <h3 className="text-center">Update KPI</h3>
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
                                            <label> targetValue:&emsp; </label>
                                                <input placeholder="targetValue" name="targetValue" className="form-control" value={this.state.targetValue} onChange={this.changetargetValueHandler}/>

                                            <label> MetricType:&emsp; </label>
                                                <select value={this.state.metricType} onChange={this.changeMetricTypeHandler}>
                      <option name="MetricType" className="form-control" >
                          Impressions
                      </option>
                      <option name="MetricType" className="form-control" >
                          ViewableImpressions
                      </option>
                      <option name="MetricType" className="form-control" >
                          Clicks
                      </option>
                      <option name="MetricType" className="form-control" >
                          CTR
                      </option>
                      <option name="MetricType" className="form-control" >
                          Reach
                      </option>
                      <option name="MetricType" className="form-control" >
                          Frequency
                      </option>
                      <option name="MetricType" className="form-control" >
                          VideoStarts
                      </option>
                      <option name="MetricType" className="form-control" >
                          VideoCompletions
                      </option>
                      <option name="MetricType" className="form-control" >
                          AvgViewTime
                      </option>
                      <option name="MetricType" className="form-control" >
                          Conversions
                      </option>
                      <option name="MetricType" className="form-control" >
                          ViewThroughConversions
                      </option>
                      <option name="MetricType" className="form-control" >
                          Spend
                      </option>
                      <option name="MetricType" className="form-control" >
                          CPM
                      </option>
                      <option name="MetricType" className="form-control" >
                          CPC
                      </option>
                      <option name="MetricType" className="form-control" >
                          CPA
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateKPI}>Save</button>
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

export default CreateKPIComponent
