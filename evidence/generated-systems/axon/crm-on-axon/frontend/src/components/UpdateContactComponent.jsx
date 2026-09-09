import React, { Component } from 'react'
import ContactService from '../services/ContactService';

class UpdateContactComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                firstName: '',
                lastName: '',
                title: '',
                email: '',
                phone: '',
                mobile: '',
                mailingAddress: '',
                preferredContactMethod: ''
        }
        this.updateContact = this.updateContact.bind(this);

        this.changefirstNameHandler = this.changefirstNameHandler.bind(this);
        this.changelastNameHandler = this.changelastNameHandler.bind(this);
        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changeemailHandler = this.changeemailHandler.bind(this);
        this.changephoneHandler = this.changephoneHandler.bind(this);
        this.changemobileHandler = this.changemobileHandler.bind(this);
        this.changemailingAddressHandler = this.changemailingAddressHandler.bind(this);
        this.changePreferredContactMethodHandler = this.changePreferredContactMethodHandler.bind(this);
    }

    componentDidMount(){
        ContactService.getContactById(this.state.id).then( (res) =>{
            let contact = res.data;
            this.setState({
                firstName: contact.firstName,
                lastName: contact.lastName,
                title: contact.title,
                email: contact.email,
                phone: contact.phone,
                mobile: contact.mobile,
                mailingAddress: contact.mailingAddress,
                preferredContactMethod: contact.preferredContactMethod
            });
        });
    }

    updateContact = (e) => {
        e.preventDefault();
        let contact = {
            contactId: this.state.id,
            firstName: this.state.firstName,
            lastName: this.state.lastName,
            title: this.state.title,
            email: this.state.email,
            phone: this.state.phone,
            mobile: this.state.mobile,
            mailingAddress: this.state.mailingAddress,
            preferredContactMethod: this.state.preferredContactMethod
        };
        console.log('contact => ' + JSON.stringify(contact));
        console.log('id => ' + JSON.stringify(this.state.id));
        ContactService.updateContact(contact).then( res => {
            this.props.history.push('/contacts');
        });
    }

    changefirstNameHandler= (event) => {
        this.setState({firstName: event.target.value});
    }
    changelastNameHandler= (event) => {
        this.setState({lastName: event.target.value});
    }
    changetitleHandler= (event) => {
        this.setState({title: event.target.value});
    }
    changeemailHandler= (event) => {
        this.setState({email: event.target.value});
    }
    changephoneHandler= (event) => {
        this.setState({phone: event.target.value});
    }
    changemobileHandler= (event) => {
        this.setState({mobile: event.target.value});
    }
    changemailingAddressHandler= (event) => {
        this.setState({mailingAddress: event.target.value});
    }
    changePreferredContactMethodHandler= (event) => {
        this.setState({preferredContactMethod: event.target.value});
    }

    cancel(){
        this.props.history.push('/contacts');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Contact</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> firstName: </label>
                                                <input placeholder="firstName" name="firstName" className="form-control" value={this.state.firstName} onChange={this.changefirstNameHandler}/>

                                            <label> lastName: </label>
                                                <input placeholder="lastName" name="lastName" className="form-control" value={this.state.lastName} onChange={this.changelastNameHandler}/>

                                            <label> title: </label>
                                                <input placeholder="title" name="title" className="form-control" value={this.state.title} onChange={this.changetitleHandler}/>

                                            <label> email: </label>
                                                <input placeholder="email" name="email" className="form-control" value={this.state.email} onChange={this.changeemailHandler}/>

                                            <label> phone: </label>
                                                <input placeholder="phone" name="phone" className="form-control" value={this.state.phone} onChange={this.changephoneHandler}/>

                                            <label> mobile: </label>
                                                <input placeholder="mobile" name="mobile" className="form-control" value={this.state.mobile} onChange={this.changemobileHandler}/>

                                            <label> mailingAddress: </label>
                                                <input placeholder="mailingAddress" name="mailingAddress" className="form-control" value={this.state.mailingAddress} onChange={this.changemailingAddressHandler}/>

                                            <label> PreferredContactMethod: </label>
                                                <select value={this.state.preferredContactMethod} onChange={this.changePreferredContactMethodHandler}>
                      <option name="PreferredContactMethod" className="form-control" >
                          Email
                      </option>
                      <option name="PreferredContactMethod" className="form-control" >
                          Phone
                      </option>
                      <option name="PreferredContactMethod" className="form-control" >
                          Mobile
                      </option>
                      <option name="PreferredContactMethod" className="form-control" >
                          SMS
                      </option>
                      <option name="PreferredContactMethod" className="form-control" >
                          InPerson
                      </option>
                      <option name="PreferredContactMethod" className="form-control" >
                          Web
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateContact}>Save</button>
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

export default UpdateContactComponent
