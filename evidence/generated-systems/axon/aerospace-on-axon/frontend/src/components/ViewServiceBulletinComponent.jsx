import React, { Component } from 'react'
import ServiceBulletinService from '../services/ServiceBulletinService'

class ViewServiceBulletinComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            serviceBulletin: {}
        }
    }

    componentDidMount(){
        ServiceBulletinService.getServiceBulletinById(this.state.id).then( res => {
            this.setState({serviceBulletin: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View ServiceBulletin Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> bulletinNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.serviceBulletin.bulletinNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Category:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.serviceBulletin.category }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewServiceBulletinComponent
