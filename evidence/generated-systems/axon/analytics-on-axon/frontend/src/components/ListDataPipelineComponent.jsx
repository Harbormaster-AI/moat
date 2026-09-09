import React, { Component } from 'react'
import DataPipelineService from '../services/DataPipelineService'

class ListDataPipelineComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                dataPipelines: []
        }
        this.addDataPipeline = this.addDataPipeline.bind(this);
        this.editDataPipeline = this.editDataPipeline.bind(this);
        this.deleteDataPipeline = this.deleteDataPipeline.bind(this);
    }

    deleteDataPipeline(id){
        DataPipelineService.deleteDataPipeline(id).then( res => {
            this.setState({dataPipelines: this.state.dataPipelines.filter(dataPipeline => dataPipeline.dataPipelineId !== id)});
        });
    }
    viewDataPipeline(id){
        this.props.history.push(`/view-dataPipeline/${id}`);
    }
    editDataPipeline(id){
        this.props.history.push(`/add-dataPipeline/${id}`);
    }

    componentDidMount(){
        DataPipelineService.getDataPipelines().then((res) => {
            this.setState({ dataPipelines: res.data});
        });
    }

    addDataPipeline(){
        this.props.history.push('/add-dataPipeline/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">DataPipeline List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addDataPipeline}> Add DataPipeline</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Schedule </th>
                                    <th> TriggerType </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.dataPipelines.map(
                                        dataPipeline => 
                                        <tr key = {dataPipeline.dataPipelineId}>
                                             <td> { dataPipeline.name } </td>
                                             <td> { dataPipeline.schedule } </td>
                                             <td> { dataPipeline.triggerType } </td>
                                             <td> { dataPipeline.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editDataPipeline(dataPipeline.dataPipelineId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteDataPipeline(dataPipeline.dataPipelineId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewDataPipeline(dataPipeline.dataPipelineId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListDataPipelineComponent
