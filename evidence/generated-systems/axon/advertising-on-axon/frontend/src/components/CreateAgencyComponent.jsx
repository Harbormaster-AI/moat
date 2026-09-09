import React, { Component } from 'react'
import AgencyService from '../services/AgencyService';

class CreateAgencyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                legalName: '',
                headquartersCountry: '',
                website: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changelegalNameHandler = this.changelegalNameHandler.bind(this);
        this.changeheadquartersCountryHandler = this.changeheadquartersCountryHandler.bind(this);
        this.changewebsiteHandler = this.changewebsiteHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            AgencyService.getAgencyById(this.state.id).then( (res) =>{
                let agency = res.data;
                this.setState({
                    name: agency.name,
                    legalName: agency.legalName,
                    headquartersCountry: agency.headquartersCountry,
                    website: agency.website
                });
            });
        }        
    }
    saveOrUpdateAgency = (e) => {
        e.preventDefault();
        let agency = {
                agencyId: this.state.id,
                name: this.state.name,
                legalName: this.state.legalName,
                headquartersCountry: this.state.headquartersCountry,
                website: this.state.website
            };
        console.log('agency => ' + JSON.stringify(agency));

        // step 5
        if(this.state.id === '_add'){
            agency.agencyId=''
            AgencyService.createAgency(agency).then(res =>{
                this.props.history.push('/agencys');
            });
        }else{
            AgencyService.updateAgency(agency).then( res => {
                this.props.history.push('/agencys');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changelegalNameHandler= (event) => {
        this.setState({legalName: event.target.value});
    }
    changeheadquartersCountryHandler= (event) => {
        this.setState({headquartersCountry: event.target.value});
    }
    changewebsiteHandler= (event) => {
        this.setState({website: event.target.value});
    }

    cancel(){
        this.props.history.push('/agencys');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Agency</h3>
        }else{
            return <h3 className="text-center">Update Agency</h3>
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

                                            <label> headquartersCountry:&emsp; </label>
                                                <input placeholder="headquartersCountry" name="headquartersCountry" className="form-control" value={this.state.headquartersCountry} onChange={this.changeheadquartersCountryHandler}/>

                                            <label> website:&emsp; </label>
                                                <input placeholder="website" name="website" className="form-control" value={this.state.website} onChange={this.changewebsiteHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateAgency}>Save</button>
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

export default CreateAgencyComponent
