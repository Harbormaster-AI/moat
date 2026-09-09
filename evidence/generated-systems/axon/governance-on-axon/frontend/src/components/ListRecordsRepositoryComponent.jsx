import React, { Component } from 'react'
import RecordsRepositoryService from '../services/RecordsRepositoryService'

class ListRecordsRepositoryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                recordsRepositorys: []
        }
        this.addRecordsRepository = this.addRecordsRepository.bind(this);
        this.editRecordsRepository = this.editRecordsRepository.bind(this);
        this.deleteRecordsRepository = this.deleteRecordsRepository.bind(this);
    }

    deleteRecordsRepository(id){
        RecordsRepositoryService.deleteRecordsRepository(id).then( res => {
            this.setState({recordsRepositorys: this.state.recordsRepositorys.filter(recordsRepository => recordsRepository.recordsRepositoryId !== id)});
        });
    }
    viewRecordsRepository(id){
        this.props.history.push(`/view-recordsRepository/${id}`);
    }
    editRecordsRepository(id){
        this.props.history.push(`/add-recordsRepository/${id}`);
    }

    componentDidMount(){
        RecordsRepositoryService.getRecordsRepositorys().then((res) => {
            this.setState({ recordsRepositorys: res.data});
        });
    }

    addRecordsRepository(){
        this.props.history.push('/add-recordsRepository/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">RecordsRepository List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addRecordsRepository}> Add RecordsRepository</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Location </th>
                                    <th> OwnerDepartment </th>
                                    <th> RepositoryType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.recordsRepositorys.map(
                                        recordsRepository => 
                                        <tr key = {recordsRepository.recordsRepositoryId}>
                                             <td> { recordsRepository.name } </td>
                                             <td> { recordsRepository.location } </td>
                                             <td> { recordsRepository.ownerDepartment } </td>
                                             <td> { recordsRepository.repositoryType } </td>
                                             <td>
                                                 <button onClick={ () => this.editRecordsRepository(recordsRepository.recordsRepositoryId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteRecordsRepository(recordsRepository.recordsRepositoryId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewRecordsRepository(recordsRepository.recordsRepositoryId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListRecordsRepositoryComponent
