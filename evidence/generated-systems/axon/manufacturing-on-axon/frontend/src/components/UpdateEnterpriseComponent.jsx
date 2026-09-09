import React, { Component } from 'react'
import EnterpriseService from '../services/EnterpriseService';

class UpdateEnterpriseComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                legalName: '',
                registrationCountry: '',
                website: '',
                taxId: ''
        }
        this.updateEnterprise = this.updateEnterprise.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changelegalNameHandler = this.changelegalNameHandler.bind(this);
        this.changeregistrationCountryHandler = this.changeregistrationCountryHandler.bind(this);
        this.changewebsiteHandler = this.changewebsiteHandler.bind(this);
        this.changetaxIdHandler = this.changetaxIdHandler.bind(this);
    }

    componentDidMount(){
        EnterpriseService.getEnterpriseById(this.state.id).then( (res) =>{
            let enterprise = res.data;
            this.setState({
                name: enterprise.name,
                legalName: enterprise.legalName,
                registrationCountry: enterprise.registrationCountry,
                website: enterprise.website,
                taxId: enterprise.taxId
            });
        });
    }

    updateEnterprise = (e) => {
        e.preventDefault();
        let enterprise = {
            enterpriseId: this.state.id,
            name: this.state.name,
            legalName: this.state.legalName,
            registrationCountry: this.state.registrationCountry,
            website: this.state.website,
            taxId: this.state.taxId
        };
        console.log('enterprise => ' + JSON.stringify(enterprise));
        console.log('id => ' + JSON.stringify(this.state.id));
        EnterpriseService.updateEnterprise(enterprise).then( res => {
            this.props.history.push('/enterprises');
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
    changetaxIdHandler= (event) => {
        this.setState({taxId: event.target.value});
    }

    cancel(){
        this.props.history.push('/enterprises');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Enterprise</h3>
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

                                            <label> taxId: </label>
                                                <input placeholder="taxId" name="taxId" className="form-control" value={this.state.taxId} onChange={this.changetaxIdHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateEnterprise}>Save</button>
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

export default UpdateEnterpriseComponent
