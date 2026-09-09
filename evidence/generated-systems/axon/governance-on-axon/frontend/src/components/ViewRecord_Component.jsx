import React, { Component } from 'react'
import Record_Service from '../services/Record_Service'

class ViewRecord_Component extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            record_: {}
        }
    }

    componentDidMount(){
        Record_Service.getRecord_ById(this.state.id).then( res => {
            this.setState({record_: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Record_ Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> title:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.record_.title }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> creationDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.record_.creationDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> RecordType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.record_.recordType }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Classification:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.record_.classification }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.record_.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewRecord_Component
