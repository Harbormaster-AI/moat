import React, { Component } from 'react'
import RunParameterService from '../services/RunParameterService'

class ListRunParameterComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                runParameters: []
        }
        this.addRunParameter = this.addRunParameter.bind(this);
        this.editRunParameter = this.editRunParameter.bind(this);
        this.deleteRunParameter = this.deleteRunParameter.bind(this);
    }

    deleteRunParameter(id){
        RunParameterService.deleteRunParameter(id).then( res => {
            this.setState({runParameters: this.state.runParameters.filter(runParameter => runParameter.runParameterId !== id)});
        });
    }
    viewRunParameter(id){
        this.props.history.push(`/view-runParameter/${id}`);
    }
    editRunParameter(id){
        this.props.history.push(`/add-runParameter/${id}`);
    }

    componentDidMount(){
        RunParameterService.getRunParameters().then((res) => {
            this.setState({ runParameters: res.data});
        });
    }

    addRunParameter(){
        this.props.history.push('/add-runParameter/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">RunParameter List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addRunParameter}> Add RunParameter</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Value </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.runParameters.map(
                                        runParameter => 
                                        <tr key = {runParameter.runParameterId}>
                                             <td> { runParameter.name } </td>
                                             <td> { runParameter.value } </td>
                                             <td>
                                                 <button onClick={ () => this.editRunParameter(runParameter.runParameterId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteRunParameter(runParameter.runParameterId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewRunParameter(runParameter.runParameterId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListRunParameterComponent
