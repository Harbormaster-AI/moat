import React, { Component } from 'react'
import CoverageDefinitionService from '../services/CoverageDefinitionService'

class ListCoverageDefinitionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                coverageDefinitions: []
        }
        this.addCoverageDefinition = this.addCoverageDefinition.bind(this);
        this.editCoverageDefinition = this.editCoverageDefinition.bind(this);
        this.deleteCoverageDefinition = this.deleteCoverageDefinition.bind(this);
    }

    deleteCoverageDefinition(id){
        CoverageDefinitionService.deleteCoverageDefinition(id).then( res => {
            this.setState({coverageDefinitions: this.state.coverageDefinitions.filter(coverageDefinition => coverageDefinition.coverageDefinitionId !== id)});
        });
    }
    viewCoverageDefinition(id){
        this.props.history.push(`/view-coverageDefinition/${id}`);
    }
    editCoverageDefinition(id){
        this.props.history.push(`/add-coverageDefinition/${id}`);
    }

    componentDidMount(){
        CoverageDefinitionService.getCoverageDefinitions().then((res) => {
            this.setState({ coverageDefinitions: res.data});
        });
    }

    addCoverageDefinition(){
        this.props.history.push('/add-coverageDefinition/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">CoverageDefinition List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCoverageDefinition}> Add CoverageDefinition</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> DefaultLimit </th>
                                    <th> DefaultDeductible </th>
                                    <th> AsMandatory </th>
                                    <th> CoverageType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.coverageDefinitions.map(
                                        coverageDefinition => 
                                        <tr key = {coverageDefinition.coverageDefinitionId}>
                                             <td> { coverageDefinition.name } </td>
                                             <td> { coverageDefinition.defaultLimit } </td>
                                             <td> { coverageDefinition.defaultDeductible } </td>
                                             <td> { coverageDefinition.asMandatory } </td>
                                             <td> { coverageDefinition.coverageType } </td>
                                             <td>
                                                 <button onClick={ () => this.editCoverageDefinition(coverageDefinition.coverageDefinitionId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCoverageDefinition(coverageDefinition.coverageDefinitionId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCoverageDefinition(coverageDefinition.coverageDefinitionId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListCoverageDefinitionComponent
