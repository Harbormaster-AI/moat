import React, { Component } from 'react'
import RecordsRepositoryService from '../services/RecordsRepositoryService';

class CreateRecordsRepositoryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                location: '',
                ownerDepartment: '',
                repositoryType: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changelocationHandler = this.changelocationHandler.bind(this);
        this.changeownerDepartmentHandler = this.changeownerDepartmentHandler.bind(this);
        this.changeRepositoryTypeHandler = this.changeRepositoryTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            RecordsRepositoryService.getRecordsRepositoryById(this.state.id).then( (res) =>{
                let recordsRepository = res.data;
                this.setState({
                    name: recordsRepository.name,
                    location: recordsRepository.location,
                    ownerDepartment: recordsRepository.ownerDepartment,
                    repositoryType: recordsRepository.repositoryType
                });
            });
        }        
    }
    saveOrUpdateRecordsRepository = (e) => {
        e.preventDefault();
        let recordsRepository = {
                recordsRepositoryId: this.state.id,
                name: this.state.name,
                location: this.state.location,
                ownerDepartment: this.state.ownerDepartment,
                repositoryType: this.state.repositoryType
            };
        console.log('recordsRepository => ' + JSON.stringify(recordsRepository));

        // step 5
        if(this.state.id === '_add'){
            recordsRepository.recordsRepositoryId=''
            RecordsRepositoryService.createRecordsRepository(recordsRepository).then(res =>{
                this.props.history.push('/recordsRepositorys');
            });
        }else{
            RecordsRepositoryService.updateRecordsRepository(recordsRepository).then( res => {
                this.props.history.push('/recordsRepositorys');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changelocationHandler= (event) => {
        this.setState({location: event.target.value});
    }
    changeownerDepartmentHandler= (event) => {
        this.setState({ownerDepartment: event.target.value});
    }
    changeRepositoryTypeHandler= (event) => {
        this.setState({repositoryType: event.target.value});
    }

    cancel(){
        this.props.history.push('/recordsRepositorys');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add RecordsRepository</h3>
        }else{
            return <h3 className="text-center">Update RecordsRepository</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> location:&emsp; </label>
                                                <input placeholder="location" name="location" className="form-control" value={this.state.location} onChange={this.changelocationHandler}/>

                                            <label> ownerDepartment:&emsp; </label>
                                                <input placeholder="ownerDepartment" name="ownerDepartment" className="form-control" value={this.state.ownerDepartment} onChange={this.changeownerDepartmentHandler}/>

                                            <label> RepositoryType:&emsp; </label>
                                                <select value={this.state.repositoryType} onChange={this.changeRepositoryTypeHandler}>
                      <option name="RepositoryType" className="form-control" >
                          DocumentManagement
                      </option>
                      <option name="RepositoryType" className="form-control" >
                          RecordsArchive
                      </option>
                      <option name="RepositoryType" className="form-control" >
                          EmailArchive
                      </option>
                      <option name="RepositoryType" className="form-control" >
                          FileShare
                      </option>
                      <option name="RepositoryType" className="form-control" >
                          ContentServices
                      </option>
                      <option name="RepositoryType" className="form-control" >
                          DataLake
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateRecordsRepository}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>
                   </div>
            </div>
        )
    }
}

export default CreateRecordsRepositoryComponent
