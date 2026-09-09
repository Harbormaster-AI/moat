import React, { Component } from 'react'
import ContactService from '../services/ContactService'

class ViewContactComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            contact: {}
        }
    }

    componentDidMount(){
        ContactService.getContactById(this.state.id).then( res => {
            this.setState({contact: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Contact Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> firstName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.contact.firstName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> lastName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.contact.lastName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> title:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.contact.title }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> email:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.contact.email }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> phone:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.contact.phone }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> mobile:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.contact.mobile }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> mailingAddress:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.contact.mailingAddress }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> PreferredContactMethod:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.contact.preferredContactMethod }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewContactComponent
