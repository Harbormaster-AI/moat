import React, { Component } from 'react'
import ContactService from '../services/ContactService'

class ListContactComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                contacts: []
        }
        this.addContact = this.addContact.bind(this);
        this.editContact = this.editContact.bind(this);
        this.deleteContact = this.deleteContact.bind(this);
    }

    deleteContact(id){
        ContactService.deleteContact(id).then( res => {
            this.setState({contacts: this.state.contacts.filter(contact => contact.contactId !== id)});
        });
    }
    viewContact(id){
        this.props.history.push(`/view-contact/${id}`);
    }
    editContact(id){
        this.props.history.push(`/add-contact/${id}`);
    }

    componentDidMount(){
        ContactService.getContacts().then((res) => {
            this.setState({ contacts: res.data});
        });
    }

    addContact(){
        this.props.history.push('/add-contact/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Contact List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addContact}> Add Contact</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> FirstName </th>
                                    <th> LastName </th>
                                    <th> Title </th>
                                    <th> Email </th>
                                    <th> Phone </th>
                                    <th> Mobile </th>
                                    <th> MailingAddress </th>
                                    <th> PreferredContactMethod </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.contacts.map(
                                        contact => 
                                        <tr key = {contact.contactId}>
                                             <td> { contact.firstName } </td>
                                             <td> { contact.lastName } </td>
                                             <td> { contact.title } </td>
                                             <td> { contact.email } </td>
                                             <td> { contact.phone } </td>
                                             <td> { contact.mobile } </td>
                                             <td> { contact.mailingAddress } </td>
                                             <td> { contact.preferredContactMethod } </td>
                                             <td>
                                                 <button onClick={ () => this.editContact(contact.contactId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteContact(contact.contactId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewContact(contact.contactId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListContactComponent
