import React, { Component } from 'react'
import MeasureService from '../services/MeasureService';

class UpdateMeasureComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                format: '',
                aggregation: ''
        }
        this.updateMeasure = this.updateMeasure.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeformatHandler = this.changeformatHandler.bind(this);
        this.changeAggregationHandler = this.changeAggregationHandler.bind(this);
    }

    componentDidMount(){
        MeasureService.getMeasureById(this.state.id).then( (res) =>{
            let measure = res.data;
            this.setState({
                name: measure.name,
                format: measure.format,
                aggregation: measure.aggregation
            });
        });
    }

    updateMeasure = (e) => {
        e.preventDefault();
        let measure = {
            measureId: this.state.id,
            name: this.state.name,
            format: this.state.format,
            aggregation: this.state.aggregation
        };
        console.log('measure => ' + JSON.stringify(measure));
        console.log('id => ' + JSON.stringify(this.state.id));
        MeasureService.updateMeasure(measure).then( res => {
            this.props.history.push('/measures');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeformatHandler= (event) => {
        this.setState({format: event.target.value});
    }
    changeAggregationHandler= (event) => {
        this.setState({aggregation: event.target.value});
    }

    cancel(){
        this.props.history.push('/measures');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Measure</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> format: </label>
                                                <input placeholder="format" name="format" className="form-control" value={this.state.format} onChange={this.changeformatHandler}/>

                                            <label> Aggregation: </label>
                                                <select value={this.state.aggregation} onChange={this.changeAggregationHandler}>
                      <option name="Aggregation" className="form-control" >
                          Sum
                      </option>
                      <option name="Aggregation" className="form-control" >
                          Average
                      </option>
                      <option name="Aggregation" className="form-control" >
                          Min
                      </option>
                      <option name="Aggregation" className="form-control" >
                          Max
                      </option>
                      <option name="Aggregation" className="form-control" >
                          Median
                      </option>
                      <option name="Aggregation" className="form-control" >
                          Count
                      </option>
                      <option name="Aggregation" className="form-control" >
                          DistinctCount
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateMeasure}>Save</button>
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

export default UpdateMeasureComponent
