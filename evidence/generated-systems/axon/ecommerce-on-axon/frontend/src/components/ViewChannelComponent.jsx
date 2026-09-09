import React, { Component } from 'react'
import ChannelService from '../services/ChannelService'

class ViewChannelComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            channel: {}
        }
    }

    componentDidMount(){
        ChannelService.getChannelById(this.state.id).then( res => {
            this.setState({channel: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Channel Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.channel.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> channelCode:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.channel.channelCode }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> locale:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.channel.locale }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> domain:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.channel.domain }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> asActive:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.channel.asActive }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> defaultCurrency:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.channel.defaultCurrency }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ChannelType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.channel.channelType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewChannelComponent
