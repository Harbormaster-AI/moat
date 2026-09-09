import React, { Component } from 'react'
import PerformanceMetricService from '../services/PerformanceMetricService';

class UpdatePerformanceMetricComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                date: '',
                value: '',
                metricType: ''
        }
        this.updatePerformanceMetric = this.updatePerformanceMetric.bind(this);

        this.changedateHandler = this.changedateHandler.bind(this);
        this.changevalueHandler = this.changevalueHandler.bind(this);
        this.changeMetricTypeHandler = this.changeMetricTypeHandler.bind(this);
    }

    componentDidMount(){
        PerformanceMetricService.getPerformanceMetricById(this.state.id).then( (res) =>{
            let performanceMetric = res.data;
            this.setState({
                date: performanceMetric.date,
                value: performanceMetric.value,
                metricType: performanceMetric.metricType
            });
        });
    }

    updatePerformanceMetric = (e) => {
        e.preventDefault();
        let performanceMetric = {
            performanceMetricId: this.state.id,
            date: this.state.date,
            value: this.state.value,
            metricType: this.state.metricType
        };
        console.log('performanceMetric => ' + JSON.stringify(performanceMetric));
        console.log('id => ' + JSON.stringify(this.state.id));
        PerformanceMetricService.updatePerformanceMetric(performanceMetric).then( res => {
            this.props.history.push('/performanceMetrics');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update PerformanceMetric</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> date: </label>
                                                <input type="date" placeholder="date" name="date" className="form-control" value={this.state.date} onChange={this.changedateHandler}/>

                                            <label> value: </label>
                                                <input placeholder="value" name="value" className="form-control" value={this.state.value} onChange={this.changevalueHandler}/>

                                            <label> MetricType: </label>
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
                                        <button className="btn btn-success" onClick={this.updatePerformanceMetric}>Save</button>
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

export default UpdatePerformanceMetricComponent
