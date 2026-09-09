import React, { Component } from 'react'
import ThirdPartyService from '../services/ThirdPartyService';

class CreateThirdPartyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                taxId: '',
                address: '',
                partyType: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changetaxIdHandler = this.changetaxIdHandler.bind(this);
        this.changeaddressHandler = this.changeaddressHandler.bind(this);
        this.changePartyTypeHandler = this.changePartyTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ThirdPartyService.getThirdPartyById(this.state.id).then( (res) =>{
                let thirdParty = res.data;
                this.setState({
                    name: thirdParty.name,
                    taxId: thirdParty.taxId,
                    address: thirdParty.address,
                    partyType: thirdParty.partyType
                });
            });
        }        
    }
    saveOrUpdateThirdParty = (e) => {
        e.preventDefault();
        let thirdParty = {
                thirdPartyId: this.state.id,
                name: this.state.name,
                taxId: this.state.taxId,
                address: this.state.address,
                partyType: this.state.partyType
            };
        console.log('thirdParty => ' + JSON.stringify(thirdParty));

        // step 5
        if(this.state.id === '_add'){
            thirdParty.thirdPartyId=''
            ThirdPartyService.createThirdParty(thirdParty).then(res =>{
                this.props.history.push('/thirdPartys');
            });
        }else{
            ThirdPartyService.updateThirdParty(thirdParty).then( res => {
                this.props.history.push('/thirdPartys');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changetaxIdHandler= (event) => {
        this.setState({taxId: event.target.value});
    }
    changeaddressHandler= (event) => {
        this.setState({address: event.target.value});
    }
    changePartyTypeHandler= (event) => {
        this.setState({partyType: event.target.value});
    }

    cancel(){
        this.props.history.push('/thirdPartys');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add ThirdParty</h3>
        }else{
            return <h3 className="text-center">Update ThirdParty</h3>
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

                                            <label> taxId:&emsp; </label>
                                                <input placeholder="taxId" name="taxId" className="form-control" value={this.state.taxId} onChange={this.changetaxIdHandler}/>

                                            <label> address:&emsp; </label>
                                                <input placeholder="address" name="address" className="form-control" value={this.state.address} onChange={this.changeaddressHandler}/>

                                            <label> PartyType:&emsp; </label>
                                                <select value={this.state.partyType} onChange={this.changePartyTypeHandler}>
                      <option name="PartyType" className="form-control" >
                          Individual
                      </option>
                      <option name="PartyType" className="form-control" >
                          Company
                      </option>
                      <option name="PartyType" className="form-control" >
                          GovernmentAgency
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateThirdParty}>Save</button>
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

export default CreateThirdPartyComponent
