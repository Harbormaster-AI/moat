import React, { Component } from 'react'
import PerformanceMetricService from '../services/PerformanceMetricService';

class CreatePerformanceMetricComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                date: '',
                value: '',
                metricType: ''
        }
        this.changedateHandler = this.changedateHandler.bind(this);
        this.changevalueHandler = this.changevalueHandler.bind(this);
        this.changeMetricTypeHandler = this.changeMetricTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            PerformanceMetricService.getPerformanceMetricById(this.state.id).then( (res) =>{
                let performanceMetric = res.data;
                this.setState({
                    date: performanceMetric.date,
                    value: performanceMetric.value,
                    metricType: performanceMetric.metricType
                });
            });
        }        
    }
    saveOrUpdatePerformanceMetric = (e) => {
        e.preventDefault();
        let performanceMetric = {
                performanceMetricId: this.state.id,
                date: this.state.date,
                value: this.state.value,
                metricType: this.state.metricType
            };
        console.log('performanceMetric => ' + JSON.stringify(performanceMetric));

        // step 5
        if(this.state.id === '_add'){
            performanceMetric.performanceMetricId=''
            PerformanceMetricService.createPerformanceMetric(performanceMetric).then(res =>{
                this.props.history.push('/performanceMetrics');
            });
        }else{
            PerformanceMetricService.updatePerformanceMetric(performanceMetric).then( res => {
                this.props.history.push('/performanceMetrics');
            });
        }
    }
    
    changedateHandler= (event) => {
        this.setState({date: event.target.value});
    }
    changevalueHandler= (event) => {
        this.setState({value: event.target.value});
    }
    changeMetricTypeHandler= (event) => {
        this.setState({metricType: event.target.value});
    }

    cancel(){
        this.props.history.push('/performanceMetrics');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add PerformanceMetric</h3>
        }else{
            return <h3 className="text-center">Update PerformanceMetric</h3>
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
                                            <label> date:&emsp; </label>
                                                <input type="date" placeholder="date" name="date" className="form-control" value={this.state.date} onChange={this.changedateHandler}/>

                                            <label> value:&emsp; </label>
                                                <input placeholder="value" name="value" className="form-control" value={this.state.value} onChange={this.changevalueHandler}/>

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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdatePerformanceMetric}>Save</button>
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

export default CreatePerformanceMetricComponent
