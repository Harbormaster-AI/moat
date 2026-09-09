import React, { Component } from 'react'
import VisualizationService from '../services/VisualizationService'

class ListVisualizationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                visualizations: []
        }
        this.addVisualization = this.addVisualization.bind(this);
        this.editVisualization = this.editVisualization.bind(this);
        this.deleteVisualization = this.deleteVisualization.bind(this);
    }

    deleteVisualization(id){
        VisualizationService.deleteVisualization(id).then( res => {
            this.setState({visualizations: this.state.visualizations.filter(visualization => visualization.visualizationId !== id)});
        });
    }
    viewVisualization(id){
        this.props.history.push(`/view-visualization/${id}`);
    }
    editVisualization(id){
        this.props.history.push(`/add-visualization/${id}`);
    }

    componentDidMount(){
        VisualizationService.getVisualizations().then((res) => {
            this.setState({ visualizations: res.data});
        });
    }

    addVisualization(){
        this.props.history.push('/add-visualization/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Visualization List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addVisualization}> Add Visualization</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Title </th>
                                    <th> Options </th>
                                    <th> ChartType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.visualizations.map(
                                        visualization => 
                                        <tr key = {visualization.visualizationId}>
                                             <td> { visualization.title } </td>
                                             <td> { visualization.options } </td>
                                             <td> { visualization.chartType } </td>
                                             <td>
                                                 <button onClick={ () => this.editVisualization(visualization.visualizationId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteVisualization(visualization.visualizationId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewVisualization(visualization.visualizationId)} className="btn btn-outline-info btn-sm">View </button>
                                             </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>

                 </div>

            </div>
        )
    }
}

export default ListVisualizationComponent
