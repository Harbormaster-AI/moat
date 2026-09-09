import React, { Component } from 'react'
import BeneficiaryService from '../services/BeneficiaryService';

class UpdateBeneficiaryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                accountIdentifier: '',
                iban: '',
                bic: '',
                address: ''
        }
        this.updateBeneficiary = this.updateBeneficiary.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeaccountIdentifierHandler = this.changeaccountIdentifierHandler.bind(this);
        this.changeibanHandler = this.changeibanHandler.bind(this);
        this.changebicHandler = this.changebicHandler.bind(this);
        this.changeaddressHandler = this.changeaddressHandler.bind(this);
    }

    componentDidMount(){
        BeneficiaryService.getBeneficiaryById(this.state.id).then( (res) =>{
            let beneficiary = res.data;
            this.setState({
                name: beneficiary.name,
                accountIdentifier: beneficiary.accountIdentifier,
                iban: beneficiary.iban,
                bic: beneficiary.bic,
                address: beneficiary.address
            });
        });
    }

    updateBeneficiary = (e) => {
        e.preventDefault();
        let beneficiary = {
            beneficiaryId: this.state.id,
            name: this.state.name,
            accountIdentifier: this.state.accountIdentifier,
            iban: this.state.iban,
            bic: this.state.bic,
            address: this.state.address
        };
        console.log('beneficiary => ' + JSON.stringify(beneficiary));
        console.log('id => ' + JSON.stringify(this.state.id));
        BeneficiaryService.updateBeneficiary(beneficiary).then( res => {
            this.props.history.push('/beneficiarys');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeaccountIdentifierHandler= (event) => {
        this.setState({accountIdentifier: event.target.value});
    }
    changeibanHandler= (event) => {
        this.setState({iban: event.target.value});
    }
    changebicHandler= (event) => {
        this.setState({bic: event.target.value});
    }
    changeaddressHandler= (event) => {
        this.setState({address: event.target.value});
    }

    cancel(){
        this.props.history.push('/beneficiarys');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Beneficiary</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> accountIdentifier: </label>
                                                <input placeholder="accountIdentifier" name="accountIdentifier" className="form-control" value={this.state.accountIdentifier} onChange={this.changeaccountIdentifierHandler}/>

                                            <label> iban: </label>
                                                <input placeholder="iban" name="iban" className="form-control" value={this.state.iban} onChange={this.changeibanHandler}/>

                                            <label> bic: </label>
                                                <input placeholder="bic" name="bic" className="form-control" value={this.state.bic} onChange={this.changebicHandler}/>

                                            <label> address: </label>
                                                <input placeholder="address" name="address" className="form-control" value={this.state.address} onChange={this.changeaddressHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateBeneficiary}>Save</button>
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

export default UpdateBeneficiaryComponent
