import React, { Component } from 'react'
import VisualizationService from '../services/VisualizationService';

class CreateVisualizationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                title: '',
                options: '',
                chartType: ''
        }
        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changeoptionsHandler = this.changeoptionsHandler.bind(this);
        this.changeChartTypeHandler = this.changeChartTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            VisualizationService.getVisualizationById(this.state.id).then( (res) =>{
                let visualization = res.data;
                this.setState({
                    title: visualization.title,
                    options: visualization.options,
                    chartType: visualization.chartType
                });
            });
        }        
    }
    saveOrUpdateVisualization = (e) => {
        e.preventDefault();
        let visualization = {
                visualizationId: this.state.id,
                title: this.state.title,
                options: this.state.options,
                chartType: this.state.chartType
            };
        console.log('visualization => ' + JSON.stringify(visualization));

        // step 5
        if(this.state.id === '_add'){
            visualization.visualizationId=''
            VisualizationService.createVisualization(visualization).then(res =>{
                this.props.history.push('/visualizations');
            });
        }else{
            VisualizationService.updateVisualization(visualization).then( res => {
                this.props.history.push('/visualizations');
            });
        }
    }
    
    changetitleHandler= (event) => {
        this.setState({title: event.target.value});
    }
    changeoptionsHandler= (event) => {
        this.setState({options: event.target.value});
    }
    changeChartTypeHandler= (event) => {
        this.setState({chartType: event.target.value});
    }

    cancel(){
        this.props.history.push('/visualizations');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Visualization</h3>
        }else{
            return <h3 className="text-center">Update Visualization</h3>
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
                                            <label> title:&emsp; </label>
                                                <input placeholder="title" name="title" className="form-control" value={this.state.title} onChange={this.changetitleHandler}/>

                                            <label> options:&emsp; </label>
                                                <input placeholder="options" name="options" className="form-control" value={this.state.options} onChange={this.changeoptionsHandler}/>

                                            <label> ChartType:&emsp; </label>
                                                <select value={this.state.chartType} onChange={this.changeChartTypeHandler}>
                      <option name="ChartType" className="form-control" >
                          Table
                      </option>
                      <option name="ChartType" className="form-control" >
                          Bar
                      </option>
                      <option name="ChartType" className="form-control" >
                          Line
                      </option>
                      <option name="ChartType" className="form-control" >
                          Area
                      </option>
                      <option name="ChartType" className="form-control" >
                          Pie
                      </option>
                      <option name="ChartType" className="form-control" >
                          Scatter
                      </option>
                      <option name="ChartType" className="form-control" >
                          Heatmap
                      </option>
                      <option name="ChartType" className="form-control" >
                          KPI
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateVisualization}>Save</button>
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

export default CreateVisualizationComponent
