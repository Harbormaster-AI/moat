import React, { Component } from 'react'
import SemanticModelService from '../services/SemanticModelService';

class UpdateSemanticModelComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                version: '',
                grain: ''
        }
        this.updateSemanticModel = this.updateSemanticModel.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeversionHandler = this.changeversionHandler.bind(this);
        this.changegrainHandler = this.changegrainHandler.bind(this);
    }

    componentDidMount(){
        SemanticModelService.getSemanticModelById(this.state.id).then( (res) =>{
            let semanticModel = res.data;
            this.setState({
                name: semanticModel.name,
                version: semanticModel.version,
                grain: semanticModel.grain
            });
        });
    }

    updateSemanticModel = (e) => {
        e.preventDefault();
        let semanticModel = {
            semanticModelId: this.state.id,
            name: this.state.name,
            version: this.state.version,
            grain: this.state.grain
        };
        console.log('semanticModel => ' + JSON.stringify(semanticModel));
        console.log('id => ' + JSON.stringify(this.state.id));
        SemanticModelService.updateSemanticModel(semanticModel).then( res => {
            this.props.history.push('/semanticModels');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update SemanticModel</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> version: </label>
                                                <input placeholder="version" name="version" className="form-control" value={this.state.version} onChange={this.changeversionHandler}/>

                                            <label> grain: </label>
                                                <input placeholder="grain" name="grain" className="form-control" value={this.state.grain} onChange={this.changegrainHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateSemanticModel}>Save</button>
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

export default UpdateSemanticModelComponent
