import React, { Component } from 'react'
import ThirdPartyService from '../services/ThirdPartyService';

class UpdateThirdPartyComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                taxId: '',
                address: '',
                partyType: ''
        }
        this.updateThirdParty = this.updateThirdParty.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changetaxIdHandler = this.changetaxIdHandler.bind(this);
        this.changeaddressHandler = this.changeaddressHandler.bind(this);
        this.changePartyTypeHandler = this.changePartyTypeHandler.bind(this);
    }

    componentDidMount(){
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

    updateThirdParty = (e) => {
        e.preventDefault();
        let thirdParty = {
            thirdPartyId: this.state.id,
            name: this.state.name,
            taxId: this.state.taxId,
            address: this.state.address,
            partyType: this.state.partyType
        };
        console.log('thirdParty => ' + JSON.stringify(thirdParty));
        console.log('id => ' + JSON.stringify(this.state.id));
        ThirdPartyService.updateThirdParty(thirdParty).then( res => {
            this.props.history.push('/thirdPartys');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update ThirdParty</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> taxId: </label>
                                                <input placeholder="taxId" name="taxId" className="form-control" value={this.state.taxId} onChange={this.changetaxIdHandler}/>

                                            <label> address: </label>
                                                <input placeholder="address" name="address" className="form-control" value={this.state.address} onChange={this.changeaddressHandler}/>

                                            <label> PartyType: </label>
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
                                        <button className="btn btn-success" onClick={this.updateThirdParty}>Save</button>
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

export default UpdateThirdPartyComponent
