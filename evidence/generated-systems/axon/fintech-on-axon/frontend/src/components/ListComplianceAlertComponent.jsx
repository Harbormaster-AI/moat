import React, { Component } from 'react'
import ComplianceAlertService from '../services/ComplianceAlertService'

class ListComplianceAlertComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                complianceAlerts: []
        }
        this.addComplianceAlert = this.addComplianceAlert.bind(this);
        this.editComplianceAlert = this.editComplianceAlert.bind(this);
        this.deleteComplianceAlert = this.deleteComplianceAlert.bind(this);
    }

    deleteComplianceAlert(id){
        ComplianceAlertService.deleteComplianceAlert(id).then( res => {
            this.setState({complianceAlerts: this.state.complianceAlerts.filter(complianceAlert => complianceAlert.complianceAlertId !== id)});
        });
    }
    viewComplianceAlert(id){
        this.props.history.push(`/view-complianceAlert/${id}`);
    }
    editComplianceAlert(id){
        this.props.history.push(`/add-complianceAlert/${id}`);
    }

    componentDidMount(){
        ComplianceAlertService.getComplianceAlerts().then((res) => {
            this.setState({ complianceAlerts: res.data});
        });
    }

    addComplianceAlert(){
        this.props.history.push('/add-complianceAlert/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ComplianceAlert List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addComplianceAlert}> Add ComplianceAlert</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> AlertCode </th>
                                    <th> RaisedAt </th>
                                    <th> Notes </th>
                                    <th> Severity </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.complianceAlerts.map(
                                        complianceAlert => 
                                        <tr key = {complianceAlert.complianceAlertId}>
                                             <td> { complianceAlert.alertCode } </td>
                                             <td> { complianceAlert.raisedAt } </td>
                                             <td> { complianceAlert.notes } </td>
                                             <td> { complianceAlert.severity } </td>
                                             <td> { complianceAlert.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editComplianceAlert(complianceAlert.complianceAlertId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteComplianceAlert(complianceAlert.complianceAlertId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewComplianceAlert(complianceAlert.complianceAlertId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListComplianceAlertComponent
