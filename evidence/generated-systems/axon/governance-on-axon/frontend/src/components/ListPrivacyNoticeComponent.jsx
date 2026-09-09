import React, { Component } from 'react'
import PrivacyNoticeService from '../services/PrivacyNoticeService'

class ListPrivacyNoticeComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                privacyNotices: []
        }
        this.addPrivacyNotice = this.addPrivacyNotice.bind(this);
        this.editPrivacyNotice = this.editPrivacyNotice.bind(this);
        this.deletePrivacyNotice = this.deletePrivacyNotice.bind(this);
    }

    deletePrivacyNotice(id){
        PrivacyNoticeService.deletePrivacyNotice(id).then( res => {
            this.setState({privacyNotices: this.state.privacyNotices.filter(privacyNotice => privacyNotice.privacyNoticeId !== id)});
        });
    }
    viewPrivacyNotice(id){
        this.props.history.push(`/view-privacyNotice/${id}`);
    }
    editPrivacyNotice(id){
        this.props.history.push(`/add-privacyNotice/${id}`);
    }

    componentDidMount(){
        PrivacyNoticeService.getPrivacyNotices().then((res) => {
            this.setState({ privacyNotices: res.data});
        });
    }

    addPrivacyNotice(){
        this.props.history.push('/add-privacyNotice/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">PrivacyNotice List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addPrivacyNotice}> Add PrivacyNotice</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Title </th>
                                    <th> Audience </th>
                                    <th> VersionLabel </th>
                                    <th> PublicationDate </th>
                                    <th> PublicationUrl </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.privacyNotices.map(
                                        privacyNotice => 
                                        <tr key = {privacyNotice.privacyNoticeId}>
                                             <td> { privacyNotice.title } </td>
                                             <td> { privacyNotice.audience } </td>
                                             <td> { privacyNotice.versionLabel } </td>
                                             <td> { privacyNotice.publicationDate } </td>
                                             <td> { privacyNotice.publicationUrl } </td>
                                             <td> { privacyNotice.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editPrivacyNotice(privacyNotice.privacyNoticeId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deletePrivacyNotice(privacyNotice.privacyNoticeId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewPrivacyNotice(privacyNotice.privacyNoticeId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListPrivacyNoticeComponent
