import React, { Component } from 'react'
import LineageNodeService from '../services/LineageNodeService'

class ViewLineageNodeComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            lineageNode: {}
        }
    }

    componentDidMount(){
        LineageNodeService.getLineageNodeById(this.state.id).then( res => {
            this.setState({lineageNode: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View LineageNode Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.lineageNode.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> qualifiedName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.lineageNode.qualifiedName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> NodeType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.lineageNode.nodeType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewLineageNodeComponent
