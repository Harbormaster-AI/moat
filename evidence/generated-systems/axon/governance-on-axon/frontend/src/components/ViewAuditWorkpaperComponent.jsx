import React, { Component } from 'react'
import AuditWorkpaperService from '../services/AuditWorkpaperService'

class ViewAuditWorkpaperComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            auditWorkpaper: {}
        }
    }

    componentDidMount(){
        AuditWorkpaperService.getAuditWorkpaperById(this.state.id).then( res => {
            this.setState({auditWorkpaper: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View AuditWorkpaper Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> workpaperRef:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.auditWorkpaper.workpaperRef }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> subject:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.auditWorkpaper.subject }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> workpaperUrl:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.auditWorkpaper.workpaperUrl }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewAuditWorkpaperComponent
