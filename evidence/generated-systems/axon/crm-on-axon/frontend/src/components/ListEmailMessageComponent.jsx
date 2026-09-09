import React, { Component } from 'react'
import EmailMessageService from '../services/EmailMessageService'

class ListEmailMessageComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                emailMessages: []
        }
        this.addEmailMessage = this.addEmailMessage.bind(this);
        this.editEmailMessage = this.editEmailMessage.bind(this);
        this.deleteEmailMessage = this.deleteEmailMessage.bind(this);
    }

    deleteEmailMessage(id){
        EmailMessageService.deleteEmailMessage(id).then( res => {
            this.setState({emailMessages: this.state.emailMessages.filter(emailMessage => emailMessage.emailMessageId !== id)});
        });
    }
    viewEmailMessage(id){
        this.props.history.push(`/view-emailMessage/${id}`);
    }
    editEmailMessage(id){
        this.props.history.push(`/add-emailMessage/${id}`);
    }

    componentDidMount(){
        EmailMessageService.getEmailMessages().then((res) => {
            this.setState({ emailMessages: res.data});
        });
    }

    addEmailMessage(){
        this.props.history.push('/add-emailMessage/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">EmailMessage List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addEmailMessage}> Add EmailMessage</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Subject </th>
                                    <th> Body </th>
                                    <th> SentAt </th>
                                    <th> MessageId </th>
                                    <th> Direction </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.emailMessages.map(
                                        emailMessage => 
                                        <tr key = {emailMessage.emailMessageId}>
                                             <td> { emailMessage.subject } </td>
                                             <td> { emailMessage.body } </td>
                                             <td> { emailMessage.sentAt } </td>
                                             <td> { emailMessage.messageId } </td>
                                             <td> { emailMessage.direction } </td>
                                             <td> { emailMessage.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editEmailMessage(emailMessage.emailMessageId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteEmailMessage(emailMessage.emailMessageId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewEmailMessage(emailMessage.emailMessageId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListEmailMessageComponent
