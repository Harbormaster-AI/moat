import React, { Component } from 'react'
import OrganizationService from '../services/OrganizationService';

class CreateOrganizationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                defaultCurrency: '',
                defaultLocale: '',
                website: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changedefaultCurrencyHandler = this.changedefaultCurrencyHandler.bind(this);
        this.changedefaultLocaleHandler = this.changedefaultLocaleHandler.bind(this);
        this.changewebsiteHandler = this.changewebsiteHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            OrganizationService.getOrganizationById(this.state.id).then( (res) =>{
                let organization = res.data;
                this.setState({
                    name: organization.name,
                    defaultCurrency: organization.defaultCurrency,
                    defaultLocale: organization.defaultLocale,
                    website: organization.website
                });
            });
        }        
    }
    saveOrUpdateOrganization = (e) => {
        e.preventDefault();
        let organization = {
                organizationId: this.state.id,
                name: this.state.name,
                defaultCurrency: this.state.defaultCurrency,
                defaultLocale: this.state.defaultLocale,
                website: this.state.website
            };
        console.log('organization => ' + JSON.stringify(organization));

        // step 5
        if(this.state.id === '_add'){
            organization.organizationId=''
            OrganizationService.createOrganization(organization).then(res =>{
                this.props.history.push('/organizations');
            });
        }else{
            OrganizationService.updateOrganization(organization).then( res => {
                this.props.history.push('/organizations');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changedefaultCurrencyHandler= (event) => {
        this.setState({defaultCurrency: event.target.value});
    }
    changedefaultLocaleHandler= (event) => {
        this.setState({defaultLocale: event.target.value});
    }
    changewebsiteHandler= (event) => {
        this.setState({website: event.target.value});
    }

    cancel(){
        this.props.history.push('/organizations');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Organization</h3>
        }else{
            return <h3 className="text-center">Update Organization</h3>
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
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> defaultCurrency:&emsp; </label>
                                                <input placeholder="defaultCurrency" name="defaultCurrency" className="form-control" value={this.state.defaultCurrency} onChange={this.changedefaultCurrencyHandler}/>

                                            <label> defaultLocale:&emsp; </label>
                                                <input placeholder="defaultLocale" name="defaultLocale" className="form-control" value={this.state.defaultLocale} onChange={this.changedefaultLocaleHandler}/>

                                            <label> website:&emsp; </label>
                                                <input placeholder="website" name="website" className="form-control" value={this.state.website} onChange={this.changewebsiteHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateOrganization}>Save</button>
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

export default CreateOrganizationComponent
