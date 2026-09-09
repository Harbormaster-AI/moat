import React, { Component } from 'react'
import Case_Service from '../services/Case_Service'

class ViewCase_Component extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            case_: {}
        }
    }

    componentDidMount(){
        Case_Service.getCase_ById(this.state.id).then( res => {
            this.setState({case_: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Case_ Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> caseNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.case_.caseNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> subject:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.case_.subject }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> description:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.case_.description }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> slaDue:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.case_.slaDue }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.case_.status }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Priority:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.case_.priority }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Origin:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.case_.origin }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Severity:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.case_.severity }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewCase_Component
