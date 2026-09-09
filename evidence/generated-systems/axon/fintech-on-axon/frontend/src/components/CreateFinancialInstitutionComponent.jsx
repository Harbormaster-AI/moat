import React, { Component } from 'react'
import FinancialInstitutionService from '../services/FinancialInstitutionService';

class CreateFinancialInstitutionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                legalName: '',
                countryOfIncorporation: '',
                bic: '',
                website: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changelegalNameHandler = this.changelegalNameHandler.bind(this);
        this.changecountryOfIncorporationHandler = this.changecountryOfIncorporationHandler.bind(this);
        this.changebicHandler = this.changebicHandler.bind(this);
        this.changewebsiteHandler = this.changewebsiteHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            FinancialInstitutionService.getFinancialInstitutionById(this.state.id).then( (res) =>{
                let financialInstitution = res.data;
                this.setState({
                    name: financialInstitution.name,
                    legalName: financialInstitution.legalName,
                    countryOfIncorporation: financialInstitution.countryOfIncorporation,
                    bic: financialInstitution.bic,
                    website: financialInstitution.website
                });
            });
        }        
    }
    saveOrUpdateFinancialInstitution = (e) => {
        e.preventDefault();
        let financialInstitution = {
                financialInstitutionId: this.state.id,
                name: this.state.name,
                legalName: this.state.legalName,
                countryOfIncorporation: this.state.countryOfIncorporation,
                bic: this.state.bic,
                website: this.state.website
            };
        console.log('financialInstitution => ' + JSON.stringify(financialInstitution));

        // step 5
        if(this.state.id === '_add'){
            financialInstitution.financialInstitutionId=''
            FinancialInstitutionService.createFinancialInstitution(financialInstitution).then(res =>{
                this.props.history.push('/financialInstitutions');
            });
        }else{
            FinancialInstitutionService.updateFinancialInstitution(financialInstitution).then( res => {
                this.props.history.push('/financialInstitutions');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changelegalNameHandler= (event) => {
        this.setState({legalName: event.target.value});
    }
    changecountryOfIncorporationHandler= (event) => {
        this.setState({countryOfIncorporation: event.target.value});
    }
    changebicHandler= (event) => {
        this.setState({bic: event.target.value});
    }
    changewebsiteHandler= (event) => {
        this.setState({website: event.target.value});
    }

    cancel(){
        this.props.history.push('/financialInstitutions');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add FinancialInstitution</h3>
        }else{
            return <h3 className="text-center">Update FinancialInstitution</h3>
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

                                            <label> countryOfIncorporation:&emsp; </label>
                                                <input placeholder="countryOfIncorporation" name="countryOfIncorporation" className="form-control" value={this.state.countryOfIncorporation} onChange={this.changecountryOfIncorporationHandler}/>

                                            <label> bic:&emsp; </label>
                                                <input placeholder="bic" name="bic" className="form-control" value={this.state.bic} onChange={this.changebicHandler}/>

                                            <label> website:&emsp; </label>
                                                <input placeholder="website" name="website" className="form-control" value={this.state.website} onChange={this.changewebsiteHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateFinancialInstitution}>Save</button>
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

export default CreateFinancialInstitutionComponent
