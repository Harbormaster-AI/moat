import React, { Component } from 'react'
import LineageNodeService from '../services/LineageNodeService'

class ListLineageNodeComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                lineageNodes: []
        }
        this.addLineageNode = this.addLineageNode.bind(this);
        this.editLineageNode = this.editLineageNode.bind(this);
        this.deleteLineageNode = this.deleteLineageNode.bind(this);
    }

    deleteLineageNode(id){
        LineageNodeService.deleteLineageNode(id).then( res => {
            this.setState({lineageNodes: this.state.lineageNodes.filter(lineageNode => lineageNode.lineageNodeId !== id)});
        });
    }
    viewLineageNode(id){
        this.props.history.push(`/view-lineageNode/${id}`);
    }
    editLineageNode(id){
        this.props.history.push(`/add-lineageNode/${id}`);
    }

    componentDidMount(){
        LineageNodeService.getLineageNodes().then((res) => {
            this.setState({ lineageNodes: res.data});
        });
    }

    addLineageNode(){
        this.props.history.push('/add-lineageNode/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">LineageNode List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addLineageNode}> Add LineageNode</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> QualifiedName </th>
                                    <th> NodeType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.lineageNodes.map(
                                        lineageNode => 
                                        <tr key = {lineageNode.lineageNodeId}>
                                             <td> { lineageNode.name } </td>
                                             <td> { lineageNode.qualifiedName } </td>
                                             <td> { lineageNode.nodeType } </td>
                                             <td>
                                                 <button onClick={ () => this.editLineageNode(lineageNode.lineageNodeId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteLineageNode(lineageNode.lineageNodeId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewLineageNode(lineageNode.lineageNodeId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListLineageNodeComponent
