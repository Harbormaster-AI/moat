import React, { Component } from 'react'
import PrivacyNoticeService from '../services/PrivacyNoticeService'

class ViewPrivacyNoticeComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            privacyNotice: {}
        }
    }

    componentDidMount(){
        PrivacyNoticeService.getPrivacyNoticeById(this.state.id).then( res => {
            this.setState({privacyNotice: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View PrivacyNotice Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> title:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.privacyNotice.title }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> audience:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.privacyNotice.audience }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> versionLabel:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.privacyNotice.versionLabel }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> publicationDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.privacyNotice.publicationDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> publicationUrl:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.privacyNotice.publicationUrl }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.privacyNotice.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewPrivacyNoticeComponent
