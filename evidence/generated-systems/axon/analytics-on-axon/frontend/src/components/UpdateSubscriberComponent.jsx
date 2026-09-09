import React, { Component } from 'react'
import SubscriberService from '../services/SubscriberService';

class UpdateSubscriberComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                address: '',
                channel: ''
        }
        this.updateSubscriber = this.updateSubscriber.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeaddressHandler = this.changeaddressHandler.bind(this);
        this.changeChannelHandler = this.changeChannelHandler.bind(this);
    }

    componentDidMount(){
        SubscriberService.getSubscriberById(this.state.id).then( (res) =>{
            let subscriber = res.data;
            this.setState({
                name: subscriber.name,
                address: subscriber.address,
                channel: subscriber.channel
            });
        });
    }

    updateSubscriber = (e) => {
        e.preventDefault();
        let subscriber = {
            subscriberId: this.state.id,
            name: this.state.name,
            address: this.state.address,
            channel: this.state.channel
        };
        console.log('subscriber => ' + JSON.stringify(subscriber));
        console.log('id => ' + JSON.stringify(this.state.id));
        SubscriberService.updateSubscriber(subscriber).then( res => {
            this.props.history.push('/subscribers');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeaddressHandler= (event) => {
        this.setState({address: event.target.value});
    }
    changeChannelHandler= (event) => {
        this.setState({channel: event.target.value});
    }

    cancel(){
        this.props.history.push('/subscribers');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Subscriber</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> address: </label>
                                                <input placeholder="address" name="address" className="form-control" value={this.state.address} onChange={this.changeaddressHandler}/>

                                            <label> Channel: </label>
                                                <select value={this.state.channel} onChange={this.changeChannelHandler}>
                      <option name="Channel" className="form-control" >
                          Email
                      </option>
                      <option name="Channel" className="form-control" >
                          SMS
                      </option>
                      <option name="Channel" className="form-control" >
                          Webhook
                      </option>
                      <option name="Channel" className="form-control" >
                          Chat
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateSubscriber}>Save</button>
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

export default UpdateSubscriberComponent
