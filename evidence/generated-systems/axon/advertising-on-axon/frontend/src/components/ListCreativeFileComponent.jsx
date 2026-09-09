import React, { Component } from 'react'
import CreativeFileService from '../services/CreativeFileService'

class ListCreativeFileComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                creativeFiles: []
        }
        this.addCreativeFile = this.addCreativeFile.bind(this);
        this.editCreativeFile = this.editCreativeFile.bind(this);
        this.deleteCreativeFile = this.deleteCreativeFile.bind(this);
    }

    deleteCreativeFile(id){
        CreativeFileService.deleteCreativeFile(id).then( res => {
            this.setState({creativeFiles: this.state.creativeFiles.filter(creativeFile => creativeFile.creativeFileId !== id)});
        });
    }
    viewCreativeFile(id){
        this.props.history.push(`/view-creativeFile/${id}`);
    }
    editCreativeFile(id){
        this.props.history.push(`/add-creativeFile/${id}`);
    }

    componentDidMount(){
        CreativeFileService.getCreativeFiles().then((res) => {
            this.setState({ creativeFiles: res.data});
        });
    }

    addCreativeFile(){
        this.props.history.push('/add-creativeFile/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">CreativeFile List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addCreativeFile}> Add CreativeFile</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Uri </th>
                                    <th> FileSizeKB </th>
                                    <th> MimeType </th>
                                    <th> Checksum </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.creativeFiles.map(
                                        creativeFile => 
                                        <tr key = {creativeFile.creativeFileId}>
                                             <td> { creativeFile.uri } </td>
                                             <td> { creativeFile.fileSizeKB } </td>
                                             <td> { creativeFile.mimeType } </td>
                                             <td> { creativeFile.checksum } </td>
                                             <td>
                                                 <button onClick={ () => this.editCreativeFile(creativeFile.creativeFileId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteCreativeFile(creativeFile.creativeFileId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewCreativeFile(creativeFile.creativeFileId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListCreativeFileComponent
