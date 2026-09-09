import React, { Component } from 'react'
import UserService from '../services/UserService'

class ViewUserComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            user: {}
        }
    }

    componentDidMount(){
        UserService.getUserById(this.state.id).then( res => {
            this.setState({user: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View User Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> username:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.user.username }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> fullName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.user.fullName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> email:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.user.email }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> locale:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.user.locale }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Role:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.user.role }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.user.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewUserComponent
