import React, { Component } from 'react'
import ModelService from '../services/ModelService'

class ListModelComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                models: []
        }
        this.addModel = this.addModel.bind(this);
        this.editModel = this.editModel.bind(this);
        this.deleteModel = this.deleteModel.bind(this);
    }

    deleteModel(id){
        ModelService.deleteModel(id).then( res => {
            this.setState({models: this.state.models.filter(model => model.modelId !== id)});
        });
    }
    viewModel(id){
        this.props.history.push(`/view-model/${id}`);
    }
    editModel(id){
        this.props.history.push(`/add-model/${id}`);
    }

    componentDidMount(){
        ModelService.getModels().then((res) => {
            this.setState({ models: res.data});
        });
    }

    addModel(){
        this.props.history.push('/add-model/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Model List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addModel}> Add Model</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> TaskDescription </th>
                                    <th> ModelType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.models.map(
                                        model => 
                                        <tr key = {model.modelId}>
                                             <td> { model.name } </td>
                                             <td> { model.taskDescription } </td>
                                             <td> { model.modelType } </td>
                                             <td>
                                                 <button onClick={ () => this.editModel(model.modelId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteModel(model.modelId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewModel(model.modelId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListModelComponent
