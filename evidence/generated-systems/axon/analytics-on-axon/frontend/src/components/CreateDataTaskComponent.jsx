import React, { Component } from 'react'
import DataTaskService from '../services/DataTaskService';

class CreateDataTaskComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                command: '',
                retries: '',
                taskType: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changecommandHandler = this.changecommandHandler.bind(this);
        this.changeretriesHandler = this.changeretriesHandler.bind(this);
        this.changeTaskTypeHandler = this.changeTaskTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            DataTaskService.getDataTaskById(this.state.id).then( (res) =>{
                let dataTask = res.data;
                this.setState({
                    name: dataTask.name,
                    command: dataTask.command,
                    retries: dataTask.retries,
                    taskType: dataTask.taskType
                });
            });
        }        
    }
    saveOrUpdateDataTask = (e) => {
        e.preventDefault();
        let dataTask = {
                dataTaskId: this.state.id,
                name: this.state.name,
                command: this.state.command,
                retries: this.state.retries,
                taskType: this.state.taskType
            };
        console.log('dataTask => ' + JSON.stringify(dataTask));

        // step 5
        if(this.state.id === '_add'){
            dataTask.dataTaskId=''
            DataTaskService.createDataTask(dataTask).then(res =>{
                this.props.history.push('/dataTasks');
            });
        }else{
            DataTaskService.updateDataTask(dataTask).then( res => {
                this.props.history.push('/dataTasks');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changecommandHandler= (event) => {
        this.setState({command: event.target.value});
    }
    changeretriesHandler= (event) => {
        this.setState({retries: event.target.value});
    }
    changeTaskTypeHandler= (event) => {
        this.setState({taskType: event.target.value});
    }

    cancel(){
        this.props.history.push('/dataTasks');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add DataTask</h3>
        }else{
            return <h3 className="text-center">Update DataTask</h3>
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

                                            <label> command:&emsp; </label>
                                                <input placeholder="command" name="command" className="form-control" value={this.state.command} onChange={this.changecommandHandler}/>

                                            <label> retries:&emsp; </label>
                                                <input type="number" placeholder="retries" name="retries" className="form-control" value={this.state.retries} onChange={this.changeretriesHandler}/>

                                            <label> TaskType:&emsp; </label>
                                                <select value={this.state.taskType} onChange={this.changeTaskTypeHandler}>
                      <option name="TaskType" className="form-control" >
                          Extract
                      </option>
                      <option name="TaskType" className="form-control" >
                          Transform
                      </option>
                      <option name="TaskType" className="form-control" >
                          Load
                      </option>
                      <option name="TaskType" className="form-control" >
                          Validate
                      </option>
                      <option name="TaskType" className="form-control" >
                          Enrich
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateDataTask}>Save</button>
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

export default CreateDataTaskComponent
