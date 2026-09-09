import React, { Component } from 'react'
import DataTaskService from '../services/DataTaskService'

class ListDataTaskComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                dataTasks: []
        }
        this.addDataTask = this.addDataTask.bind(this);
        this.editDataTask = this.editDataTask.bind(this);
        this.deleteDataTask = this.deleteDataTask.bind(this);
    }

    deleteDataTask(id){
        DataTaskService.deleteDataTask(id).then( res => {
            this.setState({dataTasks: this.state.dataTasks.filter(dataTask => dataTask.dataTaskId !== id)});
        });
    }
    viewDataTask(id){
        this.props.history.push(`/view-dataTask/${id}`);
    }
    editDataTask(id){
        this.props.history.push(`/add-dataTask/${id}`);
    }

    componentDidMount(){
        DataTaskService.getDataTasks().then((res) => {
            this.setState({ dataTasks: res.data});
        });
    }

    addDataTask(){
        this.props.history.push('/add-dataTask/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">DataTask List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addDataTask}> Add DataTask</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Command </th>
                                    <th> Retries </th>
                                    <th> TaskType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.dataTasks.map(
                                        dataTask => 
                                        <tr key = {dataTask.dataTaskId}>
                                             <td> { dataTask.name } </td>
                                             <td> { dataTask.command } </td>
                                             <td> { dataTask.retries } </td>
                                             <td> { dataTask.taskType } </td>
                                             <td>
                                                 <button onClick={ () => this.editDataTask(dataTask.dataTaskId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteDataTask(dataTask.dataTaskId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewDataTask(dataTask.dataTaskId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListDataTaskComponent
