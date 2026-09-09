import React, { Component } from 'react'
import CreativeFileService from '../services/CreativeFileService'

class ViewCreativeFileComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            creativeFile: {}
        }
    }

    componentDidMount(){
        CreativeFileService.getCreativeFileById(this.state.id).then( res => {
            this.setState({creativeFile: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View CreativeFile Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> uri:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.creativeFile.uri }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> fileSizeKB:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.creativeFile.fileSizeKB }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> mimeType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.creativeFile.mimeType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> checksum:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.creativeFile.checksum }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewCreativeFileComponent
