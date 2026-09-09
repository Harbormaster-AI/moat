import React, { Component } from 'react'
import CreativeFileService from '../services/CreativeFileService';

class UpdateCreativeFileComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                uri: '',
                fileSizeKB: '',
                mimeType: '',
                checksum: ''
        }
        this.updateCreativeFile = this.updateCreativeFile.bind(this);

        this.changeuriHandler = this.changeuriHandler.bind(this);
        this.changefileSizeKBHandler = this.changefileSizeKBHandler.bind(this);
        this.changemimeTypeHandler = this.changemimeTypeHandler.bind(this);
        this.changechecksumHandler = this.changechecksumHandler.bind(this);
    }

    componentDidMount(){
        CreativeFileService.getCreativeFileById(this.state.id).then( (res) =>{
            let creativeFile = res.data;
            this.setState({
                uri: creativeFile.uri,
                fileSizeKB: creativeFile.fileSizeKB,
                mimeType: creativeFile.mimeType,
                checksum: creativeFile.checksum
            });
        });
    }

    updateCreativeFile = (e) => {
        e.preventDefault();
        let creativeFile = {
            creativeFileId: this.state.id,
            uri: this.state.uri,
            fileSizeKB: this.state.fileSizeKB,
            mimeType: this.state.mimeType,
            checksum: this.state.checksum
        };
        console.log('creativeFile => ' + JSON.stringify(creativeFile));
        console.log('id => ' + JSON.stringify(this.state.id));
        CreativeFileService.updateCreativeFile(creativeFile).then( res => {
            this.props.history.push('/creativeFiles');
        });
    }

    changeuriHandler= (event) => {
        this.setState({uri: event.target.value});
    }
    changefileSizeKBHandler= (event) => {
        this.setState({fileSizeKB: event.target.value});
    }
    changemimeTypeHandler= (event) => {
        this.setState({mimeType: event.target.value});
    }
    changechecksumHandler= (event) => {
        this.setState({checksum: event.target.value});
    }

    cancel(){
        this.props.history.push('/creativeFiles');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update CreativeFile</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> uri: </label>
                                                <input placeholder="uri" name="uri" className="form-control" value={this.state.uri} onChange={this.changeuriHandler}/>

                                            <label> fileSizeKB: </label>
                                                <input type="number" placeholder="fileSizeKB" name="fileSizeKB" className="form-control" value={this.state.fileSizeKB} onChange={this.changefileSizeKBHandler}/>

                                            <label> mimeType: </label>
                                                <input placeholder="mimeType" name="mimeType" className="form-control" value={this.state.mimeType} onChange={this.changemimeTypeHandler}/>

                                            <label> checksum: </label>
                                                <input placeholder="checksum" name="checksum" className="form-control" value={this.state.checksum} onChange={this.changechecksumHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateCreativeFile}>Save</button>
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

export default UpdateCreativeFileComponent
