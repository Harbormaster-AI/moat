import React, { Component } from 'react'
import EnterpriseService from '../services/EnterpriseService';

class CreateEnterpriseComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                legalName: '',
                registrationCountry: '',
                website: '',
                taxId: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changelegalNameHandler = this.changelegalNameHandler.bind(this);
        this.changeregistrationCountryHandler = this.changeregistrationCountryHandler.bind(this);
        this.changewebsiteHandler = this.changewebsiteHandler.bind(this);
        this.changetaxIdHandler = this.changetaxIdHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateEnterprise = (e) => {
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

        // step 5
        if(this.state.id === '_add'){
            enterprise.enterpriseId=''
            EnterpriseService.createEnterprise(enterprise).then(res =>{
                this.props.history.push('/enterprises');
            });
        }else{
            EnterpriseService.updateEnterprise(enterprise).then( res => {
                this.props.history.push('/enterprises');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Enterprise</h3>
        }else{
            return <h3 className="text-center">Update Enterprise</h3>
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

                                            <label> registrationCountry:&emsp; </label>
                                                <input placeholder="registrationCountry" name="registrationCountry" className="form-control" value={this.state.registrationCountry} onChange={this.changeregistrationCountryHandler}/>

                                            <label> website:&emsp; </label>
                                                <input placeholder="website" name="website" className="form-control" value={this.state.website} onChange={this.changewebsiteHandler}/>

                                            <label> taxId:&emsp; </label>
                                                <input placeholder="taxId" name="taxId" className="form-control" value={this.state.taxId} onChange={this.changetaxIdHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateEnterprise}>Save</button>
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

export default CreateEnterpriseComponent
