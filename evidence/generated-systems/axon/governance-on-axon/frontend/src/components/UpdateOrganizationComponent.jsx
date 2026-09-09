import React, { Component } from 'react'
import OrganizationService from '../services/OrganizationService';

class UpdateOrganizationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                legalName: '',
                jurisdiction: '',
                industrySector: ''
        }
        this.updateOrganization = this.updateOrganization.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changelegalNameHandler = this.changelegalNameHandler.bind(this);
        this.changejurisdictionHandler = this.changejurisdictionHandler.bind(this);
        this.changeindustrySectorHandler = this.changeindustrySectorHandler.bind(this);
    }

    componentDidMount(){
        OrganizationService.getOrganizationById(this.state.id).then( (res) =>{
            let organization = res.data;
            this.setState({
                name: organization.name,
                legalName: organization.legalName,
                jurisdiction: organization.jurisdiction,
                industrySector: organization.industrySector
            });
        });
    }

    updateOrganization = (e) => {
        e.preventDefault();
        let organization = {
            organizationId: this.state.id,
            name: this.state.name,
            legalName: this.state.legalName,
            jurisdiction: this.state.jurisdiction,
            industrySector: this.state.industrySector
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
    changejurisdictionHandler= (event) => {
        this.setState({jurisdiction: event.target.value});
    }
    changeindustrySectorHandler= (event) => {
        this.setState({industrySector: event.target.value});
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

                                            <label> jurisdiction: </label>
                                                <input placeholder="jurisdiction" name="jurisdiction" className="form-control" value={this.state.jurisdiction} onChange={this.changejurisdictionHandler}/>

                                            <label> industrySector: </label>
                                                <input placeholder="industrySector" name="industrySector" className="form-control" value={this.state.industrySector} onChange={this.changeindustrySectorHandler}/>

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
