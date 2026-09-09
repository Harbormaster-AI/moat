import React, { Component } from 'react'
import BeneficiaryService from '../services/BeneficiaryService';

class CreateBeneficiaryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                share: '',
                relationship: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeshareHandler = this.changeshareHandler.bind(this);
        this.changeRelationshipHandler = this.changeRelationshipHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            BeneficiaryService.getBeneficiaryById(this.state.id).then( (res) =>{
                let beneficiary = res.data;
                this.setState({
                    name: beneficiary.name,
                    share: beneficiary.share,
                    relationship: beneficiary.relationship
                });
            });
        }        
    }
    saveOrUpdateBeneficiary = (e) => {
        e.preventDefault();
        let beneficiary = {
                beneficiaryId: this.state.id,
                name: this.state.name,
                share: this.state.share,
                relationship: this.state.relationship
            };
        console.log('beneficiary => ' + JSON.stringify(beneficiary));

        // step 5
        if(this.state.id === '_add'){
            beneficiary.beneficiaryId=''
            BeneficiaryService.createBeneficiary(beneficiary).then(res =>{
                this.props.history.push('/beneficiarys');
            });
        }else{
            BeneficiaryService.updateBeneficiary(beneficiary).then( res => {
                this.props.history.push('/beneficiarys');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Beneficiary</h3>
        }else{
            return <h3 className="text-center">Update Beneficiary</h3>
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

                                            <label> share:&emsp; </label>
                                                <input placeholder="share" name="share" className="form-control" value={this.state.share} onChange={this.changeshareHandler}/>

                                            <label> Relationship:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateBeneficiary}>Save</button>
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

export default CreateBeneficiaryComponent
