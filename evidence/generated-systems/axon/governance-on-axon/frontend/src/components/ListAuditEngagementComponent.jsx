import React, { Component } from 'react'
import AuditEngagementService from '../services/AuditEngagementService'

class ListAuditEngagementComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                auditEngagements: []
        }
        this.addAuditEngagement = this.addAuditEngagement.bind(this);
        this.editAuditEngagement = this.editAuditEngagement.bind(this);
        this.deleteAuditEngagement = this.deleteAuditEngagement.bind(this);
    }

    deleteAuditEngagement(id){
        AuditEngagementService.deleteAuditEngagement(id).then( res => {
            this.setState({auditEngagements: this.state.auditEngagements.filter(auditEngagement => auditEngagement.auditEngagementId !== id)});
        });
    }
    viewAuditEngagement(id){
        this.props.history.push(`/view-auditEngagement/${id}`);
    }
    editAuditEngagement(id){
        this.props.history.push(`/add-auditEngagement/${id}`);
    }

    componentDidMount(){
        AuditEngagementService.getAuditEngagements().then((res) => {
            this.setState({ auditEngagements: res.data});
        });
    }

    addAuditEngagement(){
        this.props.history.push('/add-auditEngagement/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">AuditEngagement List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAuditEngagement}> Add AuditEngagement</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Title </th>
                                    <th> StartDate </th>
                                    <th> EndDate </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.auditEngagements.map(
                                        auditEngagement => 
                                        <tr key = {auditEngagement.auditEngagementId}>
                                             <td> { auditEngagement.title } </td>
                                             <td> { auditEngagement.startDate } </td>
                                             <td> { auditEngagement.endDate } </td>
                                             <td> { auditEngagement.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editAuditEngagement(auditEngagement.auditEngagementId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAuditEngagement(auditEngagement.auditEngagementId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAuditEngagement(auditEngagement.auditEngagementId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAuditEngagementComponent
