import React, { Component } from 'react'
import AuditWorkpaperService from '../services/AuditWorkpaperService'

class ListAuditWorkpaperComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                auditWorkpapers: []
        }
        this.addAuditWorkpaper = this.addAuditWorkpaper.bind(this);
        this.editAuditWorkpaper = this.editAuditWorkpaper.bind(this);
        this.deleteAuditWorkpaper = this.deleteAuditWorkpaper.bind(this);
    }

    deleteAuditWorkpaper(id){
        AuditWorkpaperService.deleteAuditWorkpaper(id).then( res => {
            this.setState({auditWorkpapers: this.state.auditWorkpapers.filter(auditWorkpaper => auditWorkpaper.auditWorkpaperId !== id)});
        });
    }
    viewAuditWorkpaper(id){
        this.props.history.push(`/view-auditWorkpaper/${id}`);
    }
    editAuditWorkpaper(id){
        this.props.history.push(`/add-auditWorkpaper/${id}`);
    }

    componentDidMount(){
        AuditWorkpaperService.getAuditWorkpapers().then((res) => {
            this.setState({ auditWorkpapers: res.data});
        });
    }

    addAuditWorkpaper(){
        this.props.history.push('/add-auditWorkpaper/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">AuditWorkpaper List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAuditWorkpaper}> Add AuditWorkpaper</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> WorkpaperRef </th>
                                    <th> Subject </th>
                                    <th> WorkpaperUrl </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.auditWorkpapers.map(
                                        auditWorkpaper => 
                                        <tr key = {auditWorkpaper.auditWorkpaperId}>
                                             <td> { auditWorkpaper.workpaperRef } </td>
                                             <td> { auditWorkpaper.subject } </td>
                                             <td> { auditWorkpaper.workpaperUrl } </td>
                                             <td>
                                                 <button onClick={ () => this.editAuditWorkpaper(auditWorkpaper.auditWorkpaperId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAuditWorkpaper(auditWorkpaper.auditWorkpaperId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAuditWorkpaper(auditWorkpaper.auditWorkpaperId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAuditWorkpaperComponent
