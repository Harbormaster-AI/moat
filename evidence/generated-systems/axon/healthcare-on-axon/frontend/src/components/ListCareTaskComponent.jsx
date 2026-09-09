import React, { Component } from 'react'
import CareTaskService from '../services/CareTaskService'

class ListCareTaskComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                careTasks: []
        }
        this.addCareTask = this.addCareTask.bind(this);
        this.editCareTask = this.editCareTask.bind(this);
        this.deleteCareTask = this.deleteCareTask.bind(this);
    }

    deleteCareTask(id){
        CareTaskService.deleteCareTask(id).then( res => {
            this.setState({careTasks: this.state.careTasks.filter(careTask => careTask.careTaskId !== id)});
        });
    }
    viewCareTask(id){
        this.props.history.push(`/view-careTask/${id}`);
    }
    editCareTask(id){
        this.props.history.push(`/add-careTask/${id}`);
    }

    componentDidMount(){
        CareTaskService.getCareTasks().then((res) => {
            this.setState({ careTasks: res.data});
        });
    }

    addCareTask(){
        this.props.history.push('/add-careTask/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">CareTask List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCareTask}> Add CareTask</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Description </th>
                                    <th> DueDate </th>
                                    <th> Status </th>
                                    <th> Priority </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.careTasks.map(
                                        careTask => 
                                        <tr key = {careTask.careTaskId}>
                                             <td> { careTask.description } </td>
                                             <td> { careTask.dueDate } </td>
                                             <td> { careTask.status } </td>
                                             <td> { careTask.priority } </td>
                                             <td>
                                                 <button onClick={ () => this.editCareTask(careTask.careTaskId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCareTask(careTask.careTaskId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCareTask(careTask.careTaskId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListCareTaskComponent
