import React, { Component } from 'react'
import ChannelService from '../services/ChannelService'

class ListChannelComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                channels: []
        }
        this.addChannel = this.addChannel.bind(this);
        this.editChannel = this.editChannel.bind(this);
        this.deleteChannel = this.deleteChannel.bind(this);
    }

    deleteChannel(id){
        ChannelService.deleteChannel(id).then( res => {
            this.setState({channels: this.state.channels.filter(channel => channel.channelId !== id)});
        });
    }
    viewChannel(id){
        this.props.history.push(`/view-channel/${id}`);
    }
    editChannel(id){
        this.props.history.push(`/add-channel/${id}`);
    }

    componentDidMount(){
        ChannelService.getChannels().then((res) => {
            this.setState({ channels: res.data});
        });
    }

    addChannel(){
        this.props.history.push('/add-channel/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Channel List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addChannel}> Add Channel</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> ChannelCode </th>
                                    <th> Locale </th>
                                    <th> Domain </th>
                                    <th> AsActive </th>
                                    <th> DefaultCurrency </th>
                                    <th> ChannelType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.channels.map(
                                        channel => 
                                        <tr key = {channel.channelId}>
                                             <td> { channel.name } </td>
                                             <td> { channel.channelCode } </td>
                                             <td> { channel.locale } </td>
                                             <td> { channel.domain } </td>
                                             <td> { channel.asActive } </td>
                                             <td> { channel.defaultCurrency } </td>
                                             <td> { channel.channelType } </td>
                                             <td>
                                                 <button onClick={ () => this.editChannel(channel.channelId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteChannel(channel.channelId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewChannel(channel.channelId)} className="btn btn-outline-info btn-sm">View </button>
                                             </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>

                 </div>

            </div>
        )
    }
}

export default ListChannelComponent
