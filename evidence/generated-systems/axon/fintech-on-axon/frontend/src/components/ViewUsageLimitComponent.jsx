import React, { Component } from 'react'
import UsageLimitService from '../services/UsageLimitService'

class ViewUsageLimitComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            usageLimit: {}
        }
    }

    componentDidMount(){
        UsageLimitService.getUsageLimitById(this.state.id).then( res => {
            this.setState({usageLimit: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View UsageLimit Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.usageLimit.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> amount:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.usageLimit.amount }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> count:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.usageLimit.count }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Scope:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.usageLimit.scope }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Period:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.usageLimit.period }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewUsageLimitComponent
