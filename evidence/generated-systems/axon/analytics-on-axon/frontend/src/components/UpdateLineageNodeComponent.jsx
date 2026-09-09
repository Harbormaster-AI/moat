import React, { Component } from 'react'
import LineageNodeService from '../services/LineageNodeService';

class UpdateLineageNodeComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                qualifiedName: '',
                nodeType: ''
        }
        this.updateLineageNode = this.updateLineageNode.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changequalifiedNameHandler = this.changequalifiedNameHandler.bind(this);
        this.changeNodeTypeHandler = this.changeNodeTypeHandler.bind(this);
    }

    componentDidMount(){
        LineageNodeService.getLineageNodeById(this.state.id).then( (res) =>{
            let lineageNode = res.data;
            this.setState({
                name: lineageNode.name,
                qualifiedName: lineageNode.qualifiedName,
                nodeType: lineageNode.nodeType
            });
        });
    }

    updateLineageNode = (e) => {
        e.preventDefault();
        let lineageNode = {
            lineageNodeId: this.state.id,
            name: this.state.name,
            qualifiedName: this.state.qualifiedName,
            nodeType: this.state.nodeType
        };
        console.log('lineageNode => ' + JSON.stringify(lineageNode));
        console.log('id => ' + JSON.stringify(this.state.id));
        LineageNodeService.updateLineageNode(lineageNode).then( res => {
            this.props.history.push('/lineageNodes');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changequalifiedNameHandler= (event) => {
        this.setState({qualifiedName: event.target.value});
    }
    changeNodeTypeHandler= (event) => {
        this.setState({nodeType: event.target.value});
    }

    cancel(){
        this.props.history.push('/lineageNodes');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update LineageNode</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> qualifiedName: </label>
                                                <input placeholder="qualifiedName" name="qualifiedName" className="form-control" value={this.state.qualifiedName} onChange={this.changequalifiedNameHandler}/>

                                            <label> NodeType: </label>
                                                <select value={this.state.nodeType} onChange={this.changeNodeTypeHandler}>
                      <option name="NodeType" className="form-control" >
                          Dataset
                      </option>
                      <option name="NodeType" className="form-control" >
                          Pipeline
                      </option>
                      <option name="NodeType" className="form-control" >
                          Model
                      </option>
                      <option name="NodeType" className="form-control" >
                          Dashboard
                      </option>
                      <option name="NodeType" className="form-control" >
                          Report
                      </option>
                      <option name="NodeType" className="form-control" >
                          FeatureSet
                      </option>
                      <option name="NodeType" className="form-control" >
                          Notebook
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateLineageNode}>Save</button>
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

export default UpdateLineageNodeComponent
