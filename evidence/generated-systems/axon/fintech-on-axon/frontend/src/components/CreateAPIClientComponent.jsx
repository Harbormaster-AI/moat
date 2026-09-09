import React, { Component } from 'react'
import APIClientService from '../services/APIClientService';

class CreateAPIClientComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                clientId: '',
                redirectUri: '',
                clientType: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeclientIdHandler = this.changeclientIdHandler.bind(this);
        this.changeredirectUriHandler = this.changeredirectUriHandler.bind(this);
        this.changeClientTypeHandler = this.changeClientTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            APIClientService.getAPIClientById(this.state.id).then( (res) =>{
                let aPIClient = res.data;
                this.setState({
                    name: aPIClient.name,
                    clientId: aPIClient.clientId,
                    redirectUri: aPIClient.redirectUri,
                    clientType: aPIClient.clientType
                });
            });
        }        
    }
    saveOrUpdateAPIClient = (e) => {
        e.preventDefault();
        let aPIClient = {
                aPIClientId: this.state.id,
                name: this.state.name,
                clientId: this.state.clientId,
                redirectUri: this.state.redirectUri,
                clientType: this.state.clientType
            };
        console.log('aPIClient => ' + JSON.stringify(aPIClient));

        // step 5
        if(this.state.id === '_add'){
            aPIClient.aPIClientId=''
            APIClientService.createAPIClient(aPIClient).then(res =>{
                this.props.history.push('/aPIClients');
            });
        }else{
            APIClientService.updateAPIClient(aPIClient).then( res => {
                this.props.history.push('/aPIClients');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeclientIdHandler= (event) => {
        this.setState({clientId: event.target.value});
    }
    changeredirectUriHandler= (event) => {
        this.setState({redirectUri: event.target.value});
    }
    changeClientTypeHandler= (event) => {
        this.setState({clientType: event.target.value});
    }

    cancel(){
        this.props.history.push('/aPIClients');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add APIClient</h3>
        }else{
            return <h3 className="text-center">Update APIClient</h3>
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

                                            <label> clientId:&emsp; </label>
                                                <input placeholder="clientId" name="clientId" className="form-control" value={this.state.clientId} onChange={this.changeclientIdHandler}/>

                                            <label> redirectUri:&emsp; </label>
                                                <input placeholder="redirectUri" name="redirectUri" className="form-control" value={this.state.redirectUri} onChange={this.changeredirectUriHandler}/>

                                            <label> ClientType:&emsp; </label>
                                                <select value={this.state.clientType} onChange={this.changeClientTypeHandler}>
                      <option name="ClientType" className="form-control" >
                          Confidential
                      </option>
                      <option name="ClientType" className="form-control" >
                          Public
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateAPIClient}>Save</button>
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

export default CreateAPIClientComponent
