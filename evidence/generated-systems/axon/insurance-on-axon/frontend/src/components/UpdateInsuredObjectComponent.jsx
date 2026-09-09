import React, { Component } from 'react'
import InsuredObjectService from '../services/InsuredObjectService';

class UpdateInsuredObjectComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                description: '',
                serialOrId: '',
                primaryAddress: '',
                objectType: ''
        }
        this.updateInsuredObject = this.updateInsuredObject.bind(this);

        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
        this.changeserialOrIdHandler = this.changeserialOrIdHandler.bind(this);
        this.changeprimaryAddressHandler = this.changeprimaryAddressHandler.bind(this);
        this.changeObjectTypeHandler = this.changeObjectTypeHandler.bind(this);
    }

    componentDidMount(){
        InsuredObjectService.getInsuredObjectById(this.state.id).then( (res) =>{
            let insuredObject = res.data;
            this.setState({
                description: insuredObject.description,
                serialOrId: insuredObject.serialOrId,
                primaryAddress: insuredObject.primaryAddress,
                objectType: insuredObject.objectType
            });
        });
    }

    updateInsuredObject = (e) => {
        e.preventDefault();
        let insuredObject = {
            insuredObjectId: this.state.id,
            description: this.state.description,
            serialOrId: this.state.serialOrId,
            primaryAddress: this.state.primaryAddress,
            objectType: this.state.objectType
        };
        console.log('insuredObject => ' + JSON.stringify(insuredObject));
        console.log('id => ' + JSON.stringify(this.state.id));
        InsuredObjectService.updateInsuredObject(insuredObject).then( res => {
            this.props.history.push('/insuredObjects');
        });
    }

    changedescriptionHandler= (event) => {
        this.setState({description: event.target.value});
    }
    changeserialOrIdHandler= (event) => {
        this.setState({serialOrId: event.target.value});
    }
    changeprimaryAddressHandler= (event) => {
        this.setState({primaryAddress: event.target.value});
    }
    changeObjectTypeHandler= (event) => {
        this.setState({objectType: event.target.value});
    }

    cancel(){
        this.props.history.push('/insuredObjects');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update InsuredObject</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> description: </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

                                            <label> serialOrId: </label>
                                                <input placeholder="serialOrId" name="serialOrId" className="form-control" value={this.state.serialOrId} onChange={this.changeserialOrIdHandler}/>

                                            <label> primaryAddress: </label>
                                                <input placeholder="primaryAddress" name="primaryAddress" className="form-control" value={this.state.primaryAddress} onChange={this.changeprimaryAddressHandler}/>

                                            <label> ObjectType: </label>
                                                <select value={this.state.objectType} onChange={this.changeObjectTypeHandler}>
                      <option name="ObjectType" className="form-control" >
                          Vehicle
                      </option>
                      <option name="ObjectType" className="form-control" >
                          Property
                      </option>
                      <option name="ObjectType" className="form-control" >
                          Person
                      </option>
                      <option name="ObjectType" className="form-control" >
                          Equipment
                      </option>
                      <option name="ObjectType" className="form-control" >
                          LiabilityExposure
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateInsuredObject}>Save</button>
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

export default UpdateInsuredObjectComponent
