import React, { Component } from 'react'
import InventorySourceService from '../services/InventorySourceService';

class UpdateInventorySourceComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                domain: '',
                channel: '',
                primaryFormat: ''
        }
        this.updateInventorySource = this.updateInventorySource.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changedomainHandler = this.changedomainHandler.bind(this);
        this.changeChannelHandler = this.changeChannelHandler.bind(this);
        this.changePrimaryFormatHandler = this.changePrimaryFormatHandler.bind(this);
    }

    componentDidMount(){
        InventorySourceService.getInventorySourceById(this.state.id).then( (res) =>{
            let inventorySource = res.data;
            this.setState({
                name: inventorySource.name,
                domain: inventorySource.domain,
                channel: inventorySource.channel,
                primaryFormat: inventorySource.primaryFormat
            });
        });
    }

    updateInventorySource = (e) => {
        e.preventDefault();
        let inventorySource = {
            inventorySourceId: this.state.id,
            name: this.state.name,
            domain: this.state.domain,
            channel: this.state.channel,
            primaryFormat: this.state.primaryFormat
        };
        console.log('inventorySource => ' + JSON.stringify(inventorySource));
        console.log('id => ' + JSON.stringify(this.state.id));
        InventorySourceService.updateInventorySource(inventorySource).then( res => {
            this.props.history.push('/inventorySources');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changedomainHandler= (event) => {
        this.setState({domain: event.target.value});
    }
    changeChannelHandler= (event) => {
        this.setState({channel: event.target.value});
    }
    changePrimaryFormatHandler= (event) => {
        this.setState({primaryFormat: event.target.value});
    }

    cancel(){
        this.props.history.push('/inventorySources');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update InventorySource</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> domain: </label>
                                                <input placeholder="domain" name="domain" className="form-control" value={this.state.domain} onChange={this.changedomainHandler}/>

                                            <label> Channel: </label>
                                                <select value={this.state.channel} onChange={this.changeChannelHandler}>
                      <option name="Channel" className="form-control" >
                          Programmatic
                      </option>
                      <option name="Channel" className="form-control" >
                          Direct
                      </option>
                      <option name="Channel" className="form-control" >
                          Search
                      </option>
                      <option name="Channel" className="form-control" >
                          Social
                      </option>
                      <option name="Channel" className="form-control" >
                          Email
                      </option>
                      <option name="Channel" className="form-control" >
                          Affiliate
                      </option>
                      <option name="Channel" className="form-control" >
                          DOOH
                      </option>
                    </select>

                                            <label> PrimaryFormat: </label>
                                                <select value={this.state.primaryFormat} onChange={this.changePrimaryFormatHandler}>
                      <option name="PrimaryFormat" className="form-control" >
                          Banner
                      </option>
                      <option name="PrimaryFormat" className="form-control" >
                          Video
                      </option>
                      <option name="PrimaryFormat" className="form-control" >
                          Native
                      </option>
                      <option name="PrimaryFormat" className="form-control" >
                          Audio
                      </option>
                      <option name="PrimaryFormat" className="form-control" >
                          Interstitial
                      </option>
                      <option name="PrimaryFormat" className="form-control" >
                          RichMedia
                      </option>
                      <option name="PrimaryFormat" className="form-control" >
                          SearchText
                      </option>
                      <option name="PrimaryFormat" className="form-control" >
                          SocialPost
                      </option>
                      <option name="PrimaryFormat" className="form-control" >
                          CTVVideo
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateInventorySource}>Save</button>
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

export default UpdateInventorySourceComponent
