import React, { Component } from 'react'
import MetricService from '../services/MetricService';

class CreateMetricComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                expression: '',
                unit: '',
                metricType: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeexpressionHandler = this.changeexpressionHandler.bind(this);
        this.changeunitHandler = this.changeunitHandler.bind(this);
        this.changeMetricTypeHandler = this.changeMetricTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateMetric = (e) => {
        e.preventDefault();
        let metric = {
                metricId: this.state.id,
                name: this.state.name,
                expression: this.state.expression,
                unit: this.state.unit,
                metricType: this.state.metricType
            };
        console.log('metric => ' + JSON.stringify(metric));

        // step 5
        if(this.state.id === '_add'){
            metric.metricId=''
            MetricService.createMetric(metric).then(res =>{
                this.props.history.push('/metrics');
            });
        }else{
            MetricService.updateMetric(metric).then( res => {
                this.props.history.push('/metrics');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Metric</h3>
        }else{
            return <h3 className="text-center">Update Metric</h3>
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
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> expression:&emsp; </label>
                                                <input placeholder="expression" name="expression" className="form-control" value={this.state.expression} onChange={this.changeexpressionHandler}/>

                                            <label> unit:&emsp; </label>
                                                <input placeholder="unit" name="unit" className="form-control" value={this.state.unit} onChange={this.changeunitHandler}/>

                                            <label> MetricType:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateMetric}>Save</button>
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

export default CreateMetricComponent
