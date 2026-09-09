import React, { Component } from 'react'
import SemanticModelService from '../services/SemanticModelService'

class ListSemanticModelComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                semanticModels: []
        }
        this.addSemanticModel = this.addSemanticModel.bind(this);
        this.editSemanticModel = this.editSemanticModel.bind(this);
        this.deleteSemanticModel = this.deleteSemanticModel.bind(this);
    }

    deleteSemanticModel(id){
        SemanticModelService.deleteSemanticModel(id).then( res => {
            this.setState({semanticModels: this.state.semanticModels.filter(semanticModel => semanticModel.semanticModelId !== id)});
        });
    }
    viewSemanticModel(id){
        this.props.history.push(`/view-semanticModel/${id}`);
    }
    editSemanticModel(id){
        this.props.history.push(`/add-semanticModel/${id}`);
    }

    componentDidMount(){
        SemanticModelService.getSemanticModels().then((res) => {
            this.setState({ semanticModels: res.data});
        });
    }

    addSemanticModel(){
        this.props.history.push('/add-semanticModel/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">SemanticModel List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addSemanticModel}> Add SemanticModel</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Version </th>
                                    <th> Grain </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.semanticModels.map(
                                        semanticModel => 
                                        <tr key = {semanticModel.semanticModelId}>
                                             <td> { semanticModel.name } </td>
                                             <td> { semanticModel.version } </td>
                                             <td> { semanticModel.grain } </td>
                                             <td>
                                                 <button onClick={ () => this.editSemanticModel(semanticModel.semanticModelId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteSemanticModel(semanticModel.semanticModelId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewSemanticModel(semanticModel.semanticModelId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListSemanticModelComponent
