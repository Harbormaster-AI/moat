import React, { Component } from 'react'
import EngineTypeService from '../services/EngineTypeService'

class ListEngineTypeComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                engineTypes: []
        }
        this.addEngineType = this.addEngineType.bind(this);
        this.editEngineType = this.editEngineType.bind(this);
        this.deleteEngineType = this.deleteEngineType.bind(this);
    }

    deleteEngineType(id){
        EngineTypeService.deleteEngineType(id).then( res => {
            this.setState({engineTypes: this.state.engineTypes.filter(engineType => engineType.engineTypeId !== id)});
        });
    }
    viewEngineType(id){
        this.props.history.push(`/view-engineType/${id}`);
    }
    editEngineType(id){
        this.props.history.push(`/add-engineType/${id}`);
    }

    componentDidMount(){
        EngineTypeService.getEngineTypes().then((res) => {
            this.setState({ engineTypes: res.data});
        });
    }

    addEngineType(){
        this.props.history.push('/add-engineType/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">EngineType List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addEngineType}> Add EngineType</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> EngineModelCode </th>
                                    <th> MaxThrustKn </th>
                                    <th> Category </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.engineTypes.map(
                                        engineType => 
                                        <tr key = {engineType.engineTypeId}>
                                             <td> { engineType.engineModelCode } </td>
                                             <td> { engineType.maxThrustKn } </td>
                                             <td> { engineType.category } </td>
                                             <td>
                                                 <button onClick={ () => this.editEngineType(engineType.engineTypeId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteEngineType(engineType.engineTypeId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewEngineType(engineType.engineTypeId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListEngineTypeComponent
