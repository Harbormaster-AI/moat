import React, { Component } from 'react'
import InsuredObjectService from '../services/InsuredObjectService';

class CreateInsuredObjectComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                description: '',
                serialOrId: '',
                primaryAddress: '',
                objectType: ''
        }
        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
        this.changeserialOrIdHandler = this.changeserialOrIdHandler.bind(this);
        this.changeprimaryAddressHandler = this.changeprimaryAddressHandler.bind(this);
        this.changeObjectTypeHandler = this.changeObjectTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateInsuredObject = (e) => {
        e.preventDefault();
        let insuredObject = {
                insuredObjectId: this.state.id,
                description: this.state.description,
                serialOrId: this.state.serialOrId,
                primaryAddress: this.state.primaryAddress,
                objectType: this.state.objectType
            };
        console.log('insuredObject => ' + JSON.stringify(insuredObject));

        // step 5
        if(this.state.id === '_add'){
            insuredObject.insuredObjectId=''
            InsuredObjectService.createInsuredObject(insuredObject).then(res =>{
                this.props.history.push('/insuredObjects');
            });
        }else{
            InsuredObjectService.updateInsuredObject(insuredObject).then( res => {
                this.props.history.push('/insuredObjects');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add InsuredObject</h3>
        }else{
            return <h3 className="text-center">Update InsuredObject</h3>
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
                                            <label> description:&emsp; </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

                                            <label> serialOrId:&emsp; </label>
                                                <input placeholder="serialOrId" name="serialOrId" className="form-control" value={this.state.serialOrId} onChange={this.changeserialOrIdHandler}/>

                                            <label> primaryAddress:&emsp; </label>
                                                <input placeholder="primaryAddress" name="primaryAddress" className="form-control" value={this.state.primaryAddress} onChange={this.changeprimaryAddressHandler}/>

                                            <label> ObjectType:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateInsuredObject}>Save</button>
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

export default CreateInsuredObjectComponent
