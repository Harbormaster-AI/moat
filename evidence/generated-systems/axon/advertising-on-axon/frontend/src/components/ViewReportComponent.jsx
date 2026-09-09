import React, { Component } from 'react'
import ReportService from '../services/ReportService'

class ViewReportComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            report: {}
        }
    }

    componentDidMount(){
        ReportService.getReportById(this.state.id).then( res => {
            this.setState({report: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View Report Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> reportName:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.report.reportName }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> generatedAt:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.report.generatedAt }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> fileUrl:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.report.fileUrl }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> ReportType:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.report.reportType }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewReportComponent
