import React, { Component } from 'react'
import OrganizationService from '../services/OrganizationService';

class UpdateOrganizationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                legalName: '',
                registrationCountry: '',
                website: ''
        }
        this.updateOrganization = this.updateOrganization.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changelegalNameHandler = this.changelegalNameHandler.bind(this);
        this.changeregistrationCountryHandler = this.changeregistrationCountryHandler.bind(this);
        this.changewebsiteHandler = this.changewebsiteHandler.bind(this);
    }

    componentDidMount(){
        OrganizationService.getOrganizationById(this.state.id).then( (res) =>{
            let organization = res.data;
            this.setState({
                name: organization.name,
                legalName: organization.legalName,
                registrationCountry: organization.registrationCountry,
                website: organization.website
            });
        });
    }

    updateOrganization = (e) => {
        e.preventDefault();
        let organization = {
            organizationId: this.state.id,
            name: this.state.name,
            legalName: this.state.legalName,
            registrationCountry: this.state.registrationCountry,
            website: this.state.website
        };
        console.log('organization => ' + JSON.stringify(organization));
        console.log('id => ' + JSON.stringify(this.state.id));
        OrganizationService.updateOrganization(organization).then( res => {
            this.props.history.push('/organizations');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changelegalNameHandler= (event) => {
        this.setState({legalName: event.target.value});
    }
    changeregistrationCountryHandler= (event) => {
        this.setState({registrationCountry: event.target.value});
    }
    changewebsiteHandler= (event) => {
        this.setState({website: event.target.value});
    }

    cancel(){
        this.props.history.push('/organizations');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Organization</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> legalName: </label>
                                                <input placeholder="legalName" name="legalName" className="form-control" value={this.state.legalName} onChange={this.changelegalNameHandler}/>

                                            <label> registrationCountry: </label>
                                                <input placeholder="registrationCountry" name="registrationCountry" className="form-control" value={this.state.registrationCountry} onChange={this.changeregistrationCountryHandler}/>

                                            <label> website: </label>
                                                <input placeholder="website" name="website" className="form-control" value={this.state.website} onChange={this.changewebsiteHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateOrganization}>Save</button>
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

export default UpdateOrganizationComponent
