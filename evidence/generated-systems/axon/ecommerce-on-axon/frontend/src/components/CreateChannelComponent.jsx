import React, { Component } from 'react'
import ChannelService from '../services/ChannelService';

class CreateChannelComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                channelCode: '',
                locale: '',
                domain: '',
                asActive: '',
                defaultCurrency: '',
                channelType: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changechannelCodeHandler = this.changechannelCodeHandler.bind(this);
        this.changelocaleHandler = this.changelocaleHandler.bind(this);
        this.changedomainHandler = this.changedomainHandler.bind(this);
        this.changeasActiveHandler = this.changeasActiveHandler.bind(this);
        this.changedefaultCurrencyHandler = this.changedefaultCurrencyHandler.bind(this);
        this.changeChannelTypeHandler = this.changeChannelTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ChannelService.getChannelById(this.state.id).then( (res) =>{
                let channel = res.data;
                this.setState({
                    name: channel.name,
                    channelCode: channel.channelCode,
                    locale: channel.locale,
                    domain: channel.domain,
                    asActive: channel.asActive,
                    defaultCurrency: channel.defaultCurrency,
                    channelType: channel.channelType
                });
            });
        }        
    }
    saveOrUpdateChannel = (e) => {
        e.preventDefault();
        let channel = {
                channelId: this.state.id,
                name: this.state.name,
                channelCode: this.state.channelCode,
                locale: this.state.locale,
                domain: this.state.domain,
                asActive: this.state.asActive,
                defaultCurrency: this.state.defaultCurrency,
                channelType: this.state.channelType
            };
        console.log('channel => ' + JSON.stringify(channel));

        // step 5
        if(this.state.id === '_add'){
            channel.channelId=''
            ChannelService.createChannel(channel).then(res =>{
                this.props.history.push('/channels');
            });
        }else{
            ChannelService.updateChannel(channel).then( res => {
                this.props.history.push('/channels');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changechannelCodeHandler= (event) => {
        this.setState({channelCode: event.target.value});
    }
    changelocaleHandler= (event) => {
        this.setState({locale: event.target.value});
    }
    changedomainHandler= (event) => {
        this.setState({domain: event.target.value});
    }
    changeasActiveHandler= (event) => {
        this.setState({asActive: event.target.value});
    }
    changedefaultCurrencyHandler= (event) => {
        this.setState({defaultCurrency: event.target.value});
    }
    changeChannelTypeHandler= (event) => {
        this.setState({channelType: event.target.value});
    }

    cancel(){
        this.props.history.push('/channels');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Channel</h3>
        }else{
            return <h3 className="text-center">Update Channel</h3>
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

                                            <label> channelCode:&emsp; </label>
                                                <input placeholder="channelCode" name="channelCode" className="form-control" value={this.state.channelCode} onChange={this.changechannelCodeHandler}/>

                                            <label> locale:&emsp; </label>
                                                <input placeholder="locale" name="locale" className="form-control" value={this.state.locale} onChange={this.changelocaleHandler}/>

                                            <label> domain:&emsp; </label>
                                                <input placeholder="domain" name="domain" className="form-control" value={this.state.domain} onChange={this.changedomainHandler}/>

                                            <label> asActive:&emsp; </label>
                                                <input type="checkbox" placeholder="asActive" name="asActive" className="form-control" value={this.state.asActive} onChange={this.changeasActiveHandler}/>


                                            <label> defaultCurrency:&emsp; </label>
                                                <input placeholder="defaultCurrency" name="defaultCurrency" className="form-control" value={this.state.defaultCurrency} onChange={this.changedefaultCurrencyHandler}/>

                                            <label> ChannelType:&emsp; </label>
                                                <select value={this.state.channelType} onChange={this.changeChannelTypeHandler}>
                      <option name="ChannelType" className="form-control" >
                          Web
                      </option>
                      <option name="ChannelType" className="form-control" >
                          MobileApp
                      </option>
                      <option name="ChannelType" className="form-control" >
                          Marketplace
                      </option>
                      <option name="ChannelType" className="form-control" >
                          Social
                      </option>
                      <option name="ChannelType" className="form-control" >
                          POS
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateChannel}>Save</button>
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

export default CreateChannelComponent
