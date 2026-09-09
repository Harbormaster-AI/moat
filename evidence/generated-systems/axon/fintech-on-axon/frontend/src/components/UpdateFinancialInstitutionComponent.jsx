import React, { Component } from 'react'
import FinancialInstitutionService from '../services/FinancialInstitutionService';

class UpdateFinancialInstitutionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                legalName: '',
                countryOfIncorporation: '',
                bic: '',
                website: ''
        }
        this.updateFinancialInstitution = this.updateFinancialInstitution.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changelegalNameHandler = this.changelegalNameHandler.bind(this);
        this.changecountryOfIncorporationHandler = this.changecountryOfIncorporationHandler.bind(this);
        this.changebicHandler = this.changebicHandler.bind(this);
        this.changewebsiteHandler = this.changewebsiteHandler.bind(this);
    }

    componentDidMount(){
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

    updateFinancialInstitution = (e) => {
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
        console.log('id => ' + JSON.stringify(this.state.id));
        FinancialInstitutionService.updateFinancialInstitution(financialInstitution).then( res => {
            this.props.history.push('/financialInstitutions');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update FinancialInstitution</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> legalName: </label>
                                                <input placeholder="legalName" name="legalName" className="form-control" value={this.state.legalName} onChange={this.changelegalNameHandler}/>

                                            <label> countryOfIncorporation: </label>
                                                <input placeholder="countryOfIncorporation" name="countryOfIncorporation" className="form-control" value={this.state.countryOfIncorporation} onChange={this.changecountryOfIncorporationHandler}/>

                                            <label> bic: </label>
                                                <input placeholder="bic" name="bic" className="form-control" value={this.state.bic} onChange={this.changebicHandler}/>

                                            <label> website: </label>
                                                <input placeholder="website" name="website" className="form-control" value={this.state.website} onChange={this.changewebsiteHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateFinancialInstitution}>Save</button>
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

export default UpdateFinancialInstitutionComponent
