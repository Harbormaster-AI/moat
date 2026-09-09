import React, { Component } from 'react'
import MetricService from '../services/MetricService';

class UpdateMetricComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                expression: '',
                unit: '',
                metricType: ''
        }
        this.updateMetric = this.updateMetric.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeexpressionHandler = this.changeexpressionHandler.bind(this);
        this.changeunitHandler = this.changeunitHandler.bind(this);
        this.changeMetricTypeHandler = this.changeMetricTypeHandler.bind(this);
    }

    componentDidMount(){
        MetricService.getMetricById(this.state.id).then( (res) =>{
            let metric = res.data;
            this.setState({
                name: metric.name,
                expression: metric.expression,
                unit: metric.unit,
                metricType: metric.metricType
            });
        });
    }

    updateMetric = (e) => {
        e.preventDefault();
        let metric = {
            metricId: this.state.id,
            name: this.state.name,
            expression: this.state.expression,
            unit: this.state.unit,
            metricType: this.state.metricType
        };
        console.log('metric => ' + JSON.stringify(metric));
        console.log('id => ' + JSON.stringify(this.state.id));
        MetricService.updateMetric(metric).then( res => {
            this.props.history.push('/metrics');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeexpressionHandler= (event) => {
        this.setState({expression: event.target.value});
    }
    changeunitHandler= (event) => {
        this.setState({unit: event.target.value});
    }
    changeMetricTypeHandler= (event) => {
        this.setState({metricType: event.target.value});
    }

    cancel(){
        this.props.history.push('/metrics');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Metric</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> expression: </label>
                                                <input placeholder="expression" name="expression" className="form-control" value={this.state.expression} onChange={this.changeexpressionHandler}/>

                                            <label> unit: </label>
                                                <input placeholder="unit" name="unit" className="form-control" value={this.state.unit} onChange={this.changeunitHandler}/>

                                            <label> MetricType: </label>
                                                <select value={this.state.metricType} onChange={this.changeMetricTypeHandler}>
                      <option name="MetricType" className="form-control" >
                          Ratio
                      </option>
                      <option name="MetricType" className="form-control" >
                          Rate
                      </option>
                      <option name="MetricType" className="form-control" >
                          Count
                      </option>
                      <option name="MetricType" className="form-control" >
                          Percentage
                      </option>
                      <option name="MetricType" className="form-control" >
                          Index
                      </option>
                      <option name="MetricType" className="form-control" >
                          Score
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateMetric}>Save</button>
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

export default UpdateMetricComponent
