import React, { Component } from 'react'
import LineageNodeService from '../services/LineageNodeService';

class CreateLineageNodeComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                qualifiedName: '',
                nodeType: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changequalifiedNameHandler = this.changequalifiedNameHandler.bind(this);
        this.changeNodeTypeHandler = this.changeNodeTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            LineageNodeService.getLineageNodeById(this.state.id).then( (res) =>{
                let lineageNode = res.data;
                this.setState({
                    name: lineageNode.name,
                    qualifiedName: lineageNode.qualifiedName,
                    nodeType: lineageNode.nodeType
                });
            });
        }        
    }
    saveOrUpdateLineageNode = (e) => {
        e.preventDefault();
        let lineageNode = {
                lineageNodeId: this.state.id,
                name: this.state.name,
                qualifiedName: this.state.qualifiedName,
                nodeType: this.state.nodeType
            };
        console.log('lineageNode => ' + JSON.stringify(lineageNode));

        // step 5
        if(this.state.id === '_add'){
            lineageNode.lineageNodeId=''
            LineageNodeService.createLineageNode(lineageNode).then(res =>{
                this.props.history.push('/lineageNodes');
            });
        }else{
            LineageNodeService.updateLineageNode(lineageNode).then( res => {
                this.props.history.push('/lineageNodes');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add LineageNode</h3>
        }else{
            return <h3 className="text-center">Update LineageNode</h3>
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

                                            <label> qualifiedName:&emsp; </label>
                                                <input placeholder="qualifiedName" name="qualifiedName" className="form-control" value={this.state.qualifiedName} onChange={this.changequalifiedNameHandler}/>

                                            <label> NodeType:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateLineageNode}>Save</button>
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

export default CreateLineageNodeComponent
