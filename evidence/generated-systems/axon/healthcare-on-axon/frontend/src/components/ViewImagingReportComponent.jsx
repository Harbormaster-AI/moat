import React, { Component } from 'react'
import ImagingReportService from '../services/ImagingReportService'

class ViewImagingReportComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            imagingReport: {}
        }
    }

    componentDidMount(){
        ImagingReportService.getImagingReportById(this.state.id).then( res => {
            this.setState({imagingReport: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View ImagingReport Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> reportNumber:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.imagingReport.reportNumber }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> impression:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.imagingReport.impression }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> reportedDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.imagingReport.reportedDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> Status:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.imagingReport.status }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewImagingReportComponent
