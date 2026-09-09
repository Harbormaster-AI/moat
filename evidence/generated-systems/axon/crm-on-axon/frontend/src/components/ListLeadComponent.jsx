import React, { Component } from 'react'
import LeadService from '../services/LeadService'

class ListLeadComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                leads: []
        }
        this.addLead = this.addLead.bind(this);
        this.editLead = this.editLead.bind(this);
        this.deleteLead = this.deleteLead.bind(this);
    }

    deleteLead(id){
        LeadService.deleteLead(id).then( res => {
            this.setState({leads: this.state.leads.filter(lead => lead.leadId !== id)});
        });
    }
    viewLead(id){
        this.props.history.push(`/view-lead/${id}`);
    }
    editLead(id){
        this.props.history.push(`/add-lead/${id}`);
    }

    componentDidMount(){
        LeadService.getLeads().then((res) => {
            this.setState({ leads: res.data});
        });
    }

    addLead(){
        this.props.history.push('/add-lead/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Lead List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addLead}> Add Lead</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> FirstName </th>
                                    <th> LastName </th>
                                    <th> Company </th>
                                    <th> Email </th>
                                    <th> Phone </th>
                                    <th> Converted </th>
                                    <th> Status </th>
                                    <th> Source </th>
                                    <th> Rating </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.leads.map(
                                        lead => 
                                        <tr key = {lead.leadId}>
                                             <td> { lead.firstName } </td>
                                             <td> { lead.lastName } </td>
                                             <td> { lead.company } </td>
                                             <td> { lead.email } </td>
                                             <td> { lead.phone } </td>
                                             <td> { lead.converted } </td>
                                             <td> { lead.status } </td>
                                             <td> { lead.source } </td>
                                             <td> { lead.rating } </td>
                                             <td>
                                                 <button onClick={ () => this.editLead(lead.leadId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteLead(lead.leadId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewLead(lead.leadId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListLeadComponent
