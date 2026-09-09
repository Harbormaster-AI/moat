import React, { Component } from 'react'
import SemanticModelService from '../services/SemanticModelService';

class CreateSemanticModelComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                version: '',
                grain: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeversionHandler = this.changeversionHandler.bind(this);
        this.changegrainHandler = this.changegrainHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            SemanticModelService.getSemanticModelById(this.state.id).then( (res) =>{
                let semanticModel = res.data;
                this.setState({
                    name: semanticModel.name,
                    version: semanticModel.version,
                    grain: semanticModel.grain
                });
            });
        }        
    }
    saveOrUpdateSemanticModel = (e) => {
        e.preventDefault();
        let semanticModel = {
                semanticModelId: this.state.id,
                name: this.state.name,
                version: this.state.version,
                grain: this.state.grain
            };
        console.log('semanticModel => ' + JSON.stringify(semanticModel));

        // step 5
        if(this.state.id === '_add'){
            semanticModel.semanticModelId=''
            SemanticModelService.createSemanticModel(semanticModel).then(res =>{
                this.props.history.push('/semanticModels');
            });
        }else{
            SemanticModelService.updateSemanticModel(semanticModel).then( res => {
                this.props.history.push('/semanticModels');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeversionHandler= (event) => {
        this.setState({version: event.target.value});
    }
    changegrainHandler= (event) => {
        this.setState({grain: event.target.value});
    }

    cancel(){
        this.props.history.push('/semanticModels');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add SemanticModel</h3>
        }else{
            return <h3 className="text-center">Update SemanticModel</h3>
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

                                            <label> version:&emsp; </label>
                                                <input placeholder="version" name="version" className="form-control" value={this.state.version} onChange={this.changeversionHandler}/>

                                            <label> grain:&emsp; </label>
                                                <input placeholder="grain" name="grain" className="form-control" value={this.state.grain} onChange={this.changegrainHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateSemanticModel}>Save</button>
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

export default CreateSemanticModelComponent
