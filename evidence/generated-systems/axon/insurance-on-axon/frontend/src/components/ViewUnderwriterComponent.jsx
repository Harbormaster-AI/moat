import React, { Component } from 'react'
import UnderwriterService from '../services/UnderwriterService'

class ViewUnderwriterComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            underwriter: {}
        }
    }

    componentDidMount(){
        UnderwriterService.getUnderwriterById(this.state.id).then( res => {
            this.setState({underwriter: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Underwriter Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> firstName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.underwriter.firstName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> lastName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.underwriter.lastName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> employeeId:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.underwriter.employeeId }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> authorityLimit:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.underwriter.authorityLimit }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewUnderwriterComponent
