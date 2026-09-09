import React, { Component } from 'react'
import ExperimentVariantService from '../services/ExperimentVariantService'

class ListExperimentVariantComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                experimentVariants: []
        }
        this.addExperimentVariant = this.addExperimentVariant.bind(this);
        this.editExperimentVariant = this.editExperimentVariant.bind(this);
        this.deleteExperimentVariant = this.deleteExperimentVariant.bind(this);
    }

    deleteExperimentVariant(id){
        ExperimentVariantService.deleteExperimentVariant(id).then( res => {
            this.setState({experimentVariants: this.state.experimentVariants.filter(experimentVariant => experimentVariant.experimentVariantId !== id)});
        });
    }
    viewExperimentVariant(id){
        this.props.history.push(`/view-experimentVariant/${id}`);
    }
    editExperimentVariant(id){
        this.props.history.push(`/add-experimentVariant/${id}`);
    }

    componentDidMount(){
        ExperimentVariantService.getExperimentVariants().then((res) => {
            this.setState({ experimentVariants: res.data});
        });
    }

    addExperimentVariant(){
        this.props.history.push('/add-experimentVariant/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ExperimentVariant List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addExperimentVariant}> Add ExperimentVariant</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Allocation </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.experimentVariants.map(
                                        experimentVariant => 
                                        <tr key = {experimentVariant.experimentVariantId}>
                                             <td> { experimentVariant.name } </td>
                                             <td> { experimentVariant.allocation } </td>
                                             <td>
                                                 <button onClick={ () => this.editExperimentVariant(experimentVariant.experimentVariantId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteExperimentVariant(experimentVariant.experimentVariantId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewExperimentVariant(experimentVariant.experimentVariantId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListExperimentVariantComponent
