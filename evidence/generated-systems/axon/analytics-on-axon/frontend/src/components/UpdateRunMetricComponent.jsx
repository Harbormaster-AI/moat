import React, { Component } from 'react'
import RunMetricService from '../services/RunMetricService';

class UpdateRunMetricComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                value: ''
        }
        this.updateRunMetric = this.updateRunMetric.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changevalueHandler = this.changevalueHandler.bind(this);
    }

    componentDidMount(){
        RunMetricService.getRunMetricById(this.state.id).then( (res) =>{
            let runMetric = res.data;
            this.setState({
                name: runMetric.name,
                value: runMetric.value
            });
        });
    }

    updateRunMetric = (e) => {
        e.preventDefault();
        let runMetric = {
            runMetricId: this.state.id,
            name: this.state.name,
            value: this.state.value
        };
        console.log('runMetric => ' + JSON.stringify(runMetric));
        console.log('id => ' + JSON.stringify(this.state.id));
        RunMetricService.updateRunMetric(runMetric).then( res => {
            this.props.history.push('/runMetrics');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changevalueHandler= (event) => {
        this.setState({value: event.target.value});
    }

    cancel(){
        this.props.history.push('/runMetrics');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update RunMetric</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> value: </label>
                                                <input placeholder="value" name="value" className="form-control" value={this.state.value} onChange={this.changevalueHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateRunMetric}>Save</button>
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

export default UpdateRunMetricComponent
