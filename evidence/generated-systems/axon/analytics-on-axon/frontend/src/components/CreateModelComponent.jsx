import React, { Component } from 'react'
import ModelService from '../services/ModelService';

class CreateModelComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                taskDescription: '',
                modelType: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changetaskDescriptionHandler = this.changetaskDescriptionHandler.bind(this);
        this.changeModelTypeHandler = this.changeModelTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ModelService.getModelById(this.state.id).then( (res) =>{
                let model = res.data;
                this.setState({
                    name: model.name,
                    taskDescription: model.taskDescription,
                    modelType: model.modelType
                });
            });
        }        
    }
    saveOrUpdateModel = (e) => {
        e.preventDefault();
        let model = {
                modelId: this.state.id,
                name: this.state.name,
                taskDescription: this.state.taskDescription,
                modelType: this.state.modelType
            };
        console.log('model => ' + JSON.stringify(model));

        // step 5
        if(this.state.id === '_add'){
            model.modelId=''
            ModelService.createModel(model).then(res =>{
                this.props.history.push('/models');
            });
        }else{
            ModelService.updateModel(model).then( res => {
                this.props.history.push('/models');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changetaskDescriptionHandler= (event) => {
        this.setState({taskDescription: event.target.value});
    }
    changeModelTypeHandler= (event) => {
        this.setState({modelType: event.target.value});
    }

    cancel(){
        this.props.history.push('/models');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Model</h3>
        }else{
            return <h3 className="text-center">Update Model</h3>
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

                                            <label> taskDescription:&emsp; </label>
                                                <input placeholder="taskDescription" name="taskDescription" className="form-control" value={this.state.taskDescription} onChange={this.changetaskDescriptionHandler}/>

                                            <label> ModelType:&emsp; </label>
                                                <select value={this.state.modelType} onChange={this.changeModelTypeHandler}>
                      <option name="ModelType" className="form-control" >
                          Classification
                      </option>
                      <option name="ModelType" className="form-control" >
                          Regression
                      </option>
                      <option name="ModelType" className="form-control" >
                          Clustering
                      </option>
                      <option name="ModelType" className="form-control" >
                          Forecasting
                      </option>
                      <option name="ModelType" className="form-control" >
                          Ranking
                      </option>
                      <option name="ModelType" className="form-control" >
                          NLP
                      </option>
                      <option name="ModelType" className="form-control" >
                          ComputerVision
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateModel}>Save</button>
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

export default CreateModelComponent
