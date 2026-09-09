import React, { Component } from 'react'
import ModelVersionService from '../services/ModelVersionService'

class ListModelVersionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                modelVersions: []
        }
        this.addModelVersion = this.addModelVersion.bind(this);
        this.editModelVersion = this.editModelVersion.bind(this);
        this.deleteModelVersion = this.deleteModelVersion.bind(this);
    }

    deleteModelVersion(id){
        ModelVersionService.deleteModelVersion(id).then( res => {
            this.setState({modelVersions: this.state.modelVersions.filter(modelVersion => modelVersion.modelVersionId !== id)});
        });
    }
    viewModelVersion(id){
        this.props.history.push(`/view-modelVersion/${id}`);
    }
    editModelVersion(id){
        this.props.history.push(`/add-modelVersion/${id}`);
    }

    componentDidMount(){
        ModelVersionService.getModelVersions().then((res) => {
            this.setState({ modelVersions: res.data});
        });
    }

    addModelVersion(){
        this.props.history.push('/add-modelVersion/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ModelVersion List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addModelVersion}> Add ModelVersion</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Version </th>
                                    <th> Lifecycle </th>
                                    <th> TrainingStatus </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.modelVersions.map(
                                        modelVersion => 
                                        <tr key = {modelVersion.modelVersionId}>
                                             <td> { modelVersion.version } </td>
                                             <td> { modelVersion.lifecycle } </td>
                                             <td> { modelVersion.trainingStatus } </td>
                                             <td>
                                                 <button onClick={ () => this.editModelVersion(modelVersion.modelVersionId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteModelVersion(modelVersion.modelVersionId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewModelVersion(modelVersion.modelVersionId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListModelVersionComponent
