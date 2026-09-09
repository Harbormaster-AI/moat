import React, { Component } from 'react'
import BeneficiaryService from '../services/BeneficiaryService';

class UpdateBeneficiaryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                share: '',
                relationship: ''
        }
        this.updateBeneficiary = this.updateBeneficiary.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeshareHandler = this.changeshareHandler.bind(this);
        this.changeRelationshipHandler = this.changeRelationshipHandler.bind(this);
    }

    componentDidMount(){
        BeneficiaryService.getBeneficiaryById(this.state.id).then( (res) =>{
            let beneficiary = res.data;
            this.setState({
                name: beneficiary.name,
                share: beneficiary.share,
                relationship: beneficiary.relationship
            });
        });
    }

    updateBeneficiary = (e) => {
        e.preventDefault();
        let beneficiary = {
            beneficiaryId: this.state.id,
            name: this.state.name,
            share: this.state.share,
            relationship: this.state.relationship
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
    changeshareHandler= (event) => {
        this.setState({share: event.target.value});
    }
    changeRelationshipHandler= (event) => {
        this.setState({relationship: event.target.value});
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

                                            <label> share: </label>
                                                <input placeholder="share" name="share" className="form-control" value={this.state.share} onChange={this.changeshareHandler}/>

                                            <label> Relationship: </label>
                                                <select value={this.state.relationship} onChange={this.changeRelationshipHandler}>
                      <option name="Relationship" className="form-control" >
                          Spouse
                      </option>
                      <option name="Relationship" className="form-control" >
                          Child
                      </option>
                      <option name="Relationship" className="form-control" >
                          Parent
                      </option>
                      <option name="Relationship" className="form-control" >
                          Sibling
                      </option>
                      <option name="Relationship" className="form-control" >
                          BusinessPartner
                      </option>
                      <option name="Relationship" className="form-control" >
                          Estate
                      </option>
                      <option name="Relationship" className="form-control" >
                          Trust
                      </option>
                      <option name="Relationship" className="form-control" >
                          Other
                      </option>
                    </select>

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
