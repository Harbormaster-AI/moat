import React, { Component } from 'react'
import OrganizationService from '../services/OrganizationService';

class CreateOrganizationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                legalName: '',
                jurisdiction: '',
                industrySector: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changelegalNameHandler = this.changelegalNameHandler.bind(this);
        this.changejurisdictionHandler = this.changejurisdictionHandler.bind(this);
        this.changeindustrySectorHandler = this.changeindustrySectorHandler.bind(this);
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
                    legalName: organization.legalName,
                    jurisdiction: organization.jurisdiction,
                    industrySector: organization.industrySector
                });
            });
        }        
    }
    saveOrUpdateOrganization = (e) => {
        e.preventDefault();
        let organization = {
                organizationId: this.state.id,
                name: this.state.name,
                legalName: this.state.legalName,
                jurisdiction: this.state.jurisdiction,
                industrySector: this.state.industrySector
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

                                            <label> legalName:&emsp; </label>
                                                <input placeholder="legalName" name="legalName" className="form-control" value={this.state.legalName} onChange={this.changelegalNameHandler}/>

                                            <label> jurisdiction:&emsp; </label>
                                                <input placeholder="jurisdiction" name="jurisdiction" className="form-control" value={this.state.jurisdiction} onChange={this.changejurisdictionHandler}/>

                                            <label> industrySector:&emsp; </label>
                                                <input placeholder="industrySector" name="industrySector" className="form-control" value={this.state.industrySector} onChange={this.changeindustrySectorHandler}/>

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
