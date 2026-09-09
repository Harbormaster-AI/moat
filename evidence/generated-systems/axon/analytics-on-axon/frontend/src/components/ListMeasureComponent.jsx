import React, { Component } from 'react'
import MeasureService from '../services/MeasureService'

class ListMeasureComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                measures: []
        }
        this.addMeasure = this.addMeasure.bind(this);
        this.editMeasure = this.editMeasure.bind(this);
        this.deleteMeasure = this.deleteMeasure.bind(this);
    }

    deleteMeasure(id){
        MeasureService.deleteMeasure(id).then( res => {
            this.setState({measures: this.state.measures.filter(measure => measure.measureId !== id)});
        });
    }
    viewMeasure(id){
        this.props.history.push(`/view-measure/${id}`);
    }
    editMeasure(id){
        this.props.history.push(`/add-measure/${id}`);
    }

    componentDidMount(){
        MeasureService.getMeasures().then((res) => {
            this.setState({ measures: res.data});
        });
    }

    addMeasure(){
        this.props.history.push('/add-measure/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Measure List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addMeasure}> Add Measure</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Format </th>
                                    <th> Aggregation </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.measures.map(
                                        measure => 
                                        <tr key = {measure.measureId}>
                                             <td> { measure.name } </td>
                                             <td> { measure.format } </td>
                                             <td> { measure.aggregation } </td>
                                             <td>
                                                 <button onClick={ () => this.editMeasure(measure.measureId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteMeasure(measure.measureId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewMeasure(measure.measureId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListMeasureComponent
