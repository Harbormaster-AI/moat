import React, { Component } from 'react'
import MeasureService from '../services/MeasureService';

class CreateMeasureComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                format: '',
                aggregation: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeformatHandler = this.changeformatHandler.bind(this);
        this.changeAggregationHandler = this.changeAggregationHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            MeasureService.getMeasureById(this.state.id).then( (res) =>{
                let measure = res.data;
                this.setState({
                    name: measure.name,
                    format: measure.format,
                    aggregation: measure.aggregation
                });
            });
        }        
    }
    saveOrUpdateMeasure = (e) => {
        e.preventDefault();
        let measure = {
                measureId: this.state.id,
                name: this.state.name,
                format: this.state.format,
                aggregation: this.state.aggregation
            };
        console.log('measure => ' + JSON.stringify(measure));

        // step 5
        if(this.state.id === '_add'){
            measure.measureId=''
            MeasureService.createMeasure(measure).then(res =>{
                this.props.history.push('/measures');
            });
        }else{
            MeasureService.updateMeasure(measure).then( res => {
                this.props.history.push('/measures');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Measure</h3>
        }else{
            return <h3 className="text-center">Update Measure</h3>
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

                                            <label> format:&emsp; </label>
                                                <input placeholder="format" name="format" className="form-control" value={this.state.format} onChange={this.changeformatHandler}/>

                                            <label> Aggregation:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateMeasure}>Save</button>
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

export default CreateMeasureComponent
