import React, { Component } from 'react'
import ContactService from '../services/ContactService';

class CreateContactComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
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
        this.changefirstNameHandler = this.changefirstNameHandler.bind(this);
        this.changelastNameHandler = this.changelastNameHandler.bind(this);
        this.changetitleHandler = this.changetitleHandler.bind(this);
        this.changeemailHandler = this.changeemailHandler.bind(this);
        this.changephoneHandler = this.changephoneHandler.bind(this);
        this.changemobileHandler = this.changemobileHandler.bind(this);
        this.changemailingAddressHandler = this.changemailingAddressHandler.bind(this);
        this.changePreferredContactMethodHandler = this.changePreferredContactMethodHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateContact = (e) => {
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

        // step 5
        if(this.state.id === '_add'){
            contact.contactId=''
            ContactService.createContact(contact).then(res =>{
                this.props.history.push('/contacts');
            });
        }else{
            ContactService.updateContact(contact).then( res => {
                this.props.history.push('/contacts');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Contact</h3>
        }else{
            return <h3 className="text-center">Update Contact</h3>
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
                                            <label> firstName:&emsp; </label>
                                                <input placeholder="firstName" name="firstName" className="form-control" value={this.state.firstName} onChange={this.changefirstNameHandler}/>

                                            <label> lastName:&emsp; </label>
                                                <input placeholder="lastName" name="lastName" className="form-control" value={this.state.lastName} onChange={this.changelastNameHandler}/>

                                            <label> title:&emsp; </label>
                                                <input placeholder="title" name="title" className="form-control" value={this.state.title} onChange={this.changetitleHandler}/>

                                            <label> email:&emsp; </label>
                                                <input placeholder="email" name="email" className="form-control" value={this.state.email} onChange={this.changeemailHandler}/>

                                            <label> phone:&emsp; </label>
                                                <input placeholder="phone" name="phone" className="form-control" value={this.state.phone} onChange={this.changephoneHandler}/>

                                            <label> mobile:&emsp; </label>
                                                <input placeholder="mobile" name="mobile" className="form-control" value={this.state.mobile} onChange={this.changemobileHandler}/>

                                            <label> mailingAddress:&emsp; </label>
                                                <input placeholder="mailingAddress" name="mailingAddress" className="form-control" value={this.state.mailingAddress} onChange={this.changemailingAddressHandler}/>

                                            <label> PreferredContactMethod:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateContact}>Save</button>
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

export default CreateContactComponent
