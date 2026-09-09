import React, { Component } from 'react'
import InferenceEndpointService from '../services/InferenceEndpointService'

class ListInferenceEndpointComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                inferenceEndpoints: []
        }
        this.addInferenceEndpoint = this.addInferenceEndpoint.bind(this);
        this.editInferenceEndpoint = this.editInferenceEndpoint.bind(this);
        this.deleteInferenceEndpoint = this.deleteInferenceEndpoint.bind(this);
    }

    deleteInferenceEndpoint(id){
        InferenceEndpointService.deleteInferenceEndpoint(id).then( res => {
            this.setState({inferenceEndpoints: this.state.inferenceEndpoints.filter(inferenceEndpoint => inferenceEndpoint.inferenceEndpointId !== id)});
        });
    }
    viewInferenceEndpoint(id){
        this.props.history.push(`/view-inferenceEndpoint/${id}`);
    }
    editInferenceEndpoint(id){
        this.props.history.push(`/add-inferenceEndpoint/${id}`);
    }

    componentDidMount(){
        InferenceEndpointService.getInferenceEndpoints().then((res) => {
            this.setState({ inferenceEndpoints: res.data});
        });
    }

    addInferenceEndpoint(){
        this.props.history.push('/add-inferenceEndpoint/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">InferenceEndpoint List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addInferenceEndpoint}> Add InferenceEndpoint</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> EndpointUrl </th>
                                    <th> TrafficShare </th>
                                    <th> Mode </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.inferenceEndpoints.map(
                                        inferenceEndpoint => 
                                        <tr key = {inferenceEndpoint.inferenceEndpointId}>
                                             <td> { inferenceEndpoint.name } </td>
                                             <td> { inferenceEndpoint.endpointUrl } </td>
                                             <td> { inferenceEndpoint.trafficShare } </td>
                                             <td> { inferenceEndpoint.mode } </td>
                                             <td>
                                                 <button onClick={ () => this.editInferenceEndpoint(inferenceEndpoint.inferenceEndpointId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteInferenceEndpoint(inferenceEndpoint.inferenceEndpointId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewInferenceEndpoint(inferenceEndpoint.inferenceEndpointId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListInferenceEndpointComponent
