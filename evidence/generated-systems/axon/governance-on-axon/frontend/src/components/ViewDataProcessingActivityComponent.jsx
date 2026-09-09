import React, { Component } from 'react'
import DataProcessingActivityService from '../services/DataProcessingActivityService'

class ViewDataProcessingActivityComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
            dataProcessingActivity: {}
        }
    }

    componentDidMount(){
        DataProcessingActivityService.getDataProcessingActivityById(this.state.id).then( res => {
            this.setState({dataProcessingActivity: res.data});
        })
    }

    render() {
        return (
            <div>
                <br></br>
                <div className = "card col-md-6 offset-md-3">
                    <h3 className = "text-center"> View DataProcessingActivity Details</h3>
                    <div className = "card-body">
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> name:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dataProcessingActivity.name }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> purpose:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dataProcessingActivity.purpose }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> startDate:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dataProcessingActivity.startDate }</div>
                        </div>
                        <div className = "row">
                            <div className = "col" style={{textAlign:"right"}}><label> LawfulBasis:&emsp; </label></div>
                            <div className = "col" style={{textAlign:"left"}}> { this.state.dataProcessingActivity.lawfulBasis }</div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default ViewDataProcessingActivityComponent
