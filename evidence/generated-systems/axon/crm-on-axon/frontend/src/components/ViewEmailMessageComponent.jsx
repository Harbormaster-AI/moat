import React, { Component } from 'react'
import EmailMessageService from '../services/EmailMessageService'

class ViewEmailMessageComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            emailMessage: {}
        }
    }

    componentDidMount(){
        EmailMessageService.getEmailMessageById(this.state.id).then( res => {
            this.setState({emailMessage: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View EmailMessage Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> subject:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.emailMessage.subject }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> body:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.emailMessage.body }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> sentAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.emailMessage.sentAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> messageId:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.emailMessage.messageId }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Direction:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.emailMessage.direction }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.emailMessage.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewEmailMessageComponent
