import React, { Component } from 'react'
import EmailMessageService from '../services/EmailMessageService';

class CreateEmailMessageComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                subject: '',
                body: '',
                sentAt: '',
                messageId: '',
                direction: '',
                status: ''
        }
        this.changesubjectHandler = this.changesubjectHandler.bind(this);
        this.changebodyHandler = this.changebodyHandler.bind(this);
        this.changesentAtHandler = this.changesentAtHandler.bind(this);
        this.changemessageIdHandler = this.changemessageIdHandler.bind(this);
        this.changeDirectionHandler = this.changeDirectionHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            EmailMessageService.getEmailMessageById(this.state.id).then( (res) =>{
                let emailMessage = res.data;
                this.setState({
                    subject: emailMessage.subject,
                    body: emailMessage.body,
                    sentAt: emailMessage.sentAt,
                    messageId: emailMessage.messageId,
                    direction: emailMessage.direction,
                    status: emailMessage.status
                });
            });
        }        
    }
    saveOrUpdateEmailMessage = (e) => {
        e.preventDefault();
        let emailMessage = {
                emailMessageId: this.state.id,
                subject: this.state.subject,
                body: this.state.body,
                sentAt: this.state.sentAt,
                messageId: this.state.messageId,
                direction: this.state.direction,
                status: this.state.status
            };
        console.log('emailMessage => ' + JSON.stringify(emailMessage));

        // step 5
        if(this.state.id === '_add'){
            emailMessage.emailMessageId=''
            EmailMessageService.createEmailMessage(emailMessage).then(res =>{
                this.props.history.push('/emailMessages');
            });
        }else{
            EmailMessageService.updateEmailMessage(emailMessage).then( res => {
                this.props.history.push('/emailMessages');
            });
        }
    }
    
    changesubjectHandler= (event) => {
        this.setState({subject: event.target.value});
    }
    changebodyHandler= (event) => {
        this.setState({body: event.target.value});
    }
    changesentAtHandler= (event) => {
        this.setState({sentAt: event.target.value});
    }
    changemessageIdHandler= (event) => {
        this.setState({messageId: event.target.value});
    }
    changeDirectionHandler= (event) => {
        this.setState({direction: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/emailMessages');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add EmailMessage</h3>
        }else{
            return <h3 className="text-center">Update EmailMessage</h3>
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
                                            <label> subject:&emsp; </label>
                                                <input placeholder="subject" name="subject" className="form-control" value={this.state.subject} onChange={this.changesubjectHandler}/>

                                            <label> body:&emsp; </label>
                                                <input placeholder="body" name="body" className="form-control" value={this.state.body} onChange={this.changebodyHandler}/>

                                            <label> sentAt:&emsp; </label>
                                                <input type="time" placeholder="sentAt" name="sentAt" className="form-control" value={this.state.sentAt} onChange={this.changesentAtHandler}/>

                                            <label> messageId:&emsp; </label>
                                                <input placeholder="messageId" name="messageId" className="form-control" value={this.state.messageId} onChange={this.changemessageIdHandler}/>

                                            <label> Direction:&emsp; </label>
                                                <select value={this.state.direction} onChange={this.changeDirectionHandler}>
                      <option name="Direction" className="form-control" >
                          Inbound
                      </option>
                      <option name="Direction" className="form-control" >
                          Outbound
                      </option>
                      <option name="Direction" className="form-control" >
                          Internal
                      </option>
                    </select>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Sent
                      </option>
                      <option name="Status" className="form-control" >
                          Delivered
                      </option>
                      <option name="Status" className="form-control" >
                          Opened
                      </option>
                      <option name="Status" className="form-control" >
                          Bounced
                      </option>
                      <option name="Status" className="form-control" >
                          Failed
                      </option>
                      <option name="Status" className="form-control" >
                          Replied
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateEmailMessage}>Save</button>
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

export default CreateEmailMessageComponent
