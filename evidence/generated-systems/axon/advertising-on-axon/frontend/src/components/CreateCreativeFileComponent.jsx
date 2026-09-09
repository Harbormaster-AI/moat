import React, { Component } from 'react'
import CreativeFileService from '../services/CreativeFileService';

class CreateCreativeFileComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                uri: '',
                fileSizeKB: '',
                mimeType: '',
                checksum: ''
        }
        this.changeuriHandler = this.changeuriHandler.bind(this);
        this.changefileSizeKBHandler = this.changefileSizeKBHandler.bind(this);
        this.changemimeTypeHandler = this.changemimeTypeHandler.bind(this);
        this.changechecksumHandler = this.changechecksumHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateCreativeFile = (e) => {
        e.preventDefault();
        let creativeFile = {
                creativeFileId: this.state.id,
                uri: this.state.uri,
                fileSizeKB: this.state.fileSizeKB,
                mimeType: this.state.mimeType,
                checksum: this.state.checksum
            };
        console.log('creativeFile => ' + JSON.stringify(creativeFile));

        // step 5
        if(this.state.id === '_add'){
            creativeFile.creativeFileId=''
            CreativeFileService.createCreativeFile(creativeFile).then(res =>{
                this.props.history.push('/creativeFiles');
            });
        }else{
            CreativeFileService.updateCreativeFile(creativeFile).then( res => {
                this.props.history.push('/creativeFiles');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add CreativeFile</h3>
        }else{
            return <h3 className="text-center">Update CreativeFile</h3>
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
                                            <label> uri:&emsp; </label>
                                                <input placeholder="uri" name="uri" className="form-control" value={this.state.uri} onChange={this.changeuriHandler}/>

                                            <label> fileSizeKB:&emsp; </label>
                                                <input type="number" placeholder="fileSizeKB" name="fileSizeKB" className="form-control" value={this.state.fileSizeKB} onChange={this.changefileSizeKBHandler}/>

                                            <label> mimeType:&emsp; </label>
                                                <input placeholder="mimeType" name="mimeType" className="form-control" value={this.state.mimeType} onChange={this.changemimeTypeHandler}/>

                                            <label> checksum:&emsp; </label>
                                                <input placeholder="checksum" name="checksum" className="form-control" value={this.state.checksum} onChange={this.changechecksumHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateCreativeFile}>Save</button>
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

export default CreateCreativeFileComponent
